package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName string
	lastName string
	twitterHandler string
}

func NewPerson(firstName, lastName string) Person {
	return Person {firstName: firstName, lastName: lastName}
}

func (p Person) FullName() string {
	return fmt.Sprintf("%s, %s", p.lastName, p.firstName)
}

func (p Person) ID() string {
	return "12345"
}

func (p Person) SetTwitterHandler(twitterHandler string) error {
	if len(twitterHandler) == 0 {
		p.twitterHandler = twitterHandler
	} else if !strings.HasPrefix(twitterHandler, "@") {
		return errors.New("twitter handler must start with a @")
	} else if len(twitterHandler) < 2 {
		return errors.New("twitter handler can't be a bare @")
	}
	p.twitterHandler = twitterHandler
	return nil
}