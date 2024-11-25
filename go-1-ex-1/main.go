package main

import "fmt"

func main() {
	// Deklaration und Initialisierung der Variablen
	firstName := "Pascal"  // Vorname
	lastName := "Bossert"  // Nachname
	dayOfBirth := 16       // Geburtsdatum (Tag)
	monthOfBirth := 9      // Geburtsdatum (Monat)
	yearOfBirth := 1994    // Geburtsdatum (Jahr)
	numberOfSiblings := 2  // Anzahl Geschwister
	heightInMeters := 1.64 // Größe in Metern
	zodiacSign := '\u264D' // Sternzeichen: Jungfrau (Unicode)

	// Ausgabe der persönlichen Angaben
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
