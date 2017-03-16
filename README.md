# goerror
Create error type with custom message template.

## Usage
```     
template := goerror.NewTemplate("Required", `"{{attr}}" is required. {{attr}} should not be blank.`)
data := map[string]interface{}{"attr": "user.Name"}
err := newTempl.DefaultError(data)
```