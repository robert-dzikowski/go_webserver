package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func handleConnection(stream net.Conn) {
	bufReader := bufio.NewReader(stream)
	first_request_line, err := bufReader.ReadString('\n')
	check(err)

	var statusLine string
	var filename string

	switch first_request_line {
	case "GET / HTTP/1.1\n":
		statusLine = "HTTP/1.1 200 OK"
		filename = "hello.html"
	case "GET /sleep HTTP/1.1\n":
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

// fn handle_connection(mut stream: TcpStream) {
//     let buf_reader = BufReader::new(&mut stream);
//     let first_request_line = buf_reader.lines().next().unwrap().unwrap();

//     let (status_line, filename) = match &first_request_line[..] {
//         "GET / HTTP/1.1" => ("HTTP/1.1 200 OK", "hello.html"),
//         "GET /sleep HTTP/1.1" => {
//             thread::sleep(Duration::from_secs(10));
//             ("HTTP/1.1 200 OK", "hello-sleep.html")
//         }
//         _ => ("HTTP/1.1 404 NOT FOUND", "404.html"),
//     };

//     let contents = fs::read_to_string(filename).unwrap();
//     let length = contents.len();

//     let response = format!("{status_line}\r\nContent-Length: {length}\r\n\r\n{contents}");

//     stream.write_all(response.as_bytes()).unwrap();
// }
