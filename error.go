package errtmpl

import "strings"
import "fmt"

type (
	errorTemplate struct {
		name   string
		layout string
		data   map[string]interface{}
	}

	errorString struct {
		name    string
		message string
	}
)

func new(name, message string) error {
	return &errorString{name: name, message: message}
}

func (e *errorString) Error() string {
	return fmt.Sprintf("(#%s) %s", e.name, e.message)
}

func NewTemplate(name string, layout string, data map[string]interface{}) errorTemplate {
	return &errorTemplate{name: name, layout: layout, data: data}
}

func (tmpl *errorTemplate) Error() {
	return new(tmpl.name, tmpl.Parse())
}

func (tmpl *errorTemplate) Parse() string {
	var layout string = tmpl.layout
	for key, val := range tmpl.data {
		layout = strings.Replace(layout, key, val, -1)
	}
	return layout
}
