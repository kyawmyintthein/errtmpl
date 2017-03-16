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
	fmt.Println(tmpl.Error(data))

	animal := new(Animal)
	if animal.Name == "" {
		data := make(map[string]interface{})
		data["attr"] = "animal.Name"
		data["value"] = animal.Name
		fmt.Println(tmpl.Error(data))
	}
}
