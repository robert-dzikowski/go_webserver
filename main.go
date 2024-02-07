package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:7878")
	check(err)
	fmt.Println("Started server at localhost:7878")

	defer func() {
		fmt.Println("Closing server...")
		listener.Close()
	}()

	var wg sync.WaitGroup
	const GorNum int = 4 // Number of goroutines
	wg.Add(GorNum)

	c := make(chan func())

	for i := 1; i <= GorNum; i++ {
		go func(x int) {
			defer wg.Done()
			for job := range c { // The for range operation runs until channel is drained
				fmt.Printf("Received job for goroutine #%d\n", x)
				job()
			}
		}(i)
	}

	for i := 1; i <= 6; i++ {
		conn, e := listener.Accept()
		check(e)

		c <- returnHandler(conn)
	}

	close(c)
	wg.Wait()
}
