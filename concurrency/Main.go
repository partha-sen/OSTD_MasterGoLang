package main

import (
	"time"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWord(lineChan chan string, countChan chan int) {
  var sum int
  for line := range lineChan {
	arr:=strings.Split(line, " ")
	sum += len(arr)
  }	
  countChan<-sum
}

func startCounter(lineChan chan string, countChan chan int){
	for i := 0; i < 8; i++ {
		go countWord(lineChan, countChan)
	}
}

func readFile(lineChan chan string){
	file, err := os.Open("C:/OSTD_MasterGoLang/concurrency/sample.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}	
	scanner := bufio.NewScanner(file)	
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lineChan<-scanner.Text()
	}
	close(lineChan)
}


func main() {	

	lineChan:=make(chan string, 8)
	countChan:=make(chan int)

	startCounter(lineChan, countChan)

	start := time.Now()

	go readFile(lineChan)
	

	var total int

	for i := 0; i < 8; i++ {
		total += <-countChan
	}

	fmt.Println("Total", total)
	
	duration := time.Since(start)
	fmt.Println(duration.Nanoseconds())
}
