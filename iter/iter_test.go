package iter

import "testing"

func TestIter(t *testing.T) {
	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Repeat", func(t *testing.T) {
		got := Repeat("wow", 3)
		want := "wowwowwow"
		assertMessage(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
