package main

type Cmp struct {
	field string
}

func main() {
	a := Cmp{field: "a"}
	b := Cmp{field: "a"}
	println(a == b)
}
