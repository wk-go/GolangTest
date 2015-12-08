// text/template test
package main

import "text/template"
import "os"

func main() {

	type Inventory struct {
		Material string
		Count    int
	}
	sweaters := Inventory{"axe", 0}
	html := `
    ("test").Parse("{{.Count}} items are made of {{.Material}})"
    {{$a := .Count}}
    {{$b := 17}}
    {{$c := 18}}

    {{if eq  .Count $b}}
    oo
    {{else}}
    xx
    {{end}}

    `
	tmpl, err := template.New("test").Parse(html)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

}