function sprintf(s, tab)
	return (s:gsub('($%b{})', function(w) return tab[w:sub(3, -2)] or w end))
end
getmetatable("").__mod = sprintf

function unindent(str, args)
	str = sprintf(str, args)
	str = str:gsub("^%s*\n", ""):gsub("\n%s*$", "")
	local min_indent = nil
	local min_indent_len = math.huge
	for line in str:gmatch("[^\n]+") do
		local indent = line:match("^%s*")
		local content = line:gsub("^%s+", "")
		if #content > 0 then         -- ignore empty lines
			if #indent < min_indent_len then
				min_indent_len = #indent
				min_indent = indent
			end
		end
	end
	if min_indent and #min_indent > 0 then
		local pattern = "^" .. min_indent:gsub("([%^%$%(%)%%%.%[%]%*%+%-%?])", "%%%1")
		str = str:gsub("([^\n]+)", function(line)
			return line:gsub(pattern, "", 1)
		end)
	end
	if args.tabs and args.tabs > 0 then
		local indent_str = string.rep("\t", args.tabs)
		str = str:gsub("([^\n]+)", indent_str .. "%1")
	end
	return str
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
	local result = str:gsub("[%w]+", function(word)
		return word:sub(1, 1):upper() .. word:sub(2):lower()
	end)
	return (result:gsub("[^%w]", ""))
end
