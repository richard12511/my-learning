package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main(){
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

}