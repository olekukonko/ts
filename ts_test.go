package ts

import "testing"

// TODO Add test


func TestSize(t *testing.T) {
	size := GetSize()
	if size.Col() == 0 || size.Row() == 0 {
		t.Fatalf("Screen Size Failed")
	}
}
