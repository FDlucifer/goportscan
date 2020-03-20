package main

import (
	"sync"
	"time"

	"github.com/fatih/color"
)

func selfDefineScan(line string, ports []int, startPortv int, endPortv int, timeoutv time.Duration) {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	for port := startPortv; port <= endPortv; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(line, p, timeoutv)
			if opened {
				mutex.Lock()
				ports = append(ports, p)
				mutex.Unlock()
			}
			wg.Done()
		}(port)
	}
	wg.Wait()
	color.HiMagenta("[+] host [%s] opening ports: %v\n", line, ports)
	color.HiMagenta("------------------------------------------------\n")
	writeResults(line, ports)
}
