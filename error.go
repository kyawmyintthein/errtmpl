package goerror

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type (
	ErrorTemplate struct {
		Name   string
		Layout string
	}

	Error struct {
		name    string
		message string
	}
)

func newError(name, message string) Error {
	return Error{name: name, message: message}
}

func new(name, message string) error {
	return &Error{name: name, message: message}
}

func (e *Error) Error() string {
	return fmt.Sprintf("(#%s) %s", e.name, e.message)
}

func (e *Error) IsNil() bool {
	if e.name == "" && e.message == "" {
		return true
	}
	return false
}

func (e *Error) IsNotNil() bool {
	if e.IsNil() {
		return false
	}
	return true
}

func NewTemplate(name string, layout string) ErrorTemplate {
	return ErrorTemplate{Name: name, Layout: layout}
}

func (tmpl *ErrorTemplate) DefaultError(data map[string]interface{}) error {
	return new(tmpl.Name, tmpl.Parse(data))
}

func (tmpl *ErrorTemplate) Error(data map[string]interface{}) Error {
	return newError(tmpl.Name, tmpl.Parse(data))
}

func (tmpl *ErrorTemplate) Parse(data map[string]interface{}) string {
	var layout string = tmpl.Layout
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
