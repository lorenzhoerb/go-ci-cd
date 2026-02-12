package main

import (
	"fmt"

	"github.com/lorenzhoerb/go-ci-ci/calc"
)

func main() {
	for {
		var cmd string
		var a, b int

		if _, err := fmt.Scanf("%s %d %d\n", &cmd, &a, &b); err != nil {
			fmt.Println("[ERROR]: unknown formant")
			continue
		}

		switch cmd {
		case "add":
			fmt.Println(calc.Add(a, b))
		case "sub":
			fmt.Println(calc.Sub(a, b))
		default:
			fmt.Printf("[ERROR]: unknown command: %s\n", cmd)
		}
	}
}
