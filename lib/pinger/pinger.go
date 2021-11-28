package pinger

import (
	"encoding/binary"
	"net"
	"sync"

	"github.com/go-ping/ping"
)

// A struct to hold the Host parameters.
type Host struct {
	IP       string
	Count    int
	Timeout  int
	Stats    *ping.Statistics
	IsActive bool
}

func getSingle(ip string, count int, timeout int) Host {
	p := Host{IP: ip, Count: count, Timeout: timeout}
	return p
}

// Get a start ip and an end ip and return a slice of ping structs
func getRange(from string, to string, count int, timeout int) []Host {
	// Convert the strings to IPs
	start := binary.BigEndian.Uint32(net.ParseIP(from).To4())
	end := binary.BigEndian.Uint32(net.ParseIP(to).To4())

	// Create a slice to hold the ping structs
	hosts := make([]Host, 0)
	// Loop through the IPs
	for i := start; i <= end; i++ {
		// convert back to net.IP
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		// Create a new ping struct
		h := Host{IP: ip.String(), Count: count, Timeout: timeout}
		// Add the ping struct to the slice
		hosts = append(hosts, h)
	}
	return hosts
}

// Ping a host
func startPing(h Host, results chan<- Host, wg *sync.WaitGroup) {
	pingInstance, err := ping.NewPinger(h.IP)
	pingInstance.SetPrivileged(true)

	if err != nil {
		panic(err)
	}

	pingInstance.Count = h.Count
	//fmt.Printf("Pinging %s \n", h.ip)

	pingInstance.OnFinish = func(stats *ping.Statistics) {
		h.Stats = pingInstance.Statistics()

		if h.Stats.PacketLoss == 100 {
			h.IsActive = false
		} else {
			h.IsActive = true
		}

		results <- h
	}

	pingInstance.Run() // Blocks until finished.
	wg.Done()
}

func ScanRange(from string, to string) []Host {
	hosts := getRange(from, to, 50, 1)

	jobs := make(chan Host)
	results := make(chan Host)

	go worker(jobs, results)

	for _, h := range hosts {
		jobs <- h
	}

	close(jobs)

	for a := 0; a < len(hosts); a++ {
		hosts[a] = <-results
	}

	return hosts
}

func worker(jobs <-chan Host, results chan<- Host) {
	var wg sync.WaitGroup
	for host := range jobs {
		wg.Add(1)
		go startPing(host, results, &wg)
	}
	wg.Wait()
}
