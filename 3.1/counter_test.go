package counter_test

import (
	"bytes"
	"counter"
	"testing"
)

func TestCountLines(t *testing.T) {
	t.Parallel()

	s := "one \n two \n three"

	want := 3

	c := counter.NewCounter(
		counter.WithInput(bytes.NewBufferString(s)),
	)

	got := c.Lines()

	if want != got {
		t.Errorf("Wanted %d, Got %d", want, got)
	}

}
