package lib

import (
	"os"
	"fmt"
	"os/exec"
)

func check_cert() bool {
        if _, err := os.Stat("./cert/key.pem"); err != nil {
                return false
        }
        if _, err := os.Stat("./cert/cert.pem"); err != nil {
                return false
        }
        return true
}

func Handle_Cert() {
        if check_cert() {
                fmt.Println("Certs already created.")
         } else {
                fmt.Println("Certs missing, creating them...")
                exec.Command("/bin/sh", "-c", "mkdir ./cert; openssl req -x509 -newkey rsa:2048 -keyout './cert/key.pem' -out './cert/cert.pem' -days 3000 -nodes  -subj '/C=IT/ST=Italy/L=Milan/O=go-http-server/CN=http-server'").Run()
        }
}
