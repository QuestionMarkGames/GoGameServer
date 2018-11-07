package main

import (
	"fmt"
	"time"
)

func main() {
	age := 40
	for true {
		fmt.Printf("age=%d\n", age)
		fmt.Println("game starting")
		time.Sleep(time.Hour * 24 * 365)
		fmt.Println("game end\n")
		time.Sleep(time.Hour * 24)
		fmt.Println("game restarting")

		age++
	}
}
