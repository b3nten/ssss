local function uuid()
	local template = 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'
	return string.gsub(template, '[xy]', function(c)
		local v = (c == 'x') and math.random(0, 0xf) or math.random(8, 0xb)
		return string.format('%x', v)
	end)
end

local function make_type(base_type)
	local result = {
		metadata = {},
	}
	for k, v in pairs(base_type) do
		result[k] = v
	end
	return setmetatable(result, {
		__call = function(self, metadata)
			for k, v in pairs(metadata) do
				self.metadata[k] = v
			end
			return result
		end,
		__index = function(self, key)
			self.metadata.id = key
			return self
		end,
	})
end

bool = make_type({ type = "primitive", name = "bool" })
int8 = make_type({ type = "primitive", name = "int8" })
uint8 = make_type({ type = "primitive", name = "uint8" })
int16 = make_type({ type = "primitive", name = "int16" })
uint16 = make_type({ type = "primitive", name = "uint16" })
int32 = make_type({ type = "primitive", name = "int32" })
uint32 = make_type({ type = "primitive", name = "uint32" })
int64 = make_type({ type = "primitive", name = "int64" })
uint64 = make_type({ type = "primitive", name = "uint64" })
f32 = make_type({ type = "primitive", name = "f32" })
f64 = make_type({ type = "primitive", name = "f64" })
str = make_type({ type = "primitive", name = "string" })

function list(type)
	return make_type({ type = "list", of = type })
end

function map(keyType, valueType)
	return make_type({ type = "map", from = keyType, to = valueType })
end

function struct(fields)
	local s = {
		type = "struct",
		fields = fields,
		uuid = uuid(),
		metadata = {},
	}
	return setmetatable(s, {
		__call = function(self, metadata)
			for k, v in pairs(metadata) do
				self.metadata[k] = v
			end
			return self
		end,
		__index = function(self, key)
			self.metadata.id = key
			return self
		end,
	})
end
