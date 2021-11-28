package main

import (
	"fmt"
	"gopinger/lib/pinger"
	"time"
)

func main() {
	//s := pinger.NewSingle("192.168.8.1", 10, 1)
	startTime := time.Now()

	results := pinger.ScanRange("192.168.8.1", "192.168.8.254")

	for _, h := range results {
		if h.IsActive {
			fmt.Printf("IP:%s \t MinRtt:%d \t MaxRtt: %d \t AvgRtt:%d \n", h.IP, h.Stats.MinRtt.Milliseconds(), h.Stats.MaxRtt.Milliseconds(), h.Stats.AvgRtt.Milliseconds())
		}
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Printf("Duration: %f\n", duration.Seconds())

	fmt.Printf("Finished pinging %d hosts", len(results))

}
