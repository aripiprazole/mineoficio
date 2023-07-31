package main

import (
	"fmt"
	"net"
)

const (
	ConnHost = "localhost"
	ConnPort = "25565"
	ConnType = "tcp"
)

func HandleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		panic("Error reading: " + err.Error())
	}
	_, _ = conn.Write([]byte("Message received."))
	_ = conn.Close()
}

func main() {
	listener, err := net.Listen(ConnType, ConnHost+":"+ConnPort)
	if err != nil {
		panic(err)
	}

	// Close the listener when the application closes. (Defer)
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			panic(err)
		}
	}(listener)

	fmt.Printf("Listening on %s:%s", ConnHost, ConnPort)
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic("Error accepting: " + err.Error())
		}

		go HandleRequest(conn)
	}
}
