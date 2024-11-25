package main

import "fmt"

func main() {
	// Map mit Modulinformationen erstellen
	modules := map[int]string{
		104: "Programmieren Grundlagen",
		117: "Datenbanken Modellierung",
		346: "Cloud-Lösungen konzipieren und realisieren",
	}

	// Vorhandene Informationen ausgeben
	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// Änderungen an der Map vornehmen
	delete(modules, 117)                             // Modul 117 entfernen
	modules[500] = "Künstliche Intelligenz anwenden" // Neues Modul hinzufügen
	modules[104] = "Programmieren fortgeschritten"   // Modul 104 ersetzen

	// Map ausgeben
	fmt.Println(modules)
}
