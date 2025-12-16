		export class Stats {
			static get TypeID() { return 3 }

			constructor(fields) {
				if (fields) Object.assign(this, fields)
			}

							health
							mana
							stamina
							critChance
							critDamage
							resistances
			

			deserialize(bytes) {
				const b = new ByteReader(bytes)
				Stats_deserialize(b, this)
				return this;
			}

			serialize() {
				const w = new ByteWriter;
				Stats_serialize(w, this)
				return w.bytes()
			}
		}

				function Stats_serialize(b, s) {
			b.write_uint16(3)
			b.write_uint32(0);
			const structStart = b.length;
						if(s.health !== undefined) {
				b.write_uint16(0);
								b.write_int32(s.health)
			
			}
		  			if(s.mana !== undefined) {
				b.write_uint16(1);
								b.write_int32(s.mana)
			
			}
		  			if(s.stamina !== undefined) {
				b.write_uint16(2);
								b.write_int32(s.stamina)
			
			}
		  			if(s.critChance !== undefined) {
				b.write_uint16(3);
								b.write_f64(s.critChance)
			
			}
		  			if(s.critDamage !== undefined) {
				b.write_uint16(4);
								b.write_f64(s.critDamage)
			
			}
		  			if(s.resistances !== undefined) {
				b.write_uint16(5);
								b.write_uint32(0);
				const mapStart0 = b.length;
				for(const [key0, value0] of Object.entries(s.resistances)) {
									b.write_string(key0)
			
									b.write_int16(value0)
			
				}
				b.set_uint32(mapStart0 - 4, b.length - mapStart0);
			
			}
		  
			b.set_uint32(structStart - 4, b.length - structStart);
		}
	
				function Stats_deserialize(br, s) {
			const typeId = br.read_uint16()
			if (typeId !== 3) {
				throw new Error("Type ID mismatch deserializing struct Stats: expected 3, got " + typeId);
			}
			const length = br.read_uint32()
			if (length > (br.length - br.position)) {
				throw new Error("Struct Stats length exceeds buffer length");
			}
			const seenFields = new Set;
			const startPos = br.position;
			for (; br.position - startPos < length;) {
				const fieldId = br.read_uint16();
				if (seenFields.has(fieldId)) {
					throw new Error("Duplicate field ID " + fieldId + " in struct Stats");
				}
				if (fieldId > 5) {
					return;
				}
				seenFields.add(fieldId);
				switch (fieldId) {
									case 0:
									s.health = br.read_int32();
			
					break;
								case 1:
									s.mana = br.read_int32();
			
					break;
								case 2:
									s.stamina = br.read_int32();
			
					break;
								case 3:
									s.critChance = br.read_f64();
			
					break;
								case 4:
									s.critDamage = br.read_f64();
			
					break;
								case 5:
								{
				const mapLength0 = br.read_uint32();
				if (mapLength0 > (br.length-br.position)) {
					throw new Error("Invalid map length");
				}
				const mapStart0 = br.position;
				s.resistances = {};
				for (; br.position - mapStart0 < mapLength0;) {
					let key0;
									key0 = br.read_string();
			
					let value0;
									value0 = br.read_int16();
			
					s.resistances[key0] = value0;
				}
			}
			
					break;
				
				}
			}
		}
	
	
		export class Item {
			static get TypeID() { return 1 }

			constructor(fields) {
				if (fields) Object.assign(this, fields)
			}

							weight
							isQuestItem
							tags
							extraData
							id
							name
							rarity
			

			deserialize(bytes) {
				const b = new ByteReader(bytes)
				Item_deserialize(b, this)
				return this;
			}

			serialize() {
				const w = new ByteWriter;
				Item_serialize(w, this)
				return w.bytes()
			}
		}

				function Item_serialize(b, s) {
			b.write_uint16(1)
			b.write_uint32(0);
			const structStart = b.length;
						if(s.weight !== undefined) {
				b.write_uint16(3);
								b.write_f64(s.weight)
			
			}
		  			if(s.isQuestItem !== undefined) {
				b.write_uint16(4);
								b.write_bool(s.isQuestItem)
			
			}
		  			if(s.tags !== undefined) {
				b.write_uint16(5);
								b.write_uint32(0);
				const listStart0 = b.length;
				for(const item0 of s.tags) {
									b.write_string(item0)
			
				}
				b.set_uint32(listStart0 - 4, b.length - listStart0);
			
			}
		  			if(s.extraData !== undefined) {
				b.write_uint16(6);
								b.write_uint32(0);
				const mapStart0 = b.length;
				for(const [key0, value0] of Object.entries(s.extraData)) {
									b.write_string(key0)
			
									b.write_string(value0)
			
				}
				b.set_uint32(mapStart0 - 4, b.length - mapStart0);
			
			}
		  			if(s.id !== undefined) {
				b.write_uint16(0);
								b.write_uint32(s.id)
			
			}
		  			if(s.name !== undefined) {
				b.write_uint16(1);
								b.write_string(s.name)
			
			}
		  			if(s.rarity !== undefined) {
				b.write_uint16(2);
								b.write_uint8(s.rarity)
			
			}
		  
			b.set_uint32(structStart - 4, b.length - structStart);
		}
	
				function Item_deserialize(br, s) {
			const typeId = br.read_uint16()
			if (typeId !== 1) {
				throw new Error("Type ID mismatch deserializing struct Item: expected 1, got " + typeId);
			}
			const length = br.read_uint32()
			if (length > (br.length - br.position)) {
				throw new Error("Struct Item length exceeds buffer length");
			}
			const seenFields = new Set;
			const startPos = br.position;
			for (; br.position - startPos < length;) {
				const fieldId = br.read_uint16();
				if (seenFields.has(fieldId)) {
					throw new Error("Duplicate field ID " + fieldId + " in struct Item");
				}
				if (fieldId > 6) {
					return;
				}
				seenFields.add(fieldId);
				switch (fieldId) {
									case 3:
									s.weight = br.read_f64();
			
					break;
								case 4:
									s.isQuestItem = br.read_bool();
			
					break;
								case 5:
								{
				const listLength0 = br.read_uint32();
				if (listLength0 > (br.length - br.position)) {
					throw new Error("Invalid list length");
				}
				const listStart0 = br.position;
				s.tags = [];
				for (; br.position - listStart0 < listLength0;) {
					let item0;
									item0 = br.read_string();
			
					s.tags.push(item0);
				}
			}
			
					break;
								case 6:
								{
				const mapLength0 = br.read_uint32();
				if (mapLength0 > (br.length-br.position)) {
					throw new Error("Invalid map length");
				}
				const mapStart0 = br.position;
				s.extraData = {};
				for (; br.position - mapStart0 < mapLength0;) {
					let key0;
									key0 = br.read_string();
			
					let value0;
									value0 = br.read_string();
			
					s.extraData[key0] = value0;
				}
			}
			
					break;
								case 0:
									s.id = br.read_uint32();
			
					break;
								case 1:
									s.name = br.read_string();
			
					break;
								case 2:
									s.rarity = br.read_uint8();
			
					break;
				
				}
			}
		}
	
	
		export class EquipmentSlot {
			static get TypeID() { return 35339 }

			constructor(fields) {
				if (fields) Object.assign(this, fields)
			}

							slotName
							item
			

			deserialize(bytes) {
				const b = new ByteReader(bytes)
				EquipmentSlot_deserialize(b, this)
				return this;
			}

			serialize() {
				const w = new ByteWriter;
				EquipmentSlot_serialize(w, this)
				return w.bytes()
			}
		}

				function EquipmentSlot_serialize(b, s) {
			b.write_uint16(35339)
			b.write_uint32(0);
			const structStart = b.length;
						if(s.slotName !== undefined) {
				b.write_uint16(0);
								b.write_string(s.slotName)
			
			}
		  			if(s.item !== undefined) {
				b.write_uint16(1);
								Item_serialize(b, s.item)
			
			}
		  
			b.set_uint32(structStart - 4, b.length - structStart);
		}
	
				function EquipmentSlot_deserialize(br, s) {
			const typeId = br.read_uint16()
			if (typeId !== 35339) {
				throw new Error("Type ID mismatch deserializing struct EquipmentSlot: expected 35339, got " + typeId);
			}
			const length = br.read_uint32()
			if (length > (br.length - br.position)) {
				throw new Error("Struct EquipmentSlot length exceeds buffer length");
			}
			const seenFields = new Set;
			const startPos = br.position;
			for (; br.position - startPos < length;) {
				const fieldId = br.read_uint16();
				if (seenFields.has(fieldId)) {
					throw new Error("Duplicate field ID " + fieldId + " in struct EquipmentSlot");
				}
				if (fieldId > 1) {
					return;
				}
				seenFields.add(fieldId);
				switch (fieldId) {
									case 0:
									s.slotName = br.read_string();
			
					break;
								case 1:
									s.item = new Item();
				Item_deserialize(br, s.item);
			
					break;
				
				}
			}
		}
	
	
		export class Character {
			static get TypeID() { return 16560 }

			constructor(fields) {
				if (fields) Object.assign(this, fields)
			}

							position
							arbitraryData
							equipment
							friends
							name
							id
							stats
							inventory
							companions
							skillProgress
			

			deserialize(bytes) {
				const b = new ByteReader(bytes)
				Character_deserialize(b, this)
				return this;
			}

			serialize() {
				const w = new ByteWriter;
				Character_serialize(w, this)
				return w.bytes()
			}
		}

				function Character_serialize(b, s) {
			b.write_uint16(16560)
			b.write_uint32(0);
			const structStart = b.length;
						if(s.position !== undefined) {
				b.write_uint16(2);
								Vector3_serialize(b, s.position)
			
			}
		  			if(s.arbitraryData !== undefined) {
				b.write_uint16(9);
								b.write_uint32(0);
				const mapStart0 = b.length;
				for(const [key0, value0] of Object.entries(s.arbitraryData)) {
									b.write_string(key0)
			
									b.write_uint32(0);
				const mapStart1 = b.length;
				for(const [key1, value1] of Object.entries(value0)) {
									b.write_string(key1)
			
									b.write_uint32(0);
				const mapStart2 = b.length;
				for(const [key2, value2] of Object.entries(value1)) {
									b.write_string(key2)
			
									b.write_int8(value2)
			
				}
				b.set_uint32(mapStart2 - 4, b.length - mapStart2);
			
				}
				b.set_uint32(mapStart1 - 4, b.length - mapStart1);
			
				}
				b.set_uint32(mapStart0 - 4, b.length - mapStart0);
			
			}
		  			if(s.equipment !== undefined) {
				b.write_uint16(5);
								b.write_uint32(0);
				const mapStart0 = b.length;
				for(const [key0, value0] of Object.entries(s.equipment)) {
									b.write_string(key0)
			
									EquipmentSlot_serialize(b, value0)
			
				}
				b.set_uint32(mapStart0 - 4, b.length - mapStart0);
			
			}
		  			if(s.friends !== undefined) {
				b.write_uint16(7);
								b.write_uint32(0);
				const listStart0 = b.length;
				for(const item0 of s.friends) {
									Character_serialize(b, item0)
			
				}
				b.set_uint32(listStart0 - 4, b.length - listStart0);
			
			}
		  			if(s.name !== undefined) {
				b.write_uint16(1);
								b.write_string(s.name)
			
			}
		  			if(s.id !== undefined) {
				b.write_uint16(0);
								b.write_uint64(s.id)
			
			}
		  			if(s.stats !== undefined) {
				b.write_uint16(3);
								Stats_serialize(b, s.stats)
			
			}
		  			if(s.inventory !== undefined) {
				b.write_uint16(4);
								b.write_uint32(0);
				const listStart0 = b.length;
				for(const item0 of s.inventory) {
									b.write_uint32(0);
				const listStart1 = b.length;
				for(const item1 of item0) {
									Item_serialize(b, item1)
			
				}
				b.set_uint32(listStart1 - 4, b.length - listStart1);
			
				}
				b.set_uint32(listStart0 - 4, b.length - listStart0);
			
			}
		  			if(s.companions !== undefined) {
				b.write_uint16(6);
								b.write_uint32(0);
				const mapStart0 = b.length;
				for(const [key0, value0] of Object.entries(s.companions)) {
									b.write_uint16(key0)
			
									Companion_serialize(b, value0)
			
				}
				b.set_uint32(mapStart0 - 4, b.length - mapStart0);
			
			}
		  			if(s.skillProgress !== undefined) {
				b.write_uint16(8);
								b.write_uint32(0);
				const mapStart0 = b.length;
				for(const [key0, value0] of Object.entries(s.skillProgress)) {
									b.write_string(key0)
			
									b.write_uint32(0);
				const listStart1 = b.length;
				for(const item1 of value0) {
									b.write_f64(item1)
			
				}
				b.set_uint32(listStart1 - 4, b.length - listStart1);
			
				}
				b.set_uint32(mapStart0 - 4, b.length - mapStart0);
			
			}
		  
			b.set_uint32(structStart - 4, b.length - structStart);
		}
	
				function Character_deserialize(br, s) {
			const typeId = br.read_uint16()
			if (typeId !== 16560) {
				throw new Error("Type ID mismatch deserializing struct Character: expected 16560, got " + typeId);
			}
			const length = br.read_uint32()
			if (length > (br.length - br.position)) {
				throw new Error("Struct Character length exceeds buffer length");
			}
			const seenFields = new Set;
			const startPos = br.position;
			for (; br.position - startPos < length;) {
				const fieldId = br.read_uint16();
				if (seenFields.has(fieldId)) {
					throw new Error("Duplicate field ID " + fieldId + " in struct Character");
				}
				if (fieldId > 9) {
					return;
				}
				seenFields.add(fieldId);
				switch (fieldId) {
									case 2:
									s.position = new Vector3();
				Vector3_deserialize(br, s.position);
			
					break;
								case 9:
								{
				const mapLength0 = br.read_uint32();
				if (mapLength0 > (br.length-br.position)) {
					throw new Error("Invalid map length");
				}
				const mapStart0 = br.position;
				s.arbitraryData = {};
				for (; br.position - mapStart0 < mapLength0;) {
					let key0;
									key0 = br.read_string();
			
					let value0;
								{
				const mapLength1 = br.read_uint32();
				if (mapLength1 > (br.length-br.position)) {
					throw new Error("Invalid map length");
				}
				const mapStart1 = br.position;
				value0 = {};
				for (; br.position - mapStart1 < mapLength1;) {
					let key1;
									key1 = br.read_string();
			
					let value1;
								{
				const mapLength2 = br.read_uint32();
				if (mapLength2 > (br.length-br.position)) {
					throw new Error("Invalid map length");
				}
				const mapStart2 = br.position;
				value1 = {};
				for (; br.position - mapStart2 < mapLength2;) {
					let key2;
									key2 = br.read_string();
			
					let value2;
									value2 = br.read_int8();
			
					value1[key2] = value2;
				}
			}
			
					value0[key1] = value1;
				}
			}
			
					s.arbitraryData[key0] = value0;
				}
			}
			
					break;
								case 5:
								{
				const mapLength0 = br.read_uint32();
				if (mapLength0 > (br.length-br.position)) {
					throw new Error("Invalid map length");
				}
				const mapStart0 = br.position;
				s.equipment = {};
				for (; br.position - mapStart0 < mapLength0;) {
					let key0;
									key0 = br.read_string();
			
					let value0;
									value0 = new EquipmentSlot();
				EquipmentSlot_deserialize(br, value0);
			
					s.equipment[key0] = value0;
				}
			}
			
					break;
								case 7:
								{
				const listLength0 = br.read_uint32();
				if (listLength0 > (br.length - br.position)) {
					throw new Error("Invalid list length");
				}
				const listStart0 = br.position;
				s.friends = [];
				for (; br.position - listStart0 < listLength0;) {
					let item0;
									item0 = new Character();
				Character_deserialize(br, item0);
			
					s.friends.push(item0);
				}
			}
			
					break;
								case 1:
									s.name = br.read_string();
			
					break;
								case 0:
									s.id = br.read_uint64();
			
					break;
								case 3:
									s.stats = new Stats();
				Stats_deserialize(br, s.stats);
			
					break;
								case 4:
								{
				const listLength0 = br.read_uint32();
				if (listLength0 > (br.length - br.position)) {
					throw new Error("Invalid list length");
				}
				const listStart0 = br.position;
				s.inventory = [];
				for (; br.position - listStart0 < listLength0;) {
					let item0;
								{
				const listLength1 = br.read_uint32();
				if (listLength1 > (br.length - br.position)) {
					throw new Error("Invalid list length");
				}
				const listStart1 = br.position;
				item0 = [];
				for (; br.position - listStart1 < listLength1;) {
					let item1;
									item1 = new Item();
				Item_deserialize(br, item1);
			
					item0.push(item1);
				}
			}
			
					s.inventory.push(item0);
				}
			}
			
					break;
								case 6:
								{
				const mapLength0 = br.read_uint32();
				if (mapLength0 > (br.length-br.position)) {
					throw new Error("Invalid map length");
				}
				const mapStart0 = br.position;
				s.companions = {};
				for (; br.position - mapStart0 < mapLength0;) {
					let key0;
									key0 = br.read_uint16();
			
					let value0;
									value0 = new Companion();
				Companion_deserialize(br, value0);
			
					s.companions[key0] = value0;
				}
			}
			
					break;
								case 8:
								{
				const mapLength0 = br.read_uint32();
				if (mapLength0 > (br.length-br.position)) {
					throw new Error("Invalid map length");
				}
				const mapStart0 = br.position;
				s.skillProgress = {};
				for (; br.position - mapStart0 < mapLength0;) {
					let key0;
									key0 = br.read_string();
			
					let value0;
								{
				const listLength1 = br.read_uint32();
				if (listLength1 > (br.length - br.position)) {
					throw new Error("Invalid list length");
				}
				const listStart1 = br.position;
				value0 = [];
				for (; br.position - listStart1 < listLength1;) {
					let item1;
									item1 = br.read_f64();
			
					value0.push(item1);
				}
			}
			
					s.skillProgress[key0] = value0;
				}
			}
			
					break;
				
				}
			}
		}
	
	
		export class Vector3 {
			static get TypeID() { return 2 }

			constructor(fields) {
				if (fields) Object.assign(this, fields)
			}

							x
							y
							z
			

			deserialize(bytes) {
				const b = new ByteReader(bytes)
				Vector3_deserialize(b, this)
				return this;
			}

			serialize() {
				const w = new ByteWriter;
				Vector3_serialize(w, this)
				return w.bytes()
			}
		}

				function Vector3_serialize(b, s) {
			b.write_uint16(2)
			b.write_uint32(0);
			const structStart = b.length;
						if(s.x !== undefined) {
				b.write_uint16(0);
								b.write_f64(s.x)
			
			}
		  			if(s.y !== undefined) {
				b.write_uint16(1);
								b.write_f64(s.y)
			
			}
		  			if(s.z !== undefined) {
				b.write_uint16(2);
								b.write_f64(s.z)
			
			}
		  
			b.set_uint32(structStart - 4, b.length - structStart);
		}
	
				function Vector3_deserialize(br, s) {
			const typeId = br.read_uint16()
			if (typeId !== 2) {
				throw new Error("Type ID mismatch deserializing struct Vector3: expected 2, got " + typeId);
			}
			const length = br.read_uint32()
			if (length > (br.length - br.position)) {
				throw new Error("Struct Vector3 length exceeds buffer length");
			}
			const seenFields = new Set;
			const startPos = br.position;
			for (; br.position - startPos < length;) {
				const fieldId = br.read_uint16();
				if (seenFields.has(fieldId)) {
					throw new Error("Duplicate field ID " + fieldId + " in struct Vector3");
				}
				if (fieldId > 2) {
					return;
				}
				seenFields.add(fieldId);
				switch (fieldId) {
									case 0:
									s.x = br.read_f64();
			
					break;
								case 1:
									s.y = br.read_f64();
			
					break;
								case 2:
									s.z = br.read_f64();
			
					break;
				
				}
			}
		}
	
	
		export class LootEntry {
			static get TypeID() { return 305 }

			constructor(fields) {
				if (fields) Object.assign(this, fields)
			}

							conditions
							itemId
							minQty
							maxQty
			

			deserialize(bytes) {
				const b = new ByteReader(bytes)
				LootEntry_deserialize(b, this)
				return this;
			}

			serialize() {
				const w = new ByteWriter;
				LootEntry_serialize(w, this)
				return w.bytes()
			}
		}

				function LootEntry_serialize(b, s) {
			b.write_uint16(305)
			b.write_uint32(0);
			const structStart = b.length;
						if(s.conditions !== undefined) {
				b.write_uint16(3);
								b.write_uint32(0);
				const mapStart0 = b.length;
				for(const [key0, value0] of Object.entries(s.conditions)) {
									b.write_string(key0)
			
									b.write_bool(value0)
			
				}
				b.set_uint32(mapStart0 - 4, b.length - mapStart0);
			
			}
		  			if(s.itemId !== undefined) {
				b.write_uint16(0);
								b.write_uint32(s.itemId)
			
			}
		  			if(s.minQty !== undefined) {
				b.write_uint16(1);
								b.write_uint8(s.minQty)
			
			}
		  			if(s.maxQty !== undefined) {
				b.write_uint16(2);
								b.write_uint8(s.maxQty)
			
			}
		  
			b.set_uint32(structStart - 4, b.length - structStart);
		}
	
				function LootEntry_deserialize(br, s) {
			const typeId = br.read_uint16()
			if (typeId !== 305) {
				throw new Error("Type ID mismatch deserializing struct LootEntry: expected 305, got " + typeId);
			}
			const length = br.read_uint32()
			if (length > (br.length - br.position)) {
				throw new Error("Struct LootEntry length exceeds buffer length");
			}
			const seenFields = new Set;
			const startPos = br.position;
			for (; br.position - startPos < length;) {
				const fieldId = br.read_uint16();
				if (seenFields.has(fieldId)) {
					throw new Error("Duplicate field ID " + fieldId + " in struct LootEntry");
				}
				if (fieldId > 3) {
					return;
				}
				seenFields.add(fieldId);
				switch (fieldId) {
									case 3:
								{
				const mapLength0 = br.read_uint32();
				if (mapLength0 > (br.length-br.position)) {
					throw new Error("Invalid map length");
				}
				const mapStart0 = br.position;
				s.conditions = {};
				for (; br.position - mapStart0 < mapLength0;) {
					let key0;
									key0 = br.read_string();
			
					let value0;
									value0 = br.read_bool();
			
					s.conditions[key0] = value0;
				}
			}
			
					break;
								case 0:
									s.itemId = br.read_uint32();
			
					break;
								case 1:
									s.minQty = br.read_uint8();
			
					break;
								case 2:
									s.maxQty = br.read_uint8();
			
					break;
				
				}
			}
		}
	
	
		export class Quest {
			static get TypeID() { return 16605 }

			constructor(fields) {
				if (fields) Object.assign(this, fields)
			}

							requiredPos
							nextSteps
							description
							title
							rewards
							objectives
							difficulty
							areaLayers
							id
							prerequisites
			

			deserialize(bytes) {
				const b = new ByteReader(bytes)
				Quest_deserialize(b, this)
				return this;
			}

			serialize() {
				const w = new ByteWriter;
				Quest_serialize(w, this)
				return w.bytes()
			}
		}

				function Quest_serialize(b, s) {
			b.write_uint16(16605)
			b.write_uint32(0);
			const structStart = b.length;
						if(s.requiredPos !== undefined) {
				b.write_uint16(4);
								b.write_uint32(0);
				const listStart0 = b.length;
				for(const item0 of s.requiredPos) {
									Vector3_serialize(b, item0)
			
				}
				b.set_uint32(listStart0 - 4, b.length - listStart0);
			
			}
		  			if(s.nextSteps !== undefined) {
				b.write_uint16(7);
								b.write_uint32(0);
				const mapStart0 = b.length;
				for(const [key0, value0] of Object.entries(s.nextSteps)) {
									b.write_string(key0)
			
									Quest_serialize(b, value0)
			
				}
				b.set_uint32(mapStart0 - 4, b.length - mapStart0);
			
			}
		  			if(s.description !== undefined) {
				b.write_uint16(9);
								b.write_string(s.description)
			
			}
		  			if(s.title !== undefined) {
				b.write_uint16(1);
								b.write_string(s.title)
			
			}
		  			if(s.rewards !== undefined) {
				b.write_uint16(3);
								b.write_uint32(0);
				const listStart0 = b.length;
				for(const item0 of s.rewards) {
									Item_serialize(b, item0)
			
				}
				b.set_uint32(listStart0 - 4, b.length - listStart0);
			
			}
		  			if(s.objectives !== undefined) {
				b.write_uint16(5);
								b.write_uint32(0);
				const mapStart0 = b.length;
				for(const [key0, value0] of Object.entries(s.objectives)) {
									b.write_string(key0)
			
									b.write_uint32(0);
				const listStart1 = b.length;
				for(const item1 of value0) {
									b.write_string(item1)
			
				}
				b.set_uint32(listStart1 - 4, b.length - listStart1);
			
				}
				b.set_uint32(mapStart0 - 4, b.length - mapStart0);
			
			}
		  			if(s.difficulty !== undefined) {
				b.write_uint16(2);
								b.write_uint8(s.difficulty)
			
			}
		  			if(s.areaLayers !== undefined) {
				b.write_uint16(8);
								b.write_uint32(0);
				const listStart0 = b.length;
				for(const item0 of s.areaLayers) {
									b.write_uint32(0);
				const listStart1 = b.length;
				for(const item1 of item0) {
									b.write_uint32(0);
				const listStart2 = b.length;
				for(const item2 of item1) {
									b.write_uint16(item2)
			
				}
				b.set_uint32(listStart2 - 4, b.length - listStart2);
			
				}
				b.set_uint32(listStart1 - 4, b.length - listStart1);
			
				}
				b.set_uint32(listStart0 - 4, b.length - listStart0);
			
			}
		  			if(s.id !== undefined) {
				b.write_uint16(0);
								b.write_uint32(s.id)
			
			}
		  			if(s.prerequisites !== undefined) {
				b.write_uint16(6);
								b.write_uint32(0);
				const listStart0 = b.length;
				for(const item0 of s.prerequisites) {
									Quest_serialize(b, item0)
			
				}
				b.set_uint32(listStart0 - 4, b.length - listStart0);
			
			}
		  
			b.set_uint32(structStart - 4, b.length - structStart);
		}
	
				function Quest_deserialize(br, s) {
			const typeId = br.read_uint16()
			if (typeId !== 16605) {
				throw new Error("Type ID mismatch deserializing struct Quest: expected 16605, got " + typeId);
			}
			const length = br.read_uint32()
			if (length > (br.length - br.position)) {
				throw new Error("Struct Quest length exceeds buffer length");
			}
			const seenFields = new Set;
			const startPos = br.position;
			for (; br.position - startPos < length;) {
				const fieldId = br.read_uint16();
				if (seenFields.has(fieldId)) {
					throw new Error("Duplicate field ID " + fieldId + " in struct Quest");
				}
				if (fieldId > 9) {
					return;
				}
				seenFields.add(fieldId);
				switch (fieldId) {
									case 4:
								{
				const listLength0 = br.read_uint32();
				if (listLength0 > (br.length - br.position)) {
					throw new Error("Invalid list length");
				}
				const listStart0 = br.position;
				s.requiredPos = [];
				for (; br.position - listStart0 < listLength0;) {
					let item0;
									item0 = new Vector3();
				Vector3_deserialize(br, item0);
			
					s.requiredPos.push(item0);
				}
			}
			
					break;
								case 7:
								{
				const mapLength0 = br.read_uint32();
				if (mapLength0 > (br.length-br.position)) {
					throw new Error("Invalid map length");
				}
				const mapStart0 = br.position;
				s.nextSteps = {};
				for (; br.position - mapStart0 < mapLength0;) {
					let key0;
									key0 = br.read_string();
			
					let value0;
									value0 = new Quest();
				Quest_deserialize(br, value0);
			
					s.nextSteps[key0] = value0;
				}
			}
			
					break;
								case 9:
									s.description = br.read_string();
			
					break;
								case 1:
									s.title = br.read_string();
			
					break;
								case 3:
								{
				const listLength0 = br.read_uint32();
				if (listLength0 > (br.length - br.position)) {
					throw new Error("Invalid list length");
				}
				const listStart0 = br.position;
				s.rewards = [];
				for (; br.position - listStart0 < listLength0;) {
					let item0;
									item0 = new Item();
				Item_deserialize(br, item0);
			
					s.rewards.push(item0);
				}
			}
			
					break;
								case 5:
								{
				const mapLength0 = br.read_uint32();
				if (mapLength0 > (br.length-br.position)) {
					throw new Error("Invalid map length");
				}
				const mapStart0 = br.position;
				s.objectives = {};
				for (; br.position - mapStart0 < mapLength0;) {
					let key0;
									key0 = br.read_string();
			
					let value0;
								{
				const listLength1 = br.read_uint32();
				if (listLength1 > (br.length - br.position)) {
					throw new Error("Invalid list length");
				}
				const listStart1 = br.position;
				value0 = [];
				for (; br.position - listStart1 < listLength1;) {
					let item1;
									item1 = br.read_string();
			
					value0.push(item1);
				}
			}
			
					s.objectives[key0] = value0;
				}
			}
			
					break;
								case 2:
									s.difficulty = br.read_uint8();
			
					break;
								case 8:
								{
				const listLength0 = br.read_uint32();
				if (listLength0 > (br.length - br.position)) {
					throw new Error("Invalid list length");
				}
				const listStart0 = br.position;
				s.areaLayers = [];
				for (; br.position - listStart0 < listLength0;) {
					let item0;
								{
				const listLength1 = br.read_uint32();
				if (listLength1 > (br.length - br.position)) {
					throw new Error("Invalid list length");
				}
				const listStart1 = br.position;
				item0 = [];
				for (; br.position - listStart1 < listLength1;) {
					let item1;
								{
				const listLength2 = br.read_uint32();
				if (listLength2 > (br.length - br.position)) {
					throw new Error("Invalid list length");
				}
				const listStart2 = br.position;
				item1 = [];
				for (; br.position - listStart2 < listLength2;) {
					let item2;
									item2 = br.read_uint16();
			
					item1.push(item2);
				}
			}
			
					item0.push(item1);
				}
			}
			
					s.areaLayers.push(item0);
				}
			}
			
					break;
								case 0:
									s.id = br.read_uint32();
			
					break;
								case 6:
								{
				const listLength0 = br.read_uint32();
				if (listLength0 > (br.length - br.position)) {
					throw new Error("Invalid list length");
				}
				const listStart0 = br.position;
				s.prerequisites = [];
				for (; br.position - listStart0 < listLength0;) {
					let item0;
									item0 = new Quest();
				Quest_deserialize(br, item0);
			
					s.prerequisites.push(item0);
				}
			}
			
					break;
				
				}
			}
		}
	
	
		export class Companion {
			static get TypeID() { return 21813 }

			constructor(fields) {
				if (fields) Object.assign(this, fields)
			}

							level
							bond
							name
			

			deserialize(bytes) {
				const b = new ByteReader(bytes)
				Companion_deserialize(b, this)
				return this;
			}

			serialize() {
				const w = new ByteWriter;
				Companion_serialize(w, this)
				return w.bytes()
			}
		}

				function Companion_serialize(b, s) {
			b.write_uint16(21813)
			b.write_uint32(0);
			const structStart = b.length;
						if(s.level !== undefined) {
				b.write_uint16(1);
								b.write_uint8(s.level)
			
			}
		  			if(s.bond !== undefined) {
				b.write_uint16(2);
								b.write_f64(s.bond)
			
			}
		  			if(s.name !== undefined) {
				b.write_uint16(0);
								b.write_string(s.name)
			
			}
		  
			b.set_uint32(structStart - 4, b.length - structStart);
		}
	
				function Companion_deserialize(br, s) {
			const typeId = br.read_uint16()
			if (typeId !== 21813) {
				throw new Error("Type ID mismatch deserializing struct Companion: expected 21813, got " + typeId);
			}
			const length = br.read_uint32()
			if (length > (br.length - br.position)) {
				throw new Error("Struct Companion length exceeds buffer length");
			}
			const seenFields = new Set;
			const startPos = br.position;
			for (; br.position - startPos < length;) {
				const fieldId = br.read_uint16();
				if (seenFields.has(fieldId)) {
					throw new Error("Duplicate field ID " + fieldId + " in struct Companion");
				}
				if (fieldId > 2) {
					return;
				}
				seenFields.add(fieldId);
				switch (fieldId) {
									case 1:
									s.level = br.read_uint8();
			
					break;
								case 2:
									s.bond = br.read_f64();
			
					break;
								case 0:
									s.name = br.read_string();
			
					break;
				
				}
			}
		}
	
	
		export class Loot {
			static get TypeID() { return 983 }

			constructor(fields) {
				if (fields) Object.assign(this, fields)
			}

							baseChance
							modifiers
							entries
			

			deserialize(bytes) {
				const b = new ByteReader(bytes)
				Loot_deserialize(b, this)
				return this;
			}

			serialize() {
				const w = new ByteWriter;
				Loot_serialize(w, this)
				return w.bytes()
			}
		}

				function Loot_serialize(b, s) {
			b.write_uint16(983)
			b.write_uint32(0);
			const structStart = b.length;
						if(s.baseChance !== undefined) {
				b.write_uint16(0);
								b.write_f64(s.baseChance)
			
			}
		  			if(s.modifiers !== undefined) {
				b.write_uint16(1);
								b.write_uint32(0);
				const mapStart0 = b.length;
				for(const [key0, value0] of Object.entries(s.modifiers)) {
									b.write_string(key0)
			
									b.write_f64(value0)
			
				}
				b.set_uint32(mapStart0 - 4, b.length - mapStart0);
			
			}
		  			if(s.entries !== undefined) {
				b.write_uint16(2);
								b.write_uint32(0);
				const listStart0 = b.length;
				for(const item0 of s.entries) {
									LootEntry_serialize(b, item0)
			
				}
				b.set_uint32(listStart0 - 4, b.length - listStart0);
			
			}
		  
			b.set_uint32(structStart - 4, b.length - structStart);
		}
	
				function Loot_deserialize(br, s) {
			const typeId = br.read_uint16()
			if (typeId !== 983) {
				throw new Error("Type ID mismatch deserializing struct Loot: expected 983, got " + typeId);
			}
			const length = br.read_uint32()
			if (length > (br.length - br.position)) {
				throw new Error("Struct Loot length exceeds buffer length");
			}
			const seenFields = new Set;
			const startPos = br.position;
			for (; br.position - startPos < length;) {
				const fieldId = br.read_uint16();
				if (seenFields.has(fieldId)) {
					throw new Error("Duplicate field ID " + fieldId + " in struct Loot");
				}
				if (fieldId > 2) {
					return;
				}
				seenFields.add(fieldId);
				switch (fieldId) {
									case 0:
									s.baseChance = br.read_f64();
			
					break;
								case 1:
								{
				const mapLength0 = br.read_uint32();
				if (mapLength0 > (br.length-br.position)) {
					throw new Error("Invalid map length");
				}
				const mapStart0 = br.position;
				s.modifiers = {};
				for (; br.position - mapStart0 < mapLength0;) {
					let key0;
									key0 = br.read_string();
			
					let value0;
									value0 = br.read_f64();
			
					s.modifiers[key0] = value0;
				}
			}
			
					break;
								case 2:
								{
				const listLength0 = br.read_uint32();
				if (listLength0 > (br.length - br.position)) {
					throw new Error("Invalid list length");
				}
				const listStart0 = br.position;
				s.entries = [];
				for (; br.position - listStart0 < listLength0;) {
					let item0;
									item0 = new LootEntry();
				LootEntry_deserialize(br, item0);
			
					s.entries.push(item0);
				}
			}
			
					break;
				
				}
			}
		}
	
	
		export class World {
			static get TypeID() { return 60723 }

			constructor(fields) {
				if (fields) Object.assign(this, fields)
			}

							lootTables
							worldName
							seed
							gravity
							players
							activeQuests
							zoneData
							systemFlags
			

			deserialize(bytes) {
				const b = new ByteReader(bytes)
				World_deserialize(b, this)
				return this;
			}

			serialize() {
				const w = new ByteWriter;
				World_serialize(w, this)
				return w.bytes()
			}
		}

				function World_serialize(b, s) {
			b.write_uint16(60723)
			b.write_uint32(0);
			const structStart = b.length;
						if(s.lootTables !== undefined) {
				b.write_uint16(7);
								b.write_uint32(0);
				const mapStart0 = b.length;
				for(const [key0, value0] of Object.entries(s.lootTables)) {
									b.write_string(key0)
			
									Loot_serialize(b, value0)
			
				}
				b.set_uint32(mapStart0 - 4, b.length - mapStart0);
			
			}
		  			if(s.worldName !== undefined) {
				b.write_uint16(0);
								b.write_string(s.worldName)
			
			}
		  			if(s.seed !== undefined) {
				b.write_uint16(1);
								b.write_uint64(s.seed)
			
			}
		  			if(s.gravity !== undefined) {
				b.write_uint16(2);
								b.write_f64(s.gravity)
			
			}
		  			if(s.players !== undefined) {
				b.write_uint16(3);
								b.write_uint32(0);
				const listStart0 = b.length;
				for(const item0 of s.players) {
									Character_serialize(b, item0)
			
				}
				b.set_uint32(listStart0 - 4, b.length - listStart0);
			
			}
		  			if(s.activeQuests !== undefined) {
				b.write_uint16(4);
								b.write_uint32(0);
				const listStart0 = b.length;
				for(const item0 of s.activeQuests) {
									Quest_serialize(b, item0)
			
				}
				b.set_uint32(listStart0 - 4, b.length - listStart0);
			
			}
		  			if(s.zoneData !== undefined) {
				b.write_uint16(5);
								b.write_uint32(0);
				const mapStart0 = b.length;
				for(const [key0, value0] of Object.entries(s.zoneData)) {
									b.write_string(key0)
			
									b.write_uint32(0);
				const mapStart1 = b.length;
				for(const [key1, value1] of Object.entries(value0)) {
									b.write_string(key1)
			
									b.write_uint32(value1)
			
				}
				b.set_uint32(mapStart1 - 4, b.length - mapStart1);
			
				}
				b.set_uint32(mapStart0 - 4, b.length - mapStart0);
			
			}
		  			if(s.systemFlags !== undefined) {
				b.write_uint16(6);
								b.write_uint32(0);
				const mapStart0 = b.length;
				for(const [key0, value0] of Object.entries(s.systemFlags)) {
									b.write_string(key0)
			
									b.write_bool(value0)
			
				}
				b.set_uint32(mapStart0 - 4, b.length - mapStart0);
			
			}
		  
			b.set_uint32(structStart - 4, b.length - structStart);
		}
	
				function World_deserialize(br, s) {
			const typeId = br.read_uint16()
			if (typeId !== 60723) {
				throw new Error("Type ID mismatch deserializing struct World: expected 60723, got " + typeId);
			}
			const length = br.read_uint32()
			if (length > (br.length - br.position)) {
				throw new Error("Struct World length exceeds buffer length");
			}
			const seenFields = new Set;
			const startPos = br.position;
			for (; br.position - startPos < length;) {
				const fieldId = br.read_uint16();
				if (seenFields.has(fieldId)) {
					throw new Error("Duplicate field ID " + fieldId + " in struct World");
				}
				if (fieldId > 7) {
					return;
				}
				seenFields.add(fieldId);
				switch (fieldId) {
									case 7:
								{
				const mapLength0 = br.read_uint32();
				if (mapLength0 > (br.length-br.position)) {
					throw new Error("Invalid map length");
				}
				const mapStart0 = br.position;
				s.lootTables = {};
				for (; br.position - mapStart0 < mapLength0;) {
					let key0;
									key0 = br.read_string();
			
					let value0;
									value0 = new Loot();
				Loot_deserialize(br, value0);
			
					s.lootTables[key0] = value0;
				}
			}
			
					break;
								case 0:
									s.worldName = br.read_string();
			
					break;
								case 1:
									s.seed = br.read_uint64();
			
					break;
								case 2:
									s.gravity = br.read_f64();
			
					break;
								case 3:
								{
				const listLength0 = br.read_uint32();
				if (listLength0 > (br.length - br.position)) {
					throw new Error("Invalid list length");
				}
				const listStart0 = br.position;
				s.players = [];
				for (; br.position - listStart0 < listLength0;) {
					let item0;
									item0 = new Character();
				Character_deserialize(br, item0);
			
					s.players.push(item0);
				}
			}
			
					break;
								case 4:
								{
				const listLength0 = br.read_uint32();
				if (listLength0 > (br.length - br.position)) {
					throw new Error("Invalid list length");
				}
				const listStart0 = br.position;
				s.activeQuests = [];
				for (; br.position - listStart0 < listLength0;) {
					let item0;
									item0 = new Quest();
				Quest_deserialize(br, item0);
			
					s.activeQuests.push(item0);
				}
			}
			
					break;
								case 5:
								{
				const mapLength0 = br.read_uint32();
				if (mapLength0 > (br.length-br.position)) {
					throw new Error("Invalid map length");
				}
				const mapStart0 = br.position;
				s.zoneData = {};
				for (; br.position - mapStart0 < mapLength0;) {
					let key0;
									key0 = br.read_string();
			
					let value0;
								{
				const mapLength1 = br.read_uint32();
				if (mapLength1 > (br.length-br.position)) {
					throw new Error("Invalid map length");
				}
				const mapStart1 = br.position;
				value0 = {};
				for (; br.position - mapStart1 < mapLength1;) {
					let key1;
									key1 = br.read_string();
			
					let value1;
									value1 = br.read_uint32();
			
					value0[key1] = value1;
				}
			}
			
					s.zoneData[key0] = value0;
				}
			}
			
					break;
								case 6:
								{
				const mapLength0 = br.read_uint32();
				if (mapLength0 > (br.length-br.position)) {
					throw new Error("Invalid map length");
				}
				const mapStart0 = br.position;
				s.systemFlags = {};
				for (; br.position - mapStart0 < mapLength0;) {
					let key0;
									key0 = br.read_string();
			
					let value0;
									value0 = br.read_bool();
			
					s.systemFlags[key0] = value0;
				}
			}
			
					break;
				
				}
			}
		}
	
	


