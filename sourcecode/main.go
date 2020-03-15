package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/fatih/color"
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
	color.Yellow("####Authored by lucifer11####")
	fmt.Printf("\n")
	color.Blue("[+]use max count %d cpu kernels to scan host ports...\n", num)
	time.Sleep(time.Second * 3)
	fmt.Println(os.Args)
	color.Green("[+]args len is %d\n", len(os.Args))
	if len(os.Args) == 1 {
		color.Cyan("usage example：goportscan.exe -host baidu.com -start-port 1 -end-port 65535")
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
	flagtext, err := ioutil.ReadFile("flag.txt")
	if err != nil {
		panic(err.Error())
	}
	color.Red("------------------------------------")
	fmt.Printf("\n")
	color.Red("A Tcp Fast Port Scanner")
	fmt.Printf("\n")
	color.Red("------------------------------------")
	fmt.Printf("\n")
	color.Red("%s", flagtext)
	fmt.Printf("\n")
	color.Red("My QQ:1185151867")
	fmt.Printf("\n")
	color.Red("Blog: https://fdlucifer.github.io/")
	fmt.Printf("\n")
	color.Red("####################################")
	fmt.Printf("\n")
	color.Blue("[+]target's opening ports: %v\n", ports)
	e_time := time.Now().Unix()
	color.Blue("[+]this scan used time: %d seconds\n", e_time-s_time)
	color.Red("-------------scan over--------------")
}
