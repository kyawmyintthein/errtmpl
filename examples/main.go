package main

import "fmt"
import "goerror"

type Animal struct {
	Name string
	Age  int
}
func main() {
	animal := new(Animal)
	if animal.Name == ""{
		templ := goerror.NewTemplate("Required", `"{{attr}}" is required. {{attr}} should not be blank.`)
		data := map[string]interface{}{"attr": "animal.Name"}
		err := templ.Error(data)
		fmt.Println(err.Error())
	}
}
