package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	pcqq = `https://down.qq.com/qqweb/PCQQ/PCQQ_EXE/PCQQ2020.exe`
)

func downloader() {
	resp, err := http.Get(pcqq)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	out, _ := os.Create(`pcqq2020.exe`)
	defer out.Close()

	l, _ := io.Copy(out, resp.Body)
	fmt.Println(l)

}

func main() {
	downloader()
}