class ByteWriter {
	get length() { return this.len; }

	encoder = new TextEncoder();
	buffer = new ArrayBuffer(0xFF)
	view = new Uint8Array(this.buffer, 0)
	dview = new DataView(this.buffer, 0)
	len = 0;

	write(value) {
		ByteWriter._tmp = this.length;
		this.resize(this.len + value.length);
		this.view.set(value, ByteWriter._tmp);
	}

	set_uint8(offset, value) {
		this.resize(offset + 1);
		this.dview.setUint8(offset, value, true);
	}

	set_uint16(offset, value) {
		this.resize(offset + 2);
		this.dview.setUint16(offset, value, true);
	}

	set_uint32(offset, value) {
		this.resize(offset + 4);
		this.dview.setUint32(offset, value, true);
	}

	write_bool(value) {
		ByteWriter._tmp = this.length;
		this.resize(this.len + 1);
		this.dview.setUint8(ByteWriter._tmp, value ? 1 : 0, true);
	}

	write_int8(value) {
		ByteWriter._tmp = this.length;
		this.resize(this.len + 1);
		this.dview.setInt8(ByteWriter._tmp, value, true);
	}

	write_uint8(value) {
		ByteWriter._tmp = this.length;
		this.resize(this.len + 1);
		this.dview.setUint8(ByteWriter._tmp, value, true);
	}

