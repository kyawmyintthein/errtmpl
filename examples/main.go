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

	id := 0
	reqErr := errtmpl.Required(id, "user ID")
	fmt.Println(reqErr.Error())

	cusdata := map[string]interface{}{"attr": "id", "value": id}
	tmpl1 := errtmpl.NewTemplate("ID required", "{{attr}} should no be blank. {{attr}} is {{value}}. Please check your data input.")
	cusErr := errtmpl.RequiredWithTemplate(id, cusdata, tmpl1)
	fmt.Println(cusErr.Error())
}
