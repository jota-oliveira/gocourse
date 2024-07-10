package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, World!"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to", i)

	whatWasSaid := saySomething()

	fmt.Println(whatWasSaid)
}

func saySomething() string {
	return "Something"
}
