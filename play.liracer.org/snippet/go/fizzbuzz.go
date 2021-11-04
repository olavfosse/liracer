//go:build ignore
package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		fizzedOrBuzzed := false
		if i%3 == 0 {
			fizzedOrBuzzed = true
			fmt.Print("Fizz")
		}
		if i%5 == 0 {
			fizzedOrBuzzed = true
			fmt.Print("Buzz")
		}
		if !fizzedOrBuzzed {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
