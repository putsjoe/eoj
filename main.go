package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	//"reflect"
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

	var err_c []int

	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		var leng int = length(msg)
		//fmt.Println("Length - ", length(msg))
		//fmt.Print(reflect.TypeOf(msg))
		fmt.Print(msg)
		//msg = "clo"
		//fmt.Print(reflect.TypeOf(msg))
		msgt := strings.TrimRight(msg, "\n")
		var bing bool = msgt == "clo"

		fmt.Println(bing)

		if msg == "close" {
			fmt.Println("Close connection")
			conn.Close()
		}

		if msg == "what" {
			fmt.Println(string(len(err_c)))
		}

		if leng > 8 {
			fmt.Print("Received: ", string(msg))

			newmsg := strings.ToUpper(msg)
			conn.Write([]byte(newmsg + "\n"))

			msg = ""

		} else if leng < 9 && msg != "\n" {
			fmt.Println("ERROR_LENGTH")
			_ = append(err_c, 1)

			if len(err_c) > 2 {
				conn.Close()
				os.Exit(2)
			}

		} else {
			fmt.Println("ERROR_EMPTY")
		}
	}

}
