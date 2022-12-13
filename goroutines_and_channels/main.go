package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	portocal := "http://"
	links := make(map[string]string)
	links["google"] = "google.com"
	links["facebook"] = "facebook.com"
	links["golang"] = "golang.org"
	links["amazon"] = "amazon.com"
	links["stack"] = "stackoverflow.com"

	c := make(chan string)

	for _, link := range links {
		go checkLink(portocal + link, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(2 * time.Second)
			checkLink(l, c)
		}(l)
		fmt.Println("Up")
	}
}

func checkLink(links string,  c chan string) {
	_, err := http.Get(links)
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}
	c <- links

}
