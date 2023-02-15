package golang_tdd

import "testing"

func TestHello(t *testing.T) {

	got := Hello()
	want := "Hello"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloName(t *testing.T) {
	got := HelloName("Chris")

	want := "Hello Chris"

	assertCorrectMessage(t, got, want)
}

func TestHelloDefault(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloDefault("Chris")
		want := "Hello Chris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := HelloDefault("")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
