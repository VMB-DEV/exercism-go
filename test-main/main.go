package main

import (
	"meetup"
	"time"
)

func main() {
	// Example: Find teenth Monday in May 2013
	meetup.Day(meetup.Teenth, time.Monday, time.May, 2013)
}
