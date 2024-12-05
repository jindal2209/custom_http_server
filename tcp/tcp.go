package tcp

import (
	"net"

	"exmaple.com/http"
)

func Tcp() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	inputLength, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

	http.HandleRequest(buffer, inputLength)
}

/*
	thread limit
	thread pool
	connection timeout
	tcp backlog queue config
*/
