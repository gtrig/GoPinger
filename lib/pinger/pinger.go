package pinger

import (
	"encoding/binary"
	"encoding/json"
	"net"
	"sync"
	"time"

	"github.com/go-ping/ping"
)

// A struct to hold the Host parameters.
type Host struct {
	IP       string           `json:"ip"`
	Count    int              `json:"count"`
	Timeout  int              `json:"timeout"`
	Stats    *ping.Statistics `json:"stats"`
	IsActive bool             `json:"is_active"`
}

type ScanResult struct {
	ActiveHosts []Host   `json:"active_hosts"`
	Elapsed     Duration `json:"elapsed"`
	TotalHosts  int      `json:"total_hosts"`
	TotalActive int      `json:"total_active"`
}

type Duration struct {
	time.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func getSingle(ip string, count int, timeout int) Host {
	return Host{IP: ip, Count: count, Timeout: timeout}
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
func startPing(h Host) Host {

	pingInstance, err := ping.NewPinger(h.IP)
	pingInstance.SetPrivileged(true)
	pingInstance.Count = h.Count
	pingInstance.Timeout = time.Duration(h.Timeout) * time.Second

	if err != nil {
		panic(err)
	}

	pingInstance.OnFinish = func(stats *ping.Statistics) {
		h.Stats = pingInstance.Statistics()

		if h.Stats.PacketLoss == 100 {
			h.IsActive = false
		} else {
			h.IsActive = true
		}
	}

	pingInstance.Run() // Blocks until finished.
	return h
}

func ScanSingle(ip string) ScanResult {
	startTime := time.Now()
	host := getSingle(ip, 1, 1)
	result := startPing(host)
	endTime := time.Now()
	elapsed := Duration{endTime.Sub(startTime)}

	return ScanResult{[]Host{result}, elapsed, 1, 1}
}

func ScanRange(from string, to string) ScanResult {
	startTime := time.Now()
	hosts := getRange(from, to, 1, 1)

	jobs := make(chan Host)
	results := make(chan Host)

	go worker(jobs, results)

	for _, host := range hosts {
		jobs <- host
	}

	close(jobs)
	activeHosts := make([]Host, 0)
	for a := 0; a < len(hosts); a++ {
		host := <-results
		if host.IsActive {
			activeHosts = append(activeHosts, host)
		}

	}
	endTime := time.Now()
	elapsed := Duration{endTime.Sub(startTime)}

	return ScanResult{activeHosts, elapsed, len(hosts), len(activeHosts)}
}

func worker(jobs <-chan Host, results chan<- Host) {
	var wg sync.WaitGroup

	for host := range jobs {
		wg.Add(1)
		go func(host Host, results chan<- Host) {
			defer wg.Done()
			results <- startPing(host)
		}(host, results)

	}
	wg.Wait()
}
