package main

import (
	"testing"
)

func Test_isComplex(t *testing.T) {
	if isComplex("aaabcdefg12345", 2) != false {
		t.Error("Max Recurrance Error")
	}
	if isComplex("abcabc12", 5) != false {
		t.Error("Count Characters doesnt work")
	}

}

func Test_Addto(t *testing.T) {
	var test_list []string
	test_list = append(test_list, "abcjsndjsn", "dbhjbj", "bcdbdjb")
	//var test_list_p = append(test_list, "test")
	ts := Addto(test_list, "test")
	if ts[len(ts)-1] != "test" {
		t.Error("Addto function fail")
	}

}

func Test_length(t *testing.T) {
	if length("pingpong\n") != 8 {
		t.Error("length function test01 failed.")
	}
	if length("flick\nbob\n") != 8 {
		t.Error("Length function test02 failed")
	}

}
