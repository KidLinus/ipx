package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/gomarkdown/markdown"
)

func main() {
	tmpl, err := template.ParseFiles("src/template.html")
	if err != nil {
		log.Fatal("template_parse_fail", err)
	}
	out, err := os.OpenFile("./index.html", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("ouput_file_open_fail", err)
	}
	files, err := ioutil.ReadDir("./src/pages")
	if err != nil {
		log.Fatal("pages_dir_fail", err)
	}
	menu := []MenuItem{}
	menujs, err := ioutil.ReadFile("./src/menu.json")
	if err != nil {
		log.Fatal("menu_read_fail", err)
	}
	if err := json.Unmarshal(menujs, &menu); err != nil {
		log.Fatal("menu_parse_fail", err)
	}
	pages := []Page{}
	for _, f := range files {
		md, err := ioutil.ReadFile(path.Join("./src/pages", f.Name()))
		if err != nil {
			log.Fatal("page_read_fail", path.Join("./src/pages", f.Name()), err)
		}
		html := markdown.ToHTML(md, nil, nil)
		pages = append(pages, Page{ID: f.Name(), HTML: template.HTML(html)})
	}
	if err := tmpl.Execute(out, Site{Menu: menu, Pages: pages}); err != nil {
		log.Fatal("template_execute_fail", err)
	}
}

type Site struct {
	Menu  []MenuItem
	Pages []Page
}

type MenuItem struct {
	Text string
	Page string
}

type Page struct {
	ID   string
	HTML template.HTML
}