	write_int16(value) {
		ByteWriter._tmp = this.length;
		this.resize(this.len + 2);
		this.dview.setInt16(ByteWriter._tmp, value, true);
	}

	write_uint16(value) {
		ByteWriter._tmp = this.length;
		this.resize(this.len + 2);
		this.dview.setUint16(ByteWriter._tmp, value, true);
	}

	write_int32(value) {
		ByteWriter._tmp = this.length;
		this.resize(this.len + 4);
		this.dview.setInt32(ByteWriter._tmp, value, true);
	}

	write_uint32(value) {
		ByteWriter._tmp = this.length;
		this.resize(this.len + 4);
		this.dview.setUint32(ByteWriter._tmp, value, true);
	}

	write_int64(value) {
		ByteWriter._tmp = this.length;
		this.resize(this.len + 8);
		this.dview.setBigInt64(ByteWriter._tmp, BigInt(value), true);
	}

	write_uint64(value) {
		ByteWriter._tmp = this.length;
		this.resize(this.len + 8);
		this.dview.setBigUint64(ByteWriter._tmp, BigInt(value), true);
	}

	write_f32(value) {
		ByteWriter._tmp = this.length;
		this.resize(this.len + 4);
		this.dview.setFloat32(ByteWriter._tmp, value, true);
	}

