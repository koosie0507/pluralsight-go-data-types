package organization

import (
	"fmt"
	"strconv"
)

type Identifiable interface {
	ID() string
}

type Citizen interface {
	Identifiable
	Country() string
}

type socialSecurityNumber string

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "USA"
}

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

type europeanUnionIdentifier struct {
	id      string
	country string
}

func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}

func (eui europeanUnionIdentifier) Country() string {
	return fmt.Sprintf("EU %s", eui.country)
}

func NewEuropeanUnionIdentifier(id interface{}, country string) Citizen {
	switch v := id.(type) {
	case string:
		return europeanUnionIdentifier{id: v, country: country}
	case int:
		return europeanUnionIdentifier{id: strconv.Itoa(v), country: country}
	case europeanUnionIdentifier:
		return v
	case Person:
		euId, _ := v.Citizen.(europeanUnionIdentifier)
		return euId
	default:
		panic("invalid type for initializing an EU identifier")
	}
}
