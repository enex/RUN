package main

import (
	"./scan"
	"fmt"
	"io/ioutil"
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
	data, err := ioutil.ReadFile("example.run")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := scan.New(string(data))
	r := s.Scan()
	fmt.Println(r)
}
