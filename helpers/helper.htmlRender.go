package helpers

import (
	"fmt"
	"html/template"
	"os"
	"path"
)

type HtmlContent struct {
	Url   string
	From  string
	Token string
}

func HtmlRender(filename string, data any) error {
	cwd, _ := os.Getwd()
	templateDir := path.Join(fmt.Sprintf("%s/templates/%s.html", cwd, filename))

	htmlTemplate, _ := template.ParseFiles(templateDir)
	err := htmlTemplate.ExecuteTemplate(os.Stdout, filename, data)

	return err
}
