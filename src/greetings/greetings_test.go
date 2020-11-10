package greetings

import (
	"testing"
	"regexp"
)

// TestHelloName calls geetings.HelloError with a name,
// checking for a valid return value
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := HelloError("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`HelloError("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}


// TestHelloEmpty calls greetings.Hello with an empty string, 
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := HelloError("")
    if msg != "" || err == nil {
        t.Fatalf(`HelloError("") = %q, %v, want "", error`, msg, err)
    }
}
