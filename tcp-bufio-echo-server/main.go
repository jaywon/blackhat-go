package main

import (
	"bufio"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read data from connection")
	}

	log.Printf("Received %d bytes on the wire: %s\n", len(s), s)
	log.Println("Writing data to stdout")

	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data to stdout")
	}
	writer.Flush()
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
