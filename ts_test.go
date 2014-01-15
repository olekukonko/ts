package ts

import (
	"fmt"
	"testing"
)

func ExampleGetSize() {
	size , _ := GetSize()
	fmt.Println(size.Col())  // Get Width
	fmt.Println(size.Row())  // Get Height
	fmt.Println(size.PosX()) // Get X position
	fmt.Println(size.PosY()) // Get Y position
}

func TestSize(t *testing.T) {
	size , err := GetSize()

	if err != nil {
		t.Fatal(err)
	}
	if size.Col() == 0 || size.Row() == 0 {
		t.Fatalf("Screen Size Failed")
	}
}
