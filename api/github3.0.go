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
	Repo1   = fakeSearch("repository1")
	Repo2   = fakeSearch("repository2")
	Code1   = fakeSearch("code1")
	Code2   = fakeSearch("code2")
	Issues1 = fakeSearch("issues1")
	Issues2 = fakeSearch("issues2")
)

func Github(query string) (results []Result) {
	// START OMIT
	c := make(chan Result)
	go func() { c <- First(query, Repo1, Repo2) }()
	go func() { c <- First(query, Code1, Code2) }()
	go func() { c <- First(query, Issues1, Issues2) }()
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

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) {
		c <- replicas[i](query)
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
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
