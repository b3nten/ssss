local function print_struct_serializer(name, struct)
	local function print_writer(field, input)
		if field.type.kind == "primitive" then
			return [[
				b.write_${type}(${input})
			]] % { input = input, type = field.type.name }
		elseif field.type.kind == "struct" then
			return [[
				${type}_serialize(b, ${input})
			]] % {
				input = input,
				type = pascal_case(field.type.name)
		 }
		elseif field.type.kind == "list" then
			local i = field.index or 0
			return [[
				b.write_uint32(0);
				const listStart${i} = b.length;
				for(const item${i} of ${input}) {
					${item_writer}
				}
				b.set_uint32(listStart${i} - 4, b.length - listStart${i});
			]] % {
				i = i,
				input = input,
				item_writer = print_writer({ type = field.type.of, index = i + 1 }, "item"..i)
			}
		elseif field.type.kind == "map" then
			local i = field.index or 0
			return [[
				b.write_uint32(0);
				const mapStart${i} = b.length;
				for(const [key${i}, value${i}] of Object.entries(${input})) {
					${key_writer}
					${value_writer}
				}
				b.set_uint32(mapStart${i} - 4, b.length - mapStart${i});
			]] % {
				i = i,
				input = input,
				key_writer = print_writer({ type = field.type.from, index = i + 1 }, "key"..i),
				value_writer = print_writer({ type = field.type.to, index = i + 1 }, "value"..i)
			}
		else
			return ""
		end
	end

	local function print_field_serializers()
		local out = ""
		for _, field in pairs(struct.fields) do
			out = out .. [[
			if(s.${fname} !== undefined) {
				b.write_uint16(${field_id});
				${writer}
			}
		  ]] % {
				fname = field.name,
				field_id = field.id,
				writer = print_writer(field, "s.${fname}" % { fname = field.name } )
			}
		end
		return out
	end

	return [[
		function ${sname}_serialize(b, s) {
			b.write_uint16(${type_id})
			b.write_uint32(0);
			const structStart = b.length;
			${field_serializers}
			b.set_uint32(structStart - 4, b.length - structStart);
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

	local function print_field_deserializer(field, output)
		if field.type.kind == 'primitive' then
			return [[
				${output} = br.read_${type}();
			]] % {
				output = output,
				type = field.type.name
			}
		elseif field.type.kind == "struct" then
			return [[
				${output} = new ${type}();
				${type}_deserialize(br, ${output});
			]] % {
				output = output,
				type = pascal_case(field.type.name)
			}
		elseif field.type.kind == "list" then
			local i = field.index or 0
			return [[
			{
				const listLength${i} = br.read_uint32();
				if (listLength${i} > (br.length - br.position)) {
					throw new Error("Invalid list length");
				}
				const listStart${i} = br.position;
				${output} = [];
				for (; br.position - listStart${i} < listLength${i};) {
					let item${i};
					${item_deserializer}
					${output}.push(item${i});
				}
			}
			]] % {
				i = i,
				output = output,
				item_deserializer = print_field_deserializer({ type = field.type.of, index = i + 1 }, "item"..i)
			}
		elseif field.type.kind == "map" then
			local i = field.index or 0
			return [[
			{
				const mapLength${i} = br.read_uint32();
				if (mapLength${i} > (br.length-br.position)) {
					throw new Error("Invalid map length");
				}
				const mapStart${i} = br.position;
				${output} = {};
				for (; br.position - mapStart${i} < mapLength${i};) {
					let key${i};
					${key_deserializer}
					let value${i};
					${value_deserializer}
					${output}[key${i}] = value${i};
				}
			}
			]] % {
				i = i,
				output = output,
				key_deserializer = print_field_deserializer({ type = field.type.from, index = i + 1 }, "key"..i),
				value_deserializer = print_field_deserializer({ type = field.type.to, index = i + 1 }, "value"..i)
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
					break;
				]] % {
				field_id = field.id,
				field_deserializer = print_field_deserializer(field, "s.${field_name}" % { field_name = field.name })
			}
		end
		return out
	end

	return [[
		function ${sname}_deserialize(br, s) {
			const typeId = br.read_uint16()
			if (typeId !== ${type_id}) {
				throw new Error("Type ID mismatch deserializing struct ${sname}: expected ${type_id}, got " + typeId);
			}
			const length = br.read_uint32()
			if (length > (br.length - br.position)) {
				throw new Error("Struct ${sname} length exceeds buffer length");
			}
			const seenFields = new Set;
			const startPos = br.position;
			for (; br.position - startPos < length;) {
				const fieldId = br.read_uint16();
				if (seenFields.has(fieldId)) {
					throw new Error("Duplicate field ID " + fieldId + " in struct ${sname}");
				}
				if (fieldId > ${max_field_id}) {
					return;
				}
				seenFields.add(fieldId);
				switch (fieldId) {
					${field_deserializers}
				}
			}
		}
	]] % {
		sname = pascal_case(name),
		type_id = struct.id,
		max_field_id = get_max_field_id(),
		field_deserializers = print_field_deserializers()
	}
end

local function print_struct(name, struct)
	local function print_fields()
		local out = ""
		for field_id, field in pairs(struct.fields) do
			out = out .. [[
				${fname}
			]] % {
				fname = field.name,
			}
		end
		return out
	end

	return [[
		export class ${sname} {
			static get TypeID() { return ${type_id} }

			constructor(fields) {
				if (fields) Object.assign(this, fields)
			}

			${fields}

			deserialize(bytes) {
				const b = new ByteReader(bytes)
				${sname}_deserialize(b, this)
				return this;
			}

			serialize() {
				const w = new ByteWriter;
				${sname}_serialize(w, this)
				return w.bytes()
			}
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

local function print_types()
	local function print_struct_type(name, struct)
		local function print_field(field)
			if field.type.kind == "primitive" then
				if field.type.name == "int64" or field.type.name == "uint64" then
					return "bigint"
				elseif field.type.name == "string" then
					return "string"
				elseif field.type.name == "bool" then
					return "boolean"
				else return "number"
				end
			elseif field.type.kind == "struct" then
				return pascal_case(field.type.name)
			elseif field.type.kind == "list" then
				return print_field({ type = field.type.of }) .. "[]"
			elseif field.type.kind == "map" then
				return "Record<" .. print_field({ type = field.type.from }) .. ", " .. print_field({ type = field.type.to }) .. ">"
			end
			return "any"
		end
		local function print_fields()
			local out = ""
			for _, field in pairs(struct.fields) do
				out = out .. [[
					${fname}?: ${ftype};
				]] % {
					fname = field.name,
					ftype = print_field(field)
				}
			end
			return out
		end
		return [[
			export class ${sname} {
				static readonly TypeID: number;

				constructor(fields?: Partial<Omit<${sname}, "serialize" | "deserialize">>);

				${fields}
				deserialize(bytes: Uint8Array): ${sname};
				serialize(): Uint8Array;
			}
		]] % {
			sname = pascal_case(name),
			fields = print_fields()
		}
	end
	local out = ""
	for name, struct in pairs(Schema.structs) do
		out = out .. print_struct_type(name, struct) .. "\n"
	end
	return out
end

local file = str_block {
	name = Schema.name,
	structs = print_structs(),
} [[
${structs}

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
]]

Output[Schema.name .. ".js"] = file

Output[Schema.name .. ".d.ts"] = print_types()
