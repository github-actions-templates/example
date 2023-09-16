package example

import "testing"

func TestExample(t *testing.T) {
	if Example() != "Hello, world!" {
		t.Error("Expected 'Hello, world!'")
	}
}
