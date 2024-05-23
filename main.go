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
	Name  string
	Age   int
	Cage1 Cage
	Cage2 Cage
}

type Cage struct {
	Animals []Animal
}

func (cage *Cage) addAnimal(animal Animal) {
	cage.Animals = append(cage.Animals, animal)
}

func (zookeeper *Zookeeper) CollectAnimals(animals ...Animal) {

	for _, animal := range animals {
		if animal.Meal == Meat {
			zookeeper.Cage1.addAnimal(animal)
		} else if animal.Meal == Veggie {
			zookeeper.Cage2.addAnimal(animal)
		}
	}

}

func (zookeeper Zookeeper) printAnimals() {
	fmt.Printf("Animals collected by the zookeeper: %s, he is %d years old\n", zookeeper.Name, zookeeper.Age)

	for _, animal := range append(zookeeper.Cage1.Animals, zookeeper.Cage2.Animals...) {
		fmt.Printf("%s says %s\n", animal.Name, animal.Voice)
	}

	fmt.Printf("First cage: ")
	for _, animal := range zookeeper.Cage1.Animals {
		fmt.Printf("%s  ", animal.Name)
	}

	fmt.Printf("\nSecond cage: ")
	for _, animal := range zookeeper.Cage2.Animals {
		fmt.Printf("%s  ", animal.Name)
	}
}

func main() {
	lion := Animal{Name: "Lion", Voice: "RoarR", Meal: Meat}
	monkey := Animal{Name: "Monkey", Voice: "Ooh ooh aah aah", Meal: Veggie}
	snake := Animal{Name: "Snake", Voice: "Shhhh", Meal: Veggie}
	zookeeper := Zookeeper{Name: "Alex", Age: 32}

	zookeeper.CollectAnimals(lion, monkey, snake)
	zookeeper.printAnimals()
}
