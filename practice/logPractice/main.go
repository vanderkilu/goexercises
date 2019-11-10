package main

import (
	"log"
	"io"
	"io/ioutil"
	"os"
	
)

var (

	Trace *log.Logger
	Error *log.Logger
	Warning *log.Logger
	Info *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.txt", 
	os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Failed to open the file specified")
	}

	Trace = log.New(ioutil.Discard,"Trace: ", log.Ldate | log.Ltime | log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, file), "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {
	Trace.Println("testing the trace log")
	Error.Println("testing the error log")
	Warning.Println("testing the warning log")
	Info.Println("testing the info log")
}