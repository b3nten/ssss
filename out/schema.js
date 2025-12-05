export class Bar {
	static get TypeID() { return 32026 }

	baz


	deserialize(bytes) {
		const b = new ByteReader(bytes)
		Bar_deserialize(b, this)
		return this;
	}

	serialize() {
		const w = new ByteWriter;
		Bar_serialize(w, this)
		return w.bytes()
	}
}

function Bar_serialize(b, s) {
	b.write_uint16(32026)
	b.write_uint32(0);
	const structStart = b.length;
	if (s.baz !== undefined) {
		b.write_uint16(1);
		b.write_int32(s.baz)

	}

	b.set_uint32(structStart - 4, b.length - structStart);
}

function Bar_deserialize(br) {
	const typeId = br.read_uint32()
	const length = br.read_uint32()
	if (length > br.length || length > Math.MAX_SAFE_INTEGER) {
		throw new Error("Invalid struct length");
	}
	const seenFields = new Set;
	const startLen = br.length;
	for (; br.length - startLen < length;) {
		const fieldId = br.read_uint16();
		if (seenFields.has(fieldId)) {
			throw new Error("Duplicate field ID " + fieldId + " in struct Bar");
		}
		if (fieldId > 1) {
			return;
		}
		seenFields.add(fieldId);
		switch (fieldId) {
			case 1:
				s.baz = br.read_int32();

				break;

		}
	}
}


export class Foo {
	static get TypeID() { return 32471 }

	lst2
	m
	x
	f
	b
	bar
	lst


	deserialize(bytes) {
		const b = new ByteReader(bytes)
		Foo_deserialize(b, this)
		return this;
	}

	serialize() {
		const w = new ByteWriter;
		Foo_serialize(w, this)
		return w.bytes()
	}
}

function Foo_serialize(b, s) {
	b.write_uint16(32471)
	b.write_uint32(0);
	const structStart = b.length;
	if (s.lst2 !== undefined) {
		b.write_uint16(5);
		b.write_uint32(0);
		const listStart0 = b.length;
		for (const item0 of s.lst2) {
			b.write_uint32(0);
			const listStart1 = b.length;
			for (const item1 of item0) {
				b.write_string(item1)

			}
			b.set_uint32(listStart1 - 4, b.length - listStart1);

		}
		b.set_uint32(listStart0 - 4, b.length - listStart0);

	}
	if (s.m !== undefined) {
		b.write_uint16(6);
		b.write_uint32(0);
		const mapStart0 = b.length;
		for (const [key0, value0] of Object.entries(s.m)) {
			b.write_int32(key0)

			b.write_string(value0)

		}
		b.set_uint32(mapStart0 - 4, b.length - mapStart0);

	}
	if (s.x !== undefined) {
		b.write_uint16(7);
		b.write_uint32(0);
		const mapStart0 = b.length;
		for (const [key0, value0] of Object.entries(s.x)) {
			b.write_string(key0)

			b.write_uint32(0);
			const mapStart1 = b.length;
			for (const [key1, value1] of Object.entries(value0)) {
				b.write_int32(key1)

				b.write_bool(value1)

			}
			b.set_uint32(mapStart1 - 4, b.length - mapStart1);

		}
		b.set_uint32(mapStart0 - 4, b.length - mapStart0);

	}
	if (s.f !== undefined) {
		b.write_uint16(1);
		b.write_f32(s.f)

	}
	if (s.b !== undefined) {
		b.write_uint16(2);
		b.write_bool(s.b)

	}
	if (s.bar !== undefined) {
		b.write_uint16(3);
		Bar_serialize(b, s.bar)

	}
	if (s.lst !== undefined) {
		b.write_uint16(4);
		b.write_uint32(0);
		const listStart0 = b.length;
		for (const item0 of s.lst) {
			b.write_int32(item0)

		}
		b.set_uint32(listStart0 - 4, b.length - listStart0);

	}

	b.set_uint32(structStart - 4, b.length - structStart);
}

