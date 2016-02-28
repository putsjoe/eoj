package main

import (
	"testing"
)

func Test_isComplex(t *testing.T) {
	
	if (isComplex("abcdefghi12345", 5) != true) {
		t.Error("max recur doesnt work")
	} else {
		t.Log("Recur test passed")
	}

}


