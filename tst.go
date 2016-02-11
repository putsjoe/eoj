package main

import (
	"fmt"
)

func main() {
	var a []string
	a = append(a,"a") 
	a = append(a,"b") 
	a = append(a,"c") 
	a = append(a,"d") 
	a = append(a,"e") 
	b := a[3:]
	c := a[:3]

	fmt.Println(b)
	fmt.Println(c)

}

