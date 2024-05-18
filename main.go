package main

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

// func addAnimal(animal Animal, cage Cage) Cage {
// 	cage.Animals = append(cage.Animals, animal)
// 	return cage
// }

func (cage Cage) addAnimal(animal Animal) Cage {
	cage.Animals = append(cage.Animals, animal)
	return cage
}

func (zookeeper Zookeeper) CollectAnimals(cages ...Cage) []Animal {
	var allAnimals []Animal
	for _, cage := range cages {
		allAnimals = append(allAnimals, cage.Animals...)
	}

	return allAnimals
}

func main() {
	lion := Animal{Name: "Lion", Voice: "RoarR"}
	monkey := Animal{Name: "Monkey", Voice: "Ooh ooh aah aah"}
	snake := Animal{Name: "Snake", Voice: "Shhhh"}

}
