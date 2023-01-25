package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	p := "2SL22uyTMdQOO06lW7Mzow=="
	buf, err := base64.StdEncoding.DecodeString(p)
	if err != nil {
		log.Fatal(err)
	}
	b := sha256.Sum256(buf)
	fmt.Printf("%x\n", b)

}
