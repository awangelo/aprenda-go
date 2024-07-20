package main

func main() {
	x := somarRetornos(minhaFunc)
	println(x)
}

func minhaFunc() int {
	return 1213
}

func somarRetornos(f func() int) int {
	return f() + f()
}
