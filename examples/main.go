package main

import "fmt"
import "errtmpl"

type Animal struct {
	Name string
	Age  int
}

func main() {
	tmpl := errtmpl.NewTemplate("Required", "{{attr}} is required. {{attr}} = '{{value}}'.")
	data := make(map[string]interface{})
	data["attr"] = "name"
	data["value"] = ""
	customErr := tmpl.TError(data)
	if customErr.IsNil() {
		fmt.Println(customErr)
		fmt.Println(customErr.Error())
	}
	animal := new(Animal)
	if animal.Name == "" {
		data["attr"] = "animal.Name"
		data["value"] = animal.Name
		err := tmpl.Error(data)
		fmt.Println(err)
		fmt.Println(err.Error())
	}
}
