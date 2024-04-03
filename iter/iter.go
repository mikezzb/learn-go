package iter

func Repeat(s string, times int) string {
	var repeat string
	for i := 0; i < times; i++ {
		repeat += s
	}
	return repeat
}
