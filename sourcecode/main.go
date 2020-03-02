package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"runtime"
	"sync"
	"time"
)

func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Microsecond * 1)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", host, port))
	fmt.Printf("[+]%s ip and port is %s\n", host, tcpAddr.String())
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	fmt.Printf("[+]scanning host: %s, port: %d\n", host, port)
	if err == nil {
		_ = conn.Close()
		return true
	}
	return false
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	num := runtime.NumCPU()
	fmt.Println("####Authored by lucifer11####")
	fmt.Printf("[+]use max count %d cpu kernels to scan host ports...\n", num)
	time.Sleep(time.Second * 3)
	fmt.Println(os.Args)
	fmt.Printf("[+]args len is %d\n", len(os.Args))
	if len(os.Args) == 1 {
		fmt.Println("usage example：goportscan.exe -host baidu.com -start-port 1 -end-port 65535")
		os.Exit(0)
	}
	time.Sleep(time.Second * 2)
	s_time := time.Now().Unix()
	hostname := flag.String("host", "", "hostname or ip to test")
	startPort := flag.Int("start-port", 80, "the port on whitch the scanning starts")
	endPort := flag.Int("end-port", 10000, "the port from whitch the scanning ends")
	timeout := flag.Duration("timeout", time.Millisecond*200, "timeout")
	flag.Parse()
	ports := []int{}
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	for port := *startPort; port <= *endPort; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(*hostname, p, *timeout)
			if opened {
				mutex.Lock()
				ports = append(ports, p)
				mutex.Unlock()
			}
			wg.Done()
		}(port)
	}
	wg.Wait()
	fmt.Println("------------------------------------")
	fmt.Println("A Tcp Fast Port Scanner")
	fmt.Println("------------------------------------")
	fmt.Printf("\n")
	fmt.Println("By Lucifer11")
	fmt.Println("My QQ:1185151867")
	fmt.Println("Blog: https://fdlucifer.github.io/")
	fmt.Println("####################################")
	fmt.Printf("\n")
	fmt.Printf("[+]target's opening ports: %v\n", ports)
	e_time := time.Now().Unix()
	fmt.Printf("[+]this scan used time: %d seconds\n", e_time-s_time)
}
