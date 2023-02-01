package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	lw := logWriter{}
	io.Copy(lw, resp.Body)
	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))
	fmt.Println(deneme("ali"))
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote many bytes: ", len(bs))
	return len(bs), nil

}

func deneme(aa string) string {
	return "deneme" + aa
}