	write_f64(value) {
		ByteWriter._tmp = this.length;
		this.resize(this.len + 8);
		this.dview.setFloat64(ByteWriter._tmp, value, true);
	}

	write_string(value) {
		const stringLength = value.length;
		if (stringLength > 300) {
			const encoded = this.encoder.encode(value);
			this.set_uint32(this.length, encoded.length);
			this.write(encoded);
			return;
		}
		const lengthPos = this.length;
		this.write_uint32(0, this);
		const start = this.length;
		if (stringLength === 0) {
			return;
		}
		let codePoint;
		for (let i = 0; i < stringLength; i++) {
			// decode UTF-16
			const a = value.charCodeAt(i);
			if (i + 1 === stringLength || a < 0xD800 || a >= 0xDC00) {
				codePoint = a;
			} else {
				const b2 = value.charCodeAt(++i);  // Renamed to avoid shadowing
				codePoint = (a << 10) + b2 + (0x10000 - (0xD800 << 10) - 0xDC00);
			}
			if (codePoint < 0x80) {
				this.write_uint8(codePoint, this);
			} else {
				if (codePoint < 0x800) {
					this.write_uint8(((codePoint >> 6) & 0x1F) | 0xC0, this);
				} else {
					if (codePoint < 0x10000) {
						this.write_uint8(((codePoint >> 12) & 0x0F) | 0xE0, this);
					} else {
						this.write_uint8(((codePoint >> 18) & 0x07) | 0xF0, this);
						this.write_uint8(((codePoint >> 12) & 0x3F) | 0x80, this);
					}
					this.write_uint8(((codePoint >> 6) & 0x3F) | 0x80, this);
				}
				this.write_uint8((codePoint & 0x3F) | 0x80, this);
			}
		}
		this.set_uint32(lengthPos, this.length - start);
	}

