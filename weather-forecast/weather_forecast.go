// Package weather provides tools to allow you to report the weather forecast in a city.
package weather

// CurrentCondition is a string describing the current weather conditions.
var CurrentCondition string

// CurrentLocation is a string describing the city where weather conditions are being reported.
var CurrentLocation string

// Forecast takes:
// city - The name of the city where the forecast is being reported.
// condition - The current weather conditions in city.
// It will return a string of the weather forecast showing the current weather condition in city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
