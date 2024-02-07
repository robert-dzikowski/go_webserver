// Single threaded web server
//

package st_webserver

import (
	"fmt"
	"net"
)

func stMain() {
	listener, err := net.Listen("tcp", "127.0.0.1:7878")
	check(err)
	fmt.Println("Started server at localhost:7878")

	defer func() {
		fmt.Println("Closing server...")
		listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		check(err)

		handleConnection(conn)
	}
}
