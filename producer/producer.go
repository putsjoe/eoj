package main

import (
	"bufio"
	"fmt"
	"net"
	"crypto/rand"
)

func send_data(a string, port string) {

	conn, err := net.Dial("tcp", port)
	if err != nil {
		panic(err)
	}

	// Send data
	fmt.Fprintf(conn, a + "\n")

	// Get reply
	reply, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Received:  " + reply)
	conn.Close()

}

func rand_str(str_size int) string {
	// Taken from - https://devpy.wordpress.com/2013/10/24/create-random-string-in-golang/

	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, str_size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	fmt.Println(string(bytes))
	return string(bytes)
}

func main() {
	var port string = ":1234"
	send_data(rand_str(16), port)

}
