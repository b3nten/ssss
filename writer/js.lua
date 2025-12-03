Schema = Schema

local function fmt(s, tab)
	return (s:gsub('($%b{})', function(w) return tab[w:sub(3, -2)] or w end))
end
getmetatable("").__mod = fmt

local function to_pascal_case(str)
	local result = str:gsub("[%w]+", function(word)
		return word:sub(1, 1):upper() .. word:sub(2):lower()
	end)
	return (result:gsub("[^%w]", ""))
end

local desField = "deserialize_field"
local serField = "serialize"
local staticDes = "deserialize"

local list_writers = {}
local list_deserializers = {}

local function print_list_serializer(type)
	if type.kind == "primitive" then
		return "write_${type}" % { type = type.name }
	elseif type.kind == "struct" then
		return "${stype}_${ser}" % {
			stype = to_pascal_case(type.name),
			ser = serField
		}
	elseif type.kind == "list" then
		return "make_list_writer(${w})" % { w = print_list_serializer(type.of) }
	else
		error("Unknown list element type")
	end
end

local function print_serializer_fn(field)
	if field.type.kind == "primitive" then
		return "write_${type}(it.${field}, b);" % { type = field.type.name, field = field.name }
	elseif field.type.kind == "struct" then
		return "${stype}_${ser}(it.${field}, b)" %
				{ field = field.name, stype = to_pascal_case(field.type.name), ser = serField }
	elseif field.type.kind == "list" then
		list_writers[#list_writers + 1] = "make_list_writer(${w})" % {
			w = print_list_serializer(field.type.of)
		}
		return "lw" .. #list_writers .. "(it.${field}, b)" % { field = field.name }
	else
		error("Unknown field type")
	end
end

local function print_struct_serializer(struct)
	local out = "function ${sname}_${ser}(it, b) {\n" % {
		sname = to_pascal_case(struct.name),
		ser = serField
	}
	out = out .. "\twrite_uint16(${typeid}, b);\n" % { typeid = struct.id }
	out = out .. "\tconst start_index = b.length;\n"
	out = out .. "\twrite_uint32(0, b);\n"
	for _, field in pairs(struct.fields) do
		out = out .. "\tif(typeof it.${field} !== 'undefined') {\n" % { field = field.name }
		out = out .. "\t\twrite_uint16(${fieldid}, b);\n" % { fieldid = field.id }
		out = out .. "\t\t" .. print_serializer_fn(field) .. "\n"
		out = out .. "\t}\n"
	end
	out = out .. "\tconst end_index = b.length;\n"
	out = out .. "\tb.set_uint32(start_index, end_index - (start_index + 4))\n"
	out = out .. "\treturn b;\n"
	out = out .. "}\n"
	return out
end

local function print_list_deserializer(type)
	if type.kind == "primitive" then
		return "deserialize_${type}" % { type = type.name }
	elseif type.kind == "struct" then
		return "${stype}_${sdes}" % {
			stype = to_pascal_case(type.name),
			sdes = staticDes
		}
	elseif type.kind == "list" then
		return "make_list_deserializer(${w})" % {
			w = print_list_deserializer(type.of)
		}
	else
		error("Unknown list element type")
	end
end

local function print_deserializer_fn(field)
	if field.type.kind == "primitive" then
		return "deserialize_${type}(view, offset, it, '${field}')" % { type = field.type.name, field = field.name }
	elseif field.type.kind == "struct" then
		return "${stype}_${sdes}(view, offset, it, '${field}')"
				% {
					stype = to_pascal_case(field.type.name),
					field = field.name,
					sdes = staticDes
				}
	elseif field.type.kind == "list" then
		list_deserializers[#list_deserializers + 1] = "make_list_deserializer(${w})" % {
			w = print_list_deserializer(field.type.of)
		}
		return "ld" .. #list_deserializers .. "(view, offset, it, '${field}')" % { field = field.name }
	else
		error("Unknown field type")
	end
end

