package main

import "fmt"

// Strukturen definieren
type Student struct {
	FirstName string
	LastName  string
}

type Class struct {
	Students []Student
}

type Module struct {
	Number  int
	Name    string
	Classes []Class
}

func main() {
	// Beispieldaten erstellen
	class1 := Class{
		Students: []Student{
			{FirstName: "Alice", LastName: "Müller"},
			{FirstName: "Bob", LastName: "Schmidt"},
			{FirstName: "Charlie", LastName: "Meier"},
		},
	}

	class2 := Class{
		Students: []Student{
			{FirstName: "Diana", LastName: "Keller"},
			{FirstName: "Eve", LastName: "Bach"},
			{FirstName: "Frank", LastName: "Weber"},
		},
	}

	modules := []Module{
		{Number: 104, Name: "Programmieren Grundlagen", Classes: []Class{class1, class2}},
		{Number: 117, Name: "Datenbanken Modellierung", Classes: []Class{class1}},
		{Number: 346, Name: "Cloud-Lösungen konzipieren und realisieren", Classes: []Class{class2}},
	}

	// Daten ausgeben
	for _, module := range modules {
		fmt.Printf("Modul %d: %s\n", module.Number, module.Name)
		for i, class := range module.Classes {
			fmt.Printf("  Klasse %d:\n", i+1)
			for _, student := range class.Students {
				fmt.Printf("    - %s %s\n", student.FirstName, student.LastName)
			}
		}
	}
}
