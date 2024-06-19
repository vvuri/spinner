package profibook

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// $ openssl genrsa -out server.key 2048
// $ openssl ecparam -genkey -name secp384r1 -out server.key
// $ openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
// --> server.cr, server.key

// если сертификат является самозаверяющим, для работы HTTPS-клиента нужно использовать
// параметр InsecureSkipVerify: true в структуре http.Transport

// сертификат для клиента
// openssl req -x509 -nodes -newkey rsa:2048 -keyout client.key -out
// --> client.crt, client.key

// crypto/tls - частично реализует протокол безопасности на транспортном уровне

func TestHttpsClient(t *testing.T) {
	URL := "https://google.com"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{},
	}

	client := &http.Client{Transport: tr}

	response, err := client.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	s := strings.TrimSpace(string(content))
	fmt.Println(s)
}

// ============ Server

func Default(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is an example HTTPS server!\n")
}

// --> ListenAndServeTLS:  open server.crt: The system cannot find the file specified.
func TestHttpsSimpleServer(t *testing.T) {
	PORT := ":1443"
	http.HandleFunc("/", Default)
	fmt.Println("Listening to port number", PORT)

	//  ожидает HTTPS-соединения
	err := http.ListenAndServeTLS(PORT, "server.crt", "server.key", nil)
	if err != nil {
		fmt.Println("ListenAndServeTLS: ", err)
		return
	}
}

// ============ TLS-сервера и TLS-клиента
