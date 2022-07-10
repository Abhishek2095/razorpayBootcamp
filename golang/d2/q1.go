package main

import (
	"fmt"
	"sync"
)

func countFrequency(word string) []int {
	frequencyMap := make([]int, 26)

	for _, letter := range word {
		frequencyMap[letter-rune('a')]++
	}

	return frequencyMap
}

func main() {
	var words []string = []string{"quick", "brown", "fox", "lazy", "dog"}
	const channelBufferSize = 3

	var ch = make(chan []int, channelBufferSize)
	var wg sync.WaitGroup

	wg.Add(len(words))

	for _, word := range words {
		go func(s string) {
			ch <- countFrequency(s)
			wg.Done()
		}(word)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var res = make([]int, 26)
	for frequencyMap := range ch {
		for letter, freq := range frequencyMap {
			res[letter] += freq
		}
	}

	fmt.Println(res)
}
