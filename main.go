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

func isComplex(a string, max_r int) bool {
	// Create Map to count.
	countr := make(map[rune]int)
	var cChars string = "abcdefghijklmnopqrstuvwxyz"
	var cFail int = 0
	var cABC int = 0
	var cnABC int = 0

	/*
			var entry_len int = length(a)
			fmt.Printf("Entry Length - %v \n", entry_len)

		// Check contains characters and numbers
		// Check characters dont appear more than three times each.
		Use range with the string.
	*/

	// Cycle through array
	for _, x := range a {
		countr[x]++
	}

	for k, v := range countr {
		//fmt.Printf("%v - %v ",string(k),v)
		if v > max_r {
			cFail++
		}
		if strings.Contains(cChars, string(k)) {
			//cABC++
			cABC = cABC + v
		} else {
			cnABC = cnABC + v
			//fmt.Printf("%v - Not in ABC \n", string(k))
		}
	}
	//fmt.Printf("cFail - %v ; cABC - %v ; cnABC - %v ;  \n", cFail, cABC, cnABC)

	if cFail > 0 || cABC < 4 || cnABC < 4 {
		return false
	} else {
		return true
	}
}

func isDuplicate(a string, slice []string) bool {

	//fmt.Println("isDuplicate? - ")
	//fmt.Println(slice)

	if len(a) > 0 {
		// Compare new entry to all old
		var match bool = false // To see if any matched
		for _, element := range slice {
			//fmt.Println(element + "||")
			//fmt.Println(a + "||")
			if element == a {
				match = true
				fmt.Println("Matched")
				break
			}
		}
		if match {
			return false
		} else {
			return true
			//return true
		}
	} else {
		return false
	}

}

func main() {
	var lilo []string         // Define main slice
	var clilo []string        // Define copy slice for comparing stack
	var check_copy int = 3000 // Define how many entries at the end of the slice are checked for copies
	var min_len int = 16      // Define minimum length of acceptable strings
	var max_len int = 32      // Define maximum length of acceptable strings
	var max_recur = 5         // Define maximum reoccurance of character

	var err_leng string = "ERROR_LENGTH\n"
	//var err_empty string = "ERROR_EMPTY"
	var prt string = ":1234"
	var err_c []string // Define error count
	var emp_c int = 0  // Define empty count

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

	for {
		conn.Write([]byte("eoj:: "))
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
			emp_c = 0

		}

		if leng > 12 && leng <= 28 {
			emp_c = 0
			fmt.Print("Received: ", string(msg))

			// Copy the slice to a new slice that only contains the last so many entries
			fmt.Println(len(lilo))
			fmt.Println(check_copy)
			fmt.Println("---")

			if check_copy >= len(lilo) {
				fmt.Println("copy lilo to clilo")
				clilo = lilo
			} else {
				fmt.Println("get last of slice")
				var ch_copy_amount int = len(lilo) - check_copy
				fmt.Println(ch_copy_amount)
				clilo = lilo[ch_copy_amount:]
			}

			// Check Complexity
			msg = strings.TrimRight(msg, "\n")
			var dupl bool = isDuplicate(msg, clilo)

			// If valid add to queue
			if dupl {
				var complx bool = isComplex(msg, max_recur)
				if complx {
					lilo = Addto(lilo, msg)
				} else {
					fmt.Println("Invalid Entry - COM")
					conn.Write([]byte("ERROR_COMPLEX\n"))
				}
				// Debugging
				fmt.Printf("%v -- %v \n", len(lilo), lilo)

				if len(lilo) > 5 {
					var reply string
					reply, lilo = ReturnTop(lilo)
					conn.Write([]byte(reply + "\n"))
				} else {
					conn.Write([]byte("STACK_LOW\n"))
				}

			} else {
				// If invalid, return error and close
				fmt.Println("Invalid Entry - DUP")
				conn.Write([]byte("DUPLICATE_ERROR\n"))

			}

		} else if leng == 0 {
			fmt.Print("*")
			//	conn.Write([]byte(err_empty+"\n"))
			emp_c = emp_c + 1

			if emp_c == 3 {
				fmt.Println("\nToo many empty - Disconnecting Session")
				//conn.Write([]byte(err_empty+"\n"))
				conn.Close()
				conn, err = ln.Accept()
				emp_c = 0
			}

		} else if leng <= min_len && msg != "\n" || leng > max_len && msg != "\n" {
			emp_c = 0
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
			fmt.Println("ERROR_UNKNOWN")
		}
	}

}
