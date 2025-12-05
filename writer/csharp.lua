local cstype_map = {
	int8 = "sbyte",
	uint8 = "byte",
	int16 = "short",
	uint16 = "ushort",
	int32 = "int",
	uint32 = "uint",
	int64 = "long",
	uint64 = "ulong",
	f32 = "float",
	f64 = "double",
	bool = "bool",
	string = "string",
}

local ctype_to_reader = {
	int8 = "ReadSByte",
	uint8 = "ReadByte",
	int16 = "ReadInt16",
	uint16 = "ReadUInt16",
	int32 = "ReadInt32",
	uint32 = "ReadUInt32",
	int64 = "ReadInt64",
	uint64 = "ReadUInt64",
	f32 = "ReadSingle",
	f64 = "ReadDouble",
	bool = "ReadBoolean",
}

local uses_value = {
	int8 = true,
	uint8 = true,
	int16 = true,
	uint16 = true,
	int32 = true,
	uint32 = true,
	int64 = true,
	uint64 = true,
	f32 = true,
	f64 = true,
	bool = true,
}

local function cs_type(field_type)
	if field_type.kind == "primitive" then
		return cstype_map[field_type.name]
	elseif field_type.kind == "struct" then
		return pascal_case(field_type.name)
	elseif field_type.kind == "list" then
		return sprintf("List<${type}>", { type = cs_type(field_type.of) })
	elseif field_type.kind == "map" then
		return sprintf("Dictionary<${key_type}, ${value_type}>", {
			key_type = cs_type(field_type.from),
			value_type = cs_type(field_type.to),
		})
	else
		error("Unknown field type kind: " .. tostring(field_type.kind))
	end
end

local function print_field_serialization(name, field)
	if field.type.kind == "primitive" then
		if field.type.name == "string" then
			return str_block {
				fname = name,
			} [[
				var bytes = System.Text.Encoding.UTF8.GetBytes(${fname});
				w.Write((uint)bytes.Length);
				w.Write(bytes);
			]]
		end
		return str_block {
			fname = name,
			uses_value = uses_value[field.type.name] and not field.index and ".Value" or "",
		} ("w.Write(${fname}${uses_value});")
	elseif field.type.kind == "struct" then
		return str_block {
			fname = name,
			ftype = pascal_case(field.type.name),
		} ("_${ftype}.Serialize(${fname}, w);")
	elseif field.type.kind == "map" then
		local key_type = field.type.from
		local value_type = field.type.to
		local index = field.index or 0
		return str_block {
			fname = name,
			index = field.index or 0,
			key_serialization = print_field_serialization("kv" .. tostring(index) .. ".Key", { type = key_type, index = index + 1 }),
			value_serialization = print_field_serialization("kv" .. tostring(index) .. ".Value", { type = value_type, index = index + 1 }),
		} [[
			var length${index} = w.BaseStream.Position;
			w.Write((uint)0);
			foreach (var kv${index} in ${fname})
			{
				${key_serialization}
				${value_serialization}
			}
			var end${index} = w.BaseStream.Position;
			w.Seek((int)length${index}, SeekOrigin.Begin);
			w.Write((uint)(end${index} - length${index} - 4));
			w.Seek(0, SeekOrigin.End);
		]]
	elseif field.type.kind == "list" then
		local of_type = field.type.of
		local index = field.index or 0
		return str_block {
			fname = name,
			index = field.index or 0,
			element_serialization = print_field_serialization("e" .. tostring(index), { type = of_type, index = index + 1 }),
		} [[
			var length${index} = w.BaseStream.Position;
			w.Write((uint)0);
			for (int i${index} = 0; i${index} < ${fname}.Count; i${index}++)
			{
				var e${index} = ${fname}[i${index}];
				${element_serialization}
			}
			var end${index} = w.BaseStream.Position;
			w.Seek((int)length${index}, SeekOrigin.Begin);
			w.Write((uint)(end${index} - length${index} - 4));
			w.Seek(0, SeekOrigin.End);
		]]
	end
end

local function print_fields_serialization(struct)
	local out = ""
	for field_name, field in pairs(struct.fields) do
		out = out .. str_block {
			fname = pascal_case(field_name),
			fid = field.id,
			fcontent = print_field_serialization("it." .. pascal_case(field_name), field),
			check = uses_value[field.type.name] and "it." .. pascal_case(field_name) .. ".HasValue" or "it." .. pascal_case(field_name) .. " != null",
		} [[
			if (it.${fname} != null)
			{
				w.Write((ushort)${fid});
				${fcontent}
			}]]
	end
	return out
end

local function print_serializer(struct_name, struct)
	return str_block {
		sname = pascal_case(struct_name),
		field_serialization = print_fields_serialization(struct)
	} [[
		static public void Serialize(${sname} it, BinaryWriter w)
		{
			w.Write(${sname}.TypeId);
      var lengthPos = w.BaseStream.Position;
      w.Write((UInt32)0);
      ${field_serialization}
      var endPos = w.BaseStream.Position;
			w.Seek((int)lengthPos, SeekOrigin.Begin);
			w.Write((UInt32)(endPos - lengthPos - 4));
			w.Seek(0, SeekOrigin.End);
		}
	]]
end