local function print_deserializer_switch(struct)
	local out = "function ${sname}_${name}(it, view, fieldID, offset) {\n" %
			{ name = desField, sname = to_pascal_case(struct.name) }
	out = out .. "\tswitch(fieldID) {\n"
	for _, field in pairs(struct.fields) do
		out = out .. "\t\tcase ${fieldid}: return ${fn}\n" % {
			fieldid = field.id,
			fn = print_deserializer_fn(field)
		}
	end
	out = out .. "\t\tdefault:\n"
	out = out .. "\t\t\treturn unknown_field;\n"
	out = out .. "\t}\n"
	out = out .. "}\n"
	return out
end

local function print_static_deserializer(struct)
	local out = "function ${sname}_deserialize(view, offset, struct, field){" % { sname = to_pascal_case(struct.name) }
	out = out .. "\n\tconst s = new ${sname}();" % { sname = to_pascal_case(struct.name) }
	out = out .. "\n\toffset = parse_struct(s, ${sname}_deserialize_field, view, offset);" % {
		sname = to_pascal_case(struct.name)
	}
	out = out .. "\n\tstruct[field] = s;"
	out = out .. "\n\treturn offset;"
	out = out .. "\n}\n"
	return out
end

local function print_struct(struct)
	local out = "export class ${name} {\n" % { name = to_pascal_case(struct.name) }
	out = out .. "\tstatic get TypeID() { return ${typeid}; }" % { typeid = struct.id }
	out = out .. "\n\tconstructor(props = {}){ Object.assign(this, props) }"
	out = out .. "\n\ttoBytes() { return ${sname}_serialize(this, new ByteBuffer()).bytes(); }" % {
		sname = to_pascal_case(struct.name)
	}
	out = out .. "\n\tfromBytes(bytes) {"
	out = out .. "\n\t\tif (!('buffer' in bytes)) bytes = new Uint8Array(bytes);"
	out = out ..
			"\n\t\tparse_struct(this, ${method}, new DataView(bytes.buffer, bytes.byteOffset, bytes.byteLength), 0);\n"
			% { method = to_pascal_case(struct.name) .. "_" .. desField }
	out = out .. "\t\treturn this;\n"
	out = out .. "\t}\n"
	out = out .. "}\n"
	out = out .. print_struct_serializer(struct) .. "\n"
	out = out .. print_static_deserializer(struct) .. "\n"
	out = out .. print_deserializer_switch(struct) .. "\n"
	return out
end

local function print_deserialize(structs)
	local out = ""
	out = out .. "export function deserialize(bytes) {"
	out = out .. "\n\tif (!('buffer' in bytes)) bytes = new Uint8Array(bytes);"
	out = out .. "\n\tconst view = new DataView(bytes.buffer, bytes.byteOffset, bytes.byteLength);"
	out = out .. "\n\tconst typeID = view.getUint16(0, true);"
	out = out .. "\n\tswitch(typeID) {"
	for _, v in pairs(structs) do
		out = out .. "\n\t\tcase ${typeid}: return new ${sname}().fromBytes(bytes);" % {
			typeid = v.id,
			sname = to_pascal_case(v.name)
		}
	end
	out = out .. "\n\t\tdefault: throw new Error(`Unknown TypeID: ${typeID}`);"
	out = out .. "\n\t}"
	out = out .. "\n}\n"
	return out
end

local function print_ts_type(type)
	if type.kind == "primitive" then
		if type.name == "bool" then
			return "boolean"
		elseif type.name == "string" then
			return "string"
		else
			return "number"
		end
	elseif type.kind == "struct" then
		return to_pascal_case(type.name)
	elseif type.kind == "list" then
		return print_ts_type(type.of) .. "[]"
	else
		error("Unknown type for TS defs")
	end
end

