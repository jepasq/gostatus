// Package gostatus implements a microservices monitoring web application
// written in go.
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"path/filepath"

	"text/template"
	"github.com/pkg/browser"
)

// Package constants
const (
	keyServerAddr = "serverAddr"  // the key server address
)
// Defines where the template partials are searched for
var templateDirs = []string{"../content", "../content/partials"}

/// Return the intro text
func getIntroText() (string) {
	programName := os.Args[0]

	return "Welcome to gostatus v0.0.0-8 (" + programName + ")"
}

/// Return a configuration directory (~/.gostatus)  content
func getCfgDirectory(content string) (string) {
	ud, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting User's home dir: '%s'\n", err)
	}
	return ud + "/.config/gostatus/" + content
}

/// Print usage text to terminal
func usage() {
	fmt.Println("Usage :");
	fmt.Println("\n  --h|-?|--help Print this usage text and exit.");
}

func getAllFiles() ([]string) {
	var allFiles []string
	for _, dir := range templateDirs {
		files2, _ := ioutil.ReadDir(dir)
		for _, file := range files2 {
			filename := file.Name()
			if strings.HasSuffix(filename, ".tmpl") {
				filePath := filepath.Join(dir, filename)
				allFiles = append(allFiles, filePath)
			}
		}
	}
	return allFiles
}	

func getTemplates(tmp string) (templates *template.Template, err error) {
	var allFiles = getAllFiles();
	templates, err = template.New(tmp).ParseFiles(allFiles...)
	return
}


/**
  * w the HTTP reponse writer
  * t The template name relative to template directory
 */
func writeTemplate(w http.ResponseWriter, t string) {
	templates, err := getTemplates(t);
	if err != nil {
		fmt.Printf("Failed to get templates: '%s'\n", err)
	}
	
	err = templates.Execute(w, t)
	if err != nil {
		fmt.Printf("can't execute template '%s': '%s'\n", t, err)
		fmt.Printf("Known templates are : '%s'\n",
			strings.Join(getAllFiles(), ", "))
	}
}


func main() {
	fmt.Println(getIntroText())

	arr := SidebarArray();
	fmt.Println(arr);
	
	/*argLength := len(os.Args[1:])
	fmt.Printf("Arg length is %d\n", argLength) 
	*/

	fmt.Printf("Using interface type '%s'\n", IfacetypeToStr(Web))
	
	for _, a := range os.Args[1:] { // Argument index, argument
		//   fmt.Printf("Arg %d is %s\n", i+1, a)
		if (a == "--help" || a== "-?" || a== "-h") {
			usage()
			os.Exit(0)
		}
	}
	
	// Basic HTTP server
	fmt.Println("Listening http://localhost:3333 and opening browser...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/form",  getForm)

	// Trying to handle static content
	mux.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))
	
	// Auto-open browser at startup
	err := browser.OpenURL("localhost:3333");
	if err != nil {
		fmt.Printf("can't find template file: '%s'\n", err)
	}

	ctx, cancelCtx := context.WithCancel(context.Background())
	serverOne := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}
	go func() {
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
		cancelCtx()
	}()
	<-ctx.Done()
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	// Ok for '/', root can be handled, everything else is 404
	// It can works with only this test because getRoot() is a kind
	// of fallback for all not-yet-handled URLs.
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}	

	ctx := r.Context()
	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")

	fmt.Printf("%s: got / request. first(%t)=%s, second(%t)=%s\n",
		ctx.Value(keyServerAddr),
		hasFirst, first,
		hasSecond, second)

	writeTemplate(w, "root.tmpl");
}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("could not read body: '%s'\n", err)
	}
	
	fmt.Printf("%s: got /hello request\nbody:\n%s",
		ctx.Value(keyServerAddr),
		body)
	
	writeTemplate(w, "hello.tmpl");
}

/// Used to pass variable to the Post page
type PostVars struct {
	Name string
}

func getForm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got /form request\n", ctx.Value(keyServerAddr))

	myName := r.PostFormValue("myName")
	if myName == "" {
		myName = "&lt;empty&gt;"
		/*
		w.Header().Set("x-missing-field", "myName")
		w.WriteHeader(http.StatusBadRequest)
		return
		*/
	}

	template, err := getTemplates("form.tmpl")
	if err != nil {
		fmt.Printf("can't find template file: '%s'\n", err)
	}

	mn := PostVars{
		Name: myName,
	}
	
	err = template.Execute(w, mn)
	if err != nil {
		fmt.Printf("Error processing template 'form.tmpl': '%s'\n", err)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
