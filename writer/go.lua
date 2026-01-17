local function go_type(field)
	local t = field.name
	if t == "bool" then
		return "bool"
	elseif t == "int8" then
		return "int8"
	elseif t == "uint8" then
		return "uint8"
	elseif t == "int16" then
		return "int16"
	elseif t == "uint16" then
		return "uint16"
	elseif t == "int32" then
		return "int32"
	elseif t == "uint32" then
		return "uint32"
	elseif t == "int64" then
		return "int64"
	elseif t == "uint64" then
		return "uint64"
	elseif t == "f32" then
		return "float32"
	elseif t == "f64" then
		return "float64"
	elseif t == "string" then
		return "string"
	else
		error("unsupported field type: " .. tostring(t))
	end
end

local function print_struct_serializer(name, struct)
	local function print_writer(field, input)
		if field.type.kind == "primitive" then
			return [[
				err = b.Write(${input})
				if err != nil {
					return err
				}
			]] % { input = input }
		elseif field.type.kind == "struct" then
			return [[
				err = ${input}.serialize(b)
				if err != nil {
					return err
				}
			]] % { input = input }
		elseif field.type.kind == "list" then
			local i = field.index or 0
			return [[
				err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen${i} := b.Len()
				for _, item${i} := range ${input} {
					${item_writer}
				}
				listLen${i} := b.Len() - startLen${i}
				err = b.WriteLengthAt(listLen${i}, startLen${i}-4)
				if err != nil {
					return err
				}
			]] % {
				i = i,
				input = input,
				item_writer = print_writer({ type = field.type.of, index = i + 1 }, "item" .. i)
			}
		elseif field.type.kind == "map" then
			local i = field.index or 0
			return [[
				err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen${i} := b.Len()
				for key${i}, value${i} := range ${input} {
					${key_writer}
					${value_writer}
				}
				mapLen${i} := b.Len() - startLen${i}
				err = b.WriteLengthAt(mapLen${i}, startLen${i}-4)
				if err != nil {
					return err
				}
			]] % {
				i = i,
				input = input,
				key_writer = print_writer({ type = field.type.from, index = i + 1 }, "key" .. i),
				value_writer = print_writer({ type = field.type.to, index = i + 1 }, "value" .. i)
			}
		else
			return ""
		end
	end

	local function print_field_serializers()
		local out = ""
		for _, field in pairs(struct.fields) do
			local input = "s." .. pascal_case(field.name)
			if field.type.kind == "list" or field.type.kind == "map" then
				input = "*" .. input
			end
			out = out .. [[
			  if s.${fname} != nil {
				  err = b.WriteFieldId(${field_id})
				  if err != nil {
					  return err
				  }
				  ${writer}
			  }
		  ]] % {
				fname = pascal_case(field.name),
				field_id = field.id,
				writer = print_writer(field, input)
			}
		end
		return out
	end

	return [[
		func (s *${sname}) serialize(b *ByteWriter) error {
			err := b.WriteTypeId(${type_id})
			if err != nil {
				return err
			}
			err = b.WriteLength(0)
			if err != nil {
				return err
			}
			startLen := b.Len()
			${field_serializers}
			structLen := b.Len() - startLen
			err = b.WriteLengthAt(structLen, startLen-4)
			return err
		}
	]] % {
		type_id = struct.id,
		sname = pascal_case(name),
		field_serializers = print_field_serializers()
	}
end

