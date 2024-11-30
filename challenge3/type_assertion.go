package main

import (
	"fmt"
)

type Developer struct {
	Name string
	Age  int
}

func GetDeveloper(name interface{}, age interface{}) Developer {
	newDeveloper := Developer{name.(string), age.(int)}
	return newDeveloper
}

func main() {

	var name interface{} = "Elliot"
	var age interface{} = 26

	dynamicDev := GetDeveloper(name, age)
	fmt.Printf("This is variable name: %v with the type: %T \n", dynamicDev.Name, dynamicDev.Name)
	fmt.Printf("This is variable age: %v with the type: %T \n", dynamicDev.Age, dynamicDev.Age)

}
