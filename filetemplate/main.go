package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name  string
	Hours int
}

type Courses []Course

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, Courses{
		{"Go", 20},
		{"Python", 30},
		{"Java", 40},
	})

	if err != nil {
		panic(err)
	}
}
