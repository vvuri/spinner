package main

import (
	"io"
	"log"
	"net"
)

const siteReal = "vvuri.ru:80"

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", siteReal)
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")
	}
	defer dst.Close()

	// Выполняется в горутине для предотвращения блокировки io.Copy
	go func() {
		// Копируем вывод источника в получателя
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	// Копирование вывода получателя обратно в источник
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// Прослушивание локального порта 80
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}
