package saver

import (
	"log"
)

func ItemSaver() chan any {
	out := make(chan any)
	go func() {
		itemCount := 0
		for {
			item := <-out
			save(itemCount, item)
			itemCount++
		}
	}()
	return out
}

func save(no int, item any) {
	log.Printf("Item Saver: got item #%d: %v", no, item)
}
