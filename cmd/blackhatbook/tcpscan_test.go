package blackhatbook

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"testing"
)

func TestTcpScanNil(t *testing.T) {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err != nil {
		fmt.Println("Port is closed")
	}
	assert.Nil(t, err, "Connection successful")
}

func TestTcpScan80(t *testing.T) {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("Connection successful")
	}
}

func TestTcpScanFirt1024Ports(t *testing.T) {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// порт закрыт или отфильтрован
			continue
		}
		conn.Close()

		//t.Logf("%d open\n", i)
		fmt.Printf("%d open\n", i)
	}
}

func TestTcpScanFirt1024PortsParallel(t *testing.T) {
	var wg sync.WaitGroup

	for i := 1; i <= 1024; i++ {
		wg.Add(1)

		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)

		wg.Wait()
	}
}

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

func TestRequest(t *testing.T) {
	form := url.Values{}
	form.Add("foo", "bar")
	var client http.Client
	req, err := http.NewRequest(
		"PUT",
		"https://www.google.com/robots.txt",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		t.Error("Request error: ", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("Response error: ", err)
	}
	t.Log(resp.Header, resp.Body)
}

func TestRequestPut(t *testing.T) {
	form := url.Values{}
	form.Add("foo", "bar")
	var client http.Client
	req, err := http.NewRequest(
		"PUT",
		"https://www.google.com/robots.txt",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		t.Error("Request error: ", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("Response error: ", err)
	}
	t.Log(resp.Header, resp.Body)
}

func TestRequestGet(t *testing.T) {
	var client http.Client
	req, err := http.NewRequest(
		"GET",
		"https://habr.com/robots.txt",
		nil,
	)
	if err != nil {
		t.Error("Request error: ", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("Response error: ", err)
	}
	t.Log(resp.Header)
	t.Log(resp.Body)
	t.Log(resp)
}

func TestRequestGet2(t *testing.T) {
	resp, err := http.Get("https://habr.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	// Вывод статуса HTTP
	fmt.Println(resp.Status)

	// Считывание и отображение тела ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	resp.Body.Close()
}
