package main

import "datatypes/organization"

func main() {
	var p = organization.NewPerson("Gicu", "Piticu")
	println(p.ID())
	println(p.FullName())
}
