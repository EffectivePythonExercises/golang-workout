package main

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.IntN(len(charset))]
	}
	return string(b)
}
func BenchmarkFixed(b *testing.B) {
	funcs := []struct {
		name string
		do   func(*string)
	}{
		{name: "PlusOp", do: FixedLengthAddOps},
		{name: "InPlaceAdd", do: FixedLengthInPlaceAddOps},
	}

	cases := []struct {
		name    string
		a, b, c string
	}{
		{"CharLen 2", randString(2), randString(2), randString(2)},
		{"CharLen 10", randString(10), randString(10), randString(10)},
		{"CharLen 100", randString(100), randString(100), randString(100)},
	}

	for _, tc := range cases {
		for _, f := range funcs {
			var result string
			b.Run(fmt.Sprintf("[%s] %s", f.name, tc.name), func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					// f.do(&result, tc.a, tc.b, tc.c)
					f.do(&result)
				}
				b.StopTimer()
			})
		}
	}
}
