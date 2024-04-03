package main

import "testing"

func TestHello(t *testing.T) {

	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Mike", "")
		want := "Hello Mike"
		assertMessage(t, got, want)
	})

	t.Run("Say hello world if no name is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		assertMessage(t, got, want)
	})

}
