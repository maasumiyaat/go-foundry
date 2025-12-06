package waitgroup

import (
	"sync"
)

func WaitPrintln(str string) string {
	var wg sync.WaitGroup
	result := make(chan string, 1)
	wg.Add(1)

	go func() {
		defer wg.Done()
		result <- str
	}()

	wg.Wait()
	return <-result
}

func ChannelPrintln(str string) string {
	result := make(chan string, 1)

	go func() {
		result <- str
	}()

	return <-result
}