local function print_field_deserialization(name, field)
	if field.type.kind == "primitive" then
		if field.type.name == "string" then
			return str_block {
				fname = name,
			} [[
				{
					uint strLen = r.ReadUInt32();
					var strBytes = r.ReadBytes((int)strLen);
					${fname} = System.Text.Encoding.UTF8.GetString(strBytes);
				}
			]]
		end
		return str_block {
			fname = name,
			ftype = ctype_to_reader[field.type.name],
		} ("${fname} = r.${ftype}();")
	elseif field.type.kind == "struct" then
		return str_block {
			fname = name,
			ftype = pascal_case(field.type.name),
		} [[
			{
				${ftype} obj = new();
				_${ftype}.Deserialize(obj, r);
				${fname} = obj;
			}
		]]
	elseif field.type.kind == "map" then
		local key_type = field.type.from
		local value_type = field.type.to
		local index = field.index or 0
		return str_block {
			name = name,
			index = field.index or 0,
			key_type = key_type,
			value_type = value_type,
			ktype = cs_type(key_type),
			vtype = cs_type(value_type),
			key_des_fn = print_field_deserialization("key" .. tostring(index), { type = key_type, index = index + 1 }),
			value_des_fn = print_field_deserialization("value" .. tostring(index), { type = value_type, index = index + 1 }),
		} [[
			{
				uint mapLength${index} = r.ReadUInt32();
				long startPos${index} = r.BaseStream.Position;
				var map${index} = new System.Collections.Generic.Dictionary<${ktype}, ${vtype}>();
				while (r.BaseStream.Position - startPos${index} < mapLength${index})
				{
					${ktype} key${index};
					${vtype} value${index};
					${key_des_fn}
					${value_des_fn}
					map${index}.Add(key${index}, value${index});
				}
				${name} = map${index};
			}
		]]
	elseif field.type.kind == "list" then
		local of_type = field.type.of
		local index = field.index or 0
		return str_block {
			name = name,
			index = field.index or 0,
			of_type = of_type,
			cstype = cs_type(of_type),
			des_fn = print_field_deserialization("e" .. tostring(index), { type = of_type, index = index + 1 }),
		} [[
			{
				uint listLength${index} = r.ReadUInt32();
				long startPos${index} = r.BaseStream.Position;
				var list${index} = new System.Collections.Generic.List<${cstype}>();
				while (r.BaseStream.Position - startPos${index} < listLength${index})
				{
					${cstype} e${index};
					${des_fn}
					list${index}.Add(e${index});
				}
				${name} = list${index};
			}
		]]
	end
end

local function print_deserializer_cases(struct)
	local cases = ""
	for field_name, field in pairs(struct.fields) do
		local fname = "it." .. pascal_case(field_name)
		cases = cases .. str_block {
			fname = fname,
			fid = field.id,
			fcontent = print_field_deserialization(fname, field),
		} [[
			case ${fid}:
				${fcontent}
				break;
		]]
	end
	return cases
end

local function print_deserializer(struct_name, struct)
	return str_block {
		sname = pascal_case(struct_name),
		cases = print_deserializer_cases(struct),
	} [[
		static public void Deserialize(${sname} it, BinaryReader r)
		{
			ushort typeId = r.ReadUInt16();
			if (typeId != ${sname}.TypeId)
			{
				throw new Exception($"TypeId mismatch: expected ${sname}.TypeId but got {typeId}");
			}
			uint length = r.ReadUInt32();
			long startPos = r.BaseStream.Position;
			while (r.BaseStream.Position - startPos < length)
			{
				ushort fieldId = r.ReadUInt16();
				switch (fieldId)
				{
				${cases}
				default:
					r.BaseStream.Seek(startPos + length, SeekOrigin.Begin);
					return;
				}
			}
		}
	]]
end

local function print_class_fields(struct_name, struct)
	local fields = {}
	for field_name, field in pairs(struct.fields) do
		table.insert(
			fields,
			#fields + 1,
			sprintf(
				"public ${type}? ${name};",
				{ type = cs_type(field.type), name = pascal_case(field_name) }
			)
		)
	end
	return table.concat(fields, "\n")
end

local function print_struct(struct_name, struct)
	return str_block {
		name = pascal_case(Schema.name),
		sname = pascal_case(struct_name),
		typeid = struct.id,
		class_fields = print_class_fields(struct_name, struct),
		ser_fn = print_serializer(struct_name, struct),
		deser_fn = print_deserializer(struct_name, struct),
	} [[
		class ${sname} : I${name}<${sname}>
		{
			${class_fields}

			public readonly static ushort TypeId = ${typeid};

			public static ${sname} CreateFromBytes(byte[] data)
			{
				${sname} it = new ${sname}();
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_${sname}.Deserialize(it, r);
				}
				return it;
			}

			public ${sname} Deserialize(byte[] data)
			{
				using (MemoryStream ms = new MemoryStream(data))
				using (BinaryReader r = new BinaryReader(ms))
				{
					_${sname}.Deserialize(this, r);
				}
				return this;
			}

			public byte[] Serialize()
			{
				using (MemoryStream ms = new MemoryStream())
				using (BinaryWriter w = new BinaryWriter(ms))
				{
					_${sname}.Serialize(this, w);
					return ms.ToArray();
				}
			}
		}

		file class _${sname}
		{
			${ser_fn}
			${deser_fn}
		}
	]]
end

-- BUILD FILE

local file = str_block {
	name = pascal_case(Schema.name),
	version = Schema.version,
} [[
// Auto-generated code for schema: ${name} v${version}

namespace ${name};

interface I${name}<TSelf> where TSelf : I${name}<TSelf>
{
	public byte[] Serialize();
	public TSelf Deserialize(byte[] data);
}
]]


for struct_name, struct in pairs(Schema.structs) do
	file = file .. print_struct(struct_name, struct)
end

-- WRITE OUTPUT
Output["${name}.cs" % { name = pascal_case(Schema.name) }] = file
