package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Mikita")

	got := buffer.String()
	want := "Hello, Mikita"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