	bytes() {
		return new Uint8Array(this.buffer, 0, this.len);
	}

	resize = (length) => {
		if (this.len < length) {
			this.len = length;
			if (this.view.length < length) {
				const newBuffer = new ArrayBuffer(Math.max(this.view.length * 2, length));
				const newView = new Uint8Array(newBuffer);
				newView.set(this.view, 0);
				this.buffer = newBuffer;
				this.view = newView;
				this.dview = new DataView(this.buffer, 0);
			}
		}
	}
}

class ByteReader {
	constructor(buffer) {
  this.buffer = buffer
  this.view = new DataView(
      this.buffer.buffer,
      this.buffer.byteOffset,
      this.buffer.byteLength
  );
  this.position = 0;
  this.length = buffer.length
}

	read_bool() {
		if (this.position + 1 > this.length) {
			throw new Error("Read past end of buffer");
		}
		return this.view.getUint8(this.position++, true) !== 0;
	}

	read_int8() {
		if (this.position + 1 > this.length) {
			throw new Error("Read past end of buffer");
		}
		return this.view.getInt8(this.position++, true);
	}

	read_uint8() {
		if (this.position + 1 > this.length) {
			throw new Error("Read past end of buffer");
		}
		return this.view.getUint8(this.position++, true);
	}

	read_int16() {
		if (this.position + 2 > this.length) {
			throw new Error("Read past end of buffer");
		}
		const value = this.view.getInt16(this.position, true);
		this.position += 2;
		return value;
	}

