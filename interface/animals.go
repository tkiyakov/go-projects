package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println("Gay")
}

func (c Cat) Speak() {
	fmt.Println("Maw")
}

func MakeSound(s Speaker) {
	s.Speak()
}



func main() { 
	dog := Dog{Name: "Шарик"}
	cat := Cat{Name: "Мурзик"}

	fmt.Println(dog.Name)
	MakeSound(dog)
	fmt.Println(cat.Name)
	MakeSound(cat)
}