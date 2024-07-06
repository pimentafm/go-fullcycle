package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name  string
	Hours int
}

func main() {
	course := Course{"Go", 20}
	t := template.Must(template.New("CourseTemplate").Parse("Course Name: {{.Name}} Hours: {{.Hours}}"))

	err := t.Execute(os.Stdout, course)

	if err != nil {
		panic(err)
	}
}
