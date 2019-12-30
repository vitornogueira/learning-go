package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, expected string) {
		t.Helper()
		if got != expected {
			t.Errorf("expected %q but got %q", expected, got)
		}
	}

	t.Run("repeat string with times suplied", func(t *testing.T) {
		got := Repeat("a", 10)
		expected := "aaaaaaaaaa"
		assertCorrectMessage(t, got, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	characters := Repeat("a", 6)
	fmt.Println(characters)
	// Output: aaaaaa
}
