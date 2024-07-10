package main

import "log"


func main() {
	// For loop
	for i := 0; i < 5; i++ {
		if i == 3 {
			break;
		} else if i == 2 {
			log.Println("Almost there")
		} else {
			log.Println(i)
		}
	}

	// Bellow there is a list of animals, let's iterate over it
	animals := []string{"dog", "cat", "bird", "fish"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	// we still can use an underscore to ignore the index if we will not need it
	for _, animal := range animals {
		log.Println(animal)
	}

	// We can range over maps
	animalsMap := make(map[string]string)
	animalsMap["cat"] = "meow"
	animalsMap["dog"] = "woof"

	for animal, sound := range animalsMap {
		log.Println(animal, sound)
	}

	// It is possible to iterate over a string
	// In this case, each letter will be printed,
	// same sintax, so I won't do it.
}

// Equivalent code in JavaScript
// for (const [animal, sound] of Object.entries(animalsMap)) {
// 	console.log(animal, sound);
// }
