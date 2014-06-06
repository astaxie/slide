// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"runtime"
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
	c := make(chan Result)
	go func() { c <- Repo(query) }()
	go func() { c <- Code(query) }()
	go func() { c <- Issues(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Github("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
