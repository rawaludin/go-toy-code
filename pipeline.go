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

func splitWords(words string, downstream chan string) {
	for _, word := range strings.Fields(words) {
		downstream <- word
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

	c3 := make(chan string)
	go splitWords("Sometimes it’s easier to operate on words than on sentences. Write a pipeline element that takes strings, splits them up into words (you can use the Fields function from the strings package), and sends all the words, one by one, to the next pipeline stage.", c3)
	printGopher(c3)
}
