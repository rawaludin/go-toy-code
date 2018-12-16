package main

import (
	"fmt"
	"net/url"
)

func main() {
	urls := []string{
		"http://www.google.com",
		"http://in va lid.com",
		"http://www.facebook.com",
	}

	for _, u := range urls {
		valid, err := url.Parse(u)
		if err != nil {
			fmt.Printf("%#v\n", err)
			continue
		}
		fmt.Println(valid)
	}
}
