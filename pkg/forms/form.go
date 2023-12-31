package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](" +
	"?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

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

func (form *Form) PermittedValues(field string, opts ...string) {
	value := form.Get(field)

	if value == "" {
		return
	}

	for _, opt := range opts {
		if value == opt {
			return
		}
	}

	form.Errors.Add(field, "This field is invalid")
}

func (form *Form) Valid() bool {
	return len(form.Errors) == 0
}

func (form *Form) MinLength(field string, length int) {
	value := form.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) < length {
		form.Errors.Add(field, fmt.Sprintf("This field is too short. (Min length = %d)", length))
	}
}

func (form *Form) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := form.Get(field)
	if value == "" {
		return
	}
	if !pattern.MatchString(value) {
		form.Errors.Add(field, "This field is invalid")
	}
}
