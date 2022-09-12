package helpers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
)

type HtmlContent struct {
	Url   string
	From  string
	Token string
}

func HtmlRender(rw http.ResponseWriter, filename string, data any) error {
	cwd, _ := os.Getwd()
	templateDir := path.Join(fmt.Sprintf("%s/templates/%s.html", cwd, filename))

	htmlTemplate, _ := template.ParseFiles(templateDir)
	err := htmlTemplate.Execute(rw, data)

	return err
}
