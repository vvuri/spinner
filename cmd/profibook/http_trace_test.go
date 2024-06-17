package profibook

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
	"os"
	"testing"
)

// прохождение HTTP-запроса
func TestTraceHttp(t *testing.T) {
	URL := "http://yandex.ru" // "http://loaclhost:8004/"

	// позволяет отправить запрос на сервер и получить ответ
	client := http.Client{}

	req, _ := http.NewRequest("GET", URL, nil)

	// Когда происходит одно из событий, выполняется соответствующий код.
	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			fmt.Println("First response byte!")
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		ConnectStart: func(network, addr string) {
			fmt.Println("Dial start")
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Println("Dial done")
		},
		WroteHeaders: func() {
			fmt.Println("Wrote headers")
		},
	}

	// WithClientTrace - возвращает новый контекст
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	fmt.Println("Requesting data from server!")

	// http.DefaultTransport.RoundTrip() служит оберткой для объекта http.DefaultTransport.RoundTrip,
	// чтобы он отслеживал текущий запрос.
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// фактическое выполнение запроса к веб-серверу
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, response.Body)
}
