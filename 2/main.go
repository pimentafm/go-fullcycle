package main

const a = "Hello, World!"

var (
	b bool
	c int
	d string
	e float64
	f complex128
)

func main() {
	b = true

	g := 42
	h := 22.45

	println(a, b, c, d, e, f, g, h)
}
