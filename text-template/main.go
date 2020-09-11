package main

// https://www.do1618.com/archives/893/go-template%E5%AD%A6%E4%B9%A0%E6%A0%B7%E4%BE%8B/

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	fmt.Println("Hello, playground")

	const templ = `Here is what they said
{{ range $key, $value := . }}
# {{ $key }} desc: {{ $value.Summary }}
{{ $key }}: {{ $value.Word }}
{{end}}`

	x := map[string]Words{
		"Danny": {
			Summary: "Danny",
			Word:    "The guy really talked about my first time out with you",
		},
		"Doug": {
			Summary: "Doug",
			Word:    "Well he said I'm really amazing, I did not believe at first",
		},
	}

	t, err := template.New("index.html").Parse(templ)
	if err != nil {
		fmt.Println("Could not parse template:", err)
		return

	}
	_ = t.Execute(os.Stdout, x)

}

type Words struct {
	Summary string
	Word    string
}
