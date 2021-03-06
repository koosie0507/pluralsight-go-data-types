package main

import (
	"datatypes/organization"
	"fmt"
)

type Cmp struct {
	field string
}

func main() {
	p := organization.NewPerson("Gicu", "Piticu", organization.NewEuropeanUnionIdentifier(3, "France"))

	cmp1 := Cmp{field: "a"}
	cmp2 := Cmp{field: "a"}
	fmt.Println(cmp1 == cmp2)

	peopleByName := map[organization.Name][]organization.Person{}
	peopleByName[p.Name] = []organization.Person{p}
	fmt.Println(peopleByName)
}
