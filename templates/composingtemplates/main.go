package main

import (
	"net/http"
	"os"
	"text/template"
)

type Course struct {
	Name  string
	Hours int
}

type Courses []Course

func main() {
	templates := []string{"header.html", "content.html", "footer.html"}
	t := template.Must(template.New("content.html").ParseFiles(templates...))
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
