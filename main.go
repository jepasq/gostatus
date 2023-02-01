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

const keyServerAddr = "serverAddr"
var templateDirs = []string{"content"}

/// Return the intro text
func getIntroText() (string) {
	programName := os.Args[0]

	return "Welcome to gostatus v0.0.0-3 (" + programName + ")"
}

func getTemplates() (templates *template.Template, err error) {
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

	templates, err = template.New("root.tmpl").ParseFiles(allFiles...)
	return
}


/**
  * w the HTTP reponse writer
  * t The templater name relative to template pdirectory
 */
func writeTemplate(w http.ResponseWriter, t string) {
	templates, err := getTemplates();
	if err != nil {
		fmt.Printf("Failed to get templates: '%s'\n", err)
	}
	err = templates.Execute(w, t)
	if err != nil {
		fmt.Printf("can't execute template '%s': '%s'\n", t, err)
	}
}


func main() {
	// Capture connection properties.
	/*    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "recordings",
    }
	*/
	// Get a database handle.
	/*    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }
    
    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
	*/
	
	fmt.Println(getIntroText())
	

	argLength := len(os.Args[1:])
	fmt.Printf("Arg length is %d\n", argLength) 

	for i, a := range os.Args[1:] {
		fmt.Printf("Arg %d is %s\n", i+1, a) 
	}
	
	// Basic HTTP server
	fmt.Println("Listening http://localhost:3333 and opening browser...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/form",  getForm)

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

	writeTemplate(w, "root");
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
	
	writeTemplate(w, "hello");
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

	template, err := template.ParseFiles("content/form.tmpl")
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
