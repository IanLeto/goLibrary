package utils

import (
	"testing"
)

func fetch1(s, xs []string) {
	for _, s2 := range s {
		for i := 0; i < len(xs); i++ {
			if s2 == xs[i] {
				xs[i] = xs[len(xs)-1]
				xs[len(xs)-1] = ""
				xs = xs[:len(xs)-1]
				break
			}
		}
	}
}

func fetch2(s, xs []string) {
	for _, s2 := range s {
		for _, x := range xs {
			if s2 == x {
				break
			}
		}
	}
}
func BenchmarkList(b *testing.B) {
	s := []string{"12", "ddsds", "xcc", "x2cc"}
	xs := []string{"ddsds", "x2cc", "12", "xcc"}
	for i := 0; i < b.N; i++ {
		fetch1(s, xs)
	}

}

func BenchmarkList2(b *testing.B) {
	s := []string{"12", "ddsds", "xcc", "x2cc"}
	xs := []string{"ddsds", "x2cc", "12", "xcc"}
	for i := 0; i < b.N; i++ {
		fetch2(s, xs)
	}

}
