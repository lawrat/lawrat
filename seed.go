package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	age := 17
	if age > 18 {
		fmt.Println("je suis majeur")
	} else if age == 18 {
		fmt.Println("je viens tout juste d'etre adulte")
	} else {
		fmt.Println("je suis mineur ")
	}

}
