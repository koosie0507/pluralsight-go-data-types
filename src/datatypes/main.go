package main

type Cmp struct {
	field string
	arr   []string
}

func main() {
	a := Cmp{field: "a", arr: []string{"1", "2"}}
	b := Cmp{field: "a", arr: []string{"1", "2"}}
	println(a == b)
}
