package lib

import "net/http"

//Simple Directory Listing
func DirList(localfolder, webpath string) (err error) {
	fs := http.FileServer(http.Dir(localfolder))
	http.Handle(webpath, http.StripPrefix(webpath, fs))
	return
}
