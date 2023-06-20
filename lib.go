package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func returnHandler(stream net.Conn) func() {
	return func() {
		handleConnection(stream)
	}
}

func handleConnection(stream net.Conn) {
	bufReader := bufio.NewReader(stream)
	first_request_line, err := bufReader.ReadString('\n')
	check(err)
	first_request_line = strings.TrimSuffix(first_request_line, "\r\n")
	fmt.Println("Request: ", first_request_line)

	var statusLine string
	var filename string

	switch first_request_line {
	case "GET / HTTP/1.1": // GET / HTTP/1.1
		statusLine = "HTTP/1.1 200 OK"
		filename = "hello.html"
	case "GET /sleep HTTP/1.1": // GET /sleep HTTP/1.1
		time.Sleep(10 * time.Second)
		statusLine = "HTTP/1.1 200 OK"
		filename = "hello-sleep.html"
	default:
		statusLine = "HTTP/1.1 404 NOT FOUND"
		filename = "404.html"
	}

	contents, err := os.ReadFile(filename)
	check(err)

	length := len(contents)

	response := fmt.Sprintf("%s\r\nContent-Length: %d\r\n\r\n%s", statusLine, length, contents)

	_, err = stream.Write([]byte(response))
	check(err)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
