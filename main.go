package main

import (
	"fmt"
	"os"
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
	
	fmt.Println("Connected!")
}
