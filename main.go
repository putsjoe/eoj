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

func Addto(slice []string, m string) []string {
	m = strings.TrimRight(m, "\n")
	slice = append(slice, m)
	return slice
}

func ReturnTop(slice []string) (string, []string) {
	var first string = slice[0]
	slice = append(slice[:0], slice[1:]...)
	return first, slice
}

func main() {
	var lilo []string //Define main slice

	var err_leng string = "ERROR_LENGTH\n"
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

		if leng > 16 {
			fmt.Print("Received: ", string(msg))

			// Check Complexity

			// If valid add to queue

			//err_c = err_count(err_c)
			lilo = Addto(lilo, msg)

			// If invalid, return error and close
			fmt.Println(len(lilo))
			fmt.Println(lilo)

			if len(lilo) > 5 {
				var reply string
				reply, lilo = ReturnTop(lilo)
				conn.Write([]byte(reply + "\n"))
			} else {
				conn.Write([]byte("STACK_LOW\n"))
			}

			/*newmsg := strings.ToUpper(msg)
			conn.Write([]byte(newmsg)) */

		} else if leng <= 16 && msg != "\n" {
			fmt.Println(err_leng)
			conn.Write([]byte(err_leng))

			// Add to error count
			err_c = err_count(err_c)

			if len(err_c) > 2 { // If third length error
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
