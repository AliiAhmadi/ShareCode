package forms

type errors map[string][]string

// Add() function used for add error messages for a field.
func (es errors) Add(field string, message string) {
	es[field] = append(es[field], message)
}

// Get() function used for getting first error message for a field.
func (es errors) Get(field string) string {
	errorSlice := es[field]

	if len(errorSlice) == 0 {
		return ""
	}

	return errorSlice[0]
}
