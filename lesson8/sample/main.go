package main

import "html/template"

func main() {
	tpl := template.Must(template.New("somthi").Parse("here is the text in the template."))
}
