function sprintf(s, tab)
	return (s:gsub('($%b{})', function(w) return tab[w:sub(3, -2)] or w end))
end
getmetatable("").__mod = sprintf

local getIndentPreffix = function(str)
  local level = math.huge
  local minPreffix = ""
  local len
  for preffix in str:gmatch("\n( +)") do
    len = #preffix
    if len < level then
      level = len
      minPreffix = preffix
    end
  end
  return minPreffix
end

local unindent = function(str, args)
  str = str:gsub(" +$", ""):gsub("^ +", "") -- remove spaces at start and end
  local preffix = getIndentPreffix(str)
  return sprintf((str:gsub("\n" .. preffix, "\n"):gsub("\n$", "")), args)
end

function str_block(args)
	if type(args) == "string"
	then
		return unindent(args, {})
	else
		return function(str)
			return unindent(str, args or {})
		end
	end
end

function pascal_case(str)
	local result = str:gsub("([^%w])", " ")
	               :gsub("(%u)", " %1")
	               :gsub("^%s+", "")
	result = result:gsub("[%w]+", function(word)
		return word:sub(1, 1):upper() .. word:sub(2):lower()
	end)
	return (result:gsub("[^%w]", ""))
end
