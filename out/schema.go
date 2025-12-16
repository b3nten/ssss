package schema

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

		type Character struct {
						Id *uint64
							Position *Vector3
							Friends *				[]Character
			
							SkillProgress *				map[string]				[]float64
			
			
							ArbitraryData *				map[string]				map[string]				map[string]int8
			
			
			
							Name *string
							Inventory *				[]				[]Item
			
			
							Companions *				map[uint16]Companion
			
							Stats *Stats
							Equipment *				map[string]EquipmentSlot
			
			
		}

		func (Character) TypeId() uint16 {
			return 16560
		}

				func (s *Character) serialize(b *ByteWriter) error {
			err := b.WriteTypeId(16560)
			if err != nil {
				return err
			}
			err = b.WriteLength(0)
			if err != nil {
				return err
			}
			startLen := b.Len()
						  if s.Id != nil {
				  err = b.WriteFieldId(0)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Id)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Position != nil {
				  err = b.WriteFieldId(2)
				  if err != nil {
					  return err
				  }
									err = s.Position.serialize(b)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Friends != nil {
				  err = b.WriteFieldId(7)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for _, item0 := range *s.Friends {
									err = item0.serialize(b)
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
		  			  if s.SkillProgress != nil {
				  err = b.WriteFieldId(8)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.SkillProgress {
									err = b.Write(key0)
				if err != nil {
					return err
				}
			
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen1 := b.Len()
				for _, item1 := range value0 {
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
				mapLen0 := b.Len() - startLen0
				err = b.WriteLengthAt(mapLen0, startLen0-4)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.ArbitraryData != nil {
				  err = b.WriteFieldId(9)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.ArbitraryData {
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
			
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen2 := b.Len()
				for key2, value2 := range value1 {
									err = b.Write(key2)
				if err != nil {
					return err
				}
			
									err = b.Write(value2)
				if err != nil {
					return err
				}
			
				}
				mapLen2 := b.Len() - startLen2
				err = b.WriteLengthAt(mapLen2, startLen2-4)
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
		  			  if s.Name != nil {
				  err = b.WriteFieldId(1)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Name)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Inventory != nil {
				  err = b.WriteFieldId(4)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for _, item0 := range *s.Inventory {
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen1 := b.Len()
				for _, item1 := range item0 {
									err = item1.serialize(b)
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
		  			  if s.Companions != nil {
				  err = b.WriteFieldId(6)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.Companions {
									err = b.Write(key0)
				if err != nil {
					return err
				}
			
									err = value0.serialize(b)
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
		  			  if s.Stats != nil {
				  err = b.WriteFieldId(3)
				  if err != nil {
					  return err
				  }
									err = s.Stats.serialize(b)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Equipment != nil {
				  err = b.WriteFieldId(5)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.Equipment {
									err = b.Write(key0)
				if err != nil {
					return err
				}
			
									err = value0.serialize(b)
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
		  
			structLen := b.Len() - startLen
			err = b.WriteLengthAt(structLen, startLen-4)
			return err
		}
	
				func (s *Character) deserialize(br *ByteReader) error {
			typeId, err := br.ReadTypeId()
			if err != nil || typeId != 16560 {
				return fmt.Errorf("unexpected type id: expected %d, got %d", 16560, typeId)
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
				if fieldId > 9 {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
									case 0:
					 s.Id = new(uint64); err = br.Read(s.Id) 
								case 2:
									s.Position = &Vector3{}
				err = s.Position.deserialize(br)
			
								case 7:
									listLen0, err := br.ReadLength()
				if err != nil || listLen0 < 0 || listLen0 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen0)
				}
				startPos0 := br.Offset()
				s.Friends = &[]Character{}
				for br.Offset() < startPos0+listLen0 {
					var item0 *Character
									item0 = &Character{}
				err = item0.deserialize(br)
			
					if err != nil {
						return err
					}
					*s.Friends = append(*s.Friends, *item0)
				}
			
								case 8:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startPos0 := br.Offset()
				s.SkillProgress = &map[string][]float64{}
				for br.Offset() < startPos0+mapLen0 {
					var key0 *string
					 key0 = new(string); err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *[]float64
									listLen1, err := br.ReadLength()
				if err != nil || listLen1 < 0 || listLen1 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen1)
				}
				startPos1 := br.Offset()
				value0 = &[]float64{}
				for br.Offset() < startPos1+listLen1 {
					var item1 *float64
					 item1 = new(float64); err = br.Read(item1) 
					if err != nil {
						return err
					}
					*value0 = append(*value0, *item1)
				}
			
					if err != nil {
						return err
					}
					(*s.SkillProgress)[*key0] = *value0
				}
			
								case 9:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startPos0 := br.Offset()
				s.ArbitraryData = &map[string]map[string]map[string]int8{}
				for br.Offset() < startPos0+mapLen0 {
					var key0 *string
					 key0 = new(string); err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *map[string]map[string]int8
									mapLen1, err := br.ReadLength()
				if err != nil || mapLen1 < 0 || mapLen1 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen1)
				}
				startPos1 := br.Offset()
				value0 = &map[string]map[string]int8{}
				for br.Offset() < startPos1+mapLen1 {
					var key1 *string
					 key1 = new(string); err = br.Read(key1) 
					if err != nil {
						return err
					}
					var value1 *map[string]int8
									mapLen2, err := br.ReadLength()
				if err != nil || mapLen2 < 0 || mapLen2 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen2)
				}
				startPos2 := br.Offset()
				value1 = &map[string]int8{}
				for br.Offset() < startPos2+mapLen2 {
					var key2 *string
					 key2 = new(string); err = br.Read(key2) 
					if err != nil {
						return err
					}
					var value2 *int8
					 value2 = new(int8); err = br.Read(value2) 
					if err != nil {
						return err
					}
					(*value1)[*key2] = *value2
				}
			
					if err != nil {
						return err
					}
					(*value0)[*key1] = *value1
				}
			
					if err != nil {
						return err
					}
					(*s.ArbitraryData)[*key0] = *value0
				}
			
								case 1:
					 s.Name = new(string); err = br.Read(s.Name) 
								case 4:
									listLen0, err := br.ReadLength()
				if err != nil || listLen0 < 0 || listLen0 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen0)
				}
				startPos0 := br.Offset()
				s.Inventory = &[][]Item{}
				for br.Offset() < startPos0+listLen0 {
					var item0 *[]Item
									listLen1, err := br.ReadLength()
				if err != nil || listLen1 < 0 || listLen1 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen1)
				}
				startPos1 := br.Offset()
				item0 = &[]Item{}
				for br.Offset() < startPos1+listLen1 {
					var item1 *Item
									item1 = &Item{}
				err = item1.deserialize(br)
			
					if err != nil {
						return err
					}
					*item0 = append(*item0, *item1)
				}
			
					if err != nil {
						return err
					}
					*s.Inventory = append(*s.Inventory, *item0)
				}
			
								case 6:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startPos0 := br.Offset()
				s.Companions = &map[uint16]Companion{}
				for br.Offset() < startPos0+mapLen0 {
					var key0 *uint16
					 key0 = new(uint16); err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *Companion
									value0 = &Companion{}
				err = value0.deserialize(br)
			
					if err != nil {
						return err
					}
					(*s.Companions)[*key0] = *value0
				}
			
								case 3:
									s.Stats = &Stats{}
				err = s.Stats.deserialize(br)
			
								case 5:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startPos0 := br.Offset()
				s.Equipment = &map[string]EquipmentSlot{}
				for br.Offset() < startPos0+mapLen0 {
					var key0 *string
					 key0 = new(string); err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *EquipmentSlot
									value0 = &EquipmentSlot{}
				err = value0.deserialize(br)
			
					if err != nil {
						return err
					}
					(*s.Equipment)[*key0] = *value0
				}
			
				
				}
				if err != nil {
					return err
				}
			}
			return nil
		}
	
	
		type Quest struct {
						Prerequisites *				[]Quest
			
							AreaLayers *				[]				[]				[]uint16
			
			
			
							Description *string
							Id *uint32
							NextSteps *				map[string]Quest
			
							RequiredPos *				[]Vector3
			
							Title *string
							Difficulty *uint8
							Rewards *				[]Item
			
							Objectives *				map[string]				[]string
			
			
			
		}

		func (Quest) TypeId() uint16 {
			return 16605
		}

				func (s *Quest) serialize(b *ByteWriter) error {
			err := b.WriteTypeId(16605)
			if err != nil {
				return err
			}
			err = b.WriteLength(0)
			if err != nil {
				return err
			}
			startLen := b.Len()
						  if s.Prerequisites != nil {
				  err = b.WriteFieldId(6)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for _, item0 := range *s.Prerequisites {
									err = item0.serialize(b)
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
		  			  if s.AreaLayers != nil {
				  err = b.WriteFieldId(8)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for _, item0 := range *s.AreaLayers {
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen1 := b.Len()
				for _, item1 := range item0 {
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen2 := b.Len()
				for _, item2 := range item1 {
									err = b.Write(item2)
				if err != nil {
					return err
				}
			
				}
				listLen2 := b.Len() - startLen2
				err = b.WriteLengthAt(listLen2, startLen2-4)
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
		  			  if s.Description != nil {
				  err = b.WriteFieldId(9)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Description)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Id != nil {
				  err = b.WriteFieldId(0)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Id)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.NextSteps != nil {
				  err = b.WriteFieldId(7)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.NextSteps {
									err = b.Write(key0)
				if err != nil {
					return err
				}
			
									err = value0.serialize(b)
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
		  			  if s.RequiredPos != nil {
				  err = b.WriteFieldId(4)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for _, item0 := range *s.RequiredPos {
									err = item0.serialize(b)
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
		  			  if s.Title != nil {
				  err = b.WriteFieldId(1)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Title)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Difficulty != nil {
				  err = b.WriteFieldId(2)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Difficulty)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Rewards != nil {
				  err = b.WriteFieldId(3)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for _, item0 := range *s.Rewards {
									err = item0.serialize(b)
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
		  			  if s.Objectives != nil {
				  err = b.WriteFieldId(5)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.Objectives {
									err = b.Write(key0)
				if err != nil {
					return err
				}
			
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen1 := b.Len()
				for _, item1 := range value0 {
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
				mapLen0 := b.Len() - startLen0
				err = b.WriteLengthAt(mapLen0, startLen0-4)
				if err != nil {
					return err
				}
			
			  }
		  
			structLen := b.Len() - startLen
			err = b.WriteLengthAt(structLen, startLen-4)
			return err
		}
	
				func (s *Quest) deserialize(br *ByteReader) error {
			typeId, err := br.ReadTypeId()
			if err != nil || typeId != 16605 {
				return fmt.Errorf("unexpected type id: expected %d, got %d", 16605, typeId)
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
				if fieldId > 9 {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
									case 6:
									listLen0, err := br.ReadLength()
				if err != nil || listLen0 < 0 || listLen0 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen0)
				}
				startPos0 := br.Offset()
				s.Prerequisites = &[]Quest{}
				for br.Offset() < startPos0+listLen0 {
					var item0 *Quest
									item0 = &Quest{}
				err = item0.deserialize(br)
			
					if err != nil {
						return err
					}
					*s.Prerequisites = append(*s.Prerequisites, *item0)
				}
			
								case 8:
									listLen0, err := br.ReadLength()
				if err != nil || listLen0 < 0 || listLen0 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen0)
				}
				startPos0 := br.Offset()
				s.AreaLayers = &[][][]uint16{}
				for br.Offset() < startPos0+listLen0 {
					var item0 *[][]uint16
									listLen1, err := br.ReadLength()
				if err != nil || listLen1 < 0 || listLen1 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen1)
				}
				startPos1 := br.Offset()
				item0 = &[][]uint16{}
				for br.Offset() < startPos1+listLen1 {
					var item1 *[]uint16
									listLen2, err := br.ReadLength()
				if err != nil || listLen2 < 0 || listLen2 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen2)
				}
				startPos2 := br.Offset()
				item1 = &[]uint16{}
				for br.Offset() < startPos2+listLen2 {
					var item2 *uint16
					 item2 = new(uint16); err = br.Read(item2) 
					if err != nil {
						return err
					}
					*item1 = append(*item1, *item2)
				}
			
					if err != nil {
						return err
					}
					*item0 = append(*item0, *item1)
				}
			
					if err != nil {
						return err
					}
					*s.AreaLayers = append(*s.AreaLayers, *item0)
				}
			
								case 9:
					 s.Description = new(string); err = br.Read(s.Description) 
								case 0:
					 s.Id = new(uint32); err = br.Read(s.Id) 
								case 7:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startPos0 := br.Offset()
				s.NextSteps = &map[string]Quest{}
				for br.Offset() < startPos0+mapLen0 {
					var key0 *string
					 key0 = new(string); err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *Quest
									value0 = &Quest{}
				err = value0.deserialize(br)
			
					if err != nil {
						return err
					}
					(*s.NextSteps)[*key0] = *value0
				}
			
								case 4:
									listLen0, err := br.ReadLength()
				if err != nil || listLen0 < 0 || listLen0 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen0)
				}
				startPos0 := br.Offset()
				s.RequiredPos = &[]Vector3{}
				for br.Offset() < startPos0+listLen0 {
					var item0 *Vector3
									item0 = &Vector3{}
				err = item0.deserialize(br)
			
					if err != nil {
						return err
					}
					*s.RequiredPos = append(*s.RequiredPos, *item0)
				}
			
								case 1:
					 s.Title = new(string); err = br.Read(s.Title) 
								case 2:
					 s.Difficulty = new(uint8); err = br.Read(s.Difficulty) 
								case 3:
									listLen0, err := br.ReadLength()
				if err != nil || listLen0 < 0 || listLen0 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen0)
				}
				startPos0 := br.Offset()
				s.Rewards = &[]Item{}
				for br.Offset() < startPos0+listLen0 {
					var item0 *Item
									item0 = &Item{}
				err = item0.deserialize(br)
			
					if err != nil {
						return err
					}
					*s.Rewards = append(*s.Rewards, *item0)
				}
			
								case 5:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startPos0 := br.Offset()
				s.Objectives = &map[string][]string{}
				for br.Offset() < startPos0+mapLen0 {
					var key0 *string
					 key0 = new(string); err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *[]string
									listLen1, err := br.ReadLength()
				if err != nil || listLen1 < 0 || listLen1 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen1)
				}
				startPos1 := br.Offset()
				value0 = &[]string{}
				for br.Offset() < startPos1+listLen1 {
					var item1 *string
					 item1 = new(string); err = br.Read(item1) 
					if err != nil {
						return err
					}
					*value0 = append(*value0, *item1)
				}
			
					if err != nil {
						return err
					}
					(*s.Objectives)[*key0] = *value0
				}
			
				
				}
				if err != nil {
					return err
				}
			}
			return nil
		}
	
	
		type EquipmentSlot struct {
						SlotName *string
							Item *Item
			
		}

		func (EquipmentSlot) TypeId() uint16 {
			return 35339
		}

				func (s *EquipmentSlot) serialize(b *ByteWriter) error {
			err := b.WriteTypeId(35339)
			if err != nil {
				return err
			}
			err = b.WriteLength(0)
			if err != nil {
				return err
			}
			startLen := b.Len()
						  if s.SlotName != nil {
				  err = b.WriteFieldId(0)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.SlotName)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Item != nil {
				  err = b.WriteFieldId(1)
				  if err != nil {
					  return err
				  }
									err = s.Item.serialize(b)
				if err != nil {
					return err
				}
			
			  }
		  
			structLen := b.Len() - startLen
			err = b.WriteLengthAt(structLen, startLen-4)
			return err
		}
	
				func (s *EquipmentSlot) deserialize(br *ByteReader) error {
			typeId, err := br.ReadTypeId()
			if err != nil || typeId != 35339 {
				return fmt.Errorf("unexpected type id: expected %d, got %d", 35339, typeId)
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
				if fieldId > 1 {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
									case 0:
					 s.SlotName = new(string); err = br.Read(s.SlotName) 
								case 1:
									s.Item = &Item{}
				err = s.Item.deserialize(br)
			
				
				}
				if err != nil {
					return err
				}
			}
			return nil
		}
	
	
		type LootEntry struct {
						MinQty *uint8
							MaxQty *uint8
							Conditions *				map[string]bool
			
							ItemId *uint32
			
		}

		func (LootEntry) TypeId() uint16 {
			return 305
		}

				func (s *LootEntry) serialize(b *ByteWriter) error {
			err := b.WriteTypeId(305)
			if err != nil {
				return err
			}
			err = b.WriteLength(0)
			if err != nil {
				return err
			}
			startLen := b.Len()
						  if s.MinQty != nil {
				  err = b.WriteFieldId(1)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.MinQty)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.MaxQty != nil {
				  err = b.WriteFieldId(2)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.MaxQty)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Conditions != nil {
				  err = b.WriteFieldId(3)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.Conditions {
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
		  			  if s.ItemId != nil {
				  err = b.WriteFieldId(0)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.ItemId)
				if err != nil {
					return err
				}
			
			  }
		  
			structLen := b.Len() - startLen
			err = b.WriteLengthAt(structLen, startLen-4)
			return err
		}
	
				func (s *LootEntry) deserialize(br *ByteReader) error {
			typeId, err := br.ReadTypeId()
			if err != nil || typeId != 305 {
				return fmt.Errorf("unexpected type id: expected %d, got %d", 305, typeId)
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
				if fieldId > 3 {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
									case 1:
					 s.MinQty = new(uint8); err = br.Read(s.MinQty) 
								case 2:
					 s.MaxQty = new(uint8); err = br.Read(s.MaxQty) 
								case 3:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startPos0 := br.Offset()
				s.Conditions = &map[string]bool{}
				for br.Offset() < startPos0+mapLen0 {
					var key0 *string
					 key0 = new(string); err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *bool
					 value0 = new(bool); err = br.Read(value0) 
					if err != nil {
						return err
					}
					(*s.Conditions)[*key0] = *value0
				}
			
								case 0:
					 s.ItemId = new(uint32); err = br.Read(s.ItemId) 
				
				}
				if err != nil {
					return err
				}
			}
			return nil
		}
	
	
		type World struct {
						ActiveQuests *				[]Quest
			
							ZoneData *				map[string]				map[string]uint32
			
			
							SystemFlags *				map[string]bool
			
							LootTables *				map[string]Loot
			
							WorldName *string
							Seed *uint64
							Gravity *float64
							Players *				[]Character
			
			
		}

		func (World) TypeId() uint16 {
			return 60723
		}

				func (s *World) serialize(b *ByteWriter) error {
			err := b.WriteTypeId(60723)
			if err != nil {
				return err
			}
			err = b.WriteLength(0)
			if err != nil {
				return err
			}
			startLen := b.Len()
						  if s.ActiveQuests != nil {
				  err = b.WriteFieldId(4)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for _, item0 := range *s.ActiveQuests {
									err = item0.serialize(b)
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
		  			  if s.ZoneData != nil {
				  err = b.WriteFieldId(5)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.ZoneData {
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
		  			  if s.SystemFlags != nil {
				  err = b.WriteFieldId(6)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.SystemFlags {
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
		  			  if s.LootTables != nil {
				  err = b.WriteFieldId(7)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.LootTables {
									err = b.Write(key0)
				if err != nil {
					return err
				}
			
									err = value0.serialize(b)
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
		  			  if s.WorldName != nil {
				  err = b.WriteFieldId(0)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.WorldName)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Seed != nil {
				  err = b.WriteFieldId(1)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Seed)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Gravity != nil {
				  err = b.WriteFieldId(2)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Gravity)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Players != nil {
				  err = b.WriteFieldId(3)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for _, item0 := range *s.Players {
									err = item0.serialize(b)
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
	
				func (s *World) deserialize(br *ByteReader) error {
			typeId, err := br.ReadTypeId()
			if err != nil || typeId != 60723 {
				return fmt.Errorf("unexpected type id: expected %d, got %d", 60723, typeId)
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
				if fieldId > 7 {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
									case 4:
									listLen0, err := br.ReadLength()
				if err != nil || listLen0 < 0 || listLen0 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen0)
				}
				startPos0 := br.Offset()
				s.ActiveQuests = &[]Quest{}
				for br.Offset() < startPos0+listLen0 {
					var item0 *Quest
									item0 = &Quest{}
				err = item0.deserialize(br)
			
					if err != nil {
						return err
					}
					*s.ActiveQuests = append(*s.ActiveQuests, *item0)
				}
			
								case 5:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startPos0 := br.Offset()
				s.ZoneData = &map[string]map[string]uint32{}
				for br.Offset() < startPos0+mapLen0 {
					var key0 *string
					 key0 = new(string); err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *map[string]uint32
									mapLen1, err := br.ReadLength()
				if err != nil || mapLen1 < 0 || mapLen1 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen1)
				}
				startPos1 := br.Offset()
				value0 = &map[string]uint32{}
				for br.Offset() < startPos1+mapLen1 {
					var key1 *string
					 key1 = new(string); err = br.Read(key1) 
					if err != nil {
						return err
					}
					var value1 *uint32
					 value1 = new(uint32); err = br.Read(value1) 
					if err != nil {
						return err
					}
					(*value0)[*key1] = *value1
				}
			
					if err != nil {
						return err
					}
					(*s.ZoneData)[*key0] = *value0
				}
			
								case 6:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startPos0 := br.Offset()
				s.SystemFlags = &map[string]bool{}
				for br.Offset() < startPos0+mapLen0 {
					var key0 *string
					 key0 = new(string); err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *bool
					 value0 = new(bool); err = br.Read(value0) 
					if err != nil {
						return err
					}
					(*s.SystemFlags)[*key0] = *value0
				}
			
								case 7:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startPos0 := br.Offset()
				s.LootTables = &map[string]Loot{}
				for br.Offset() < startPos0+mapLen0 {
					var key0 *string
					 key0 = new(string); err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *Loot
									value0 = &Loot{}
				err = value0.deserialize(br)
			
					if err != nil {
						return err
					}
					(*s.LootTables)[*key0] = *value0
				}
			
								case 0:
					 s.WorldName = new(string); err = br.Read(s.WorldName) 
								case 1:
					 s.Seed = new(uint64); err = br.Read(s.Seed) 
								case 2:
					 s.Gravity = new(float64); err = br.Read(s.Gravity) 
								case 3:
									listLen0, err := br.ReadLength()
				if err != nil || listLen0 < 0 || listLen0 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen0)
				}
				startPos0 := br.Offset()
				s.Players = &[]Character{}
				for br.Offset() < startPos0+listLen0 {
					var item0 *Character
									item0 = &Character{}
				err = item0.deserialize(br)
			
					if err != nil {
						return err
					}
					*s.Players = append(*s.Players, *item0)
				}
			
				
				}
				if err != nil {
					return err
				}
			}
			return nil
		}
	
	
		type Vector3 struct {
						X *float64
							Y *float64
							Z *float64
			
		}

		func (Vector3) TypeId() uint16 {
			return 2
		}

				func (s *Vector3) serialize(b *ByteWriter) error {
			err := b.WriteTypeId(2)
			if err != nil {
				return err
			}
			err = b.WriteLength(0)
			if err != nil {
				return err
			}
			startLen := b.Len()
						  if s.X != nil {
				  err = b.WriteFieldId(0)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.X)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Y != nil {
				  err = b.WriteFieldId(1)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Y)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Z != nil {
				  err = b.WriteFieldId(2)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Z)
				if err != nil {
					return err
				}
			
			  }
		  
			structLen := b.Len() - startLen
			err = b.WriteLengthAt(structLen, startLen-4)
			return err
		}
	
				func (s *Vector3) deserialize(br *ByteReader) error {
			typeId, err := br.ReadTypeId()
			if err != nil || typeId != 2 {
				return fmt.Errorf("unexpected type id: expected %d, got %d", 2, typeId)
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
				if fieldId > 2 {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
									case 0:
					 s.X = new(float64); err = br.Read(s.X) 
								case 1:
					 s.Y = new(float64); err = br.Read(s.Y) 
								case 2:
					 s.Z = new(float64); err = br.Read(s.Z) 
				
				}
				if err != nil {
					return err
				}
			}
			return nil
		}
	
	
		type Stats struct {
						CritDamage *float64
							Resistances *				map[string]int16
			
							Health *int32
							Mana *int32
							Stamina *int32
							CritChance *float64
			
		}

		func (Stats) TypeId() uint16 {
			return 3
		}

				func (s *Stats) serialize(b *ByteWriter) error {
			err := b.WriteTypeId(3)
			if err != nil {
				return err
			}
			err = b.WriteLength(0)
			if err != nil {
				return err
			}
			startLen := b.Len()
						  if s.CritDamage != nil {
				  err = b.WriteFieldId(4)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.CritDamage)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Resistances != nil {
				  err = b.WriteFieldId(5)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.Resistances {
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
		  			  if s.Health != nil {
				  err = b.WriteFieldId(0)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Health)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Mana != nil {
				  err = b.WriteFieldId(1)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Mana)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Stamina != nil {
				  err = b.WriteFieldId(2)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Stamina)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.CritChance != nil {
				  err = b.WriteFieldId(3)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.CritChance)
				if err != nil {
					return err
				}
			
			  }
		  
			structLen := b.Len() - startLen
			err = b.WriteLengthAt(structLen, startLen-4)
			return err
		}
	
				func (s *Stats) deserialize(br *ByteReader) error {
			typeId, err := br.ReadTypeId()
			if err != nil || typeId != 3 {
				return fmt.Errorf("unexpected type id: expected %d, got %d", 3, typeId)
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
				if fieldId > 5 {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
									case 4:
					 s.CritDamage = new(float64); err = br.Read(s.CritDamage) 
								case 5:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startPos0 := br.Offset()
				s.Resistances = &map[string]int16{}
				for br.Offset() < startPos0+mapLen0 {
					var key0 *string
					 key0 = new(string); err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *int16
					 value0 = new(int16); err = br.Read(value0) 
					if err != nil {
						return err
					}
					(*s.Resistances)[*key0] = *value0
				}
			
								case 0:
					 s.Health = new(int32); err = br.Read(s.Health) 
								case 1:
					 s.Mana = new(int32); err = br.Read(s.Mana) 
								case 2:
					 s.Stamina = new(int32); err = br.Read(s.Stamina) 
								case 3:
					 s.CritChance = new(float64); err = br.Read(s.CritChance) 
				
				}
				if err != nil {
					return err
				}
			}
			return nil
		}
	
	
		type Item struct {
						Id *uint32
							Name *string
							Rarity *uint8
							Weight *float64
							IsQuestItem *bool
							Tags *				[]string
			
							ExtraData *				map[string]string
			
			
		}

		func (Item) TypeId() uint16 {
			return 1
		}

				func (s *Item) serialize(b *ByteWriter) error {
			err := b.WriteTypeId(1)
			if err != nil {
				return err
			}
			err = b.WriteLength(0)
			if err != nil {
				return err
			}
			startLen := b.Len()
						  if s.Id != nil {
				  err = b.WriteFieldId(0)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Id)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Name != nil {
				  err = b.WriteFieldId(1)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Name)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Rarity != nil {
				  err = b.WriteFieldId(2)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Rarity)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Weight != nil {
				  err = b.WriteFieldId(3)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Weight)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.IsQuestItem != nil {
				  err = b.WriteFieldId(4)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.IsQuestItem)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Tags != nil {
				  err = b.WriteFieldId(5)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for _, item0 := range *s.Tags {
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
		  			  if s.ExtraData != nil {
				  err = b.WriteFieldId(6)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.ExtraData {
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
		  
			structLen := b.Len() - startLen
			err = b.WriteLengthAt(structLen, startLen-4)
			return err
		}
	
				func (s *Item) deserialize(br *ByteReader) error {
			typeId, err := br.ReadTypeId()
			if err != nil || typeId != 1 {
				return fmt.Errorf("unexpected type id: expected %d, got %d", 1, typeId)
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
				if fieldId > 6 {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
									case 0:
					 s.Id = new(uint32); err = br.Read(s.Id) 
								case 1:
					 s.Name = new(string); err = br.Read(s.Name) 
								case 2:
					 s.Rarity = new(uint8); err = br.Read(s.Rarity) 
								case 3:
					 s.Weight = new(float64); err = br.Read(s.Weight) 
								case 4:
					 s.IsQuestItem = new(bool); err = br.Read(s.IsQuestItem) 
								case 5:
									listLen0, err := br.ReadLength()
				if err != nil || listLen0 < 0 || listLen0 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen0)
				}
				startPos0 := br.Offset()
				s.Tags = &[]string{}
				for br.Offset() < startPos0+listLen0 {
					var item0 *string
					 item0 = new(string); err = br.Read(item0) 
					if err != nil {
						return err
					}
					*s.Tags = append(*s.Tags, *item0)
				}
			
								case 6:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startPos0 := br.Offset()
				s.ExtraData = &map[string]string{}
				for br.Offset() < startPos0+mapLen0 {
					var key0 *string
					 key0 = new(string); err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *string
					 value0 = new(string); err = br.Read(value0) 
					if err != nil {
						return err
					}
					(*s.ExtraData)[*key0] = *value0
				}
			
				
				}
				if err != nil {
					return err
				}
			}
			return nil
		}
	
	
		type Loot struct {
						Modifiers *				map[string]float64
			
							Entries *				[]LootEntry
			
							BaseChance *float64
			
		}

		func (Loot) TypeId() uint16 {
			return 983
		}

				func (s *Loot) serialize(b *ByteWriter) error {
			err := b.WriteTypeId(983)
			if err != nil {
				return err
			}
			err = b.WriteLength(0)
			if err != nil {
				return err
			}
			startLen := b.Len()
						  if s.Modifiers != nil {
				  err = b.WriteFieldId(1)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.Modifiers {
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
		  			  if s.Entries != nil {
				  err = b.WriteFieldId(2)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for _, item0 := range *s.Entries {
									err = item0.serialize(b)
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
		  			  if s.BaseChance != nil {
				  err = b.WriteFieldId(0)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.BaseChance)
				if err != nil {
					return err
				}
			
			  }
		  
			structLen := b.Len() - startLen
			err = b.WriteLengthAt(structLen, startLen-4)
			return err
		}
	
				func (s *Loot) deserialize(br *ByteReader) error {
			typeId, err := br.ReadTypeId()
			if err != nil || typeId != 983 {
				return fmt.Errorf("unexpected type id: expected %d, got %d", 983, typeId)
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
				if fieldId > 2 {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
									case 1:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startPos0 := br.Offset()
				s.Modifiers = &map[string]float64{}
				for br.Offset() < startPos0+mapLen0 {
					var key0 *string
					 key0 = new(string); err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *float64
					 value0 = new(float64); err = br.Read(value0) 
					if err != nil {
						return err
					}
					(*s.Modifiers)[*key0] = *value0
				}
			
								case 2:
									listLen0, err := br.ReadLength()
				if err != nil || listLen0 < 0 || listLen0 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen0)
				}
				startPos0 := br.Offset()
				s.Entries = &[]LootEntry{}
				for br.Offset() < startPos0+listLen0 {
					var item0 *LootEntry
									item0 = &LootEntry{}
				err = item0.deserialize(br)
			
					if err != nil {
						return err
					}
					*s.Entries = append(*s.Entries, *item0)
				}
			
								case 0:
					 s.BaseChance = new(float64); err = br.Read(s.BaseChance) 
				
				}
				if err != nil {
					return err
				}
			}
			return nil
		}
	
	
		type Companion struct {
						Name *string
							Level *uint8
							Bond *float64
			
		}

		func (Companion) TypeId() uint16 {
			return 21813
		}

				func (s *Companion) serialize(b *ByteWriter) error {
			err := b.WriteTypeId(21813)
			if err != nil {
				return err
			}
			err = b.WriteLength(0)
			if err != nil {
				return err
			}
			startLen := b.Len()
						  if s.Name != nil {
				  err = b.WriteFieldId(0)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Name)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Level != nil {
				  err = b.WriteFieldId(1)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Level)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Bond != nil {
				  err = b.WriteFieldId(2)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Bond)
				if err != nil {
					return err
				}
			
			  }
		  
			structLen := b.Len() - startLen
			err = b.WriteLengthAt(structLen, startLen-4)
			return err
		}
	
				func (s *Companion) deserialize(br *ByteReader) error {
			typeId, err := br.ReadTypeId()
			if err != nil || typeId != 21813 {
				return fmt.Errorf("unexpected type id: expected %d, got %d", 21813, typeId)
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
				if fieldId > 2 {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
									case 0:
					 s.Name = new(string); err = br.Read(s.Name) 
								case 1:
					 s.Level = new(uint8); err = br.Read(s.Level) 
								case 2:
					 s.Bond = new(float64); err = br.Read(s.Bond) 
				
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
					case 16560:
				s := &Character{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, 16560, nil
					case 16605:
				s := &Quest{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, 16605, nil
					case 35339:
				s := &EquipmentSlot{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, 35339, nil
					case 305:
				s := &LootEntry{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, 305, nil
					case 60723:
				s := &World{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, 60723, nil
					case 2:
				s := &Vector3{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, 2, nil
					case 3:
				s := &Stats{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, 3, nil
					case 1:
				s := &Item{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, 1, nil
					case 983:
				s := &Loot{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, 983, nil
					case 21813:
				s := &Companion{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, 21813, nil
		
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