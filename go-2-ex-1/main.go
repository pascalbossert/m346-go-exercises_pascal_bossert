package main

import "fmt"

// Strukturen definieren
type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   int
	Month int
	Year  int
}

type Profile struct {
	Name             FullName
	BirthDate        BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	// Initialisierung der Variablen
	var me = Profile{
		Name: FullName{
			FirstName: "Pascal",
			LastName:  "Bossert",
		},
		BirthDate: BirthDate{
			Day:   16,
			Month: 9,
			Year:  1994,
		},
		NumberOfSiblings: 2,
		ZodiacSign:       '\u264D', // Jungfrau
	}

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings++ // Geschwisterzahl um eins erhöhen
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
