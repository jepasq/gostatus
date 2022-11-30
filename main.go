package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"net/http"
)

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
	http.HandleFunc("/", getRoot)
	err := http.ListenAndServe(":3333", nil)
	fmt.Println("Listening localhost:3333")
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
