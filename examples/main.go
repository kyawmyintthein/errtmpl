package main

import "fmt"
import "goerror"

type Animal struct {
	Name string
	Age  int
}

func main() {
	var errors []error
	animal := new(Animal)
	templ := goerror.NewTemplate("Required", `"{{attr}}" is required. {{attr}} should not be blank.`)

	if animal.Name == "" {
		data := map[string]interface{}{"attr": "animal.Name"}
		err := templ.Error(data)
		fmt.Println(err.Error())
		errors = append(errors, err)
	}

	if animal.Age == 0 {
		templ := goerror.NewTemplate("Required", `"{{attr}}" is required. {{attr}} should not be blank.`)
		data := map[string]interface{}{"attr": "animal.Age"}
		err := templ.Error(data)
		fmt.Println(err.Error())
		errors = append(errors, err)
	}

	httpError := goerror.NewHttpErrors(403, errors)
	fmt.Println(httpError)
}
