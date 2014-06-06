// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

var (
	Repo   = fakeSearch("repository")
	Code   = fakeSearch("code")
	Issues = fakeSearch("issues")
)

func Github(query string) (results []Result) {
	// START OMIT
	c := make(chan Result)
	go func() { c <- Repo(query) }()
	go func() { c <- Code(query) }()
	go func() { c <- Issues(query) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
	// STOP OMIT
}

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Github("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
