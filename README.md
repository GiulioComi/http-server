This simple HTTPS server provides download (with directory listing feature) and upload functionalities to ease the process of sharing files between computers/phones/tablets that are in the same network. It also comes handy for pentesting purposes under certain circumtances :-).

## Features
* Supports TLS
* Directory listing for the files to download
* Upload files to the shared directory
* Finds local network ip address
* Provide some security checks (whitelist to prevent path traversal, token to prevent unauthorized uploads)

### Install
```
go get https://github.com/GiulioComi/http-server
```

## Usage

```
Usage of serve:
  -f string
    	The folder to serve (default /tmp/http-server)
  -p string
    	The port to run on (default "8086")
  -u string
    	The webpath where the upload form is hosted (default "upload/")
  -w string
    	The root webpath (default "/")
  -h prints this help
```

# Example of console output

```
./http-server 
Certs already created.
Local ip address: 10.0.2.15
Serving local folder /tmp/http-server/ on "https://10.0.2.15:8086/"
File Upload form available at "https://10.0.2.15:8086/upload"
Token to use for upload:	"81ffc3175283db2c4cf0"
Press Control+C to stop

```

## Example to upload files without GUI access to the machine

```curl -i -s -k  -X $'POST' \
    -H $'Host: 10.0.2.15:8086' -H $'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8' -H $'Accept-Language: en-US,en;q=0.5' -H $'Accept-Encoding: gzip, deflate' -H $'Referer: https://[URL]/upload' -H $'DNT: 1' -H $'Connection: close' -H $'Upgrade-Insecure-Requests: 1' -H $'Content-Type: multipart/form-data; boundary=---------------------------606768120669673484490201340' -H $'Content-Length: 140654' \
    --data-binary $'-----------------------------\x0d\x0aContent-Disposition: form-data; name=\"uploadfile\"; filename=\"[FILENAME]\"\x0d\x0aContent-Type: [CONTENT-TYPE]\x0d\x0a\x0d\x0a[CONTENT]\x0d\x0a-----------------------------\x0d\x0aContent-Disposition: form-data; name=\"token\"\x0d\x0a\x0d\x0ad0cd53e404919e8509f9\x0d\x0a-------------------------------\x0d\x0a' \
    $'https://[URL]/upload'```

## Credits
Thanks to empijei for initial inspiration (https://github.com/empijei/serve).
