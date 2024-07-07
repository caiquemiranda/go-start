package main

import "fmt"

const ebulicaoFahrenheit = 212.0

func main() {

	temperaturaFahrenheit := ebulicaoFahrenheit
	temperaturaCelsius := (temperaturaFahrenheit - 32) * 5 / 9

	temperaturaFahrenheit = temperaturaCelsius*1.8 + 32

	fmt.Printf("Ponto de ebulição da água em Fahrenheit: %.2f\n", ebulicaoFahrenheit)
	fmt.Printf("Ponto de ebulição da água em Celsius: %.2f\n", temperaturaCelsius)
}
