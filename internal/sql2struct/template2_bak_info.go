package sql2struct

import (
	"os"
	"strings"
	"text/template"
)

func main() {
	funcMap := template.FuncMap{
		"strupper": upper,
	}
	t1 := template.New("test1")
	tmpl, err := t1.Funcs(funcMap).Parse(`{{strupper .}}`)
	if err != nil {
		panic(err)
	}
	_ = tmpl.Execute(os.Stdout, "go programming")
}

func upper(str string) string {
	return strings.ToUpper(str)
}