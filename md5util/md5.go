package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	m, err := Md5File("file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}

func Md5File(filename string) (s string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		// log.Fatal(err)
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		// log.Fatal(err)
		return "", nil
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil

}
