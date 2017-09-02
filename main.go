package main

import (
	"time"
	"log"
)

func worker() {
	ol := time.NewTicker(10 * time.Second)
	hb := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ol.C:
			log.Printf("Tick every 10s")
		case <-hb.C:
			log.Printf("Tick every 1s")
		}
	}
}

func main() {
	worker()
}
