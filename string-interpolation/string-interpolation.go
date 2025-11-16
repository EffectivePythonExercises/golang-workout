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
	DELIM      string = "_" // delimiter
	REPETITION uint   = 1_000_000
)

var FIXED_STRINGS = [...]string{"Hello", "Goody", "Buddy"}

func main() {
	FixedLenInterpolationBenchmarks()
	VariableLenInterpolationBenchmarks()
}

func FixedLenInterpolationBenchmarks() {
	fmt.Println("================================ FIXED LENGTH ================================")

	Tsukuyomi(FixedLengthAddOps)

	Tsukuyomi(FixedLengthInPlaceAddOps)

	Tsukuyomi(FixedLengthUseBuffer)

	Tsukuyomi(FixedLengthSprintf)

	Tsukuyomi(FixedLengthJoin)
}

func VariableLenInterpolationBenchmarks() {
	fmt.Println("================================ VARIABLE LENGTH ================================")

	MugenTsukuyomi(VariableLengthAddOps)

	MugenTsukuyomi(VariableLengthUseBuffer)

	MugenTsukuyomi(VariableLengthSprintf)

	MugenTsukuyomi(VariableLengthJoin)

}

func Tsukuyomi(fn func(*string)) {
	start := time.Now()

	flec := reflect.ValueOf(fn).Pointer()
	pc := runtime.FuncForPC(flec)
	funcname := pc.Name()

	fmt.Printf("==== %s ====\n", funcname)

	var result string
	for range REPETITION {
		fn(&result)
	}

	elapsed := time.Since(start)

	fmt.Println("    * Result:", result)
	fmt.Printf("    * Function %-55s took %v\n\n", funcname, elapsed)
}

func MugenTsukuyomi(fn func(...string) string) {
	start := time.Now()

	flec := reflect.ValueOf(fn).Pointer()
	pc := runtime.FuncForPC(flec)
	funcname := pc.Name()

	fmt.Printf("==== %s ====\n", funcname)

	arguments := []string{"a", "b", "AAAA", "dgs", "'$$ssdfs22'", "'9jsdfjsdf'", "'(TEST)'", "'^^^^^'"}
	var result string

	for range REPETITION {
		result = fn(arguments...)
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

func VariableLengthAddOps(ss ...string) string {
	if len(ss) == 0 {
		return ""
	}

	concatenated := ""

	for i, s := range ss {
		concatenated += s
		if i != len(ss)-1 {
			concatenated += DELIM
		}
	}
	return concatenated
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

func VariableLengthUseBuffer(ss ...string) string {
	if len(ss) == 0 {
		return ""
	}

	var buffer bytes.Buffer

	for i, s := range ss {
		buffer.WriteString(s)
		if i != len(ss)-1 {
			buffer.WriteString(DELIM)
		}
	}
	concatenated := buffer.String()
	return concatenated
}

func FixedLengthSprintf(result *string) {
	var s1, s2, s3 string
	s1 = FIXED_STRINGS[0]
	s2 = FIXED_STRINGS[1]
	s3 = FIXED_STRINGS[2]

	concatenated := fmt.Sprintf("%s%s%s%s%s", s1, DELIM, s2, DELIM, s3)

	*result = concatenated
}

func VariableLengthSprintf(ss ...string) string {
	if len(ss) == 0 {
		return ""
	}

	var items = make([]any, len(ss)*2-1)

	k := 0
	for i := 0; i < len(items); i++ {
		if i%2 == 0 {
			items[i] = ss[k]
			k += 1
		} else {
			items[i] = DELIM
		}

	}
	foramtString := strings.Repeat("%s", len(items))
	concatenated := fmt.Sprintf(foramtString, items...)
	return concatenated
}

func FixedLengthJoin(result *string) {
	concatenated := strings.Join(FIXED_STRINGS[:], DELIM)

	*result = concatenated
}

func VariableLengthJoin(ss ...string) string {
	if len(ss) == 0 {
		return ""
	}

	concatenated := strings.Join(ss, DELIM)
	return concatenated
}
