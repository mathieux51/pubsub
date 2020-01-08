package main

import (
	"testing"
)

var testCases = []struct {
	expected int
}{
	{
		expected: 3,
	},
}

func TestReadMessages(t *testing.T) {
	for _, tc := range testCases {
		// fmt.Printf("%+v\n", tc)
		if got, _ := readMessages(); len(got) != tc.expected {
			t.Errorf("readMessages() = %v; want %v", len(got), tc.expected)
		}
	}

}
