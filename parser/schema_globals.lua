local function make_type(base_type)
	return setmetatable(base_type, {
		__call = function(self, metadata)
			local new = { metadata = {} }
			for k, v in pairs(self) do
				new[k] = v
			end
			for k, v in pairs(metadata) do
				new.metadata[k] = v
			end
			return new
		end,
		__index = function(self, key)
			local new = { metadata = {} }
			for k, v in pairs(self) do
				new[k] = v
			end
			new.metadata.id = key
			return new
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
		metadata = { name = nil },
	}
	return setmetatable(s, {
		__call = function(self, metadata)
			for k, v in pairs(metadata) do
				self.metadata[k] = v
			end
			return self
		end,
		__index = function(self, key)
			local new = {}
			for k, v in pairs(self) do
				new[k] = v
			end
			new.metadata.id = key
			return new
		end,
	})
end

__Structs = {}

setmetatable(_G, {
    __index = function(t, key)
        if __Structs[key] then
						return __Structs[key]
				else
						return rawget(t, key)
				end
    end,
    __newindex = function(t, key, val)
        if type(val) == "table" and val.type == "struct" then
            if val.metadata.name == nil then
                val.metadata.name = key
            end
            __Structs[key] = val
				else
						rawset(t, key, val)
        end
    end
})
