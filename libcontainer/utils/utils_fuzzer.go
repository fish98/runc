package utils

import (
	// "bytes"
	// "testing"
	// "strings"

	"golang.org/x/sys/unix"
)

func IsDivisbleBy(n int, divisibleby int) bool {
	return (n % divisibleby) == 0
}

// func FuzzSearchLables(data []) int {
// // First with one lable set
// 	if len(data) == 0 {
// 		return -1
// 	}
// 	if !IsDivisbleBy(len(data), 2) {
// 		return -1
// 	}

// 	var divided [][]byte

// 	chunkSize := len(data) / 2

// 	for i := 0; i < len(data); i += chunkSize {
// 		end := i + chunkSize

// 		divided = append(divided, data[i:end])
// 	}

	

// }