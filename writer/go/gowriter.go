package gowriter

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"

	"github.com/b3nten/ssss/parser"
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
	sb.WriteString(fmt.Sprintf("type %s struct { ", toPascalCase(sv.Name)))
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
			capitalizeNext = true
			continue
		}
		if unicode.IsUpper(r) && i > 0 {
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

var postlude = `
type Serializable interface {
	TypeID() uint16
	toBytes(*bytes.Buffer)
}

func MarshalBinary[K any, KT interface {
	*K
	Serializable
}](value KT) ([]byte, error) {
	buf := bytes.Buffer{}
	value.toBytes(&buf)
	return buf.Bytes(), nil
}

func Ptr[K any](v K) *K {
	return &v
}

var UnknownFieldError = errors.New("unknown field index")

const lenSize = 4
const idSize = 2

type deserializer[K any] func(data []byte, offset int) (value K, len int, err error)

type serializer[K any] func(value K, b *bytes.Buffer) int

func deserializeBool(data []byte, offset int) (bool, int, error) {
	slice := data[offset:]
	if len(slice) < 1 {
		return false, 0, fmt.Errorf("insufficient data for bool")
	}
	val := slice[0] != 0
	return val, 1, nil
}

func deserializeInt8(data []byte, offset int) (int8, int, error) {
	slice := data[offset:]
	if len(slice) < 1 {
		return 0, 0, fmt.Errorf("insufficient data for int8")
	}
	val := int8(slice[0])
	return val, 1, nil
}

func deserializeUint8(data []byte, offset int) (uint8, int, error) {
	slice := data[offset:]
	if len(slice) < 1 {
		return 0, 0, fmt.Errorf("insufficient data for uint8")
	}
	val := uint8(slice[0])
	return val, 1, nil
}

func deserializeInt16(data []byte, offset int) (int16, int, error) {
	slice := data[offset:]
	if len(slice) < 2 {
		return 0, 0, fmt.Errorf("insufficient data for int16")
	}
	val := int16(slice[0])<<8 | int16(slice[1])
	return val, 2, nil
}

func deserializeUint16(data []byte, offset int) (uint16, int, error) {
	slice := data[offset:]
	if len(slice) < 2 {
		return 0, 0, fmt.Errorf("insufficient data for uint16")
	}
	val := uint16(slice[0])<<8 | uint16(slice[1])
	return val, 2, nil
}

func deserializeInt32(data []byte, offset int) (int32, int, error) {
	slice := data[offset:]
	if len(slice) < 4 {
		return 0, 0, fmt.Errorf("insufficient data for int32")
	}
	val := int32(slice[0])<<24 | int32(slice[1])<<16 | int32(slice[2])<<8 | int32(slice[3])
	return val, 4, nil
}

func deserializeUint32(data []byte, offset int) (uint32, int, error) {
	slice := data[offset:]
	if len(slice) < 4 {
		return 0, 0, fmt.Errorf("insufficient data for uint32")
	}
	val := uint32(slice[0])<<24 | uint32(slice[1])<<16 | uint32(slice[2])<<8 | uint32(slice[3])
	return val, 4, nil
}

func deserializeString(data []byte, offset int) (string, int, error) {
	slice := data[offset:]
	if len(slice) < 4 {
		return "", 0, fmt.Errorf("insufficient data for string length")
	}
	strLen := int(binary.LittleEndian.Uint32(slice))
	if strLen == 0 {
		return "", 4, nil
	}
	if len(slice[4:]) < strLen {
		return "", 0, fmt.Errorf("insufficient data for string content")
	}
	val := string(slice[4 : 4+strLen])
	return val, 4 + strLen, nil
}

func newListDeserializer[K any](elemDeserializer deserializer[K]) deserializer[[]K] {
	return func(data []byte, offset int) (value []K, l int, err error) {
		slice := data[offset:]
		if len(slice) < 4 {
			return nil, 0, fmt.Errorf("insufficient data for list length")
		}
		listLen := int(binary.LittleEndian.Uint32(slice))
		i := 4
		value = make([]K, 0)
		for i < listLen+4 {
			elem, n, err := elemDeserializer(data, offset+i)
			if err != nil {
				return nil, 0, err
			}
			value = append(value, elem)
			i += n
		}
		return value, listLen + 4, nil
	}
}

func serializeBool(value bool, b *bytes.Buffer) int {
	var boolByte byte = 0
	if value {
		boolByte = 1
	}
	b.WriteByte(boolByte)
	return 1
}

func serializeInt8(value int8, b *bytes.Buffer) int {
	b.WriteByte(byte(value))
	return 1
}

func serializeUint8(value uint8, b *bytes.Buffer) int {
	b.WriteByte(byte(value))
	return 1
}

func serializeInt16(value int16, b *bytes.Buffer) int {
	binary.Write(b, binary.LittleEndian, value)
	return 2
}

func serializeUint16(value uint16, b *bytes.Buffer) int {
	binary.Write(b, binary.LittleEndian, value)
	return 2
}

func serializeInt32(value int32, b *bytes.Buffer) int {
	binary.Write(b, binary.LittleEndian, value)
	return 4
}

func serializeUint32(value uint32, b *bytes.Buffer) int {
	binary.Write(b, binary.LittleEndian, value)
	return 4
}

func serializeString(value string, b *bytes.Buffer) int {
	strLen := uint32(len(value))
	binary.Write(b, binary.LittleEndian, strLen)
	b.WriteString(value)
	return 4 + len(value)
}

func serializeStruct[K Serializable](value K, b *bytes.Buffer) int {
	startLen := len(b.Bytes())
	value.toBytes(b)
	return len(b.Bytes()) - startLen
}

func newListSerializer[K any](elemSerializer serializer[K]) serializer[[]K] {
	return func(value []K, b *bytes.Buffer) int {
		serializeUint32(0, b)
		i := len(b.Bytes())
		for _, item := range value {
			elemSerializer(item, b)
		}
		binary.LittleEndian.PutUint32(b.Bytes()[i-lenSize:], uint32(len(b.Bytes())-i))
		return len(b.Bytes()) - i + lenSize
	}
}

func parse(parser func(data []byte, fieldIndex uint16, offset int) (int, error), bytes []byte) (int, error) {
	if len(bytes) < idSize+lenSize {
		return 0, fmt.Errorf("data too short: need at least 6 bytes for id and length")
	}
	dataSize := binary.LittleEndian.Uint32(bytes[idSize:])
	totalSize := int(dataSize) + idSize + lenSize
	if len(bytes) < totalSize {
		return 0, fmt.Errorf("data too short: expected %d bytes, got %d", totalSize, len(bytes))
	}
	for i := idSize + lenSize; i < totalSize; {
		fieldIndex, err := getField(bytes[i:])
		if err != nil {
			return totalSize, err
		}
		i += 2
		next, err := parser(bytes, fieldIndex, i)
		if err != nil {
			if errors.Is(err, UnknownFieldError) {
				return totalSize, nil
			} else {
				return totalSize, err
			}
		}
		i += next
	}
	return totalSize, nil
}

func getField(b []byte) (uint16, error) {
	if len(b) < 2 {
		return 0, fmt.Errorf("insufficient data for field index")
	}
	return binary.LittleEndian.Uint16(b), nil
}
`