	read_uint16() {
		if (this.position + 2 > this.length) {
			throw new Error("Read past end of buffer");
		}
		const value = this.view.getUint16(this.position, true);
		this.position += 2;
		return value;
	}

	read_int32() {
		if (this.position + 4 > this.length) {
			throw new Error("Read past end of buffer");
		}
		const value = this.view.getInt32(this.position, true);
		this.position += 4;
		return value;
	}

	read_uint32() {
		if (this.position + 4 > this.length) {
			throw new Error("Read past end of buffer");
		}
		const value = this.view.getUint32(this.position, true);
		this.position += 4;
		return value;
	}

	read_int64() {
		if (this.position + 8 > this.length) {
			throw new Error("Read past end of buffer");
		}
		const value = this.view.getBigInt64(this.position, true);
		this.position += 8;
		return value;
	}

	read_uint64() {
		if (this.position + 8 > this.length) {
			throw new Error("Read past end of buffer");
		}
		const value = this.view.getBigUint64(this.position, true);
		this.position += 8;
		return value;
	}

	read_f32() {
		if (this.position + 4 > this.length) {
			throw new Error("Read past end of buffer");
		}
		const value = this.view.getFloat32(this.position, true);
		this.position += 4;
		return value;
	}

	read_f64() {
		if (this.position + 8 > this.length) {
			throw new Error("Read past end of buffer");
		}
		const value = this.view.getFloat64(this.position, true);
		this.position += 8;
		return value;
	}

