package integers

import (
	"fmt"
	"testing"
)

func IntegersTest(t *testing.T) {
	assertMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Add shall work", func(t *testing.T) {
		got := Add(5, 7)
		want := 5 + 7
		assertMessage(t, got, want)
	})

}

func ExampleAdd() {
	got := Add(2, 3)
	fmt.Println(got)
	// Output: 5
}
