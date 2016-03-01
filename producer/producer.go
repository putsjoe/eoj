package main 

import (
	"net"
	"fmt"
	"bufio"
	)

func main() {
	var port string = ":1234"
	
	fmt.Println("Sending on port ", port)
	
	conn, err := net.Dial("tcp", port)
	if err != nil {
			panic(err)
	}
	
	for {
		
		// Send data
		fmt.Fprintf(conn, "abcdef123" + "\n")
		
		// Get reply
		reply, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Recieved:  " + reply)

	}

	

}

