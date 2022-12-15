package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"io"
	"net"
	"net/http"
)

const keyServerAddr = "serverAddr"

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
	
	programName := os.Args[0]
	fmt.Println(programName)

	argLength := len(os.Args[1:])
	fmt.Printf("Arg length is %d\n", argLength) 

	for i, a := range os.Args[1:] {
		fmt.Printf("Arg %d is %s\n", i+1, a) 
	}
	
	// Basic HTTP server
	fmt.Println("Listening localhost:3333")
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)


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
	ctx := r.Context()
	fmt.Printf("%s: got / request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, `<HTML>
<BODY>This is my website!\n You can also go to ".
		"<a href='/hello'>HELLO</a> page.</BODY></HTML>`)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, `<HTML>
<BODY><p>You're in the HELLO page!</p><p><a href="..">back</a></BODY></HTML>`)
}
