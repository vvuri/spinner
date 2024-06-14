package profibook

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func TestHttpHello(t *testing.T) {
	http.HandleFunc("/hello", hello)
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/hello", nil)
	http.DefaultServeMux.ServeHTTP(writer, request)
	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
	if expected, actual := "Hello World!", writer.Body.String(); expected != actual {
		t.Errorf("Response body is %v", actual)
	}
}

// -------------------------------------------------------
