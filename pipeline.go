package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

func removeDuplicate(upstream, downstream chan string) {
	pushed := make(map[string]bool)
	for item := range upstream {
		if !pushed[item] {
			pushed[item] = true
			downstream <- item
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	c2 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	go removeDuplicate(c1, c2)
	printGopher(c2)
}
