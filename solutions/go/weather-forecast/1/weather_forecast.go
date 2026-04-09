// Package weather proporciona herramientas para el pronóstico del clima.
package weather


var (
    // CurrentCondition representa el estado actual del clima.
	CurrentCondition string
    // CurrentLocation indica la ciudad o región del pronóstico.
	CurrentLocation  string
)

// Forecast devuelve un resumen del clima para una ubicación específica.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
