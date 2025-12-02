local function uuid()
    local template ='xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'
    return string.gsub(template, '[xy]', function (c)
        local v = (c == 'x') and math.random(0, 0xf) or math.random(8, 0xb)
        return string.format('%x', v)
    end)
end

-- Helper to make a type callable for metadata
local function make_type(base_type)
    return setmetatable(base_type, {
        __call = function(self, metadata)
            local result = {}
            for k, v in pairs(self) do
                result[k] = v
            end
            result.metadata = metadata
            return result
        end
    })
end

-- Primitives
bool = make_type({ type = "primitive", name = "bool" })
int8 = make_type({ type = "primitive", name = "int8" })
uint8 = make_type({ type = "primitive", name = "uint8" })
int16 = make_type({ type = "primitive", name = "int16" })
uint16 = make_type({ type = "primitive", name = "uint16" })
int32 = make_type({ type = "primitive", name = "int32" })
uint32 = make_type({ type = "primitive", name = "uint32" })
int64 = make_type({ type = "primitive", name = "int64" })
uint64 = make_type({ type = "primitive", name = "uint64" })
str = make_type({ type = "primitive", name = "string" })

function struct(fields)
    local s = {
        type = "struct",
        fields = fields,
        uuid = uuid(),
    }
    return setmetatable(s, {
        __call = function(self, metadata)
            local result = {}
            for k, v in pairs(self) do
                result[k] = v
            end
            result.metadata = metadata
            return result
        end
    })
end

function list(type)
    return make_type({type = "list", of = type})
end

function map(keyType, valueType)
    return make_type({type = "map", key = keyType, value = valueType})
end
