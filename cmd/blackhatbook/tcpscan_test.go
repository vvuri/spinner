package blackhatbook

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net"
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
