package hello_test

import (
	"bytes"
	"hello"
	"testing"
)

func TestPrintHelloMessageToWriter(t *testing.T) {
	t.Parallel()

	fakeTerminal := &bytes.Buffer{}

	p := &hello.Printer{
		Output: fakeTerminal,
	}

	p.Print()

	want := "Hello World"
	got := fakeTerminal.String()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
