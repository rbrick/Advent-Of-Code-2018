package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

var counter *Counter

type (
	Operation int
	Counter struct {
		input   *bufio.Scanner
		workers chan Worker

		done chan bool

		wg sync.WaitGroup

		m            sync.Mutex
		currentCount int64
	}

	Worker struct {
		operation Operation
		delta     int64
	}
)

const (
	Add Operation = iota
	Sub
)

func (c *Counter) readLine() {
	s := c.input.Text()

	var operation Operation

	switch s[0] {
	case '+':
		operation = Add
	case '-':
		operation = Sub
	}

	i, _ := strconv.Atoi(s[1:])
	w := Worker{
		operation: operation,
		delta:     int64(i),
	}
	c.workers <- w
}

func (c *Counter) exec() {
	for  {
		select {
		case x := <- c.workers:
			c.m.Lock()
			switch x.operation {
			case Add:
				c.currentCount += x.delta
			case Sub:
				c.currentCount -= x.delta
			}
			c.m.Unlock()
		case <- c.done:
			fmt.Println(c.currentCount)
			c.wg.Done()
		}
	}
}

func init() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	counter = &Counter{
		input:        bufio.NewScanner(f),
		workers: make(chan Worker),
		done: make(chan bool),
		m:            sync.Mutex{},
		currentCount: 0,
	}
}

func main() {
	go counter.exec()
	go func() {
		for counter.input.Scan() {
			counter.readLine()
		}
		counter.done <- true
	}()

	counter.wg.Add(1)
	counter.wg.Wait()

	// my solution is 516
}