function Foo_deserialize(br) {
	const typeId = br.read_uint32()
	const length = br.read_uint32()
	if (length > br.length || length > Math.MAX_SAFE_INTEGER) {
		throw new Error("Invalid struct length");
	}
	const seenFields = new Set;
	const startLen = br.length;
	for (; br.length - startLen < length;) {
		const fieldId = br.read_uint16();
		if (seenFields.has(fieldId)) {
			throw new Error("Duplicate field ID " + fieldId + " in struct Foo");
		}
		if (fieldId > 7) {
			return;
		}
		seenFields.add(fieldId);
		switch (fieldId) {
			case 5:
				{
					const listLength0 = br.read_uint32();
					const listStart0 = br.position;
					s.lst2 = [];
					for (; br.position - listStart0 < listLength0;) {
						let item0;
						{
							const listLength1 = br.read_uint32();
							const listStart1 = br.position;
							item0 = [];
							for (; br.position - listStart1 < listLength1;) {
								let item1;
								item1 = br.read_string();

								item0.push(item1);
							}
						}

						s.lst2.push(item0);
					}
				}

				break;
			case 6:
				{
					const mapLength0 = br.read_uint32();
					const mapStart0 = br.position;
					s.m = {};
					for (; br.position - mapStart0 < mapLength0;) {
						let key0;
						key0 = br.read_int32();

						let value0;
						value0 = br.read_string();

						s.m[key0] = value0;
					}
				}

				break;
			case 7:
				{
					const mapLength0 = br.read_uint32();
					const mapStart0 = br.position;
					s.x = {};
					for (; br.position - mapStart0 < mapLength0;) {
						let key0;
						key0 = br.read_string();

						let value0;
						{
							const mapLength1 = br.read_uint32();
							const mapStart1 = br.position;
							value0 = {};
							for (; br.position - mapStart1 < mapLength1;) {
								let key1;
								key1 = br.read_int32();

								let value1;
								value1 = br.read_bool();

								value0[key1] = value1;
							}
						}

						s.x[key0] = value0;
					}
				}

				break;
			case 1:
				s.f = br.read_f32();

				break;
			case 2:
				s.b = br.read_bool();

				break;
			case 3:
				s.bar = new Bar();
				Bar_deserialize(br, s.bar);

				break;
			case 4:
				{
					const listLength0 = br.read_uint32();
					const listStart0 = br.position;
					s.lst = [];
					for (; br.position - listStart0 < listLength0;) {
						let item0;
						item0 = br.read_int32();

						s.lst.push(item0);
					}
				}

				break;

		}
	}
}




class ByteWriter {
	get length() { return this.len; }

	encoder = new TextEncoder();
	buffer = new ArrayBuffer(0xFFF)
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

	write_f32(value) {
		ByteWriter._tmp = this.length;
		this.resize(this.len + 4);
		this.dview.setFloat32(ByteWriter._tmp, value, true);
	}

	write_f64(value) {
		ByteWriter._tmp = this.length;
		this.resize(this.len + 8);
		this.dview.setFloat64(ByteWriter._ByteWriter._ByteWriter._tmp, value, true);
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
		write_uint32(0, this);
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
				write_uint8(codePoint, this);
			} else {
				if (codePoint < 0x800) {
					write_uint8(((codePoint >> 6) & 0x1F) | 0xC0, this);
				} else {
					if (codePoint < 0x10000) {
						write_uint8(((codePoint >> 12) & 0x0F) | 0xE0, this);
					} else {
						write_uint8(((codePoint >> 18) & 0x07) | 0xF0, this);
						write_uint8(((codePoint >> 12) & 0x3F) | 0x80, this);
					}
					write_uint8(((codePoint >> 6) & 0x3F) | 0x80, this);
				}
				write_uint8((codePoint & 0x3F) | 0x80, this);
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
		this.view = new DataView(this.buffer)
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
		if (length > 300) {
			const bytes = new Uint8Array(this.buffer, this.position, length);
			if (!ByteReader.decoder) ByteReader.decoder = new TextDecoder;
			struct[field] = ByteReader.decoder.decode(bytes);
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
		}
		return result;
	}
}
