package main

import (
	//	"fmt"
	"testing"
	"regexp"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

/// intro text copntains 'Welcome' word
func TestIntroWelcome(t *testing.T) {
	name := "Welcome"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg := getIntroText()
	if !want.MatchString(msg) {
		t.Fatalf(`Can't find Welcome text =
%q, want match for %#q, nil`, msg, want)
	}
}

/// We can render correctly the hello page.
/// see https://blog.questionable.services/article/testing-http-handlers-go/
func TestHelloPageRendering(t *testing.T) {
	name := "You're in the HELLO page!"
	want := regexp.MustCompile(`\b`+name+`\b`)
	
	// Create a request to pass to our handler. We don't have any query
	// parameters for now, so we'll pass 'nil' as the third parameter.
	_ /*req*/ , err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter)
	// to record the response.
	rr := httptest.NewRecorder()
	writeTemplate(rr, "hello")
	//	t.Fatalf(`HELLO page rendering failed %q`, want,)
	res := rr.Result()	
	data, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(data))
	if err != nil {
	    t.Errorf("expected error to be nil got %v", err)
	}
	
	if !want.MatchString(string(data)) {
		t.Fatalf(`HELLO page rendering failed`)
	}
}
