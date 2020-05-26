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
	eu := organization.NewEuropeanUnionIdentifier("123-45-6789", "Germany")

	println(ssn == eu)
	fmt.Printf("%T\n", ssn)
	fmt.Printf("%T\n", eu)
}
