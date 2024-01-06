package iteration

import (
	"fmt"
	"testing"
)

func TestIterate(t *testing.T) {
	repeated := Iterate("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleIterate() {
	repeated := Iterate("a", 0)
	fmt.Printf(repeated)
	// Output: "aaaaa"
}

func BenchmarkIterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterate("a", 0)
	}
}
