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
