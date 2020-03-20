#### Tcp Fast Port Scanner

-----------------------------------------------------

``` bash
 ____          _               _  __           _ _
| __ ) _   _  | |   _   _  ___(_)/ _| ___ _ __/ / |
|  _ \| | | | | |  | | | |/ __| | |_ / _ \ '__| | |
| |_) | |_| | | |__| |_| | (__| |  _|  __/ |  | | |
|____/ \__, | |_____\__,_|\___|_|_|  \___|_|  |_|_|
       |___/

```

-----------------------------------------------------
 - By Lucifer11
 - My QQ:1185151867
 - [My Blog](https://fdlucifer.github.io/)


## help

``` bash
D:\Go\testgofiles\oldboygo\day12\goportscan>goportscan.exe -h
Usage of goportscan.exe:
  -allports string
        select this option to scan all open ports in 65535 ports (default "n")
  -end-port int
        the port from whitch the scanning ends (default 10000)
  -ipfile string
        the file stored ip list (default "iplist.txt")
  -start-port int
        the port on whitch the scanning starts (default 1)
  -timeout duration
        timeout (default 500ms)
```

## 批量扫ip65535全端口

``` bash
D:\Go\testgofiles\oldboygo\day12\goportscan>goportscan.exe -allports y
...
[+] scanning host: [163.com], port: [64660]
[+] scanning host: [163.com], port: [64642]
[+] scanning host: [163.com], port: [65094]
[+] scanning host: [163.com], port: [64655]
[+] scanning host: [163.com], port: [65173]
[+] scanning host: [163.com], port: [65311]
[+] scanning host: [163.com], port: [64934]
[+] scanning host: [163.com], port: [60555]
[+] scanning host: [163.com], port: [60259]
[+] scanning host: [163.com], port: [64640]
[+] host [163.com] opening ports: [443 80]
------------------------------------------------
[+] [4399.com] Resolved address is [118.184.184.70]
```

 - ![](/pics/allports.jpg)


## 批量扫ip3389端口

``` bash
D:\Go\testgofiles\oldboygo\day12\goportscan>goportscan.exe -start-port 3389 -end-port 3389
......
[+] [60.31.205.76] Resolved address is [60.31.205.76]
[+] scanning host: [60.31.205.76], port: [3389]
[+] host [60.31.205.76] opening ports: []
------------------------------------------------
[+] [119.23.85.137] Resolved address is [119.23.85.137]
[+] scanning host: [119.23.85.137], port: [3389]
[+] host [119.23.85.137] opening ports: []
------------------------------------------------
[+] [39.104.21.7] Resolved address is [39.104.21.7]
[+] scanning host: [39.104.21.7], port: [3389]
[+] host [39.104.21.7] opening ports: [3389]
------------------------------------------------
[+] [106.12.118.25] Resolved address is [106.12.118.25]
[+] scanning host: [106.12.118.25], port: [3389]
[+] host [106.12.118.25] opening ports: []
------------------------------------------------
[+] iplist scan finished!
```

 - ![](/pics/singleport.jpg)


### 优点特色

 - 批量快速快速扫描ip文件列表，并遍历65535全端口
 - ip和域名都能扫
 - 扫描结果以ip:[port]形式放入resultip.txt文件中
 - 支持批量ip单端口扫描
 - 持续添加新功能

### 注意
 - 如果工具打开报错请把flag.txt和go build生成的goportscan.exe放在同一目录下即可