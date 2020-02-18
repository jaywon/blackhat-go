package main

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "sudokrew.com:443")
	if err != nil {
		log.Fatalln("Could not connect to proxy. Womp womp...")
	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	lisetener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln("Unable to bind to port successfully")
	}

	for {
		conn, err := lisetener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go handle(conn)
	}
}
