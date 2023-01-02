package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

const keyServerAddr = "serverAddr"

/// Return the intro text
func getIntroText() (string) {
	programName := os.Args[0]

	return "Welcome to gostatus v0.0.0-2 (" + programName + ")"
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
	fmt.Println("Listening http://localhost:3333")
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/form",  getForm)

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
	
	io.WriteString(w, `<HTML>
<BODY>This is my website!<br>
You can also go to <a href='/hello'>HELLO</a> page.<br>
Or test the <a href='/form'>FORM</a> posting page<br>
<br>
Yoy may also want to test this page with parameters
<a href="?first=aze">here</a> and 
<a href="?first=aze&second=zer">here</a>.
</BODY></HTML>`)
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
	
	io.WriteString(w, `<HTML>
<BODY><p>You're in the HELLO page!</p>
<p>And read logs, we're reading full request body!<br>
To test this, contactthis server with curl like :
<pre>curl -X POST -d 'This is the body' 'http://localhost:3333/hello</pre>

</p>
<a href="..">back</a></BODY></HTML>`)
}

func getForm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got /form request\n", ctx.Value(keyServerAddr))

	myName := r.PostFormValue("myName")
	if myName == "" {
		myName = "<empty>"
		/*
		w.Header().Set("x-missing-field", "myName")
		w.WriteHeader(http.StatusBadRequest)
		return
		*/
	}
	io.WriteString(w, fmt.Sprintf(`<html><body>
myName form value is %s!<br>
<form action="" method="post">
<fieldset>
<legend>Change the posted name :</legend>
  <input name="myName"></input>
  <input type="submit" value="Post!">
</fieldset>
</form>

`, myName))
	
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
