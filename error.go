package errtmpl

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type (
	errorTemplate struct {
		name   string
		layout string
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

func NewTemplate(name string, layout string) errorTemplate {
	return errorTemplate{name: name, layout: layout}
}

func (tmpl *errorTemplate) Error(data map[string]interface{}) error {
	return new(tmpl.name, tmpl.Parse(data))
}

func (tmpl *errorTemplate) Parse(data map[string]interface{}) string {
	var layout string = tmpl.layout
	for key, val := range data {
		switch t := val.(type) {
		case string:
			layout = strings.Replace(layout, fmt.Sprintf("{{%s}}", key), t, -1)
		case int:
			str := strconv.Itoa(t)
			layout = strings.Replace(layout, fmt.Sprintf("{{%s}}", key), str, -1)
		default:
			s := reflect.ValueOf(t)
			layout = strings.Replace(layout, fmt.Sprintf("{{%s}}", key), fmt.Sprintf("%s", s), -1)
		}
	}
	return layout
}
