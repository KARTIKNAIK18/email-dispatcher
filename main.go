package main

import (
	"github.com/KARTIKNAIK18/email-dispatcher/Pipeline"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	queue := make(chan pipeline.Emails)

	wg.Add(1)
	go pipeline.LoadRecipient("./contact.csv", queue, &wg)

	WorkerCount := 5
	for i := 0; i <= WorkerCount; i++ {
		wg.Add(1)

		go pipeline.EmailWorker(i, queue, &wg)
	}
	wg.Wait()
}
