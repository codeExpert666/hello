package hello_test

import (
	"fmt"
	"testing"

	"github.com/codeExpert666/hello"
)

func TestHello(t *testing.T) {
	data := "Jack"
	expected := fmt.Sprintf("Hello, %s!", data)
	result := hello.Hello(data)

	if result != expected {
		t.Fatalf("expected result %s, but got %s", expected, result)
	}
}
