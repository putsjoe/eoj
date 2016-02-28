package main

import (
	"fmt"
	"strings"
)

func tst() {
	var a []string
	a = append(a, "a")
	a = append(a, "b")
	a = append(a, "c")
	a = append(a, "d")
	a = append(a, "e")
	b := a[3:]
	c := a[:3]

	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(" -- -- ")

	var sting string = "aaabbcccddddeeeee"

	// Make map
	countr := make(map[rune]int)

	for _, d := range sting {
		countr[d]++
	}

	for k, v := range countr {
		//fmt.Println(string(k) + " - " + string(v))
		fmt.Printf("%v - %v ", k, v)
		if v > 3 {
			fmt.Printf("BingBong\n")
		} else {
			fmt.Printf("\n")
		}
	}
	fmt.Println(" -- -- -- ")

	var abc string = "abcdefghijklmnopqrstuvwxyz"
	var xem string = "g"

	if strings.Contains(abc, xem) {
		fmt.Println(xem + " in ABC")
	}


}
