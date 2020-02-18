package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("localhost:%d", j)

			conn, err := net.Dial("tcp", address)

			if err != nil {
				return
			}

			conn.Close()
			fmt.Printf("Port %d is open\n", j)
		}(i)
	}
	wg.Wait()
}
