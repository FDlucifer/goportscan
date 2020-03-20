package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
)

func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Microsecond * 1)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	fmt.Printf("[+] scanning host: [%s], port: [%d]\n", host, port)
	if err == nil {
		_ = conn.Close()
		return true
	}
	return false
}

func getIP(host string) {
	addr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		color.Red("[-] Resolvtion error...\n", err.Error())
	}
	color.Cyan("[+] [%s] Resolved address is [%s]\n", host, addr.String())
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	num := runtime.NumCPU()
	color.Yellow("[+] ####Authored by lucifer11####\n")
	color.Blue("[+] use max count %d cpu kernels to scan host ports...\n", num)
	fmt.Println(os.Args)
	color.Green("[+] args len is %d\n", len(os.Args))
	if len(os.Args) == 10 {
		color.Cyan("[+] use -h to get help options")
		os.Exit(0)
	}
	flagtext, err := ioutil.ReadFile("flag.txt")
	if err != nil {
		panic(err.Error())
	}
	color.Red("------------------------------------------------\n")
	color.Red("A Tcp Fast Port Scanner\n")
	color.Red("------------------------------------------------\n")
	color.Red("%s\n", flagtext)
	color.Red("My QQ:1185151867\n")
	color.Red("Blog: https://fdlucifer.github.io/\n")
	color.Red("################################################\n")
	time.Sleep(time.Second * 1)
	iplist := flag.String("ipfile", "iplist.txt", "the file stored ip list")
	startPort := flag.Int("start-port", 1, "the port on whitch the scanning starts")
	endPort := flag.Int("end-port", 10000, "the port from whitch the scanning ends")
	timeout := flag.Duration("timeout", time.Millisecond*500, "timeout")
	allPorts := flag.String("allports", "n", "select this option to scan all open ports in 65535 ports")
	flag.Parse()
	ports := []int{}
	fileName := *iplist
	file, err := os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("[-] Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var size = stat.Size()
	color.Green("[+] start read iplists [%s] and scan, iplist file size: [%d]\n", *iplist, size)
	time.Sleep(time.Second * 1)

	buf := bufio.NewReader(file)
	for {
		//ch := make(chan string, 1)
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		getIP(line)
		time.Sleep(time.Second * 1)
		/*ch <- line
		inputip, ok := <-ch
		if ok == false {
			fmt.Println("chan is closed")
			break
		}
		close(ch)*/
		startPortv := *startPort
		endPortv := *endPort
		timeoutv := *timeout
		allPortsv := *allPorts

		if allPortsv == "y" {
			scanAllPorts(line, ports, timeoutv)
		} else {
			selfDefineScan(line, ports, startPortv, endPortv, timeoutv)
		}
		//time.Sleep(time.Second * 1)
		ports = ports[:0]
		if err != nil {
			if err == io.EOF {
				fmt.Println("[+] iplist scan finished!")
				break
			} else {
				fmt.Println("[-] Read file error!", err)
				return
			}
		}
	}
}
