package ts


import (
	"fmt"
	"testing"
)

func ExampleGetSize() {
	size := GetSize()
	fmt.Println(size.Col()) // Get Width
	fmt.Println(size.Row()) // Get Height
	fmt.Println(size.PosX())   // Get X position
	fmt.Println(size.PosY())   // Get Y position
}

func TestSize(t *testing.T) {
	size := GetSize()
	if size.Col() == 0 || size.Row() == 0 {
		t.Fatalf("Screen Size Failed")
	}
}
