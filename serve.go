package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"./lib"
)

var localport = flag.String("p", "8086", "The port to run on")
var localfolder = flag.String("f", "/tmp/http-server/", "The folder to serve (default PWD)")
var uploadpath = flag.String("u", "upload", "The webpath where the upload form is hosted")
var webpath = flag.String("w", "/", "The root webpath")
var certpath = "./cert/"
var token = ""

func init() {
	flag.Parse()
	flag.Arg(1)
	if !strings.HasPrefix(*localport, ":") {
		*localport = ":" + *localport
	}
}

func main() {
	setup()
	token = lib.Generate_Token()
	display_info(token)

	err := lib.DirList(*localfolder, *webpath)
	if err != nil {
		log.Fatal(err)
	}


	http.HandleFunc(*webpath+*uploadpath, lib.UploaderEndpoint(Name(), *uploadpath, *webpath, token))

	err = http.ListenAndServeTLS(*localport, certpath+"cert.pem", certpath+"key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Name returns the full URL for the local port and webpath passed from the command line
func Name() string {
	b := bytes.Buffer{}
	_, _ = b.WriteString("https://")
	_, _ = b.WriteString(lib.MyName())
	_, _ = b.WriteString(*localport)
	_, _ = b.WriteString(*webpath)
	return b.String()
}

func display_info(token string) {
	fmt.Println("Local ip address: " + lib.MyIP().String())
	fmt.Printf("Serving local folder %v on \"%s\"\n", *localfolder, Name())
	fmt.Printf("File Upload form available at \"%s%s\"\n", Name(), *uploadpath)
	fmt.Printf("Token to use for upload:	\"%s\"\n", token)
	fmt.Println("Press Control+C to stop")
}

func setup() {
	lib.Handle_Cert()
	lib.PrepareWorkingDir()
}

