package main

import "fmt"
import "goerror"

type Animal struct {
	Name string
	Age  int
}

func main() {
	tmpl := goerror.NewTemplate("32", "{{attr}} is required. {{attr}} = '{{value}}'.")
	data := map[string]interface{}{"attr": "name", "value": ""}
	err := tmpl.Error(data)
	if err.IsNotNil() {
		fmt.Println(err.Error())
	}

	animal := new(Animal)
	if animal.Name == "" {
		newTempl := goerror.NewTemplate("Required", `"{{attr}}" is required. {{attr}} should not be blank.`)
		data := map[string]interface{}{"attr": "animal.Name"}
		err := newTempl.DefaultError(data)
		fmt.Println(err.Error())
	}
}