local function print_ts_defs(structs)
	local out = ""
	for _, struct in pairs(structs) do
		out = out .. "export class ${sname} {\n" % { sname = to_pascal_case(struct.name) }
		for _, field in pairs(struct.fields) do
			out = out .. "\t${field}?: ${type};\n" % {
				field = field.name,
				type = print_ts_type(field.type)
			}
		end
		out = out .. "\tconstructor(props?: Omit<Partial<${sname}>, 'fromBytes' | 'toBytes'>);" % { sname = to_pascal_case(struct.name) }
		out = out .. "\n\tstatic readonly TypeID: number;\n"
		out = out .. "\ttoBytes(): Uint8Array;\n"
		out = out ..
				"\tfromBytes(bytes: ArrayBuffer | ArrayBufferView): ${sname};\n" % { sname = to_pascal_case(struct.name) }
		out = out .. "}\n\n"
	end
	out = out .. "export function deserialize(bytes: ArrayBuffer | ArrayBufferView): "
	for _, struct in pairs(structs) do
		out = out .. "${sname} | " % { sname = to_pascal_case(struct.name) }
	end
	out = out:sub(1, -4) .. ";\n\n"
	return out
end

local function print_prelude(name, version)
	local out = "// Auto-generated code for schema: ${name} v${version}\n\n" % {
		name = name,
		version = version
	}
	return out
end

local function print_list_items()
	local out = ""
	for i, v in pairs(list_writers) do
		out = out .. "const lw${i} = ${fn};\n" %	{ i = i, fn = v }
	end
	for i, v in pairs(list_deserializers) do
		out = out .. "const ld${i} = ${fn};\n" %	{ i = i, fn = v }
	end
	return out
end

-- CODEGEN STEP

local js_file = print_prelude(Schema.name, Schema.version)
for _, v in pairs(Schema.structs) do
	js_file = js_file .. print_struct(v)
end
js_file = js_file .. print_deserialize(Schema.structs)
js_file = js_file .. "\n" .. print_list_items()

local ts_file = print_prelude(Schema.name, Schema.version) .. print_ts_defs(Schema.structs)

