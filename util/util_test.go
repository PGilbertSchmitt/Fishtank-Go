package util

import "testing"

func TestHash32(t *testing.T) {
	var goodHash = []struct {
		in  Vector // Input
		out uint32 // Expected Output
	}{
		{Vector{1, 4}, 1810974526},
		{Vector{5, 7}, 1391114847},
		{Vector{-4, 9}, 1447020315},
	}

	for _, v := range goodHash {
		out := v.out
		actual := Hash32(v.in)
		if out != actual {
			t.Errorf("Hash32(): Expected %v, Received %v", out, actual)
		}
	}
}
