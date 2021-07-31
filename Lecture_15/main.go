package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//net = network
	fmt.Println("Local server is running on port 8888")
	nl, err := net.Listen("tcp", ":8888") //ip:port

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //1
	}

	conn, err := nl.Accept() //Layer 5 session layer

	if err != nil {
		fmt.Println(err.Error())

	}

	bs := make([]byte, 1024)

	n, e := conn.Read(bs)
	if e != nil {
		fmt.Println(e.Error())

	}

	fmt.Println(n)

	reqstr := string(bs)
	fmt.Println(reqstr)

	body := `<!DOCTYPE html>
	<html>
	<head>
	<title></title>
	</head>
	<body>
	
	<h1>My First Heading</h1>

	
	</body>
	</html>`

	// fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	// fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	// fmt.Fprint(conn, "Content-Type: text/html\r\n")
	// fmt.Fprint(conn, "\r\n")
	// fmt.Fprint(conn, body)

	resp := "HTTP/1.1 200 OK\r\n" +
		"Content-Length: %d\r\n" +
		"Content-Type: text/html\r\n" +
		"\r\n%s"

	msg := fmt.Sprintf(resp, len(body), body)
	fmt.Println(msg)
	conn.Write([]byte(msg))

}
