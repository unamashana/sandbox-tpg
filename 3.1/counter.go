package counter

import (
	"bufio"
	"io"
	"os"
)

type counter struct {
	input io.Reader
}

type option func(counter) counter

func (c *counter) Lines() int {
	lines := 0
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		lines++
	}

	return lines
}

func WithInput(i io.Reader) option {
	return func(c counter) counter {
		c.input = i
		return c
	}
}

func NewCounter(opts ...option) *counter {
	c := counter{
		input: os.Stdout,
	}

	for _, opt := range opts {
		c = opt(c)
	}

	return &c
}

func Lines() int {
	return NewCounter().Lines()
}
