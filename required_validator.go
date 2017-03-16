package errtmpl

const (
	RequiredDefaultErrorTemplate = "{{attr}} is required."
)

/*
 * name   : Required
 * desc   : Validate empty or null value in golang type such as int, string, float, struct
 * return : Return custom error type TError
 */
func Required(i interface{}, attr string) errorString {
	requiredValidator := NewValidator("Required", RequiredDefaultErrorTemplate, func(v interface{}) bool {
		return isEmpty(v)
	})
	data := map[string]interface{}{"attr": attr}
	return requiredValidator.Validate(i, data)
}

/*
 * name   : RequiredDefaultError
 * desc   : Validate empty or null value in golang type such as int, string, float, struct
 * return : Return custom error type TError
 */
func RequiredDefaultError(i interface{}, attr string) error {
	requiredValidator := NewValidator("Required", RequiredDefaultErrorTemplate, func(v interface{}) bool {
		return isEmpty(v)
	})

	data := map[string]interface{}{"attr": attr}
	return requiredValidator.ValidateDefaultError(i, data)
}

/*
 * name   : Required
 * desc   : Validate empty or null value in golang type such as int, string, float, struct
 * return : Return custom error type TError
 */
func RequiredDefaultErrorWithTemplate(i interface{}, data map[string]interface{}, tmpl errorTemplate) error {
	requiredValidator := NewValidator(tmpl.name, tmpl.layout, func(v interface{}) bool { return isEmpty(v) })
	return requiredValidator.ValidateDefaultError(i, data)
}

/*
 * name   : RequiredWithCustomTemplate
 * desc   : Validate empty or null value in golang type such as int, string, float, struct with custom *          message template
 * return : Return custom error type TError
 */
func RequiredWithTemplate(i interface{}, data map[string]interface{}, tmpl errorTemplate) errorString {
	requiredValidator := NewValidator(tmpl.name, tmpl.layout, func(v interface{}) bool {
		return isEmpty(v)
	})
	return requiredValidator.Validate(i, data)
}
