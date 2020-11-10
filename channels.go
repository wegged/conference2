package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
)

func main() {
	url := "http://engadget.com"
	content := readUrlToString(url)

	result := letterFrequencyConcurrent(content)
	print(result)
}

func readUrlToString(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}


func letterFrequencySequential(content string) map[rune]int {
	return countLetters(content)
}

func letterFrequencyConcurrent(content string) map[rune]int {
	wg := sync.WaitGroup{}
	chunkSize := (len(content) + runtime.NumCPU() -1)/ runtime.NumCPU()
	c := make(chan map[rune]int)
	for i:=0;i<len(content);i+=chunkSize {
		end := i +chunkSize
		if end>len(content) {
			end = len(content)
		}
		wg.Add(1)
		go func(i,end int) {
			c <- countLetters(content[i:end])
			wg.Done()
		}(i,end)
	}



	//for _, s := range strings.Fields(content) {
	//	wg.Add(1)
	//	go func() {
	//		c <- countLetters(s)
	//		wg.Done()
	//	}()
	//}

	go func() {
		wg.Wait()
		close(c)
	}()

	result := aggregate(c)
	return result
}



func countLetters(s string) map[rune]int{
	result := make(map[rune]int,256)
	for _, i := range s {
		result[i]++
	}
	return result
}

func print(m map[rune]int) {
	for i, i2 := range m {
		fmt.Printf("%s: %d\n",string(i),i2)
	}
}

func aggregate(c chan map[rune]int) map[rune]int {
	result  := make(map[rune]int,256)
	for m := range c {
		for k, v := range m {
			result[k] += v
		}
	}
	return result
}