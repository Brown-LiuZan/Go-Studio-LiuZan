package main

import (
	"./temperature"
	"fmt"
)

func temperatureTest() {
	fmt.Println("*** test of temperature package: begin ***")

	var c temperature.Celsius = 25.0
	fmt.Println("Absolute zero in Celsius:", temperature.AbsoluteZeroC)
	fmt.Println("Freezing point in Celsius:", temperature.FreezingC)
	fmt.Println("Boiling point in Celsius:", temperature.BoilingC)
	fmt.Println("Room temperature in Celsius:", c)
	fmt.Println("Room temperature in Fahrenheit:", c.ToFahrenheit())
	fmt.Println("Room temperature in Kelvin:", c.ToKelvin())

	var k temperature.Kelvin = temperature.AbsoluteZeroC.ToKelvin()
	fmt.Println("Absolute zero in Kelvin:", k)
	c.FromKelvin(k)
	fmt.Println("Absolute zeor in Celsius:", c)

	var f temperature.Fahrenheit = temperature.BoilingC.ToFahrenheit()
	fmt.Println("Boiling point in Fahrenheit:", f)
	c.FromFahrenheit(f)
	fmt.Println("Boiling point in Celsius:", c)

	fmt.Println("*** test of temperature package: end ***")
	fmt.Println()
}

func main() {
	temperatureTest()
}
