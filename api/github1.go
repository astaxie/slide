// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string

// START1 OMIT
func Github(query string) (results []Result) {
	results = append(results, Repo(query))
	results = append(results, Code(query))
	results = append(results, Issues(query))
	return
}

// STOP1 OMIT

// START2 OMIT
var (
	Repo   = fakeSearch("repository")
	Code   = fakeSearch("code")
	Issues = fakeSearch("issues")
)

type Search func(query string) Result // HL

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// STOP2 OMIT

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Github("golang") // HL
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