	read_string() {
		const length = this.read_uint32();
		if (length === 0) return ""

		if (length > (this.length - this.position)) {
			throw new Error("String is longer than remaining buffer");
		}

		if (length > 300) {
			const encoded = this.buffer.slice(this.position, this.position + length);
			this.position += length;
			const decoder = new TextDecoder();
			return decoder.decode(encoded);
		} else {
			const end = this.position + length;
			if (end > this.length) {
				throw new Error("Read past end of buffer");
			}
			let result = "";
			let codePoint;
			while (this.position < end) {
				const a = this.read_uint8();
				if (a < 0xC0) {
					codePoint = a;
				} else {
					const b = this.read_uint8();
					if (a < 0xE0) {
						codePoint = ((a & 0x1F) << 6) | (b & 0x3F);
					} else {
						const c = this.read_uint8();
						if (a < 0xF0) {
							codePoint = ((a & 0x0F) << 12) | ((b & 0x3F) << 6) | (c & 0x3F);
						} else {
							const d = this.read_uint8();
							codePoint = ((a & 0x07) << 18) | ((b & 0x3F) << 12) | ((c & 0x3F) << 6) | (d & 0x3F);
						}
					}
				}
				if (codePoint < 0x10000) {
					result += String.fromCharCode(codePoint);
				} else {
					codePoint -= 0x10000;
					result += String.fromCharCode((codePoint >> 10) + 0xD800, (codePoint & ((1 << 10) - 1)) + 0xDC00);
				}
			}
			this.position = end;
			return result;
		}
	}
}