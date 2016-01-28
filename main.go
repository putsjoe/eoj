package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	var prt string = ":1234"
	fmt.Println("Listening on port ", prt)

	ln, _ := net.Listen("tcp", prt)

	conn, _ := ln.Accept()

	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Received: ", string(msg))

		newmsg := strings.ToUpper(msg)
		conn.Write([]byte(newmsg + "\n"))

		msg = ""
	}

}
