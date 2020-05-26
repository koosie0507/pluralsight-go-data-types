package main

import "datatypes/organization"

func main() {
	var p = organization.NewPerson("Gicu", "Piticu", organization.NewEuropeanUnionIdentifier("123-45-6789", "Germany"))
	println(p.ID(), p.FullName(), p.Country())
	err := p.SetTwitterHandler("@abc")
	if err != nil {
		println(err.Error())
	} else {
		println(p.TwitterHandler())
		println(p.TwitterHandler().RedirectUrl())
	}
}
