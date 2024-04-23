package structs

import "errors"

type Person struct {
	firstName string // private fields
	lastName  string
	age       int
}

// constructor
func NewPerson(fn, ln string, a int) Person {
	return Person{
		firstName: fn,
		lastName:  ln,
		age:       a,
	}
}

// receiver functions
func (p Person) Walk() string {
	return p.firstName + p.lastName + " likes walking"
}

// pointer based receiver function
func (p *Person) SetFirstName(f string) error {
	if len(f) < 3 {
		return errors.New("invalid firstname")
	}
	p.firstName = f
	return nil
}

// getter function
func (p *Person) FirstName() string {
	return p.firstName
}
