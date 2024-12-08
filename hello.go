package hello

import "fmt"

// Hello returns hello message
func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
