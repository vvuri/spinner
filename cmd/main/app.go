package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type Entry struct {
	Number int
	Double int
	Square int
}

var DATA []Entry
var tFile string

func pageHello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Simple app for Kind <br> {url}/first <br> {url}/second"))
	if err != nil {
		log.Println(err)
	}
}

func pageFirst(w http.ResponseWriter, r *http.Request) {
	log.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)
	tmpl := template.Must(template.ParseFiles("./web/layout.html"))
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(w, data)
}

func pageSecond(w http.ResponseWriter, r *http.Request) {
	log.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)
	parseT, err := template.ParseGlob(tFile)
	if err != nil {
		log.Printf("Error parse template %s", err)
	}
	myT := template.Must(parseT, err)
	// первый параметр — это переменная, в которой хранится соединение с HTTP клиентом,
	// второй параметр — файл шаблона, который будет использоваться для  форматирования данных,
	// третий параметр — срез структур с данными.
	myT.ExecuteTemplate(w, tFile, DATA)
}

func initData() {
	tFile = "web/template.html"

	for i := range [5]int{} {
		temp := Entry{Number: i, Double: i + 1, Square: i + 2}
		DATA = append(DATA, temp)
	}
	log.Println(DATA)
}

func startServer() {
	http.HandleFunc("/", pageHello)
	http.HandleFunc("/first", pageFirst)
	http.HandleFunc("/second", pageSecond)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	log.Println("Server started")
	printDir()
	//- initData()
	//- startServer()
	// wget -qO- http://localhost:8080
}

// TODO: Починить /second т.к. не работает

// TODO: Add Metrics
