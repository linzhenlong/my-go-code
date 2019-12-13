package jiecheng

func Add(n int) int {
	var res int
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			res = res + i*j
		}
	}
	return res
}
