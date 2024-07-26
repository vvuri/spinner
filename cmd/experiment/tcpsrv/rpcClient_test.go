package tcpsrv

import (
	"fmt"
	"math"
	"net"
	"net/rpc"
	"testing"
)

type MyFloats struct {
	A1, A2 float64
}

type MyInterface interface {
	Multiply(arguments *MyFloats, reply *float64) error
	Power(arguments *MyFloats, reply *float64) error
}

func Power(x, y float64) float64 {
	return math.Pow(x, y)
}

func (t *MyInterface) Multiply(arguments *MyFloats, reply *float64) error {
	*reply = arguments.A1 * arguments.A2
	return nil
}

func (t *MyInterface) Power(arguments *MyFloats, reply *float64) error {
	*reply = Power(arguments.A1, arguments.A2)
	return nil
}

func rpcServer(port string) {
	myInterface := new(MyInterface)
	rpc.Register(myInterface)
	t, err := net.ResolveTCPAddr("tcp4", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := net.ListenTCP("tcp4", t)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		fmt.Printf("%s\n", c.RemoteAddr())
		rpc.ServeConn(c)
	}
}

func rpcClient(url string) {
	c, err := rpc.Dial("tcp", url)
	if err != nil {
		fmt.Println(err)
		return
	}

	args := MyFloats{16, -0.5}
	var reply float64
	err = c.Call("MyInterface.Multiply", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Multiply): %f\n", reply)

	err = c.Call("MyInterface.Power", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Power): %f\n", reply)
}

func TestRPCServer(t *testing.T) {
	go rpcServer("5052")
	rpcClient("localhost:5052")
}
