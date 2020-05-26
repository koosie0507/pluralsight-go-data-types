package main

import "datatypes/organization"

func main() {
	var p = organization.NewPerson("Gicu", "Piticu")
	println(p.ID())
	println(p.FullName())
	err := p.SetTwitterHandler("@abc")
	if err != nil {
		println(err.Error())
	} else {
		println(p.TwitterHandler())
		println(p.TwitterHandler().RedirectUrl())
	}
}
