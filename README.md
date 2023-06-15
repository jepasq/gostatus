# gostatus

A microservices monitoring application written in *go*.

# Dependencies

To build and run this project, you need the Go programming language core
compiler.

To install it, on *Manjaro* :

	sudo pacman install go
	
on *Debian* :

	sudo apt-get install golang-go

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
