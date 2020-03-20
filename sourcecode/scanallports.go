package main

import (
	"sync"
	"time"

	"github.com/fatih/color"
)

func scanAllPorts(line string, ports []int, timeoutv time.Duration) {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	for port := 1; port <= 5000; port++ {
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
	for port := 5001; port <= 10000; port++ {
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
	for port := 10001; port <= 15000; port++ {
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
	for port := 15001; port <= 20000; port++ {
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
	for port := 20001; port <= 25000; port++ {
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
	for port := 25001; port <= 30000; port++ {
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
	for port := 30001; port <= 35000; port++ {
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
	for port := 35001; port <= 40000; port++ {
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
	for port := 40001; port <= 45000; port++ {
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
	for port := 45001; port <= 50000; port++ {
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
	for port := 50001; port <= 55000; port++ {
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
	for port := 55001; port <= 60000; port++ {
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
	for port := 60001; port <= 65535; port++ {
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
	color.HiYellow("[+] host [%s] opening ports: %v\n", line, ports)
	color.HiYellow("------------------------------------------------\n")
	writeResults(line, ports)
}
