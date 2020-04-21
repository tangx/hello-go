package tmpl

import (
	htmltemplate "html/template"
	"os"
	"text/template"
)

var (
	txt = `WARNING: 11 < 18 `
)

type Person struct {
	Name string
	Age  int
}

func Gen() {

	tmpl, _ := template.New("test").Parse(txt)
	tmpl.Execute(os.Stdout, nil) // WARNING: 11 < 18

	htmltmpl, _ := htmltemplate.New("test").Parse(txt)
	htmltmpl.Execute(os.Stdout, nil) // WARNING: 11 &lt; 18

}
