package profibook

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func CheckStatusOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `Fine!`)
}

func StatusNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func WWWServer() {
	PORT := ":8001"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Using default port number: ", PORT)
	} else {
		PORT = ":" + arguments[1]
	}

	http.HandleFunc("/CheckStatusOK", CheckStatusOK)
	http.HandleFunc("/StatusNotFound", StatusNotFound)
	http.HandleFunc("/", MyHandler)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestCheckStatusOK(t *testing.T) {
	go WWWServer()

	req, err := http.NewRequest("GET", "/CheckStatusOK", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CheckStatusOK)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned %v", status)
	}
	expect := `Fine!`
	if rr.Body.String() != expect {
		t.Errorf("handler returned %v", rr.Body.String())
	}
}

func TestStatusNotFound(t *testing.T) {
	go WWWServer()

	req, err := http.NewRequest("GET", "/StatusNotFound", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StatusNotFound)
	handler.ServeHTTP(rr, req)
	status := rr.Code
	if status != http.StatusNotFound {
		t.Errorf("handler returned %v", status)
	}
}
