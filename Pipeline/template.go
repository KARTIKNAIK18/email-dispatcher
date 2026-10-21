package pipeline

import (
	"bytes"
	"text/template"
)

func ExecuteTemplate(email Emails) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}

	var output bytes.Buffer
	if err := t.Execute(&output, email); err != nil {
		return "", err
	}

	return output.String(), nil
}
