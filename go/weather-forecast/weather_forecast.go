//Package weather provides  the weather forecast.
package weather
//CurrentCondition represents the condition.
var CurrentCondition string
//CurrentLocation represents the location.
var CurrentLocation string
//Forecast returns a string value with currentLocation and CurrentCondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
