//Package weather provides  the weather forecast
//for a specific city.
package weather

//CurrentCondition represents a certain weather condition
//as a string value.
var CurrentCondition string

//CurrentLocation represents a certain location
//as a string value.
var CurrentLocation string

//Forecast returns a concatenated string value as sentence
//that shows the weather condition for a given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
