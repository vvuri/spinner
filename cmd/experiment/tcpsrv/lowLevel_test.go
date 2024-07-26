package tcpsrv

import (
	"fmt"
	"net"
	"os"
	"syscall"
	"testing"
)

func TestICMP(t *testing.T) {
	netaddr, err := net.ResolveIPAddr("ip4", "127.0.0.1")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := net.ListenIP("ip4:icmp", netaddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFrom(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("% X\n", buffer[0:n])
}

func TestSyscall(t *testing.T) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		fmt.Println("Error in syscall.Socket:", err)
		return
	}

	f := os.NewFile(uintptr(fd), "captureICMP")
	if f == nil {
		fmt.Println("Error in os.NewFile:", err)
		return
	}

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_RCVBUF, 256)
	if err != nil {
		fmt.Println("Error in syscall.Socket:", err)
		return
	}

	for {
		buf := make([]byte, 1024)
		numRead, err := f.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("% X\n", buf[:numRead])
	}
}
