package tools

// BenchmarkAbs-4                  1000000000               0.510 ns/op
// BenchmarkAbs2-4                 1000000000               0.460 ns/op

// 绝对值
func Abs(a int64) int64 {
	return (a ^ a>>31) - a>>31
}

// 绝对值
func Abs2(n int64) int64 {
	return (n ^ n>>63) - n>>63
}
