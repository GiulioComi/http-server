package lib

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"io/ioutil"
)


//Prevent path traversal attacks with a regex based on whitelist characters
func sanitize(filename_input string) bool {
	safe_filename := regexp.MustCompile(`^([[:alnum:]])+\.([[:alnum:]])+$`)
	return safe_filename.MatchString(filename_input)
}

// UploaderEndpoint handles file uploading.                                                                                                                                                                  
// It responds to GET requests with the file upload form, and to POST                                                                                                                                        
// requests with the actual uploading.  
func UploaderEndpoint(name, path, webpath, token string) http.HandlerFunc {

	var form,err = ioutil.ReadFile("./resources/upload_form.html")
  	if err != nil {
  	      log.Fatal(err)
  	}

	webpath = name + webpath
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			data := make([]byte, 8)
			_, err := rand.Read(data)

		t := template.Must(template.New("uploadform").Parse(string(form)))
			fill := struct {
				Endpoint string
			}{}
			fill.Endpoint = path
			//TODO use a cookie instead
			err = t.Execute(w, fill)
			if err != nil {
				log.Println(err)
			}
		} else if r.Method == "POST" {
			_ = r.ParseMultipartForm(32 << 20)
			file, handler, err := r.FormFile("uploadfile")
			token_input := r.Form["token"]
			if  token_input[0] == token {
			if sanitize(handler.Filename) {
				if err != nil {
					fmt.Println(err)
					return
				}
				defer func() { _ = file.Close() }()
				f, err := os.OpenFile("/tmp/http-server/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer func() { _ = f.Close() }()
				n, err := io.Copy(f, file)
				if err != nil {
					fmt.Fprintf(w, "Errors occurred")
					log.Println(err)
					return
				}
				fmt.Fprintf(w, "<h1> Uploaded %d bytes</h1><a href='"+webpath+"'>Back to dirlist</a><br><a href='"+path+"'>Back to upload form</a>", n)
			} else {
				fmt.Fprintf(w, "<h1>Invalid filename</h1>")
			}
			} else {
				fmt.Fprintf(w, "<h1>Invalid token</h1>")
			}
		} else {
			http.Error(w, "Invalid method.", http.StatusMethodNotAllowed)
		}
	}
}
