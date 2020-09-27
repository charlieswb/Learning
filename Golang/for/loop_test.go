package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("x", 7)
	expected := strings.Repeat("x", 7)

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	text := Repeat("a", 5)
	fmt.Println(text)
	// Output: aaaaa

}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("j", 7)
	}
}