-- APPEND INCLDUES AND SET js_file
local include = [[
let tmp;

class ByteBuffer {
	get length() { return this.len; }
	encoder = new TextEncoder();
	buffer = new ArrayBuffer(0xFFF)
	view = new Uint8Array(this.buffer, 0)
	dview = new DataView(this.buffer, 0)
	len = 0;

	write(value) {
		tmp = this.length;
		this.resize(this.len + value.length);
		this.view.set(value, tmp);
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

function write_bool(value, b) {
	tmp = b.length;
	b.resize(b.len + 1);
	b.dview.setUint8(tmp, value ? 1 : 0, true);
}

function write_int8(value, b) {
	tmp = b.length;
	b.resize(b.len + 1);
	b.dview.setInt8(tmp, value, true);
}

function write_uint8(value, b) {
	tmp = b.length;
	b.resize(b.len + 1);
	b.dview.setUint8(tmp, value, true);
}

function write_int16(value, b) {
	tmp = b.length;
	b.resize(b.len + 2);
	b.dview.setInt16(tmp, value, true);
}

function write_uint16(value, b) {
	tmp = b.length;
	b.resize(b.len + 2);
	b.dview.setUint16(tmp, value, true);
}

function write_int32(value, b) {
	tmp = b.length;
	b.resize(b.len + 4);
	b.dview.setInt32(tmp, value, true);
}

function write_uint32(value, b) {
	tmp = b.length;
	b.resize(b.len + 4);
	b.dview.setUint32(tmp, value, true);
}

function write_f32(value, b) {
	tmp = b.length;
	b.resize(b.len + 4);
	b.dview.setFloat32(tmp, value, true);
}

function write_f64(value, b) {
	tmp = b.length;
	b.resize(b.len + 8);
	b.dview.setFloat64(tmp, value, true);
}

function write_string(value, b) {
	const stringLength = value.length;
	if (stringLength > 300) {
		const encoded = b.encoder.encode(value);
		b.set_uint32(b.length, encoded.length);
		b.write(encoded);
		return;
	}
	const lengthPos = b.length;
	write_uint32(0, b);
	const start = b.length;
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
			write_uint8(codePoint, b);
		} else {
			if (codePoint < 0x800) {
				write_uint8(((codePoint >> 6) & 0x1F) | 0xC0, b);
			} else {
				if (codePoint < 0x10000) {
					write_uint8(((codePoint >> 12) & 0x0F) | 0xE0, b);
				} else {
					write_uint8(((codePoint >> 18) & 0x07) | 0xF0, b);
					write_uint8(((codePoint >> 12) & 0x3F) | 0x80, b);
				}
				write_uint8(((codePoint >> 6) & 0x3F) | 0x80, b);
			}
			write_uint8((codePoint & 0x3F) | 0x80, b);
		}
	}
	b.set_uint32(lengthPos, b.length - start);
}

function make_list_writer(s) {
	return (value, b) => {
		write_uint32(0, b);
		const sizeIndex = b.len;
		for (const item of value) s(item, b);
		b.set_uint32(sizeIndex - 4, b.len - sizeIndex);
	}
}

function deserialize_bool(view, offset, struct, field) {
	struct[field] = view.getUint8(offset, true) !== 0;
	return offset + 1;
}

function deserialize_int8(data, offset, struct, field) {
	struct[field] = data.getInt8(offset, true);
	return offset + 1;
}

function deserialize_uint8(data, offset, struct, field) {
	struct[field] = data.getUint8(offset, true);
	return offset + 1;
}

function deserialize_int16(data, offset, struct, field) {
	struct[field] = data.getInt16(offset, true);
	return offset + 2;
}

function deserialize_uint16(data, offset, struct, field) {
	struct[field] = data.getUint16(offset, true);
	return offset + 2;
}

function deserialize_int32(data, offset, struct, field) {
	struct[field] = data.getInt32(offset, true);
	return offset + 4;
}

function deserialize_uint32(data, offset, struct, field) {
	struct[field] = data.getUint32(offset, true);
	return offset + 4;
}

function deserialize_f32(data, offset, struct, field) {
	struct[field] = data.getFloat32(offset, true);
	return offset + 4;
}

function deserialize_f64(data, offset, struct, field) {
	struct[field] = data.getFloat64(offset, true);
	return offset + 8;
}

const text_decoder = new TextDecoder();
function deserialize_string(data, offset, struct, field) {
	const length = data.getUint32(offset, true);
	offset += 4;
	if (length > 300) {
		const bytes = new Uint8Array(data.buffer, data.byteOffset + offset, length);
		struct[field] = text_decoder.decode(bytes);
		return offset + length;
	} else {
		const end = offset + length;
		let result = "";
		let codePoint;
		while (offset < end) {
			const a = data.getUint8(offset++);
			if (a < 0xC0) {
				codePoint = a;
			} else {
				const b = data.getUint8(offset++);
				if (a < 0xE0) {
					codePoint = ((a & 0x1F) << 6) | (b & 0x3F);
				} else {
					const c = data.getUint8(offset++);
					if (a < 0xF0) {
						codePoint = ((a & 0x0F) << 12) | ((b & 0x3F) << 6) | (c & 0x3F);
					} else {
						const d = data.getUint8(offset++);
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
		offset = end;
		struct[field] = result;
		return offset;
	}
}

function make_list_deserializer(item_deserializer) {
	return (data, offset, struct, field) => {
		const length = data.getUint32(offset, true);
		offset += 4;
		const endOffset = offset + length;
		const list = [];
		let i = 0;
		while (offset < endOffset) {
			offset = item_deserializer(data, offset, list, i);
			i++;
		}
		struct[field] = list;
		return offset;
	};
}

function parse_struct(struct, field_method, view, offset) {
	const typeID = view.getUint16(offset, true);
	offset += 2;
	if (typeID !== struct.constructor.TypeID) {
		throw new Error(`Type Mismatch: Expected ${struct.constructor.TypeID} got ${typeID} for ${struct.constructor.name}`);
	}
	const length = view.getUint32(offset, true);
	const totalSize = length + 6;
	offset += 4;
	const endOffset = offset + length;
	while (offset < endOffset) {
		const fieldID = view.getUint16(offset, true);
		offset += 2;
		const next = field_method(struct, view, fieldID, offset)
		if (next === unknown_field) {
			return totalSize;
		}
		offset = next;
	}
	return offset;
}

const unknown_field = new Error("Unknown Field")
]]

Output = {}

Output[Schema.name .. ".js"] = js_file .. "\n" .. include .. "\n"
Output[Schema.name .. ".d.ts"] = ts_file
