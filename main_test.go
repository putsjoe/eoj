package main

import (
	"testing"
)

func Test_isComplex(t *testing.T) {
	if (isComplex("aaabcdefg12345", 2) != false) {
		t.Error("Max Recurrance Error")
	}
	if (isComplex("abcabc12", 5) != false) {
		t.Error("Count Characters doesnt work")
	}

}


