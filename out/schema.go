package schema

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

		type Bar struct {
						Baz *int32
			
		}

		func (Bar) TypeId() uint16 {
			return 32026
		}

				func (s *Bar) serialize(b *ByteWriter) error {
			err := b.WriteTypeId(32026)
			if err != nil {
				return err
			}
			err = b.WriteLength(0)
			if err != nil {
				return err
			}
			startLen := b.Len()
						  if s.Baz != nil {
				  err = b.WriteFieldId(1)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Baz)
				if err != nil {
					return err
				}
			
			  }
		  
			structLen := b.Len() - startLen
			err = b.WriteLengthAt(structLen, startLen-4)
			return err
		}
	
				func (s *Bar) deserialize(br *ByteReader) error {
			typeId, err := br.ReadTypeId()
			if err != nil || typeId != 32026 {
				return fmt.Errorf("unexpected type id: expected %d, got %d", 32026, typeId)
			}
			length, err := br.ReadLength()
			if err != nil || length > br.Len() || length > math.MaxInt32 {
				return fmt.Errorf("invalid struct length: %d", length)
			}
			seenFields := make(map[uint16]bool)
			startLen := br.Len()
			for br.Len() > startLen-int(length) {
				fieldId, err := br.ReadFieldId()
				if err != nil || seenFields[fieldId] {
					return fmt.Errorf("error reading field id or duplicate field id: %d", fieldId)
				}
				if fieldId > 1 {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
									case 1:
					 err = br.Read(s.Baz) 
				
				}
				if err != nil {
					return err
				}
			}
			return nil
		}
	
	
		type Foo struct {
						Lst2 *				[]				[]string
			
			
							M *				map[int32]string
			
							X *				map[string]				map[int32]bool
			
			
							F *float32
							B *bool
							Bar *Bar
							Lst *				[]int32
			
			
		}

		func (Foo) TypeId() uint16 {
			return 32471
		}

				func (s *Foo) serialize(b *ByteWriter) error {
			err := b.WriteTypeId(32471)
			if err != nil {
				return err
			}
			err = b.WriteLength(0)
			if err != nil {
				return err
			}
			startLen := b.Len()
						  if s.Lst2 != nil {
				  err = b.WriteFieldId(5)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for _, item0 := range *s.Lst2 {
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen1 := b.Len()
				for _, item1 := range item0 {
									err = b.Write(item1)
				if err != nil {
					return err
				}
			
				}
				listLen1 := b.Len() - startLen1
				err = b.WriteLengthAt(listLen1, startLen1-4)
				if err != nil {
					return err
				}
			
				}
				listLen0 := b.Len() - startLen0
				err = b.WriteLengthAt(listLen0, startLen0-4)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.M != nil {
				  err = b.WriteFieldId(6)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.M {
									err = b.Write(key0)
				if err != nil {
					return err
				}
			
									err = b.Write(value0)
				if err != nil {
					return err
				}
			
				}
				mapLen0 := b.Len() - startLen0
				err = b.WriteLengthAt(mapLen0, startLen0-4)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.X != nil {
				  err = b.WriteFieldId(7)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.X {
									err = b.Write(key0)
				if err != nil {
					return err
				}
			
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen1 := b.Len()
				for key1, value1 := range value0 {
									err = b.Write(key1)
				if err != nil {
					return err
				}
			
									err = b.Write(value1)
				if err != nil {
					return err
				}
			
				}
				mapLen1 := b.Len() - startLen1
				err = b.WriteLengthAt(mapLen1, startLen1-4)
				if err != nil {
					return err
				}
			
				}
				mapLen0 := b.Len() - startLen0
				err = b.WriteLengthAt(mapLen0, startLen0-4)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.F != nil {
				  err = b.WriteFieldId(1)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.F)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.B != nil {
				  err = b.WriteFieldId(2)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.B)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Bar != nil {
				  err = b.WriteFieldId(3)
				  if err != nil {
					  return err
				  }
									err = s.Bar.serialize(b)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Lst != nil {
				  err = b.WriteFieldId(4)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for _, item0 := range *s.Lst {
									err = b.Write(item0)
				if err != nil {
					return err
				}
			
				}
				listLen0 := b.Len() - startLen0
				err = b.WriteLengthAt(listLen0, startLen0-4)
				if err != nil {
					return err
				}
			
			  }
		  
			structLen := b.Len() - startLen
			err = b.WriteLengthAt(structLen, startLen-4)
			return err
		}
	
				func (s *Foo) deserialize(br *ByteReader) error {
			typeId, err := br.ReadTypeId()
			if err != nil || typeId != 32471 {
				return fmt.Errorf("unexpected type id: expected %d, got %d", 32471, typeId)
			}
			length, err := br.ReadLength()
			if err != nil || length > br.Len() || length > math.MaxInt32 {
				return fmt.Errorf("invalid struct length: %d", length)
			}
			seenFields := make(map[uint16]bool)
			startLen := br.Len()
			for br.Len() > startLen-int(length) {
				fieldId, err := br.ReadFieldId()
				if err != nil || seenFields[fieldId] {
					return fmt.Errorf("error reading field id or duplicate field id: %d", fieldId)
				}
				if fieldId > 7 {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
									case 5:
									listLen0, err := br.ReadLength()
				if err != nil || listLen0 < 0 || listLen0 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen0)
				}
				startLen0 := br.Len()
				s.Lst2 = &[][]string{}
				for br.Len() > startLen0-listLen0 {
					var item0 *[]string
									listLen1, err := br.ReadLength()
				if err != nil || listLen1 < 0 || listLen1 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen1)
				}
				startLen1 := br.Len()
				item0 = &[]string{}
				for br.Len() > startLen1-listLen1 {
					var item1 *string
					 err = br.Read(item1) 
					if err != nil {
						return err
					}
					*item0 = append(*item0, *item1)
				}
			
					if err != nil {
						return err
					}
					*s.Lst2 = append(*s.Lst2, *item0)
				}
			
								case 6:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startLen0 := br.Len()
				s.M = &map[int32]string{}
				for br.Len() > startLen0-mapLen0 {
					var key0 *int32
					 err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *string
					 err = br.Read(value0) 
					if err != nil {
						return err
					}
					(*s.M)[*key0] = *value0
				}
			
								case 7:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startLen0 := br.Len()
				s.X = &map[string]map[int32]bool{}
				for br.Len() > startLen0-mapLen0 {
					var key0 *string
					 err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *map[int32]bool
									mapLen1, err := br.ReadLength()
				if err != nil || mapLen1 < 0 || mapLen1 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen1)
				}
				startLen1 := br.Len()
				value0 = &map[int32]bool{}
				for br.Len() > startLen1-mapLen1 {
					var key1 *int32
					 err = br.Read(key1) 
					if err != nil {
						return err
					}
					var value1 *bool
					 err = br.Read(value1) 
					if err != nil {
						return err
					}
					(*value0)[*key1] = *value1
				}
			
					if err != nil {
						return err
					}
					(*s.X)[*key0] = *value0
				}
			
								case 1:
					 err = br.Read(s.F) 
								case 2:
					 err = br.Read(s.B) 
								case 3:
									s.Bar = &Bar{}
				err = s.Bar.deserialize(br)
			
								case 4:
									listLen0, err := br.ReadLength()
				if err != nil || listLen0 < 0 || listLen0 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen0)
				}
				startLen0 := br.Len()
				s.Lst = &[]int32{}
				for br.Len() > startLen0-listLen0 {
					var item0 *int32
					 err = br.Read(item0) 
					if err != nil {
						return err
					}
					*s.Lst = append(*s.Lst, *item0)
				}
			
				
				}
				if err != nil {
					return err
				}
			}
			return nil
		}
	
	


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
					case 32026:
				s := &Bar{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, 32026, nil
					case 32471:
				s := &Foo{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, 32471, nil
		
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
	case *bool, *int16, *uint16, *int32, *uint32, *float32, *float64:
		err = binary.Write(&bw.b, binary.LittleEndian, v)
	case []byte:
		err = bw.WriteLength(len(v))
		_, err = bw.b.Write(v)
	case string:
		strBytes := []byte(v)
		err = bw.WriteLength(len(strBytes))
		_, err = bw.b.Write(strBytes)
	default:
		return fmt.Errorf("unsupported data type: %T", v)
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
		return fmt.Errorf("unsupported data type: %T", v)
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