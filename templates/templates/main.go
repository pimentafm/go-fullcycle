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
	tmp := template.New("CourseTemplate")
	tmp, _ = tmp.Parse("Course Name: {{.Name}} Hours: {{.Hours}}")

	err := tmp.ExecuteTemplate(os.Stdout, "CourseTemplate", course)
	if err != nil {
		panic(err)
	}
}
