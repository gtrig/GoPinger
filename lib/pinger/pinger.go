package pinger

import (
	"encoding/binary"
	"net"
	"sync"

	"github.com/go-ping/ping"
)

// A struct to hold the Host parameters.
type Host struct {
	ip       string
	count    int
	timeout  int
	stats    *ping.Statistics
	isActive bool
}

func NewSingle(ip string, count int, timeout int) Host {
	p := Host{ip: ip, count: count, timeout: timeout}
	return p
}

// Get a start ip and an end ip and return a slice of ping structs
func NewRange(from string, to string, count int, timeout int) []Host {
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
		h := Host{ip: ip.String(), count: count, timeout: timeout}
		// Add the ping struct to the slice
		hosts = append(hosts, h)
	}
	return hosts
}

// Ping the IP
func (h *Host) InitiatePing(wg *sync.WaitGroup) {
	pingInstance, err := ping.NewPinger(h.ip)
	pingInstance.SetPrivileged(true)

	if err != nil {
		panic(err)
	}

	pingInstance.Count = h.count
	//fmt.Printf("Pinging %s \n", h.ip)

	pingInstance.OnFinish = func(stats *ping.Statistics) {
		h.stats = pingInstance.Statistics()

		if h.stats.PacketLoss == 100 {
			h.isActive = false
		} else {
			h.isActive = true
		}
	}

	pingInstance.Run() // Blocks until finished.
	wg.Done()
}

func ScanRange(from string, to string) []Host {
	hosts := NewRange(from, to, 20, 1)

	jobs := make(chan Host, len(hosts))
	results := make(chan Host, len(hosts))

	var waitgroup sync.WaitGroup

	for _, h := range hosts {
		go h.InitiatePing()
	}

	waitgroup.Wait()

	return hosts
}

// func worker(jobs <-chan Host, results chan<- Host) {
// 	for j := range jobs {
// 		j.InitiatePing(&waitgroup)
// 		results <- j
// 	}
// }
