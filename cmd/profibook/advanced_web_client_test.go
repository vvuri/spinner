package profibook

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestAdvancedServer(t *testing.T) {
	URL, err := url.Parse("http://ya.ru")
	if err != nil {
		fmt.Println("Error in parsing:", err)
		return
	}

	c := &http.Client{
		Timeout: 15 * time.Second,
	}
	// возвращает объект http.Request с указанием метода, URL-адреса и, возможно, тела запроса
	request, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		fmt.Println("Get:", err)
		return
	}
	// отправляет HTTP-запрос (http.Request) посредством http.Client
	//и получает HTTP-ответ (http.Response)
	httpData, err := c.Do(request)
	if err != nil {
		fmt.Println("Error in Do():", err)
		return
	}

	fmt.Println("Status code:", httpData.Status)

	header, _ := httputil.DumpResponse(httpData, false)
	fmt.Print(string(header))

	contentType := httpData.Header.Get("Content-Type")
	characterSet := strings.SplitAfter(contentType, "charset=")
	if len(characterSet) > 1 {
		fmt.Println("Character Set:", characterSet[1])
	}
	if httpData.ContentLength == -1 {
		fmt.Println("ContentLength is unknown!")
	} else {
		fmt.Println("ContentLength:", httpData.ContentLength)
	}

	// поиск ответа сервера
	length := 0
	var buffer [1024]byte
	r := httpData.Body
	for {
		n, err := r.Read(buffer[0:])
		if err != nil {
			fmt.Println(err)
			break
		}
		length = length + n
	}
	fmt.Println("Calculated response data length:", length)
}
