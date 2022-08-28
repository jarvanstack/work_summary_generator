package main

import (
	"fmt"
	"testing"
)

func Test_readData(t *testing.T) {
	s := readData()
	fmt.Printf("s: %v\n", s)
}
