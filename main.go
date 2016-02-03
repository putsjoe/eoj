package main

import (
	"bufio"
	"fmt"
	"net"
	//	"os"
	//"reflect"
	"strings"
	"unicode/utf8"
)

func length(a string) int {
	a = strings.Replace(a, "\n", "", -1)
	return utf8.RuneCountInString(a)
}

func err_count(slice []string) []string {
	slice = append(slice, "a")
	return slice
}

func main() {
	var prt string = ":1234"
	fmt.Println("Listening on port ", prt)

	ln, err := net.Listen("tcp", prt)
	if err != nil {
		panic(err)
	}

	// Accept connection
	conn, err := ln.Accept()
	if err != nil {
		conn.Close()

		// Handle Error
	}

	var err_c []string
	//err_c := make([]string, 1)

	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		var leng int = length(msg)
		msgt := strings.TrimRight(msg, "\n")

		//fmt.Print(msg)

		//var bing bool = msgt == "clo"
		//fmt.Println(bing)

		if msgt == "close" {
			fmt.Println("Close connection")
			conn.Close()
			msg = ""

		}

		if leng > 8 {
			fmt.Print("Received: ", string(msg))

			newmsg := strings.ToUpper(msg)
			conn.Write([]byte(newmsg))

			msg = ""

		} else if leng < 9 && msg != "\n" {
			fmt.Println("ERROR_LENGTH")

			err_c = err_count(err_c)

			if len(err_c) > 2 {
				conn.Close()
				// Next line stops the program shitting itself.
				conn, err = ln.Accept()
				// Next line resets the error count
				err_c = err_c[:0]

			}

		} else {
			fmt.Println("ERROR_EMPTY")
		}
	}

}