local function print_struct_deserializer(name, struct)
	local function get_max_field_id()
		local max_id = 0
		for _, field in pairs(struct.fields) do
			if field.id > max_id then
				max_id = field.id
			end
		end
		return max_id
	end

	local function print_type(type)
		if type.kind == "primitive" then
			return go_type(type)
		elseif type.kind == "struct" then
			return pascal_case(type.name)
		elseif type.kind == "list" then
			return "[]${elem_type}" % {
				elem_type = print_type(type.of)
			}
		elseif type.kind == "map" then
			return "map[${key_type}]${value_type}" % {
				key_type = print_type(type.from),
				value_type = print_type(type.to)
			}
		end
		return ""
	end

	local function print_field_deserializer(field, output)
		if field.type.kind == 'primitive' then
			return [[ ${output} = new(${gtype}); err = br.Read(${output}) ]] % { output = output, gtype = go_type(field.type) }
		elseif field.type.kind == "struct" then
			return [[
				${output} = &${stype}{}
				err = ${output}.deserialize(br)
			]] % {
				output = output,
				stype = pascal_case(field.type.name)
			}
		elseif field.type.kind == "list" then
			local i = field.index or 0
			return [[
				listLen${i}, err := br.ReadLength()
				if err != nil || listLen${i} < 0 || listLen${i} > br.Remaining() {
					return fmt.Errorf("invalid list length: %d", listLen${i})
				}
				startPos${i} := br.Offset()
				${output} = &${list_type}{}
				for br.Offset() < startPos${i}+listLen${i} {
					var item${i} *${elem_type}
					${item_deserializer}
					if err != nil {
						return err
					}
					*${output} = append(*${output}, *item${i})
				}
			]] % {
				i = i,
				output = output,
				list_type = print_type(field.type),
				elem_type = print_type(field.type.of),
				item_deserializer = print_field_deserializer({ type = field.type.of, index = i + 1 }, "item" .. i)
			}
		elseif field.type.kind == "map" then
			local i = field.index or 0
			return [[
				mapLen${i}, err := br.ReadLength()
				if err != nil || mapLen${i} < 0 || mapLen${i} > br.Remaining() {
					return fmt.Errorf("invalid map length: %d", mapLen${i})
				}
				startPos${i} := br.Offset()
				${output} = &${map_type}{}
				for br.Offset() < startPos${i}+mapLen${i} {
					var key${i} *${key_type}
					${key_deserializer}
					if err != nil {
						return err
					}
					var value${i} *${value_type}
					${value_deserializer}
					if err != nil {
						return err
					}
					(*${output})[*key${i}] = *value${i}
				}
			]] % {
				i = i,
				output = output,
				map_type = print_type(field.type),
				key_type = print_type(field.type.from),
				value_type = print_type(field.type.to),
				key_deserializer = print_field_deserializer({ type = field.type.from, index = i + 1 }, "key" .. i),
				value_deserializer = print_field_deserializer({ type = field.type.to, index = i + 1 }, "value" .. i)
			}
		end
		return ""
	end

	local function print_field_deserializers()
		local out = ""
		for _, field in pairs(struct.fields) do
			out = out .. [[
				case ${field_id}:
					${field_deserializer}
				]] % {
				field_id = field.id,
				field_deserializer = print_field_deserializer(field, "s.${field_name}" % { field_name = pascal_case(field.name) })
			}
		end
		return out
	end

	return [[
		func (s *${sname}) deserialize(br *ByteReader) error {
			typeId, err := br.ReadTypeId()
			if err != nil || typeId != ${type_id} {
				return fmt.Errorf("unexpected type id: expected %d, got %d", ${type_id}, typeId)
			}
			length, err := br.ReadLength()
			if err != nil || length > br.Len() || length > math.MaxInt32 {
				return fmt.Errorf("invalid struct length: %d", length)
			}
			seenFields := make(map[uint16]bool)
			startPos := br.Offset()
			for br.Offset() < startPos + length {
				fieldId, err := br.ReadFieldId()
				if err != nil {
					return err
				}
				if seenFields[fieldId] {
					return fmt.Errorf("duplicate field id: %d", fieldId)
				}
				if fieldId > ${max_field_id} {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
					${field_deserializers}
				}
				if err != nil {
					return err
				}
			}
			return nil
		}
	]] % {
		sname = pascal_case(name),
		type_id = struct.id,
		max_field_id = get_max_field_id(),
		field_deserializers = print_field_deserializers()
	}
end

local function print_struct(name, struct)
	--  Print a struct definition with types
	local function print_struct_type(type)
		if type.kind == "primitive" then
			return go_type(type)
		elseif type.kind == "struct" then
			return pascal_case(type.name)
		elseif type.kind == "list" then
			return [[
				[]${elem_type}
			]] % {
				elem_type = print_struct_type(type.of)
			}
		elseif type.kind == "map" then
			return [[
				map[${key_type}]${value_type}
			]] % {
				key_type = print_struct_type(type.from),
				value_type = print_struct_type(type.to)
			}
		else
			error("unsupported field type kind: " .. tostring(type.kind))
		end
	end

	local function print_fields()
		local out = ""
		for field_id, field in pairs(struct.fields) do
			out = out .. [[
				${fname} *${ftype}
			]] % {
				fname = pascal_case(field.name),
				ftype = print_struct_type(field.type)
			}
		end
		return out
	end

	return [[
		type ${sname} struct {
		${fields}
		}

		func (${sname}) TypeId() uint16 {
			return ${type_id}
		}

		${serializer}
		${deserializer}
	]] % {
		sname = pascal_case(name),
		fields = print_fields(),
		type_id = struct.id,
		serializer = print_struct_serializer(name, struct),
		deserializer = print_struct_deserializer(name, struct)
	}
