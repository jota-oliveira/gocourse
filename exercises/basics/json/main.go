package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Age int `json:"age"`
	HasDog bool `json:"has_dog"`
}

func main() {
	myJson := `
		[
			{
				"first_name": "Clark",
				"last_name": "Kent",
				"age": 29,
				"has_dog": true
			},
			{
				"first_name": "Bruce",
				"last_name": "Wayne",
				"age": 32,
				"has_dog": false
			}
		]
	`

	// It means that we are going to unmarshal a slice of Person, it's the similar of javascript command JSON.parse
	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling JSON", err)
	}

	log.Printf("Unmarshalled: %v", unmarshalled)

	// writing an json from an struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Barry"
	m1.LastName = "Allen"
	m1.Age = 25
	m1.HasDog = true

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Daiana"
	m2.LastName = "Prince"
	m2.Age = 240
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "  ")

	if err != nil {
		log.Println("Error marshalling JSON", err)
	}

	log.Printf("Marshalled: %s", newJson)
}
