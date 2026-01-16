package meetup

import (
	"fmt"
	"time"
)

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	fmt.Println("Inputs:", wSched, wDay, month, year)
	// Your logic here
	return 0
}
