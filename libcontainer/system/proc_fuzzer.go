package system

import (
	// "errors"
	// "math/bits"
	// "os"
	// "reflect"
	// "strconv"
	"strings"
	// "testing"
)

func FuzzParseStat(data []byte) int {
	_, _ := parseStat(strings.NewReader(string(data)))
	return 1
}