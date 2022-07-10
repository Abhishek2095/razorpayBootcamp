package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ratingGenerator(studentId <-chan int, rating chan<- int) {
	for range studentId {
		time.Sleep(time.Duration(rand.Intn(2)))
		rating <- rand.Intn(101)
	}
}

func main() {
	const workers = 10
	const studentCount = 200

	channelBufferSize := workers
	var studentId = make(chan int, channelBufferSize)
	var ratings = make(chan int, channelBufferSize)

	for i := 0; i < workers; i++ {
		go ratingGenerator(studentId, ratings)
	}

	go func() {
		defer close(studentId)
		for i := 0; i < studentCount; i++ {
			studentId <- i
		}
	}()

	totalRating := 0
	for i := 0; i < studentCount; i++ {
		totalRating += <-ratings
	}

	fmt.Println("Average Rating : ", totalRating/studentCount)
}
