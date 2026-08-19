package main

import (
	"fmt"
	"strings"
)

func add(r rune) rune {
	return r + 1
}

func main() {

	fmt.Println(strings.Map(add, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add, "Admix"))    // "Benjy"

}
