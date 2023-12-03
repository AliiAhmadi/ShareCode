package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (form *Form) Required(fields ...string) {
	for _, field := range fields {
		value := form.Get(field)

		if strings.TrimSpace(value) == "" {
			form.Errors.Add(field, "This field can not be blank.")
		}
	}
}

func (form *Form) MaxLength(field string, length int) {
	value := form.Get(field)

	if utf8.RuneCountInString(value) > length {
		form.Errors.Add(field, fmt.Sprintf("This field is too long. (Max length = %d)", length))
	}
}

// func (form *Form)
