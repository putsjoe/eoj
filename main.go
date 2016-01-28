package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"unicode/utf8"
)

func length(a string) int {
	a = strings.Replace(a, "\n", "", -1)
	return utf8.RuneCountInString(a)
}

func main() {
	var prt string = ":1234"
	fmt.Println("Listening on port ", prt)

	ln, _ := net.Listen("tcp", prt)

	conn, _ := ln.Accept()

	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Println("Length - ", length(msg))

		if msg != "\n" {
			fmt.Print("Received: ", string(msg))

			newmsg := strings.ToUpper(msg)
			conn.Write([]byte(newmsg + "\n"))

			msg = ""
		} else {
			fmt.Println("Empty")
		}
	}

}
