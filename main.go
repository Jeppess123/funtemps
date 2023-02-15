package main

import (
	"flag"
	"fmt"

	"github.com/Jeppess123/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var out string
var funfacts string
var cel float64
var kel float64

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.

/*
   Her er eksempler på hvordan man implementerer parsing av flagg.
   For eksempel, kommando
       funtemps -F 0 -out C
   skal returnere output: 0°F er -17.78°C
*/

func init() {
	flag.Float64Var(&fahr, "F", 0.0, "temperature in degrees Fahrenheit")
	flag.StringVar(&out, "out", "C", "calculate temperature in C - Celsius, F - Fahrenheit, K - Kelvin")
	// flag.StringVar(&funfacts, "funfacts", "sun", ""fun-facts" about sun - The Sun, luna - The Moon, and terra - The Earth")
	flag.Float64Var(&kel, "K", 0.0, "temperature in degrees Kelvin")
	flag.Float64Var(&cel, "C", 0.0, "temperature in degrees Celsius")
}

func main() {

	flag.Parse()

	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.

	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.

	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())

	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen

	if out == "C" && isFlagPassed("F") {
		cel := conv.FarhenheitToCelsius(fahr)
		fmt.Printf("%.2f°F is %.2f°C\n", fahr, cel)
	}

	if out == "F" && isFlagPassed("C") {
		fahr := conv.CelsiusToFahrenheit(cel)
		fmt.Printf("%.2f°C is %.2f°F\n", cel, fahr)
	}

	if out == "K" && isFlagPassed("C") {
		kel := conv.CelsiusToKelvin(cel)
		fmt.Printf("%.2f°C is %.2f°K\n", cel, kel)
	}

	if out == "C" && isFlagPassed("K") {
		cel := conv.KelvinToCelsius(kel)
		fmt.Printf("%.2f°K is %.2f°C\n", kel, cel)
	}

	if out == "F" && isFlagPassed("K") {
		fahr := conv.KelvinToFarhenheit(kel)
		fmt.Printf("%.2f°K is %.2f°F\n", kel, fahr)
	}

	if out == "K" && isFlagPassed("F") {
		kel := conv.FahrenheitToKelvin(fahr)
		fmt.Printf("%.2f°F is %.2f°K\n", fahr, kel)
	}

	if funfacts == "sun" {
		fmt.Println("The Sun is the star at the center of the Solar System. It is a nearly perfect spherical ball of hot plasma, with internal convective motion that generates a magnetic field via a dynamo process. It is by far the most important source of energy for life on Earth.")
	}

	fmt.Println(fahr, out, funfacts)

	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())

	fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk
	if out == "C" && isFlagPassed("F") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
		fmt.Println("0°F er -17.78°C")
	}

}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
