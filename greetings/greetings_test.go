package greetings

import (
	"regexp"
	"strings"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Gladys")) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// // TestHelloEmpty calls greetings.Hello with an empty string,
// // checking for an error.

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

//* Table Driven Test

func TestHello(t *testing.T) {
	
	tests := []struct {
		name string
		wantErr bool
	} {
		{"Hyunwoo", false},
		{"Eunbi", false},
		{"", true},
	}

	for _, tt := range tests {

		msg, err := Hello(tt.name)

		if tt.wantErr {
			if err == nil {
				t.Errorf("Hello(%q) expected error", tt.name)
			}

			continue
		}

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !strings.Contains(msg, tt.name) {
			t.Errorf("Hello(%q) = %q", tt.name, msg)
		}
	}

}