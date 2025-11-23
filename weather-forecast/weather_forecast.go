// Package weather
// explicit comment.
package weather

var (
	// CurrentCondition blablabla.
	CurrentCondition string
	// CurrentLocation blablabla.
	CurrentLocation string
)

// Forecast blablabla.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
