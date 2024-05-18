package main

type Animal struct {
	Name  string
	Voice string
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
