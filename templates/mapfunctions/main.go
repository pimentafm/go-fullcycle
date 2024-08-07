package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Course struct {
	Name  string
	Hours int
}

type Courses []Course

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{"header.html", "content.html", "footer.html"}
	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))
	err := t.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Python", 30},
		{"Java", 10},
	})
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", nil)
}
