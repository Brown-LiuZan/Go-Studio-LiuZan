package temperature

import "fmt"

/*** type definition: Celsius ***/
type Celsius float32

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f °C", c)
}

func (c Celsius) ToKelvin() Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

func (c *Celsius) FromKelvin(k Kelvin) {
	*c = Celsius(k) + AbsoluteZeroC
}

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9.0/5.0 + 32.0)
}

func (c *Celsius) FromFahrenheit(f Fahrenheit) {
	*c = Celsius((f - 32.0) * 5.0 / 9.0)
}

/*** type definition: Kelvin ***/
type Kelvin float32

func (k Kelvin) String() string {
	return fmt.Sprintf("%.2f K", k)
}

/*** type definition: Fahrenheit ***/
type Fahrenheit float32

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f °F", f)
}
