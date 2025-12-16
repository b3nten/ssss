// Auto-generated code for schema: Schema v1

namespace Schema;

interface ISchema<TSelf> where TSelf : ISchema<TSelf>
{
	public byte[] Serialize();
	public TSelf Deserialize(byte[] data);
}		class Loot : ISchema<Loot>
		{
			public Dictionary<string, double>? Modifiers;
public List<LootEntry>? Entries;
public double? BaseChance;

			public readonly static ushort TypeId = 983;

			public static Loot CreateFromBytes(byte[] data)
			{
				Loot it = new Loot();
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_Loot.Deserialize(it, r);
				}
				return it;
			}

			public Loot Deserialize(byte[] data)
			{
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_Loot.Deserialize(this, r);
				}
				return this;
			}

			public byte[] Serialize()
			{
				using (MemoryStream ms = new MemoryStream())
				using (BinaryWriter w = new BinaryWriter(ms))
				{
					_Loot.Serialize(this, w);
					return ms.ToArray();
				}
			}
		}

		file class _Loot
		{
					static public void Serialize(Loot it, BinaryWriter w)
		{
			w.Write(Loot.TypeId);
var lengthPos = w.BaseStream.Position;
w.Write((UInt32)0);
			if (it.Modifiers != null)
			{
				w.Write((ushort)1);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			foreach (var kv0 in it.Modifiers)
			{
								var bytes1 = System.Text.Encoding.UTF8.GetBytes(kv0.Key);
				w.Write((uint)bytes1.Length);
				w.Write(bytes1);
			
				w.Write(kv0.Value);
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.Entries != null)
			{
				w.Write((ushort)2);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			for (int i0 = 0; i0 < it.Entries.Count; i0++)
			{
				var e0 = it.Entries[i0];
				_LootEntry.Serialize(e0, w);
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.BaseChance != null)
			{
				w.Write((ushort)0);
				w.Write(it.BaseChance.Value);
			}
var endPos = w.BaseStream.Position;
			w.Seek((int)lengthPos, SeekOrigin.Begin);
			w.Write((UInt32)(endPos - lengthPos - 4));
			w.Seek(0, SeekOrigin.End);
		}
	
					static public void Deserialize(Loot it, BinaryReader r)
		{
			ushort typeId = r.ReadUInt16();
			if (typeId != Loot.TypeId)
			{
				throw new Exception($"TypeId mismatch: expected Loot.TypeId but got {typeId}");
			}
			uint length = r.ReadUInt32();
			long startPos = r.BaseStream.Position;
			while (r.BaseStream.Position - startPos < length)
			{
				ushort fieldId = r.ReadUInt16();
				switch (fieldId)
				{
							case 1:
							{
				uint mapLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var map0 = new System.Collections.Generic.Dictionary<string, double>();
				while (r.BaseStream.Position - startPos0 < mapLength0)
				{
					string key0;
					double value0;
									{
					uint strLen2 = r.ReadUInt32();
					var strBytes2 = r.ReadBytes((int)strLen2);
					key0 = System.Text.Encoding.UTF8.GetString(strBytes2);
				}
			
					value0 = r.ReadDouble();
					map0.Add(key0, value0);
				}
				it.Modifiers = map0;
			}
		
				break;
					case 2:
							{
				uint listLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var list0 = new System.Collections.Generic.List<LootEntry>();
				while (r.BaseStream.Position - startPos0 < listLength0)
				{
					LootEntry e0;
								{
				LootEntry obj = new();
				_LootEntry.Deserialize(obj, r);
				e0 = obj;
			}
		
					list0.Add(e0);
				}
				it.Entries = list0;
			}
		
				break;
					case 0:
				it.BaseChance = r.ReadDouble();
				break;
		
				default:
					r.BaseStream.Seek(startPos + length, SeekOrigin.Begin);
					return;
				}
			}
		}
	
		}
			class LootEntry : ISchema<LootEntry>
		{
			public byte? MinQty;
public byte? MaxQty;
public Dictionary<string, bool>? Conditions;
public uint? ItemId;

			public readonly static ushort TypeId = 305;

			public static LootEntry CreateFromBytes(byte[] data)
			{
				LootEntry it = new LootEntry();
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_LootEntry.Deserialize(it, r);
				}
				return it;
			}

			public LootEntry Deserialize(byte[] data)
			{
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_LootEntry.Deserialize(this, r);
				}
				return this;
			}

			public byte[] Serialize()
			{
				using (MemoryStream ms = new MemoryStream())
				using (BinaryWriter w = new BinaryWriter(ms))
				{
					_LootEntry.Serialize(this, w);
					return ms.ToArray();
				}
			}
		}

		file class _LootEntry
		{
					static public void Serialize(LootEntry it, BinaryWriter w)
		{
			w.Write(LootEntry.TypeId);
var lengthPos = w.BaseStream.Position;
w.Write((UInt32)0);
			if (it.MinQty != null)
			{
				w.Write((ushort)1);
				w.Write(it.MinQty.Value);
			}			if (it.MaxQty != null)
			{
				w.Write((ushort)2);
				w.Write(it.MaxQty.Value);
			}			if (it.Conditions != null)
			{
				w.Write((ushort)3);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			foreach (var kv0 in it.Conditions)
			{
								var bytes3 = System.Text.Encoding.UTF8.GetBytes(kv0.Key);
				w.Write((uint)bytes3.Length);
				w.Write(bytes3);
			
				w.Write(kv0.Value);
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.ItemId != null)
			{
				w.Write((ushort)0);
				w.Write(it.ItemId.Value);
			}
var endPos = w.BaseStream.Position;
			w.Seek((int)lengthPos, SeekOrigin.Begin);
			w.Write((UInt32)(endPos - lengthPos - 4));
			w.Seek(0, SeekOrigin.End);
		}
	
					static public void Deserialize(LootEntry it, BinaryReader r)
		{
			ushort typeId = r.ReadUInt16();
			if (typeId != LootEntry.TypeId)
			{
				throw new Exception($"TypeId mismatch: expected LootEntry.TypeId but got {typeId}");
			}
			uint length = r.ReadUInt32();
			long startPos = r.BaseStream.Position;
			while (r.BaseStream.Position - startPos < length)
			{
				ushort fieldId = r.ReadUInt16();
				switch (fieldId)
				{
							case 1:
				it.MinQty = r.ReadByte();
				break;
					case 2:
				it.MaxQty = r.ReadByte();
				break;
					case 3:
							{
				uint mapLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var map0 = new System.Collections.Generic.Dictionary<string, bool>();
				while (r.BaseStream.Position - startPos0 < mapLength0)
				{
					string key0;
					bool value0;
									{
					uint strLen4 = r.ReadUInt32();
					var strBytes4 = r.ReadBytes((int)strLen4);
					key0 = System.Text.Encoding.UTF8.GetString(strBytes4);
				}
			
					value0 = r.ReadBoolean();
					map0.Add(key0, value0);
				}
				it.Conditions = map0;
			}
		
				break;
					case 0:
				it.ItemId = r.ReadUInt32();
				break;
		
				default:
					r.BaseStream.Seek(startPos + length, SeekOrigin.Begin);
					return;
				}
			}
		}
	
		}
			class Quest : ISchema<Quest>
		{
			public string? Title;
public byte? Difficulty;
public List<Vector3>? RequiredPos;
public Dictionary<string, List<string>>? Objectives;
public List<Quest>? Prerequisites;
public Dictionary<string, Quest>? NextSteps;
public List<List<List<ushort>>>? AreaLayers;
public uint? Id;
public string? Description;
public List<Item>? Rewards;

			public readonly static ushort TypeId = 16605;

			public static Quest CreateFromBytes(byte[] data)
			{
				Quest it = new Quest();
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_Quest.Deserialize(it, r);
				}
				return it;
			}

			public Quest Deserialize(byte[] data)
			{
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_Quest.Deserialize(this, r);
				}
				return this;
			}

			public byte[] Serialize()
			{
				using (MemoryStream ms = new MemoryStream())
				using (BinaryWriter w = new BinaryWriter(ms))
				{
					_Quest.Serialize(this, w);
					return ms.ToArray();
				}
			}
		}

		file class _Quest
		{
					static public void Serialize(Quest it, BinaryWriter w)
		{
			w.Write(Quest.TypeId);
var lengthPos = w.BaseStream.Position;
w.Write((UInt32)0);
			if (it.Title != null)
			{
				w.Write((ushort)1);
								var bytes5 = System.Text.Encoding.UTF8.GetBytes(it.Title);
				w.Write((uint)bytes5.Length);
				w.Write(bytes5);
			
			}			if (it.Difficulty != null)
			{
				w.Write((ushort)2);
				w.Write(it.Difficulty.Value);
			}			if (it.RequiredPos != null)
			{
				w.Write((ushort)4);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			for (int i0 = 0; i0 < it.RequiredPos.Count; i0++)
			{
				var e0 = it.RequiredPos[i0];
				_Vector3.Serialize(e0, w);
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.Objectives != null)
			{
				w.Write((ushort)5);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			foreach (var kv0 in it.Objectives)
			{
								var bytes6 = System.Text.Encoding.UTF8.GetBytes(kv0.Key);
				w.Write((uint)bytes6.Length);
				w.Write(bytes6);
			
							var length1 = w.BaseStream.Position;
			w.Write((uint)0);
			for (int i1 = 0; i1 < kv0.Value.Count; i1++)
			{
				var e1 = kv0.Value[i1];
								var bytes7 = System.Text.Encoding.UTF8.GetBytes(e1);
				w.Write((uint)bytes7.Length);
				w.Write(bytes7);
			
			}
			var end1 = w.BaseStream.Position;
			w.Seek((int)length1, SeekOrigin.Begin);
			w.Write((uint)(end1 - length1 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.Prerequisites != null)
			{
				w.Write((ushort)6);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			for (int i0 = 0; i0 < it.Prerequisites.Count; i0++)
			{
				var e0 = it.Prerequisites[i0];
				_Quest.Serialize(e0, w);
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.NextSteps != null)
			{
				w.Write((ushort)7);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			foreach (var kv0 in it.NextSteps)
			{
								var bytes8 = System.Text.Encoding.UTF8.GetBytes(kv0.Key);
				w.Write((uint)bytes8.Length);
				w.Write(bytes8);
			
				_Quest.Serialize(kv0.Value, w);
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.AreaLayers != null)
			{
				w.Write((ushort)8);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			for (int i0 = 0; i0 < it.AreaLayers.Count; i0++)
			{
				var e0 = it.AreaLayers[i0];
							var length1 = w.BaseStream.Position;
			w.Write((uint)0);
			for (int i1 = 0; i1 < e0.Count; i1++)
			{
				var e1 = e0[i1];
							var length2 = w.BaseStream.Position;
			w.Write((uint)0);
			for (int i2 = 0; i2 < e1.Count; i2++)
			{
				var e2 = e1[i2];
				w.Write(e2);
			}
			var end2 = w.BaseStream.Position;
			w.Seek((int)length2, SeekOrigin.Begin);
			w.Write((uint)(end2 - length2 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}
			var end1 = w.BaseStream.Position;
			w.Seek((int)length1, SeekOrigin.Begin);
			w.Write((uint)(end1 - length1 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.Id != null)
			{
				w.Write((ushort)0);
				w.Write(it.Id.Value);
			}			if (it.Description != null)
			{
				w.Write((ushort)9);
								var bytes9 = System.Text.Encoding.UTF8.GetBytes(it.Description);
				w.Write((uint)bytes9.Length);
				w.Write(bytes9);
			
			}			if (it.Rewards != null)
			{
				w.Write((ushort)3);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			for (int i0 = 0; i0 < it.Rewards.Count; i0++)
			{
				var e0 = it.Rewards[i0];
				_Item.Serialize(e0, w);
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}
var endPos = w.BaseStream.Position;
			w.Seek((int)lengthPos, SeekOrigin.Begin);
			w.Write((UInt32)(endPos - lengthPos - 4));
			w.Seek(0, SeekOrigin.End);
		}
	
					static public void Deserialize(Quest it, BinaryReader r)
		{
			ushort typeId = r.ReadUInt16();
			if (typeId != Quest.TypeId)
			{
				throw new Exception($"TypeId mismatch: expected Quest.TypeId but got {typeId}");
			}
			uint length = r.ReadUInt32();
			long startPos = r.BaseStream.Position;
			while (r.BaseStream.Position - startPos < length)
			{
				ushort fieldId = r.ReadUInt16();
				switch (fieldId)
				{
							case 1:
								{
					uint strLen10 = r.ReadUInt32();
					var strBytes10 = r.ReadBytes((int)strLen10);
					it.Title = System.Text.Encoding.UTF8.GetString(strBytes10);
				}
			
				break;
					case 2:
				it.Difficulty = r.ReadByte();
				break;
					case 4:
							{
				uint listLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var list0 = new System.Collections.Generic.List<Vector3>();
				while (r.BaseStream.Position - startPos0 < listLength0)
				{
					Vector3 e0;
								{
				Vector3 obj = new();
				_Vector3.Deserialize(obj, r);
				e0 = obj;
			}
		
					list0.Add(e0);
				}
				it.RequiredPos = list0;
			}
		
				break;
					case 5:
							{
				uint mapLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var map0 = new System.Collections.Generic.Dictionary<string, List<string>>();
				while (r.BaseStream.Position - startPos0 < mapLength0)
				{
					string key0;
					List<string> value0;
									{
					uint strLen11 = r.ReadUInt32();
					var strBytes11 = r.ReadBytes((int)strLen11);
					key0 = System.Text.Encoding.UTF8.GetString(strBytes11);
				}
			
								{
				uint listLength1 = r.ReadUInt32();
				long startPos1 = r.BaseStream.Position;
				var list1 = new System.Collections.Generic.List<string>();
				while (r.BaseStream.Position - startPos1 < listLength1)
				{
					string e1;
									{
					uint strLen12 = r.ReadUInt32();
					var strBytes12 = r.ReadBytes((int)strLen12);
					e1 = System.Text.Encoding.UTF8.GetString(strBytes12);
				}
			
					list1.Add(e1);
				}
				value0 = list1;
			}
		
					map0.Add(key0, value0);
				}
				it.Objectives = map0;
			}
		
				break;
					case 6:
							{
				uint listLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var list0 = new System.Collections.Generic.List<Quest>();
				while (r.BaseStream.Position - startPos0 < listLength0)
				{
					Quest e0;
								{
				Quest obj = new();
				_Quest.Deserialize(obj, r);
				e0 = obj;
			}
		
					list0.Add(e0);
				}
				it.Prerequisites = list0;
			}
		
				break;
					case 7:
							{
				uint mapLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var map0 = new System.Collections.Generic.Dictionary<string, Quest>();
				while (r.BaseStream.Position - startPos0 < mapLength0)
				{
					string key0;
					Quest value0;
									{
					uint strLen13 = r.ReadUInt32();
					var strBytes13 = r.ReadBytes((int)strLen13);
					key0 = System.Text.Encoding.UTF8.GetString(strBytes13);
				}
			
								{
				Quest obj = new();
				_Quest.Deserialize(obj, r);
				value0 = obj;
			}
		
					map0.Add(key0, value0);
				}
				it.NextSteps = map0;
			}
		
				break;
					case 8:
							{
				uint listLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var list0 = new System.Collections.Generic.List<List<List<ushort>>>();
				while (r.BaseStream.Position - startPos0 < listLength0)
				{
					List<List<ushort>> e0;
								{
				uint listLength1 = r.ReadUInt32();
				long startPos1 = r.BaseStream.Position;
				var list1 = new System.Collections.Generic.List<List<ushort>>();
				while (r.BaseStream.Position - startPos1 < listLength1)
				{
					List<ushort> e1;
								{
				uint listLength2 = r.ReadUInt32();
				long startPos2 = r.BaseStream.Position;
				var list2 = new System.Collections.Generic.List<ushort>();
				while (r.BaseStream.Position - startPos2 < listLength2)
				{
					ushort e2;
					e2 = r.ReadUInt16();
					list2.Add(e2);
				}
				e1 = list2;
			}
		
					list1.Add(e1);
				}
				e0 = list1;
			}
		
					list0.Add(e0);
				}
				it.AreaLayers = list0;
			}
		
				break;
					case 0:
				it.Id = r.ReadUInt32();
				break;
					case 9:
								{
					uint strLen14 = r.ReadUInt32();
					var strBytes14 = r.ReadBytes((int)strLen14);
					it.Description = System.Text.Encoding.UTF8.GetString(strBytes14);
				}
			
				break;
					case 3:
							{
				uint listLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var list0 = new System.Collections.Generic.List<Item>();
				while (r.BaseStream.Position - startPos0 < listLength0)
				{
					Item e0;
								{
				Item obj = new();
				_Item.Deserialize(obj, r);
				e0 = obj;
			}
		
					list0.Add(e0);
				}
				it.Rewards = list0;
			}
		
				break;
		
				default:
					r.BaseStream.Seek(startPos + length, SeekOrigin.Begin);
					return;
				}
			}
		}
	
		}
			class EquipmentSlot : ISchema<EquipmentSlot>
		{
			public string? SlotName;
public Item? Item;

			public readonly static ushort TypeId = 35339;

			public static EquipmentSlot CreateFromBytes(byte[] data)
			{
				EquipmentSlot it = new EquipmentSlot();
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_EquipmentSlot.Deserialize(it, r);
				}
				return it;
			}

			public EquipmentSlot Deserialize(byte[] data)
			{
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_EquipmentSlot.Deserialize(this, r);
				}
				return this;
			}

			public byte[] Serialize()
			{
				using (MemoryStream ms = new MemoryStream())
				using (BinaryWriter w = new BinaryWriter(ms))
				{
					_EquipmentSlot.Serialize(this, w);
					return ms.ToArray();
				}
			}
		}

		file class _EquipmentSlot
		{
					static public void Serialize(EquipmentSlot it, BinaryWriter w)
		{
			w.Write(EquipmentSlot.TypeId);
var lengthPos = w.BaseStream.Position;
w.Write((UInt32)0);
			if (it.SlotName != null)
			{
				w.Write((ushort)0);
								var bytes15 = System.Text.Encoding.UTF8.GetBytes(it.SlotName);
				w.Write((uint)bytes15.Length);
				w.Write(bytes15);
			
			}			if (it.Item != null)
			{
				w.Write((ushort)1);
				_Item.Serialize(it.Item, w);
			}
var endPos = w.BaseStream.Position;
			w.Seek((int)lengthPos, SeekOrigin.Begin);
			w.Write((UInt32)(endPos - lengthPos - 4));
			w.Seek(0, SeekOrigin.End);
		}
	
					static public void Deserialize(EquipmentSlot it, BinaryReader r)
		{
			ushort typeId = r.ReadUInt16();
			if (typeId != EquipmentSlot.TypeId)
			{
				throw new Exception($"TypeId mismatch: expected EquipmentSlot.TypeId but got {typeId}");
			}
			uint length = r.ReadUInt32();
			long startPos = r.BaseStream.Position;
			while (r.BaseStream.Position - startPos < length)
			{
				ushort fieldId = r.ReadUInt16();
				switch (fieldId)
				{
							case 0:
								{
					uint strLen16 = r.ReadUInt32();
					var strBytes16 = r.ReadBytes((int)strLen16);
					it.SlotName = System.Text.Encoding.UTF8.GetString(strBytes16);
				}
			
				break;
					case 1:
							{
				Item obj = new();
				_Item.Deserialize(obj, r);
				it.Item = obj;
			}
		
				break;
		
				default:
					r.BaseStream.Seek(startPos + length, SeekOrigin.Begin);
					return;
				}
			}
		}
	
		}
			class Companion : ISchema<Companion>
		{
			public string? Name;
public byte? Level;
public double? Bond;

			public readonly static ushort TypeId = 21813;

			public static Companion CreateFromBytes(byte[] data)
			{
				Companion it = new Companion();
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_Companion.Deserialize(it, r);
				}
				return it;
			}

			public Companion Deserialize(byte[] data)
			{
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_Companion.Deserialize(this, r);
				}
				return this;
			}

			public byte[] Serialize()
			{
				using (MemoryStream ms = new MemoryStream())
				using (BinaryWriter w = new BinaryWriter(ms))
				{
					_Companion.Serialize(this, w);
					return ms.ToArray();
				}
			}
		}

		file class _Companion
		{
					static public void Serialize(Companion it, BinaryWriter w)
		{
			w.Write(Companion.TypeId);
var lengthPos = w.BaseStream.Position;
w.Write((UInt32)0);
			if (it.Name != null)
			{
				w.Write((ushort)0);
								var bytes17 = System.Text.Encoding.UTF8.GetBytes(it.Name);
				w.Write((uint)bytes17.Length);
				w.Write(bytes17);
			
			}			if (it.Level != null)
			{
				w.Write((ushort)1);
				w.Write(it.Level.Value);
			}			if (it.Bond != null)
			{
				w.Write((ushort)2);
				w.Write(it.Bond.Value);
			}
var endPos = w.BaseStream.Position;
			w.Seek((int)lengthPos, SeekOrigin.Begin);
			w.Write((UInt32)(endPos - lengthPos - 4));
			w.Seek(0, SeekOrigin.End);
		}
	
					static public void Deserialize(Companion it, BinaryReader r)
		{
			ushort typeId = r.ReadUInt16();
			if (typeId != Companion.TypeId)
			{
				throw new Exception($"TypeId mismatch: expected Companion.TypeId but got {typeId}");
			}
			uint length = r.ReadUInt32();
			long startPos = r.BaseStream.Position;
			while (r.BaseStream.Position - startPos < length)
			{
				ushort fieldId = r.ReadUInt16();
				switch (fieldId)
				{
							case 0:
								{
					uint strLen18 = r.ReadUInt32();
					var strBytes18 = r.ReadBytes((int)strLen18);
					it.Name = System.Text.Encoding.UTF8.GetString(strBytes18);
				}
			
				break;
					case 1:
				it.Level = r.ReadByte();
				break;
					case 2:
				it.Bond = r.ReadDouble();
				break;
		
				default:
					r.BaseStream.Seek(startPos + length, SeekOrigin.Begin);
					return;
				}
			}
		}
	
		}
			class World : ISchema<World>
		{
			public Dictionary<string, bool>? SystemFlags;
public Dictionary<string, Loot>? LootTables;
public string? WorldName;
public ulong? Seed;
public double? Gravity;
public List<Character>? Players;
public List<Quest>? ActiveQuests;
public Dictionary<string, Dictionary<string, uint>>? ZoneData;

			public readonly static ushort TypeId = 60723;

			public static World CreateFromBytes(byte[] data)
			{
				World it = new World();
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_World.Deserialize(it, r);
				}
				return it;
			}

			public World Deserialize(byte[] data)
			{
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_World.Deserialize(this, r);
				}
				return this;
			}

			public byte[] Serialize()
			{
				using (MemoryStream ms = new MemoryStream())
				using (BinaryWriter w = new BinaryWriter(ms))
				{
					_World.Serialize(this, w);
					return ms.ToArray();
				}
			}
		}

		file class _World
		{
					static public void Serialize(World it, BinaryWriter w)
		{
			w.Write(World.TypeId);
var lengthPos = w.BaseStream.Position;
w.Write((UInt32)0);
			if (it.SystemFlags != null)
			{
				w.Write((ushort)6);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			foreach (var kv0 in it.SystemFlags)
			{
								var bytes19 = System.Text.Encoding.UTF8.GetBytes(kv0.Key);
				w.Write((uint)bytes19.Length);
				w.Write(bytes19);
			
				w.Write(kv0.Value);
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.LootTables != null)
			{
				w.Write((ushort)7);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			foreach (var kv0 in it.LootTables)
			{
								var bytes20 = System.Text.Encoding.UTF8.GetBytes(kv0.Key);
				w.Write((uint)bytes20.Length);
				w.Write(bytes20);
			
				_Loot.Serialize(kv0.Value, w);
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.WorldName != null)
			{
				w.Write((ushort)0);
								var bytes21 = System.Text.Encoding.UTF8.GetBytes(it.WorldName);
				w.Write((uint)bytes21.Length);
				w.Write(bytes21);
			
			}			if (it.Seed != null)
			{
				w.Write((ushort)1);
				w.Write(it.Seed.Value);
			}			if (it.Gravity != null)
			{
				w.Write((ushort)2);
				w.Write(it.Gravity.Value);
			}			if (it.Players != null)
			{
				w.Write((ushort)3);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			for (int i0 = 0; i0 < it.Players.Count; i0++)
			{
				var e0 = it.Players[i0];
				_Character.Serialize(e0, w);
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.ActiveQuests != null)
			{
				w.Write((ushort)4);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			for (int i0 = 0; i0 < it.ActiveQuests.Count; i0++)
			{
				var e0 = it.ActiveQuests[i0];
				_Quest.Serialize(e0, w);
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.ZoneData != null)
			{
				w.Write((ushort)5);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			foreach (var kv0 in it.ZoneData)
			{
								var bytes22 = System.Text.Encoding.UTF8.GetBytes(kv0.Key);
				w.Write((uint)bytes22.Length);
				w.Write(bytes22);
			
							var length1 = w.BaseStream.Position;
			w.Write((uint)0);
			foreach (var kv1 in kv0.Value)
			{
								var bytes23 = System.Text.Encoding.UTF8.GetBytes(kv1.Key);
				w.Write((uint)bytes23.Length);
				w.Write(bytes23);
			
				w.Write(kv1.Value);
			}
			var end1 = w.BaseStream.Position;
			w.Seek((int)length1, SeekOrigin.Begin);
			w.Write((uint)(end1 - length1 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}
var endPos = w.BaseStream.Position;
			w.Seek((int)lengthPos, SeekOrigin.Begin);
			w.Write((UInt32)(endPos - lengthPos - 4));
			w.Seek(0, SeekOrigin.End);
		}
	
					static public void Deserialize(World it, BinaryReader r)
		{
			ushort typeId = r.ReadUInt16();
			if (typeId != World.TypeId)
			{
				throw new Exception($"TypeId mismatch: expected World.TypeId but got {typeId}");
			}
			uint length = r.ReadUInt32();
			long startPos = r.BaseStream.Position;
			while (r.BaseStream.Position - startPos < length)
			{
				ushort fieldId = r.ReadUInt16();
				switch (fieldId)
				{
							case 6:
							{
				uint mapLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var map0 = new System.Collections.Generic.Dictionary<string, bool>();
				while (r.BaseStream.Position - startPos0 < mapLength0)
				{
					string key0;
					bool value0;
									{
					uint strLen24 = r.ReadUInt32();
					var strBytes24 = r.ReadBytes((int)strLen24);
					key0 = System.Text.Encoding.UTF8.GetString(strBytes24);
				}
			
					value0 = r.ReadBoolean();
					map0.Add(key0, value0);
				}
				it.SystemFlags = map0;
			}
		
				break;
					case 7:
							{
				uint mapLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var map0 = new System.Collections.Generic.Dictionary<string, Loot>();
				while (r.BaseStream.Position - startPos0 < mapLength0)
				{
					string key0;
					Loot value0;
									{
					uint strLen25 = r.ReadUInt32();
					var strBytes25 = r.ReadBytes((int)strLen25);
					key0 = System.Text.Encoding.UTF8.GetString(strBytes25);
				}
			
								{
				Loot obj = new();
				_Loot.Deserialize(obj, r);
				value0 = obj;
			}
		
					map0.Add(key0, value0);
				}
				it.LootTables = map0;
			}
		
				break;
					case 0:
								{
					uint strLen26 = r.ReadUInt32();
					var strBytes26 = r.ReadBytes((int)strLen26);
					it.WorldName = System.Text.Encoding.UTF8.GetString(strBytes26);
				}
			
				break;
					case 1:
				it.Seed = r.ReadUInt64();
				break;
					case 2:
				it.Gravity = r.ReadDouble();
				break;
					case 3:
							{
				uint listLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var list0 = new System.Collections.Generic.List<Character>();
				while (r.BaseStream.Position - startPos0 < listLength0)
				{
					Character e0;
								{
				Character obj = new();
				_Character.Deserialize(obj, r);
				e0 = obj;
			}
		
					list0.Add(e0);
				}
				it.Players = list0;
			}
		
				break;
					case 4:
							{
				uint listLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var list0 = new System.Collections.Generic.List<Quest>();
				while (r.BaseStream.Position - startPos0 < listLength0)
				{
					Quest e0;
								{
				Quest obj = new();
				_Quest.Deserialize(obj, r);
				e0 = obj;
			}
		
					list0.Add(e0);
				}
				it.ActiveQuests = list0;
			}
		
				break;
					case 5:
							{
				uint mapLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var map0 = new System.Collections.Generic.Dictionary<string, Dictionary<string, uint>>();
				while (r.BaseStream.Position - startPos0 < mapLength0)
				{
					string key0;
					Dictionary<string, uint> value0;
									{
					uint strLen27 = r.ReadUInt32();
					var strBytes27 = r.ReadBytes((int)strLen27);
					key0 = System.Text.Encoding.UTF8.GetString(strBytes27);
				}
			
								{
				uint mapLength1 = r.ReadUInt32();
				long startPos1 = r.BaseStream.Position;
				var map1 = new System.Collections.Generic.Dictionary<string, uint>();
				while (r.BaseStream.Position - startPos1 < mapLength1)
				{
					string key1;
					uint value1;
									{
					uint strLen28 = r.ReadUInt32();
					var strBytes28 = r.ReadBytes((int)strLen28);
					key1 = System.Text.Encoding.UTF8.GetString(strBytes28);
				}
			
					value1 = r.ReadUInt32();
					map1.Add(key1, value1);
				}
				value0 = map1;
			}
		
					map0.Add(key0, value0);
				}
				it.ZoneData = map0;
			}
		
				break;
		
				default:
					r.BaseStream.Seek(startPos + length, SeekOrigin.Begin);
					return;
				}
			}
		}
	
		}
			class Character : ISchema<Character>
		{
			public Dictionary<string, Dictionary<string, Dictionary<string, sbyte>>>? ArbitraryData;
public Dictionary<string, List<double>>? SkillProgress;
public Vector3? Position;
public List<List<Item>>? Inventory;
public Dictionary<string, EquipmentSlot>? Equipment;
public Dictionary<ushort, Companion>? Companions;
public string? Name;
public ulong? Id;
public Stats? Stats;
public List<Character>? Friends;

			public readonly static ushort TypeId = 16560;

			public static Character CreateFromBytes(byte[] data)
			{
				Character it = new Character();
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_Character.Deserialize(it, r);
				}
				return it;
			}

			public Character Deserialize(byte[] data)
			{
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_Character.Deserialize(this, r);
				}
				return this;
			}

			public byte[] Serialize()
			{
				using (MemoryStream ms = new MemoryStream())
				using (BinaryWriter w = new BinaryWriter(ms))
				{
					_Character.Serialize(this, w);
					return ms.ToArray();
				}
			}
		}

		file class _Character
		{
					static public void Serialize(Character it, BinaryWriter w)
		{
			w.Write(Character.TypeId);
var lengthPos = w.BaseStream.Position;
w.Write((UInt32)0);
			if (it.ArbitraryData != null)
			{
				w.Write((ushort)9);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			foreach (var kv0 in it.ArbitraryData)
			{
								var bytes29 = System.Text.Encoding.UTF8.GetBytes(kv0.Key);
				w.Write((uint)bytes29.Length);
				w.Write(bytes29);
			
							var length1 = w.BaseStream.Position;
			w.Write((uint)0);
			foreach (var kv1 in kv0.Value)
			{
								var bytes30 = System.Text.Encoding.UTF8.GetBytes(kv1.Key);
				w.Write((uint)bytes30.Length);
				w.Write(bytes30);
			
							var length2 = w.BaseStream.Position;
			w.Write((uint)0);
			foreach (var kv2 in kv1.Value)
			{
								var bytes31 = System.Text.Encoding.UTF8.GetBytes(kv2.Key);
				w.Write((uint)bytes31.Length);
				w.Write(bytes31);
			
				w.Write(kv2.Value);
			}
			var end2 = w.BaseStream.Position;
			w.Seek((int)length2, SeekOrigin.Begin);
			w.Write((uint)(end2 - length2 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}
			var end1 = w.BaseStream.Position;
			w.Seek((int)length1, SeekOrigin.Begin);
			w.Write((uint)(end1 - length1 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.SkillProgress != null)
			{
				w.Write((ushort)8);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			foreach (var kv0 in it.SkillProgress)
			{
								var bytes32 = System.Text.Encoding.UTF8.GetBytes(kv0.Key);
				w.Write((uint)bytes32.Length);
				w.Write(bytes32);
			
							var length1 = w.BaseStream.Position;
			w.Write((uint)0);
			for (int i1 = 0; i1 < kv0.Value.Count; i1++)
			{
				var e1 = kv0.Value[i1];
				w.Write(e1);
			}
			var end1 = w.BaseStream.Position;
			w.Seek((int)length1, SeekOrigin.Begin);
			w.Write((uint)(end1 - length1 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.Position != null)
			{
				w.Write((ushort)2);
				_Vector3.Serialize(it.Position, w);
			}			if (it.Inventory != null)
			{
				w.Write((ushort)4);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			for (int i0 = 0; i0 < it.Inventory.Count; i0++)
			{
				var e0 = it.Inventory[i0];
							var length1 = w.BaseStream.Position;
			w.Write((uint)0);
			for (int i1 = 0; i1 < e0.Count; i1++)
			{
				var e1 = e0[i1];
				_Item.Serialize(e1, w);
			}
			var end1 = w.BaseStream.Position;
			w.Seek((int)length1, SeekOrigin.Begin);
			w.Write((uint)(end1 - length1 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.Equipment != null)
			{
				w.Write((ushort)5);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			foreach (var kv0 in it.Equipment)
			{
								var bytes33 = System.Text.Encoding.UTF8.GetBytes(kv0.Key);
				w.Write((uint)bytes33.Length);
				w.Write(bytes33);
			
				_EquipmentSlot.Serialize(kv0.Value, w);
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.Companions != null)
			{
				w.Write((ushort)6);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			foreach (var kv0 in it.Companions)
			{
				w.Write(kv0.Key);
				_Companion.Serialize(kv0.Value, w);
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.Name != null)
			{
				w.Write((ushort)1);
								var bytes34 = System.Text.Encoding.UTF8.GetBytes(it.Name);
				w.Write((uint)bytes34.Length);
				w.Write(bytes34);
			
			}			if (it.Id != null)
			{
				w.Write((ushort)0);
				w.Write(it.Id.Value);
			}			if (it.Stats != null)
			{
				w.Write((ushort)3);
				_Stats.Serialize(it.Stats, w);
			}			if (it.Friends != null)
			{
				w.Write((ushort)7);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			for (int i0 = 0; i0 < it.Friends.Count; i0++)
			{
				var e0 = it.Friends[i0];
				_Character.Serialize(e0, w);
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}
var endPos = w.BaseStream.Position;
			w.Seek((int)lengthPos, SeekOrigin.Begin);
			w.Write((UInt32)(endPos - lengthPos - 4));
			w.Seek(0, SeekOrigin.End);
		}
	
					static public void Deserialize(Character it, BinaryReader r)
		{
			ushort typeId = r.ReadUInt16();
			if (typeId != Character.TypeId)
			{
				throw new Exception($"TypeId mismatch: expected Character.TypeId but got {typeId}");
			}
			uint length = r.ReadUInt32();
			long startPos = r.BaseStream.Position;
			while (r.BaseStream.Position - startPos < length)
			{
				ushort fieldId = r.ReadUInt16();
				switch (fieldId)
				{
							case 9:
							{
				uint mapLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var map0 = new System.Collections.Generic.Dictionary<string, Dictionary<string, Dictionary<string, sbyte>>>();
				while (r.BaseStream.Position - startPos0 < mapLength0)
				{
					string key0;
					Dictionary<string, Dictionary<string, sbyte>> value0;
									{
					uint strLen35 = r.ReadUInt32();
					var strBytes35 = r.ReadBytes((int)strLen35);
					key0 = System.Text.Encoding.UTF8.GetString(strBytes35);
				}
			
								{
				uint mapLength1 = r.ReadUInt32();
				long startPos1 = r.BaseStream.Position;
				var map1 = new System.Collections.Generic.Dictionary<string, Dictionary<string, sbyte>>();
				while (r.BaseStream.Position - startPos1 < mapLength1)
				{
					string key1;
					Dictionary<string, sbyte> value1;
									{
					uint strLen36 = r.ReadUInt32();
					var strBytes36 = r.ReadBytes((int)strLen36);
					key1 = System.Text.Encoding.UTF8.GetString(strBytes36);
				}
			
								{
				uint mapLength2 = r.ReadUInt32();
				long startPos2 = r.BaseStream.Position;
				var map2 = new System.Collections.Generic.Dictionary<string, sbyte>();
				while (r.BaseStream.Position - startPos2 < mapLength2)
				{
					string key2;
					sbyte value2;
									{
					uint strLen37 = r.ReadUInt32();
					var strBytes37 = r.ReadBytes((int)strLen37);
					key2 = System.Text.Encoding.UTF8.GetString(strBytes37);
				}
			
					value2 = r.ReadSByte();
					map2.Add(key2, value2);
				}
				value1 = map2;
			}
		
					map1.Add(key1, value1);
				}
				value0 = map1;
			}
		
					map0.Add(key0, value0);
				}
				it.ArbitraryData = map0;
			}
		
				break;
					case 8:
							{
				uint mapLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var map0 = new System.Collections.Generic.Dictionary<string, List<double>>();
				while (r.BaseStream.Position - startPos0 < mapLength0)
				{
					string key0;
					List<double> value0;
									{
					uint strLen38 = r.ReadUInt32();
					var strBytes38 = r.ReadBytes((int)strLen38);
					key0 = System.Text.Encoding.UTF8.GetString(strBytes38);
				}
			
								{
				uint listLength1 = r.ReadUInt32();
				long startPos1 = r.BaseStream.Position;
				var list1 = new System.Collections.Generic.List<double>();
				while (r.BaseStream.Position - startPos1 < listLength1)
				{
					double e1;
					e1 = r.ReadDouble();
					list1.Add(e1);
				}
				value0 = list1;
			}
		
					map0.Add(key0, value0);
				}
				it.SkillProgress = map0;
			}
		
				break;
					case 2:
							{
				Vector3 obj = new();
				_Vector3.Deserialize(obj, r);
				it.Position = obj;
			}
		
				break;
					case 4:
							{
				uint listLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var list0 = new System.Collections.Generic.List<List<Item>>();
				while (r.BaseStream.Position - startPos0 < listLength0)
				{
					List<Item> e0;
								{
				uint listLength1 = r.ReadUInt32();
				long startPos1 = r.BaseStream.Position;
				var list1 = new System.Collections.Generic.List<Item>();
				while (r.BaseStream.Position - startPos1 < listLength1)
				{
					Item e1;
								{
				Item obj = new();
				_Item.Deserialize(obj, r);
				e1 = obj;
			}
		
					list1.Add(e1);
				}
				e0 = list1;
			}
		
					list0.Add(e0);
				}
				it.Inventory = list0;
			}
		
				break;
					case 5:
							{
				uint mapLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var map0 = new System.Collections.Generic.Dictionary<string, EquipmentSlot>();
				while (r.BaseStream.Position - startPos0 < mapLength0)
				{
					string key0;
					EquipmentSlot value0;
									{
					uint strLen39 = r.ReadUInt32();
					var strBytes39 = r.ReadBytes((int)strLen39);
					key0 = System.Text.Encoding.UTF8.GetString(strBytes39);
				}
			
								{
				EquipmentSlot obj = new();
				_EquipmentSlot.Deserialize(obj, r);
				value0 = obj;
			}
		
					map0.Add(key0, value0);
				}
				it.Equipment = map0;
			}
		
				break;
					case 6:
							{
				uint mapLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var map0 = new System.Collections.Generic.Dictionary<ushort, Companion>();
				while (r.BaseStream.Position - startPos0 < mapLength0)
				{
					ushort key0;
					Companion value0;
					key0 = r.ReadUInt16();
								{
				Companion obj = new();
				_Companion.Deserialize(obj, r);
				value0 = obj;
			}
		
					map0.Add(key0, value0);
				}
				it.Companions = map0;
			}
		
				break;
					case 1:
								{
					uint strLen40 = r.ReadUInt32();
					var strBytes40 = r.ReadBytes((int)strLen40);
					it.Name = System.Text.Encoding.UTF8.GetString(strBytes40);
				}
			
				break;
					case 0:
				it.Id = r.ReadUInt64();
				break;
					case 3:
							{
				Stats obj = new();
				_Stats.Deserialize(obj, r);
				it.Stats = obj;
			}
		
				break;
					case 7:
							{
				uint listLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var list0 = new System.Collections.Generic.List<Character>();
				while (r.BaseStream.Position - startPos0 < listLength0)
				{
					Character e0;
								{
				Character obj = new();
				_Character.Deserialize(obj, r);
				e0 = obj;
			}
		
					list0.Add(e0);
				}
				it.Friends = list0;
			}
		
				break;
		
				default:
					r.BaseStream.Seek(startPos + length, SeekOrigin.Begin);
					return;
				}
			}
		}
	
		}
			class Stats : ISchema<Stats>
		{
			public int? Mana;
public int? Stamina;
public double? CritChance;
public double? CritDamage;
public Dictionary<string, short>? Resistances;
public int? Health;

			public readonly static ushort TypeId = 3;

			public static Stats CreateFromBytes(byte[] data)
			{
				Stats it = new Stats();
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_Stats.Deserialize(it, r);
				}
				return it;
			}

			public Stats Deserialize(byte[] data)
			{
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_Stats.Deserialize(this, r);
				}
				return this;
			}

			public byte[] Serialize()
			{
				using (MemoryStream ms = new MemoryStream())
				using (BinaryWriter w = new BinaryWriter(ms))
				{
					_Stats.Serialize(this, w);
					return ms.ToArray();
				}
			}
		}

		file class _Stats
		{
					static public void Serialize(Stats it, BinaryWriter w)
		{
			w.Write(Stats.TypeId);
var lengthPos = w.BaseStream.Position;
w.Write((UInt32)0);
			if (it.Mana != null)
			{
				w.Write((ushort)1);
				w.Write(it.Mana.Value);
			}			if (it.Stamina != null)
			{
				w.Write((ushort)2);
				w.Write(it.Stamina.Value);
			}			if (it.CritChance != null)
			{
				w.Write((ushort)3);
				w.Write(it.CritChance.Value);
			}			if (it.CritDamage != null)
			{
				w.Write((ushort)4);
				w.Write(it.CritDamage.Value);
			}			if (it.Resistances != null)
			{
				w.Write((ushort)5);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			foreach (var kv0 in it.Resistances)
			{
								var bytes41 = System.Text.Encoding.UTF8.GetBytes(kv0.Key);
				w.Write((uint)bytes41.Length);
				w.Write(bytes41);
			
				w.Write(kv0.Value);
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.Health != null)
			{
				w.Write((ushort)0);
				w.Write(it.Health.Value);
			}
var endPos = w.BaseStream.Position;
			w.Seek((int)lengthPos, SeekOrigin.Begin);
			w.Write((UInt32)(endPos - lengthPos - 4));
			w.Seek(0, SeekOrigin.End);
		}
	
					static public void Deserialize(Stats it, BinaryReader r)
		{
			ushort typeId = r.ReadUInt16();
			if (typeId != Stats.TypeId)
			{
				throw new Exception($"TypeId mismatch: expected Stats.TypeId but got {typeId}");
			}
			uint length = r.ReadUInt32();
			long startPos = r.BaseStream.Position;
			while (r.BaseStream.Position - startPos < length)
			{
				ushort fieldId = r.ReadUInt16();
				switch (fieldId)
				{
							case 1:
				it.Mana = r.ReadInt32();
				break;
					case 2:
				it.Stamina = r.ReadInt32();
				break;
					case 3:
				it.CritChance = r.ReadDouble();
				break;
					case 4:
				it.CritDamage = r.ReadDouble();
				break;
					case 5:
							{
				uint mapLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var map0 = new System.Collections.Generic.Dictionary<string, short>();
				while (r.BaseStream.Position - startPos0 < mapLength0)
				{
					string key0;
					short value0;
									{
					uint strLen42 = r.ReadUInt32();
					var strBytes42 = r.ReadBytes((int)strLen42);
					key0 = System.Text.Encoding.UTF8.GetString(strBytes42);
				}
			
					value0 = r.ReadInt16();
					map0.Add(key0, value0);
				}
				it.Resistances = map0;
			}
		
				break;
					case 0:
				it.Health = r.ReadInt32();
				break;
		
				default:
					r.BaseStream.Seek(startPos + length, SeekOrigin.Begin);
					return;
				}
			}
		}
	
		}
			class Item : ISchema<Item>
		{
			public string? Name;
public byte? Rarity;
public double? Weight;
public bool? IsQuestItem;
public List<string>? Tags;
public Dictionary<string, string>? ExtraData;
public uint? Id;

			public readonly static ushort TypeId = 1;

			public static Item CreateFromBytes(byte[] data)
			{
				Item it = new Item();
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_Item.Deserialize(it, r);
				}
				return it;
			}

			public Item Deserialize(byte[] data)
			{
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_Item.Deserialize(this, r);
				}
				return this;
			}

			public byte[] Serialize()
			{
				using (MemoryStream ms = new MemoryStream())
				using (BinaryWriter w = new BinaryWriter(ms))
				{
					_Item.Serialize(this, w);
					return ms.ToArray();
				}
			}
		}

		file class _Item
		{
					static public void Serialize(Item it, BinaryWriter w)
		{
			w.Write(Item.TypeId);
var lengthPos = w.BaseStream.Position;
w.Write((UInt32)0);
			if (it.Name != null)
			{
				w.Write((ushort)1);
								var bytes43 = System.Text.Encoding.UTF8.GetBytes(it.Name);
				w.Write((uint)bytes43.Length);
				w.Write(bytes43);
			
			}			if (it.Rarity != null)
			{
				w.Write((ushort)2);
				w.Write(it.Rarity.Value);
			}			if (it.Weight != null)
			{
				w.Write((ushort)3);
				w.Write(it.Weight.Value);
			}			if (it.IsQuestItem != null)
			{
				w.Write((ushort)4);
				w.Write(it.IsQuestItem.Value);
			}			if (it.Tags != null)
			{
				w.Write((ushort)5);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			for (int i0 = 0; i0 < it.Tags.Count; i0++)
			{
				var e0 = it.Tags[i0];
								var bytes44 = System.Text.Encoding.UTF8.GetBytes(e0);
				w.Write((uint)bytes44.Length);
				w.Write(bytes44);
			
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.ExtraData != null)
			{
				w.Write((ushort)6);
							var length0 = w.BaseStream.Position;
			w.Write((uint)0);
			foreach (var kv0 in it.ExtraData)
			{
								var bytes45 = System.Text.Encoding.UTF8.GetBytes(kv0.Key);
				w.Write((uint)bytes45.Length);
				w.Write(bytes45);
			
								var bytes46 = System.Text.Encoding.UTF8.GetBytes(kv0.Value);
				w.Write((uint)bytes46.Length);
				w.Write(bytes46);
			
			}
			var end0 = w.BaseStream.Position;
			w.Seek((int)length0, SeekOrigin.Begin);
			w.Write((uint)(end0 - length0 - 4));
			w.Seek(0, SeekOrigin.End);
		
			}			if (it.Id != null)
			{
				w.Write((ushort)0);
				w.Write(it.Id.Value);
			}
var endPos = w.BaseStream.Position;
			w.Seek((int)lengthPos, SeekOrigin.Begin);
			w.Write((UInt32)(endPos - lengthPos - 4));
			w.Seek(0, SeekOrigin.End);
		}
	
					static public void Deserialize(Item it, BinaryReader r)
		{
			ushort typeId = r.ReadUInt16();
			if (typeId != Item.TypeId)
			{
				throw new Exception($"TypeId mismatch: expected Item.TypeId but got {typeId}");
			}
			uint length = r.ReadUInt32();
			long startPos = r.BaseStream.Position;
			while (r.BaseStream.Position - startPos < length)
			{
				ushort fieldId = r.ReadUInt16();
				switch (fieldId)
				{
							case 1:
								{
					uint strLen47 = r.ReadUInt32();
					var strBytes47 = r.ReadBytes((int)strLen47);
					it.Name = System.Text.Encoding.UTF8.GetString(strBytes47);
				}
			
				break;
					case 2:
				it.Rarity = r.ReadByte();
				break;
					case 3:
				it.Weight = r.ReadDouble();
				break;
					case 4:
				it.IsQuestItem = r.ReadBoolean();
				break;
					case 5:
							{
				uint listLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var list0 = new System.Collections.Generic.List<string>();
				while (r.BaseStream.Position - startPos0 < listLength0)
				{
					string e0;
									{
					uint strLen48 = r.ReadUInt32();
					var strBytes48 = r.ReadBytes((int)strLen48);
					e0 = System.Text.Encoding.UTF8.GetString(strBytes48);
				}
			
					list0.Add(e0);
				}
				it.Tags = list0;
			}
		
				break;
					case 6:
							{
				uint mapLength0 = r.ReadUInt32();
				long startPos0 = r.BaseStream.Position;
				var map0 = new System.Collections.Generic.Dictionary<string, string>();
				while (r.BaseStream.Position - startPos0 < mapLength0)
				{
					string key0;
					string value0;
									{
					uint strLen49 = r.ReadUInt32();
					var strBytes49 = r.ReadBytes((int)strLen49);
					key0 = System.Text.Encoding.UTF8.GetString(strBytes49);
				}
			
									{
					uint strLen50 = r.ReadUInt32();
					var strBytes50 = r.ReadBytes((int)strLen50);
					value0 = System.Text.Encoding.UTF8.GetString(strBytes50);
				}
			
					map0.Add(key0, value0);
				}
				it.ExtraData = map0;
			}
		
				break;
					case 0:
				it.Id = r.ReadUInt32();
				break;
		
				default:
					r.BaseStream.Seek(startPos + length, SeekOrigin.Begin);
					return;
				}
			}
		}
	
		}
			class Vector3 : ISchema<Vector3>
		{
			public double? X;
public double? Y;
public double? Z;

			public readonly static ushort TypeId = 2;

			public static Vector3 CreateFromBytes(byte[] data)
			{
				Vector3 it = new Vector3();
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_Vector3.Deserialize(it, r);
				}
				return it;
			}

			public Vector3 Deserialize(byte[] data)
			{
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_Vector3.Deserialize(this, r);
				}
				return this;
			}

			public byte[] Serialize()
			{
				using (MemoryStream ms = new MemoryStream())
				using (BinaryWriter w = new BinaryWriter(ms))
				{
					_Vector3.Serialize(this, w);
					return ms.ToArray();
				}
			}
		}

		file class _Vector3
		{
					static public void Serialize(Vector3 it, BinaryWriter w)
		{
			w.Write(Vector3.TypeId);
var lengthPos = w.BaseStream.Position;
w.Write((UInt32)0);
			if (it.X != null)
			{
				w.Write((ushort)0);
				w.Write(it.X.Value);
			}			if (it.Y != null)
			{
				w.Write((ushort)1);
				w.Write(it.Y.Value);
			}			if (it.Z != null)
			{
				w.Write((ushort)2);
				w.Write(it.Z.Value);
			}
var endPos = w.BaseStream.Position;
			w.Seek((int)lengthPos, SeekOrigin.Begin);
			w.Write((UInt32)(endPos - lengthPos - 4));
			w.Seek(0, SeekOrigin.End);
		}
	
					static public void Deserialize(Vector3 it, BinaryReader r)
		{
			ushort typeId = r.ReadUInt16();
			if (typeId != Vector3.TypeId)
			{
				throw new Exception($"TypeId mismatch: expected Vector3.TypeId but got {typeId}");
			}
			uint length = r.ReadUInt32();
			long startPos = r.BaseStream.Position;
			while (r.BaseStream.Position - startPos < length)
			{
				ushort fieldId = r.ReadUInt16();
				switch (fieldId)
				{
							case 0:
				it.X = r.ReadDouble();
				break;
					case 1:
				it.Y = r.ReadDouble();
				break;
					case 2:
				it.Z = r.ReadDouble();
				break;
		
				default:
					r.BaseStream.Seek(startPos + length, SeekOrigin.Begin);
					return;
				}
			}
		}
	
		}
	