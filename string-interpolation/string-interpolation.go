package main

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

const (
	DELIM      = "_" // delimiter
	REPETITION = 1_000_000
)

var FIXED_STRINGS = [...]string{"Hello", "Goody", "Buddy"}

func main() {
	fmt.Println("================================ FIXED LENGTH ================================")

	Tsukuyomi(FixedLengthAddOps)

	Tsukuyomi(FixedLengthInPlaceAddOps)

	Tsukuyomi(FixedLengthUseBuffer)

	Tsukuyomi(FixedLengthSprintf)

	Tsukuyomi(FixedLengthJoin)

}

func Tsukuyomi(fn func(*string)) {
	start := time.Now()

	flec := reflect.ValueOf(fn).Pointer()
	pc := runtime.FuncForPC(flec)
	funcname := pc.Name()

	fmt.Printf("==== %s ====\n", funcname)

	var result string
	for i := 0; i < REPETITION; i++ {
		fn(&result)
	}

	elapsed := time.Since(start)

	fmt.Println("    * Result:", result)
	fmt.Printf("    * Function %-55s took %v\n\n", funcname, elapsed)
}

func FixedLengthAddOps(result *string) {
	var s1, s2, s3 string
	s1 = FIXED_STRINGS[0]
	s2 = FIXED_STRINGS[1]
	s3 = FIXED_STRINGS[2]

	concatenated := s1 + DELIM + s2 + DELIM + s3

	*result = concatenated
}

func FixedLengthInPlaceAddOps(result *string) {
	var s1, s2, s3 string
	s1 = FIXED_STRINGS[0]
	s2 = FIXED_STRINGS[1]
	s3 = FIXED_STRINGS[2]

	concatenated := ""
	concatenated += s1
	concatenated += DELIM
	concatenated += s2
	concatenated += DELIM
	concatenated += s3

	*result = concatenated
}

func FixedLengthUseBuffer(result *string) {
	var s1, s2, s3 string
	s1 = FIXED_STRINGS[0]
	s2 = FIXED_STRINGS[1]
	s3 = FIXED_STRINGS[2]

	var buffer bytes.Buffer

	buffer.WriteString(s1)
	buffer.WriteString(DELIM)
	buffer.WriteString(s2)
	buffer.WriteString(DELIM)
	buffer.WriteString(s3)

	concatenated := buffer.String()

	*result = concatenated
}

func FixedLengthSprintf(result *string) {
	var s1, s2, s3 string
	s1 = FIXED_STRINGS[0]
	s2 = FIXED_STRINGS[1]
	s3 = FIXED_STRINGS[2]

	concatenated := fmt.Sprintf("%s%s%s%s%s", s1, DELIM, s2, DELIM, s3)

	*result = concatenated
}

func FixedLengthJoin(result *string) {
	concatenated := strings.Join(FIXED_STRINGS[:], DELIM)

	*result = concatenated
}
