package tcpsrv

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

func tcpClient(conn string) {
	c, err := net.Dial("tcp", conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	//	for {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print(">> ")

	//text, _ := reader.ReadString('\n')
	text := "Echo test"
	fmt.Fprintf(c, text+"\n")

	message, _ := bufio.NewReader(c).ReadString('\n')
	fmt.Print("->: " + message)

	if strings.TrimSpace(string(text)) == "STOP" {
		fmt.Println("TCP client exiting...")
		return
	}
	//	}
}

func tcpClientTwo(url string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", url)
	if err != nil {
		fmt.Println("ResolveTCPAddr:", err.Error())
		return
	}

	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		fmt.Println("DialTCP:", err.Error())
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			conn.Close()
			return
		}
	}
}

func tcpServer(port string) {
	PORT := ":" + port
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}
		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}

func TestHttpsClientLocal(t *testing.T) {
	go tcpServer("5050")
	tcpClient("localhost:5050")
}

func TestHttpsClientLocalTwo(t *testing.T) {
	go tcpServer("5050")
	tcpClientTwo("localhost:5050")
}
