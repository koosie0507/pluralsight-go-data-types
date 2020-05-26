package organization

import (
	"errors"
	"fmt"
	"strings"
)

type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type Identifiable interface {
	ID() string
}

type Name struct {
	first string
	last string
}

func (n Name) FullName() string {
	return fmt.Sprintf("%s, %s", n.last, n.first)
}

type Person struct {
	Name
	twitterHandler TwitterHandler
}

func NewPerson(firstName, lastName string) Person {
	return Person {Name: Name {first: firstName, last: lastName}}
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