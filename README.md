This README would normally document whatever steps are necessary to get your application up and running.

### What is this repository for? ###
This repository is for creating skelaton for Microservice application.

### How do I get set up? ###
1. Run the command `go get "github.com/Kaustubh-pande/scaffolding-microservice"`
2. Change directory to `cd $GOPATH/src/github.com/Kaustubh-pande/scaffolding-microservice`
2. To Run this project follow below steps:
	- build the binary with this command
	`go build -o bin/main main.go`
	- Run below command to skeleton your project

	`./bin/main -path="Mentioned path where you want to create folder" -name="Name of microservice" -proto="Enter proto file name -username="provider/username"`
	
	- flags use as
		- -path : Insert full path where you want to create project.
		- -name : Your microservice name.
		- -proto : Your proto file name.