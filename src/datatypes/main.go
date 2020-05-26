package main

import (
	"datatypes/organization"
	"fmt"
)

type Cmp struct {
	field string
}

func main() {
	ssn := organization.NewSocialSecurityNumber("123-45-6789")
	eu1 := organization.NewEuropeanUnionIdentifier("123-45-6789", "Germany")
	eu2 := organization.NewEuropeanUnionIdentifier("123-45-6789", "Germany")

	println(eu1 == eu2)
	fmt.Printf("%T\n", ssn)
	fmt.Printf("%T\n", eu1)
	fmt.Printf("%T\n", eu2)
}
