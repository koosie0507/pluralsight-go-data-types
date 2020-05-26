package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	return ""
}

type Person struct {
	firstName string
	lastName string
	twitterHandler TwitterHandler
}

func NewPerson(firstName, lastName string) Person {
	return Person {firstName: firstName, lastName: lastName}
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s, %s", p.lastName, p.firstName)
}

func (p *Person) ID() string {
	return "12345"
}

func (p *Person) SetTwitterHandler(twitterHandler TwitterHandler) error {
	if len(twitterHandler) == 0 {
		p.twitterHandler = twitterHandler
	} else if !strings.HasPrefix(string(twitterHandler), "@") {
		return errors.New("twitter handler must start with a @")
	} else if len(twitterHandler) < 2 {
		return errors.New("twitter handler can't be a bare @")
	}
	p.twitterHandler = twitterHandler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}