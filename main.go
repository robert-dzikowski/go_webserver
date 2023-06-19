package main

import (
	"fmt"
	// mypack "gowebserver/internal"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:7878")
	check(err)
	fmt.Println("Started server at localhost:7878")

	defer func() {
		fmt.Println("Closing server...")
		listener.Close()
	}()

	// mypack.Fun()

	for {
		conn, err := listener.Accept()
		check(err)

		handleConnection(conn)
	}
}

// func main() {
// 	c := make(chan int)
// 	var wg sync.WaitGroup
// 	const GorNum int = 5
// 	wg.Add(GorNum)

// 	for i := 1; i <= GorNum; i++ {
// 		go func(x int) {
// 			defer wg.Done()
// 			for v := range c {
// 				fmt.Printf("Data %d from goroutine #%d\n", v, x)
// 				time.Sleep(time.Second / 4)
// 			}
// 		}(i)
// 	}

// 	for i2 := 1; i2 <= 20; i2++ {
// 		c <- i2
// 	}

// 	close(c)
// 	wg.Wait()
// }
