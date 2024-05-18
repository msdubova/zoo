package main

import "fmt"

type Meal string

const (
	Meat   Meal = "Meat"
	Veggie Meal = "Veggie"
)

type Animal struct {
	Name  string
	Voice string
	Meal  Meal
}

type Zookeeper struct {
	Name string
	Age  int
}

type Cage struct {
	Animals []Animal
}

func (cage *Cage) addAnimal(animal Animal) {
	cage.Animals = append(cage.Animals, animal)
}

func (zookeeper Zookeeper) CollectAnimals(animals ...Animal) []Animal {
	var allAnimals []Animal
	for _, animal := range animals {
		allAnimals = append(allAnimals, animal)
	}

	return allAnimals
}

func sortAnimals(animals []Animal) (cage1 Cage, cage2 Cage) {
	for _, animal := range animals {
		if animal.Meal == Meat {
			cage1.addAnimal(animal)
		} else if animal.Meal == Veggie {
			cage2.addAnimal(animal)
		}
	}

	return cage1, cage2
}

func printAnimals(cage Cage) {
	for _, animal := range cage.Animals {
		fmt.Printf("%s  ", animal.Name)
	}
}

func main() {
	lion := Animal{Name: "Lion", Voice: "RoarR", Meal: Meat}
	monkey := Animal{Name: "Monkey", Voice: "Ooh ooh aah aah", Meal: Veggie}
	snake := Animal{Name: "Snake", Voice: "Shhhh", Meal: Veggie}

	zookeeper := Zookeeper{Name: "Alex", Age: 32}
	allAnimals := zookeeper.CollectAnimals(lion, monkey, snake)

	cage1, cage2 := sortAnimals(allAnimals)

	fmt.Printf("Animals collected by the zookeeper: %s, he is %d years old", zookeeper.Name, zookeeper.Age)
	println()
	for _, animal := range allAnimals {
		fmt.Printf("%s says %s\n", animal.Name, animal.Voice)
	}

	fmt.Printf("First cage: ")
	printAnimals(cage1)
	fmt.Printf("\nSecond cage: ")
	printAnimals(cage2)
}
