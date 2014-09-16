package main

import (
	"./Scanner"
	"fmt"
)

/*
Die sprace setzt nur auf ein Prinzip auf, Pattern matching.

 - Kalmmern und Einrückungen werden als gruppierungen gesehen
 - Funktionen werden als Pattern definiert. Dieße sind ähnlich regulären
   Ausdrücken und ermöglichen das Defieren eines eigenen syntay. Die reihenfolge
   der Definitionen reflektiert die reihenfolge der match-vorgänge.

[Pattern] = [body]
*/

func main() {
	f := &File{}
	f.Read(`
test = 
	Haus = "test"
	Auto = 
		if Haus == "test"
			"ja"
			"nein"
	Baum = "nicht da"
`)
}
