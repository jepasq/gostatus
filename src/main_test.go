package main

import (
	"os"
	"fmt"
	"regexp"
	"strings"           // USES strings.Contains
	"testing"
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

/// Test with an unexisting URL
func TestWrongUrlRequestShouldFail(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query
	// parameters for now, so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/unexistingPageUrl_com", nil)
	fmt.Println(req)
	fmt.Print("Err: ")
	fmt.Println(err)
	req, err = http.NewRequest("GET", "/home", nil)
	fmt.Println(req)

	if err == nil {
		t.Fatal("Requesting a non existing page should fail")
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
	writeTemplate(rr, "hello", nil)
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

/// We can render correctly the form page.
func TestFormPageRendering(t *testing.T) {
	name := "Form"
	want := regexp.MustCompile(`\b`+name+`\b`)
	
	// Create a request to pass to our handler. We don't have any query
	// parameters for now, so we'll pass 'nil' as the third parameter.
	_ /*req*/ , err := http.NewRequest("GET", "/form", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter)
	// to record the response.
	rr := httptest.NewRecorder()
	writeTemplate(rr, "form", nil)
	//	t.Fatalf(`HELLO page rendering failed %q`, want,)
	res := rr.Result()	
	data, err := ioutil.ReadAll(res.Body)
	//	fmt.Println(string(data))
	if err != nil {
	    t.Errorf("expected error to be nil got %v", err)
	}
	
	if !want.MatchString(string(data)) {
		t.Fatalf(`FORM page rendering failed`)
	}
}

func TestGetCfgDirectory(t *testing.T) {
	cd := getCfgDirectory("")

	// Must contain appname
	dir := "gostatus"
	want := regexp.MustCompile(`\b`+dir+`\b`)
	if !want.MatchString(cd) {
		t.Fatalf(`getCfgDirectory doesn't contain appname`)
	}

	// Must contain user's home dir
	ud, err := os.UserHomeDir()
	if err != nil {
		t.Fatal( err )
	}
	if !strings.Contains(cd, ud) {
		t.Fatalf(`getCfgDirectory doesn't contain User's home dir`)
	}

	// Must contain the string passed as parameter
	cd = getCfgDirectory("aze")
	if !strings.Contains(cd, "aze") {
		t.Fatalf(`getCfgDirectory doesn't contain the param content`)
	}
	
}
