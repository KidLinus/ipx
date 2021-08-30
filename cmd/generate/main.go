package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	template, err := template.ParseFiles("src/template.html")
	if err != nil {
		log.Fatal("template_parse_fail", err)
	}
	out, err := os.OpenFile("./index.html", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("ouput_file_open_fail", err)
	}
	if err := template.Execute(out, Params{Menu: "derp", Content: "Dorp"}); err != nil {
		log.Fatal("template_execute_fail", err)
	}
}

type Params struct {
	Menu    string
	Content string
}
