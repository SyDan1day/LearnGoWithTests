package iteration

func Repeat(c string) string {
	ans := ""
	for i := 0; i < 5; i++ {
		ans += c
	}
	return ans
}
