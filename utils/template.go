package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
)

// LoadTemplate loads the email template from a file
func LoadTemplate(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// GenerateEmailBody generates the email body for any template and data
func GenerateEmailBody(templateFile string, data interface{}) (string, error) {
	// Load the template content from file
	templateContent, err := LoadTemplate(fmt.Sprintf("templates/%s.html", templateFile))
	if err != nil {
		return "", err
	}

	// Parse the template content
	tmpl, err := template.New("email").Parse(templateContent)
	if err != nil {
		return "", err
	}

	// Execute the template with the data
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
