package helpers

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

type HtmlContent struct {
	Url   string
	To    string
	Token string
}

func HtmlRender(filename string, data any) (string, error) {
	cwd, _ := os.Getwd()
	templateDir := path.Join(fmt.Sprintf("%s/templates/%s.html", cwd, filename))

	htmlTemplate, parseErr := template.ParseFiles(templateDir)
	if parseErr != nil {
		defer logrus.Errorf("Parse html template erro: %v", parseErr)
		return "", parseErr
	}

	outputWriter := new(bytes.Buffer)
	executeErr := htmlTemplate.Execute(outputWriter, data)

	if executeErr != nil {
		defer logrus.Errorf("Execute html template error: %v", executeErr)
		return "", executeErr
	}

	return outputWriter.String(), nil
}
