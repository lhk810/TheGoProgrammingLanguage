package main

import "fmt"

func main() {
	a := [...]int{3: 1, 3 + 4: 2}
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}
