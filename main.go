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
