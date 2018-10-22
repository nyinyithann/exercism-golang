package reverse

// String creates a reverse string
func String(input string) string {
	rs := []rune(input)
	for x, y := 0, len(rs)-1; x < y; x, y = x+1, y-1 {
		rs[x], rs[y] = rs[y], rs[x]
	}
	return string(rs)
}
