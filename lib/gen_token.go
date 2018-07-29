package lib

import (
	"fmt"
	"crypto/rand"
)

func Generate_Token() string {
    var b = make([]byte, 10)
    rand.Read(b)
    return fmt.Sprintf("%x", b)
}
