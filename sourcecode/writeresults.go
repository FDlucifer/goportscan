package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeResults(line string, ports []int) error {
	f, err := os.OpenFile("resultip.txt", os.O_WRONLY, 0777)
	if err != nil {
		c, err := os.Create("resultip.txt")
		if err != nil {
			fmt.Printf("[-] create result file error: %v\n", err)
		}
		return err
		c.Close()
	}
	w := bufio.NewWriter(f)
	lineStr := fmt.Sprintf("%s:%v\n", line, ports)
	if len(ports) > 0 { //扫描到的ports端口为非空时，写入ip和端口
		n, _ := f.Seek(0, 2)                   //查找文件末尾偏移量
		_, err = f.WriteAt([]byte(lineStr), n) //从文件末尾偏移量写入内容
	}
	defer f.Close()
	return w.Flush()
}
