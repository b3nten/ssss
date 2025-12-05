
bar = struct {
	baz = int32 [1],
}

foo = struct {
	-- bar = map(int32, str) [1]
	f = f32[1],
	b = bool [2],
	bar = bar [3],
	lst = list(int32) [4],
	lst2 = list(list(str)) [5],
	m = map(int32, str)[6],
	x = map(str, map(int32, bool)) [7],
}
