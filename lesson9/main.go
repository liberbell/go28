package main

import (
	"text/template"
	"time"
)

var tpl *template.Template

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}
