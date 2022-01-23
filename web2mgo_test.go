package main

import (
	"fmt"
	"testing"
)

func TestODatareq(t *testing.T) {
	var tests = []struct {
		Name string
		Data string
	}{
		{"", "\n"},
		{"", ""},
		{"\t", "one\ttwo\tthree\n"},
		{"Data for test", "Number 99999 to data test"},
		{"Yes, no", "No, or, yes"},
	}

	var prevName string
	for _, test := range tests {
		if test.Name != prevName {
			fmt.Printf("\n%s\n", test.Name)
			prevName = test.Name
		}
	}

	var prevData string
	for _, test := range tests {
		if test.Data != prevData {
			fmt.Printf("\n%s\n", test.Data)
			prevData = test.Data
		}
	}
}
