package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	b := make([]byte, 512)
	for {
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
		}

		if err != nil {
			log.Println("Unexpected error occurred. :'( ")
		}

		log.Printf("Received %d bytes on the wire: %s\n", size, string(b))

		log.Println("Writing data to stdout")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data to stdout")
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln("Cannot bind to port.")
	}
	log.Println("Listening on 0.0.0.0:8000")

	for {
		conn, err := listener.Accept()
		log.Println("Received connection...")

		if err != nil {
			log.Fatalln("Unable to accept connection. Terminated!")
		}

		go echo(conn)
	}
}
