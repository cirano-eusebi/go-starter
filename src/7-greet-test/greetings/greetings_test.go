package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value
func TestHelloName(t *testing.T) {
	// Arrange.
	name := "GLaDOS"
	want := regexp.MustCompile(`\b` + name + `\b`)
	// Act.
	msg, err := Hello(name)
	// Assert.
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("GLaDOS") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	// Arrange.
	emptyString := ""
	// Act.
	msg, err := Hello(emptyString)
	// Assert.
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
