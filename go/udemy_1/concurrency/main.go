package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
		"http://target.com",
	}	

	linkChan := make(chan string)

	for _, l := range links {
		go func(link string) { 
			time.Sleep(2 * time.Second)
			checkLink(link, linkChan)
		}(l)
	}


	for msg := range linkChan {
		go fmt.Println(msg)
	}

}

func checkLink(link string, c chan string){
	resp, err := http.Get(link)

	if err != nil {
		c <- err.Error()
	}

	c <- resp.Status + " " + link
}