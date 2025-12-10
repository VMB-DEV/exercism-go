package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(a string) string
}

func SayHello(a string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(a))
}

type Italian struct{}

func (Italian) LanguageName() string {
	return "Italian"
}
func (Italian) Greet(a string) string {
	return fmt.Sprintf("Ciao %s!", a)
}

type Portuguese struct{}

func (Portuguese) LanguageName() string {
	return "Portuguese"
}
func (Portuguese) Greet(a string) string {
	return fmt.Sprintf("Olá %s!", a)
}
