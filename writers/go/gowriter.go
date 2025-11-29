package gowriter

import (
	"github.com/b3nten/ssss/parser"
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func printType(t parser.Type) string {
	switch t.TypeKind() {
		case "primitive":
			pt := t.(parser.PrimitiveType)
			return pt.Name
		case "struct":
			st := t.(*parser.StructType)
			return toPascalCase(st.Name)
		case "list":
			lt := t.(parser.ListType)
			return fmt.Sprintf("[]%s", printType(lt.ElementType))
		default:
			return "unknown"
	}
}

func printField(f parser.Field) string {
	return fmt.Sprintf(
		"\n\t %s *%s",
		toPascalCase(f.Name),
		printType(f.Type),
	)
}

func printListSerializer(lt parser.ListType) string {
	switch lt.ElementType.TypeKind() {
		case "primitive":
			pt := lt.ElementType.(parser.PrimitiveType)
			return fmt.Sprintf("newListSerializer[%s](serialize%s)", pt.Name, capFirst(pt.Name))
		case "struct":
			st := lt.ElementType.(*parser.StructType)
			return fmt.Sprintf("newListSerializer[%s](serializeStruct[%s])", toPascalCase(st.Name), toPascalCase(st.Name))
		case "list":
			nlt := lt.ElementType.(parser.ListType)
			return fmt.Sprintf("newListSerializer[%s](%s)", printType(nlt), printListSerializer(nlt))
		default:
			return "unknown"
	}
}

func printListDeserializer(lt parser.ListType) string {
	switch lt.ElementType.TypeKind() {
		case "primitive":
			pt := lt.ElementType.(parser.PrimitiveType)
			return fmt.Sprintf("newListDeserializer[%s](deserialize%s)", pt.Name, capFirst(pt.Name))
		case "struct":
			st := lt.ElementType.(*parser.StructType)
			return fmt.Sprintf("newListDeserializer[%s](deserializeStruct[%s])", toPascalCase(st.Name), toPascalCase(st.Name))
		case "list":
			nlt := lt.ElementType.(parser.ListType)
			return fmt.Sprintf("newListDeserializer[%s](%s)", printType(nlt), printListDeserializer(nlt))
		default:
			return "unknown"
	}
}

func printStruct(sv parser.StructType) string {
	sb := strings.Builder{}

	// print struct
	sb.WriteString(fmt.Sprintf("type %s struct { ",  toPascalCase(sv.Name)))
	for _, f := range sv.Fields {
		sb.WriteString(printField(f))
	}
	sb.WriteString("\n}\n")

	// print TypeID
	sb.WriteString(fmt.Sprintf("func (%s) TypeID() uint16 { return uint16(%d) }\n", toPascalCase(sv.Name), sv.ID))

	// print toBytes
	sb.WriteString(fmt.Sprintf("func (it %s) toBytes(data *bytes.Buffer) {\n", toPascalCase(sv.Name)))
	sb.WriteString(fmt.Sprintf("\tserializeUint16(%d, data)\n", sv.ID))
	sb.WriteString("\tstartLenPos := data.Len()\n")
	sb.WriteString("\tserializeUint32(0, data)\n")
	for _, field := range sv.Fields {
		sb.WriteString(fmt.Sprintf("\tif it.%s != nil {\n", toPascalCase(field.Name)))
		sb.WriteString(fmt.Sprintf("\t\tserializeUint16(%d, data)\n", field.ID))

		switch field.Type.TypeKind() {
			case "primitive":
				pt := field.Type.(parser.PrimitiveType)
				sb.WriteString(fmt.Sprintf("\t\tserialize%s(*it.%s, data)\n", capFirst(pt.Name), toPascalCase(field.Name)))
			case "struct":
				sb.WriteString(fmt.Sprintf("\t\tserializeStruct(*it.%s, data)\n", toPascalCase(field.Name)))
			case "list":
				sb.WriteString(
					fmt.Sprintf("\t\t%s(*it.%s, data)\n",
						printListSerializer(field.Type.(parser.ListType)),
						toPascalCase(field.Name),
					),
				)
		}

		sb.WriteString("\t}\n")
	}
	sb.WriteString("\tbinary.BigEndian.PutUint32(data.Bytes()[startLenPos:], uint32(len(data.Bytes())-(startLenPos+lenSize)))")
	sb.WriteString("\n}\n")

	// print fromBytes
	sb.WriteString(fmt.Sprintf("func (it *%s) fromBytes(data []byte, fieldIndex uint16, offset int) (int, error) {\n", toPascalCase(sv.Name)))
	sb.WriteString("\tswitch fieldIndex {\n")
	for i, field := range sv.Fields {
		var name string
		switch field.Type.TypeKind() {
			case "primitive":
				pt := field.Type.(parser.PrimitiveType)
				name = "deserialize" + capFirst(pt.Name)
			case "struct":
				name = "deserialize" + "Struct[" + toPascalCase(field.Type.(*parser.StructType).Name) + "]"
			case "list":
				name = printListDeserializer(field.Type.(parser.ListType))
		}
		sb.WriteString(
			fmt.Sprintf(
				"\tcase %d: val, len, err := %s(data, offset); it.%s = &val; return len, err\n",
			 	i,
				name,
				toPascalCase(field.Name),
			),
		)
	}
	sb.WriteString("\t}\n")
	sb.WriteString("\treturn 0, UnknownFieldError\n")
	sb.WriteString("}\n")

	return sb.String()
}

func Print(s *parser.Schema, namespace string) (string, error) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("// Auto-generated code for schema: %s v%d\n\n", s.Name, s.Version))
	sb.WriteString(fmt.Sprintf("package %s\n\n", namespace))
	sb.WriteString("import(\n")
	sb.WriteString("\t\"bytes\"\n")
	sb.WriteString("\t\"encoding/binary\"\n")
	sb.WriteString("\t\"errors\"\n")
	sb.WriteString("\t\"fmt\"\n")
	sb.WriteString(")\n\n")

	for _, str := range s.Structs {
		sb.WriteString(printStruct(str) + "\n")
	}

	// Deserialize function
	sb.WriteString("func UnmarshalBinary[K any, KT interface {*K; Serializable}](b []byte, out KT) error {\n")
	sb.WriteString("\tif len(b) < idSize+lenSize { return fmt.Errorf(\"data too short to contain message header\") }\n")
	sb.WriteString("\ttypeID := uint16(binary.BigEndian.Uint16(b[0:idSize]))\n")
	sb.WriteString("\tif out.TypeID() != typeID { return fmt.Errorf(\"type ID mismatch: expected %d, got %d\", out.TypeID(), typeID) }\n")
	sb.WriteString("\tswitch v := any(out).(type) {\n")
	for _, str := range s.Structs {
		sb.WriteString(fmt.Sprintf("\tcase *%s: _, err := parse(v.fromBytes, b); return err\n", toPascalCase(str.Name)))
	}
	sb.WriteString("\tdefault: return fmt.Errorf(\"unsupported type for deserialization\")\n\t}\n}\n\n")

	// deserializeStruct function
	sb.WriteString("func deserializeStruct[K Serializable](data []byte, offset int) (K, int, error) {\n")
	sb.WriteString("\tvar val K\n")
	sb.WriteString("\tslice := data[offset:]\n")
	sb.WriteString("\tswitch v := any(val).(type) {\n")
	for _, str := range s.Structs {
		sb.WriteString(fmt.Sprintf("\tcase %s:\n", toPascalCase(str.Name)))
		sb.WriteString("\t\ti, err := parse(v.fromBytes, slice)\n")
		sb.WriteString("\t\treturn any(v).(K), i, err\n")
	}
	sb.WriteString("\tdefault:\n")
	sb.WriteString("\t\treturn val, 0, fmt.Errorf(\"unsupported struct type for deserialization\")\n")
	sb.WriteString("\t}\n}\n\n")


	postlude, err := readAfterLine("schema/gowriter/helpers", 9)
	if err != nil {
		return "", err
	}
	sb.WriteString(postlude)
	return sb.String(), nil
}

func capFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func toPascalCase(s string) string {
	if s == "" {
		return ""
	}

	var result strings.Builder
	capitalizeNext := true

	for i, r := range s {
		if r == '_' || r == '-' || r == ' ' || r == '.' {
			// Treat these as word separators
			capitalizeNext = true
			continue
		}

		if unicode.IsUpper(r) && i > 0 {
			// If we encounter an uppercase letter that's not at the start,
			// it might be camelCase, so capitalize it
			result.WriteRune(r)
			capitalizeNext = false
		} else if capitalizeNext {
			result.WriteRune(unicode.ToUpper(r))
			capitalizeNext = false
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}

	return result.String()
}

func readAfterLine(filename string, lineNum int) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result strings.Builder
	currentLine := 0

	for scanner.Scan() {
		currentLine++
		if currentLine > lineNum {
			result.WriteString(scanner.Text())
			result.WriteString("\n")
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return result.String(), nil
}
