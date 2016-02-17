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

func isComplex(a string) bool {
	// Create Map to count.
	countr := make(map[rune]int)
	var cChars string = "abcdefghijklmnopqrstuvwxyz"
	var cFail int = 0
	var cABC int = 0
	var cnABC int = 0

	// Cycle through array
	for _, x := range a {
		countr[x]++
	}

	for k,v := range countr {
		//fmt.Printf("%v - %v ",string(k),v)
		if v > 5 {
			cFail++
		}
		if strings.Contains(cChars, string(k)) == false {
			cnABC++
			fmt.Printf("%v - Not in ABC \n",string(k))
		} else {
			cABC++
		}
	}
	fmt.Printf("cFail - %v ; cABC - %v ; cnABC - %v ;  \n",cFail,cABC,cnABC)

	if cFail > 0 || cABC < 4 || cnABC < 4 {
		return false
	} else {
		return true
	}

	// Check contains characters and numbers
	/* Check characters dont appear more than three times each.
	Use range with the string.
	Use maybe a map (golang dict equiv) to keep track of number of times a character is used)
	*/
	//> Count characters

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

		}

		if leng > 12 && leng <= 28 {
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
				var complx bool = isComplex(msg)
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

			//err_c = err_count(err_c)

			/*newmsg := strings.ToUpper(msg)
			conn.Write([]byte(newmsg)) */

		} else if leng <= 16 && msg != "\n" || leng > 28 && msg != "\n" {
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
