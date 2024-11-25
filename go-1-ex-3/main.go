package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed für Zufallszahlen
	var eyes = rand.Intn(6) + 1
	var when = time.Now()

	// Ausgabe der Augenzahl auf os.Stdout
	fmt.Fprintln(os.Stdout, "The dice shows", eyes, "eyes")

	// Ausgabe der Zeit auf os.Stderr
	fmt.Fprintln(os.Stderr, "The dice was rolled at", when)

	// Hinweise für die Aufrufe:
	// Um die Augenzahl in eyes.txt und die Zeit in dice.log zu speichern:
	// go run ex3/main.go > eyes.txt 2> dice.log
}
