package place

import (
	"fmt"
	"sync"
	"testing"
)

func TestBroadcastMessage(t *testing.T) {
	t.Log("TestBroadcastMessage")

	messageChannel := make(chan string)
	var wg sync.WaitGroup

	// Create receivers
	numReceivers := 5
	for i := 0; i < numReceivers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			message := <-messageChannel
			fmt.Printf("Receiver %d received: %s\n", id, message)
		}(i)
	}

	// Send a message to all receivers
	message := "Hello, broadcast!"
	messageChannel <- message

	wg.Wait()
	close(messageChannel)
}
