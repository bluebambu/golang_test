package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i<100; i++ {
		jobs <- i
	}
	close(jobs)

	for  j:=0; j<100; j++ {
		(<-results);
	}

	fmt.Println("Hello world");
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		fmt.Println("fetching from routine id: ", goid())
		tmp := fib(n)
		results <- tmp
		fmt.Println("calculated: ", tmp)
	}
}

func fib(n int) int {
	if n<= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
