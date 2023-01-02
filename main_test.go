package main

import (
	"testing"
	"regexp"
)

/// intro text copntains 'Welcome' word
func TestIntroWelcome(t *testing.T) {
	name := "Welcome"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg := getIntroText()
	if !want.MatchString(msg) {
		t.Fatalf(`Hello("Gladys") = %q, want match for %#q, nil`, msg, want)
	}
}
