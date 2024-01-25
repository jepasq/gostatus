# gostatus

A microservices monitoring application written in *go*.

# Dependencies

## Install

To build and run this project, you need the Go programming language core
compiler.

To install it, on *Manjaro* :

	sudo pacman install go go-tools
	
on *Debian* :

	sudo apt-get install golang-go golang-golang-x-tools

The *-tools packages usually contain *godoc* package used to generate and 
browse API documentation.

## Upgrade

Once installed, you may need to upgrade dependencies. From the `src/`
directory, run :

	go list -u -m all

to list upgradable dependencies

	go get -u ./...
	
to upgrade.


# Build and run

To compile and build dependencies and binary run these commands in *src/* :

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

From the  *src* directory, run the following command :
	
	make doc

This should automatically open the http://127.0.0.1:6060/pkg/jepasq/gostatus/ 
URL in your favorite/system browser.