end

local function print_structs()
	local out = ""
	for name, struct in pairs(Schema.structs) do
		out = out .. print_struct(name, struct) .. "\n"
	end
	return out
end

local function print_type_cases()
	local out = ""
	for name, struct in pairs(Schema.structs) do
		out = out .. str_block {
			type_id = struct.id,
			sname = pascal_case(name)
		} [[
			case ${type_id}:
				s := &${sname}{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, ${type_id}, nil
		]]
	end
	return out
end

local file = str_block {
	name = Schema.name,
	structs = print_structs(),
	type_cases = print_type_cases()
} [[
package ${name}

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

${structs}

type Serializable interface {
	TypeId() uint16
	serialize(b *ByteWriter) error
	deserialize(b *ByteReader) error
}

func MarshalBytes(s Serializable) ([]byte, error) {
	b := &ByteWriter{}
	err := s.serialize(b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func UnmarshalBytes(data []byte, s Serializable) error {
	br := NewByteReader(&data)
	return s.deserialize(br)
}

func DeserializeBytes(data []byte) (value Serializable, typeID uint16, err error) {
	br := NewByteReader(&data)
	typeID, err = br.PeekTypeId()
	if err != nil {
		return nil, 0, err
	}
	switch typeID {
		${type_cases}
		default:
			return nil, 0, fmt.Errorf("unknown type id: %d", typeID)
	}
}

func GetTypeID(b []byte) uint16 {
	if len(b) < 2 {
		return 0
	}
	return binary.LittleEndian.Uint16(b[0:2])
}

type ByteWriter struct {
	b bytes.Buffer
}

func (bw *ByteWriter) Len() int {
	return bw.b.Len()
}

func (bw *ByteWriter) WriteFieldId(field int) error {
	if field < 0 {
		return fmt.Errorf("field number cannot be negative")
	}
	if field > math.MaxUint16 {
		return fmt.Errorf("field number exceeds maximum uint16 value")
	}
	return binary.Write(&bw.b, binary.LittleEndian, uint16(field))
}

func (bw *ByteWriter) WriteTypeId(field int) error {
	if field < 0 {
		return fmt.Errorf("field number cannot be negative")
	}
	if field > math.MaxUint16 {
		return fmt.Errorf("field number exceeds maximum uint16 value")
	}
	return binary.Write(&bw.b, binary.LittleEndian, uint16(field))
}

func (bw *ByteWriter) WriteLength(len int) error {
	if len < 0 {
		return fmt.Errorf("length cannot be negative")
	}
	if len > math.MaxUint32 {
		return fmt.Errorf("length exceeds maximum uint32 value")
	}
	return binary.Write(&bw.b, binary.LittleEndian, uint32(len))
}

func (bw *ByteWriter) Write(data any) error {
	var err error
	switch v := data.(type) {
	case *bool, bool, *int8, int8, *uint8, uint8, *int16, int16, *uint16, uint16, *int32, int32, *uint32, uint32, *int64, int64, *uint64, uint64, *float32, float32, *float64, float64:
		err = binary.Write(&bw.b, binary.LittleEndian, v)
	case []byte:
		err = bw.WriteLength(len(v))
		_, err = bw.b.Write(v)
	case string, *string:
		if strPtr, ok := v.(*string); ok {
			strBytes := []byte(*strPtr)
			err = bw.WriteLength(len(strBytes))
			_, err = bw.b.Write(strBytes)
		} else {
			strBytes := []byte(v.(string))
			err = bw.WriteLength(len(strBytes))
			_, err = bw.b.Write(strBytes)
		}
	default:
		return fmt.Errorf("writing unsupported data type: %T", v)
	}
	return err
}

func (bw *ByteWriter) WriteLengthAt(data int, offset int) error {
	if offset < 0 || offset+4 > bw.b.Len() {
		return fmt.Errorf("offset out of bounds")
	}
	if data > math.MaxUint32 {
		return fmt.Errorf("data exceeds maximum uint32 value")
	}
	if bw.b.Len() < offset+4 {
		return fmt.Errorf("buffer too small to write at offset")
	}
	binary.LittleEndian.PutUint32(bw.b.Bytes()[offset:], uint32(data))
	return nil
}

func (bw *ByteWriter) Bytes() []byte {
	return bw.b.Bytes()
}

type ByteReader struct {
	b *[]byte
	offset int
}

func (br *ByteReader) Len() int {
	return len(*br.b)
}

func (br *ByteReader) Offset() int {
	return br.offset
}

func (br *ByteReader) Remaining() int {
	return br.Len() - br.Offset()
}

func (br *ByteReader) Read(out any) error {
	switch v := out.(type) {
	case *bool:
		if br.Remaining() < 1 {
			return fmt.Errorf("not enough data to read bool")
		}
		*v = (*br.b)[br.offset] != 0
		br.offset += 1
	case *int8:
		if br.Remaining() < 1 {
			return fmt.Errorf("not enough data to read int8")
		}
		*v = int8((*br.b)[br.offset])
		br.offset += 1
	case *uint8:
		if br.Remaining() < 1 {
			return fmt.Errorf("not enough data to read uint8")
		}
		*v = (*br.b)[br.offset]
		br.offset += 1
	case *int16:
		if br.Remaining() < 2 {
			return fmt.Errorf("not enough data to read int16")
		}
		*v = int16(binary.LittleEndian.Uint16((*br.b)[br.offset:br.offset+2]))
		br.offset += 2
	case *uint16:
		if br.Remaining() < 2 {
			return fmt.Errorf("not enough data to read uint16")
		}
		*v = binary.LittleEndian.Uint16((*br.b)[br.offset:br.offset+2])
		br.offset += 2
	case *int32:
		if br.Remaining() < 4 {
			return fmt.Errorf("not enough data to read int32")
		}
		*v = int32(binary.LittleEndian.Uint32((*br.b)[br.offset:br.offset+4]))
		br.offset += 4
	case *uint32:
		if br.Remaining() < 4 {
			return fmt.Errorf("not enough data to read uint32")
		}
		*v = binary.LittleEndian.Uint32((*br.b)[br.offset:br.offset+4])
		br.offset += 4
	case *int64:
		if br.Remaining() < 8 {
			return fmt.Errorf("not enough data to read int64")
		}
		*v = int64(binary.LittleEndian.Uint64((*br.b)[br.offset:br.offset+8]))
		br.offset += 8
	case *uint64:
		if br.Remaining() < 8 {
			return fmt.Errorf("not enough data to read uint64")
		}
		*v = binary.LittleEndian.Uint64((*br.b)[br.offset:br.offset+8])
		br.offset += 8
	case *float32:
		if br.Remaining() < 4 {
			return fmt.Errorf("not enough data to read float32")
		}
		*v = math.Float32frombits(binary.LittleEndian.Uint32((*br.b)[br.offset:br.offset+4]))
		br.offset += 4
	case *float64:
		if br.Remaining() < 8 {
			return fmt.Errorf("not enough data to read float64")
		}
		*v = math.Float64frombits(binary.LittleEndian.Uint64((*br.b)[br.offset:br.offset+8]))
		br.offset += 8
	case *string:
		var length uint32
		if err := br.Read(&length); err != nil {
			return err
		}
		if br.Remaining() < int(length) {
			return fmt.Errorf("not enough data to read string of length %d", length)
		}
		*v = string((*br.b)[br.offset : br.offset+int(length)])
		br.offset += int(length)
	default:
		return fmt.Errorf("reading unsupported data type: %T", v)
	}
	return nil
}

func (br *ByteReader) ReadFieldId() (uint16, error) {
	var fieldId uint16
	if err := br.Read(&fieldId); err != nil {
		return 0, err
	}
	return fieldId, nil
}

func (br *ByteReader) ReadTypeId() (uint16, error) {
	var typeId uint16
	if err := br.Read(&typeId); err != nil {
		return 0, err
	}
	return typeId, nil
}

func (br *ByteReader) PeekTypeId() (uint16, error) {
	var typeId uint16
	if br.Remaining() < 2 {
		return 0, fmt.Errorf("not enough data to peek type id")
	}
	typeId = binary.LittleEndian.Uint16((*br.b)[br.offset : br.offset+2])
	return typeId, nil
}

func (br *ByteReader) ReadLength() (int, error) {
	var length uint32
	if err := br.Read(&length); err != nil {
		return 0, err
	}
	return int(length), nil
}

func NewByteReader(data *[]byte) *ByteReader {
	return &ByteReader{data, 0}
}
]]

Output[Schema.name .. ".go"] = file
