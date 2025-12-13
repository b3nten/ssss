package schema

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

		type Stats struct {
						Health *int32
							Mana *int32
							Stamina *int32
							Critchance *float32
							Critdamage *float32
							Resistances *				map[string]int16
			
			
		}

		func (Stats) TypeId() uint16 {
			return 36492
		}

				func (s *Stats) serialize(b *ByteWriter) error {
			err := b.WriteTypeId(36492)
			if err != nil {
				return err
			}
			err = b.WriteLength(0)
			if err != nil {
				return err
			}
			startLen := b.Len()
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
		  			  if s.Critchance != nil {
				  err = b.WriteFieldId(3)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Critchance)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Critdamage != nil {
				  err = b.WriteFieldId(4)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Critdamage)
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
		  
			structLen := b.Len() - startLen
			err = b.WriteLengthAt(structLen, startLen-4)
			return err
		}
	
				func (s *Stats) deserialize(br *ByteReader) error {
			typeId, err := br.ReadTypeId()
			if err != nil || typeId != 36492 {
				return fmt.Errorf("unexpected type id: expected %d, got %d", 36492, typeId)
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
				if fieldId > 5 {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
									case 0:
					 err = br.Read(s.Health) 
								case 1:
					 err = br.Read(s.Mana) 
								case 2:
					 err = br.Read(s.Stamina) 
								case 3:
					 err = br.Read(s.Critchance) 
								case 4:
					 err = br.Read(s.Critdamage) 
								case 5:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startLen0 := br.Len()
				s.Resistances = &map[string]int16{}
				for br.Len() > startLen0-mapLen0 {
					var key0 *string
					 err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *int16
					 err = br.Read(value0) 
					if err != nil {
						return err
					}
					(*s.Resistances)[*key0] = *value0
				}
			
				
				}
				if err != nil {
					return err
				}
			}
			return nil
		}
	
	
		type Item struct {
						Name *string
							Rarity *uint8
							Weight *float32
							Isquestitem *bool
							Tags *				[]string
			
							Extradata *				map[string]string
			
							Id *uint32
			
		}

		func (Item) TypeId() uint16 {
			return 32774
		}

				func (s *Item) serialize(b *ByteWriter) error {
			err := b.WriteTypeId(32774)
			if err != nil {
				return err
			}
			err = b.WriteLength(0)
			if err != nil {
				return err
			}
			startLen := b.Len()
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
		  			  if s.Isquestitem != nil {
				  err = b.WriteFieldId(4)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Isquestitem)
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
		  			  if s.Extradata != nil {
				  err = b.WriteFieldId(6)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.Extradata {
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
		  
			structLen := b.Len() - startLen
			err = b.WriteLengthAt(structLen, startLen-4)
			return err
		}
	
				func (s *Item) deserialize(br *ByteReader) error {
			typeId, err := br.ReadTypeId()
			if err != nil || typeId != 32774 {
				return fmt.Errorf("unexpected type id: expected %d, got %d", 32774, typeId)
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
				if fieldId > 6 {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
									case 1:
					 err = br.Read(s.Name) 
								case 2:
					 err = br.Read(s.Rarity) 
								case 3:
					 err = br.Read(s.Weight) 
								case 4:
					 err = br.Read(s.Isquestitem) 
								case 5:
									listLen0, err := br.ReadLength()
				if err != nil || listLen0 < 0 || listLen0 > br.Len() {
					return fmt.Errorf("invalid list length: %d", listLen0)
				}
				startLen0 := br.Len()
				s.Tags = &[]string{}
				for br.Len() > startLen0-listLen0 {
					var item0 *string
					 err = br.Read(item0) 
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
				startLen0 := br.Len()
				s.Extradata = &map[string]string{}
				for br.Len() > startLen0-mapLen0 {
					var key0 *string
					 err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *string
					 err = br.Read(value0) 
					if err != nil {
						return err
					}
					(*s.Extradata)[*key0] = *value0
				}
			
								case 0:
					 err = br.Read(s.Id) 
				
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
							Bond *float32
			
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
			startLen := br.Len()
			for br.Len() > startLen-int(length) {
				fieldId, err := br.ReadFieldId()
				if err != nil || seenFields[fieldId] {
					return fmt.Errorf("error reading field id or duplicate field id: %d", fieldId)
				}
				if fieldId > 2 {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
									case 0:
					 err = br.Read(s.Name) 
								case 1:
					 err = br.Read(s.Level) 
								case 2:
					 err = br.Read(s.Bond) 
				
				}
				if err != nil {
					return err
				}
			}
			return nil
		}
	
	
		type Loot struct {
						Basechance *float32
							Modifiers *				map[string]float32
			
			
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
						  if s.Basechance != nil {
				  err = b.WriteFieldId(0)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Basechance)
				if err != nil {
					return err
				}
			
			  }
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
									case 0:
					 err = br.Read(s.Basechance) 
								case 1:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startLen0 := br.Len()
				s.Modifiers = &map[string]float32{}
				for br.Len() > startLen0-mapLen0 {
					var key0 *string
					 err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *float32
					 err = br.Read(value0) 
					if err != nil {
						return err
					}
					(*s.Modifiers)[*key0] = *value0
				}
			
				
				}
				if err != nil {
					return err
				}
			}
			return nil
		}
	
	
		type Lootentry struct {
						Itemid *uint32
							Minqty *uint8
							Maxqty *uint8
							Conditions *				map[string]bool
			
			
		}

		func (Lootentry) TypeId() uint16 {
			return 305
		}

				func (s *Lootentry) serialize(b *ByteWriter) error {
			err := b.WriteTypeId(305)
			if err != nil {
				return err
			}
			err = b.WriteLength(0)
			if err != nil {
				return err
			}
			startLen := b.Len()
						  if s.Itemid != nil {
				  err = b.WriteFieldId(0)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Itemid)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Minqty != nil {
				  err = b.WriteFieldId(1)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Minqty)
				if err != nil {
					return err
				}
			
			  }
		  			  if s.Maxqty != nil {
				  err = b.WriteFieldId(2)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Maxqty)
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
		  
			structLen := b.Len() - startLen
			err = b.WriteLengthAt(structLen, startLen-4)
			return err
		}
	
				func (s *Lootentry) deserialize(br *ByteReader) error {
			typeId, err := br.ReadTypeId()
			if err != nil || typeId != 305 {
				return fmt.Errorf("unexpected type id: expected %d, got %d", 305, typeId)
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
				if fieldId > 3 {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
									case 0:
					 err = br.Read(s.Itemid) 
								case 1:
					 err = br.Read(s.Minqty) 
								case 2:
					 err = br.Read(s.Maxqty) 
								case 3:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startLen0 := br.Len()
				s.Conditions = &map[string]bool{}
				for br.Len() > startLen0-mapLen0 {
					var key0 *string
					 err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *bool
					 err = br.Read(value0) 
					if err != nil {
						return err
					}
					(*s.Conditions)[*key0] = *value0
				}
			
				
				}
				if err != nil {
					return err
				}
			}
			return nil
		}
	
	
		type World struct {
						Gravity *float64
							Zonedata *				map[string]				map[string]uint32
			
			
							Systemflags *				map[string]bool
			
							Loottables *				map[string]Loot
			
							Worldname *string
							Seed *uint64
			
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
		  			  if s.Zonedata != nil {
				  err = b.WriteFieldId(5)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.Zonedata {
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
		  			  if s.Systemflags != nil {
				  err = b.WriteFieldId(6)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.Systemflags {
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
		  			  if s.Loottables != nil {
				  err = b.WriteFieldId(7)
				  if err != nil {
					  return err
				  }
									err = b.WriteLength(0)
				if err != nil {
					return err
				}
				startLen0 := b.Len()
				for key0, value0 := range *s.Loottables {
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
		  			  if s.Worldname != nil {
				  err = b.WriteFieldId(0)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Worldname)
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
									case 2:
					 err = br.Read(s.Gravity) 
								case 5:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startLen0 := br.Len()
				s.Zonedata = &map[string]map[string]uint32{}
				for br.Len() > startLen0-mapLen0 {
					var key0 *string
					 err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *map[string]uint32
									mapLen1, err := br.ReadLength()
				if err != nil || mapLen1 < 0 || mapLen1 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen1)
				}
				startLen1 := br.Len()
				value0 = &map[string]uint32{}
				for br.Len() > startLen1-mapLen1 {
					var key1 *string
					 err = br.Read(key1) 
					if err != nil {
						return err
					}
					var value1 *uint32
					 err = br.Read(value1) 
					if err != nil {
						return err
					}
					(*value0)[*key1] = *value1
				}
			
					if err != nil {
						return err
					}
					(*s.Zonedata)[*key0] = *value0
				}
			
								case 6:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startLen0 := br.Len()
				s.Systemflags = &map[string]bool{}
				for br.Len() > startLen0-mapLen0 {
					var key0 *string
					 err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *bool
					 err = br.Read(value0) 
					if err != nil {
						return err
					}
					(*s.Systemflags)[*key0] = *value0
				}
			
								case 7:
									mapLen0, err := br.ReadLength()
				if err != nil || mapLen0 < 0 || mapLen0 > br.Len() {
					return fmt.Errorf("invalid map length: %d", mapLen0)
				}
				startLen0 := br.Len()
				s.Loottables = &map[string]Loot{}
				for br.Len() > startLen0-mapLen0 {
					var key0 *string
					 err = br.Read(key0) 
					if err != nil {
						return err
					}
					var value0 *Loot
									value0 = &Loot{}
				err = value0.deserialize(br)
			
					if err != nil {
						return err
					}
					(*s.Loottables)[*key0] = *value0
				}
			
								case 0:
					 err = br.Read(s.Worldname) 
								case 1:
					 err = br.Read(s.Seed) 
				
				}
				if err != nil {
					return err
				}
			}
			return nil
		}
	
	
		type Equipmentslot struct {
						Slotname *string
			
		}

		func (Equipmentslot) TypeId() uint16 {
			return 35339
		}

				func (s *Equipmentslot) serialize(b *ByteWriter) error {
			err := b.WriteTypeId(35339)
			if err != nil {
				return err
			}
			err = b.WriteLength(0)
			if err != nil {
				return err
			}
			startLen := b.Len()
						  if s.Slotname != nil {
				  err = b.WriteFieldId(0)
				  if err != nil {
					  return err
				  }
									err = b.Write(s.Slotname)
				if err != nil {
					return err
				}
			
			  }
		  
			structLen := b.Len() - startLen
			err = b.WriteLengthAt(structLen, startLen-4)
			return err
		}
	
				func (s *Equipmentslot) deserialize(br *ByteReader) error {
			typeId, err := br.ReadTypeId()
			if err != nil || typeId != 35339 {
				return fmt.Errorf("unexpected type id: expected %d, got %d", 35339, typeId)
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
				if fieldId > 0 {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
									case 0:
					 err = br.Read(s.Slotname) 
				
				}
				if err != nil {
					return err
				}
			}
			return nil
		}
	
	
		type Vector3 struct {
						Z *float32
							X *float32
							Y *float32
			
		}

		func (Vector3) TypeId() uint16 {
			return 29123
		}

				func (s *Vector3) serialize(b *ByteWriter) error {
			err := b.WriteTypeId(29123)
			if err != nil {
				return err
			}
			err = b.WriteLength(0)
			if err != nil {
				return err
			}
			startLen := b.Len()
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
		  
			structLen := b.Len() - startLen
			err = b.WriteLengthAt(structLen, startLen-4)
			return err
		}
	
				func (s *Vector3) deserialize(br *ByteReader) error {
			typeId, err := br.ReadTypeId()
			if err != nil || typeId != 29123 {
				return fmt.Errorf("unexpected type id: expected %d, got %d", 29123, typeId)
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
				if fieldId > 2 {
					return nil
				}
				seenFields[fieldId] = true
				switch fieldId {
									case 2:
					 err = br.Read(s.Z) 
								case 0:
					 err = br.Read(s.X) 
								case 1:
					 err = br.Read(s.Y) 
				
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
					case 36492:
				s := &Stats{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, 36492, nil
					case 32774:
				s := &Item{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, 32774, nil
					case 21813:
				s := &Companion{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, 21813, nil
					case 983:
				s := &Loot{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, 983, nil
					case 305:
				s := &Lootentry{}
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
					case 35339:
				s := &Equipmentslot{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, 35339, nil
					case 29123:
				s := &Vector3{}
				err = s.deserialize(br)
				if err != nil {
					return nil, 0, err
				}
				return s, 29123, nil
		
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