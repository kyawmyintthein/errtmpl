package errtmpl

type ValidatorFunc func(v interface{}) bool
type Validator struct {
	Name            string
	MessageTemplate string
	Data            map[string]interface{}
	Fn              ValidatorFunc
}

func NewValidator(name string, messageTmpl string, fn ValidatorFunc) Validator {
	return Validator{
		Name:            name,
		MessageTemplate: messageTmpl,
		Fn:              fn,
	}
}

func (v *Validator) Validate(i interface{}, data map[string]interface{}) ErrorString {
	var err ErrorString
	if v.Fn(i) {
		tmpl := NewTemplate(v.Name, v.MessageTemplate)
		return tmpl.TError(data)
	}
	return err
}

func (v *Validator) ValidateDefaultError(i interface{}, data map[string]interface{}) error {
	var err error
	if v.Fn(i) {
		tmpl := NewTemplate(v.Name, v.MessageTemplate)
		return tmpl.Error(data)
	}
	return err
}
