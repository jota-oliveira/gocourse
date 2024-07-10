package main

import (
	"fmt"
	"sort"
)

type UserType struct {
	FirstName string
	LastName string
}

func main() {
	// This is a simple example of how to use a map in Go
	// Maps are like dictionaries in Python

	// Create a map with string as key and int as value
	firstMap := make(map[string]int)

	firstMap["first"] = 1
	firstMap["second"] = 2

	fmt.Println(firstMap)
	fmt.Println(firstMap["first"])

	// We can also create a map using an structure or any other type

	secondMapWithStruct := make(map[int]UserType)

	secondMapWithStruct[1] = UserType{
		FirstName: "John",
		LastName: "Doe",
	}

	fmt.Println(secondMapWithStruct[1].FirstName, secondMapWithStruct[1].LastName)


	// There is no arrays in go, but there is slices, that is a kind of an dynamic array
	// Slices are like lists in Python

	// Create a slice with 3 elements
	firstSlice := make([]int, 3)

	firstSlice[0] = 1
	firstSlice[1] = 2
	firstSlice[2] = 3

	// The difference is that I can append elements to a slice and sort it

	firstSlice = append(firstSlice, 4)
	fmt.Println(firstSlice)

	// Option 1
	sort.Slice(firstSlice, func(i, j int) bool {
		return firstSlice[i] < firstSlice[j]
	})

	// Option 2
	sort.Ints(firstSlice)

	// We can start an slice with elements
	secondSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(secondSlice)

	// It's also possible to get only a part of a slice like this:
	fmt.Println(secondSlice[1:3]) // [2 3]

}
