package example

import "testing"

func TestExample(t *testing.T) {
	// success example
	if Example() != "Hello, world!" {
		t.Error("Expected 'Hello, world!'")
	}

	// error example
	if Example() != "Hello, world" {
		t.Error("Expected 'Hello, world!'")
	}
}
