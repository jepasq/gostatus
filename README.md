# gostatus

A microservices monitoring application written in *go*.

# Dependencies

To build and run this project, you need the Go programming language core
compiler.

To install it, on *Manjaro* :

	sudo pacman install go go-tools
	
on *Debian* :

	sudo apt-get install golang-go golang-golang-x-tools

The *-tools packages usually contain *godoc* package used to generate and 
browse API documentation.

# Build and run

To compile and build dependencies and binary :

	go build

Then, to run it :

	go run .

It will auto-open a new browser window with the *gostatus* home page.

# Unit tests

To run unit tests :

	go test
	
# Usage

To print a short usage text and exit, run :

	go run . --help

# API documenation

Run the following command :
	
	godoc -http=:6060

and open the page http://127.0.0.1:6060/pkg/jepasq/gostatus/ with your 
favorite browser.
