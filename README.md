#### Tcp Fast Port Scanner

------------------------------------

A Tcp Fast Port Scanner

------------------------------------

 ____          _               _  __           _ _
| __ ) _   _  | |   _   _  ___(_)/ _| ___ _ __/ / |
|  _ \| | | | | |  | | | |/ __| | |_ / _ \ '__| | |
| |_) | |_| | | |__| |_| | (__| |  _|  __/ |  | | |
|____/ \__, | |_____\__,_|\___|_|_|  \___|_|  |_|_|
       |___/
------------------------------------
 - By Lucifer11
 - My QQ:1185151867
 - [My Blog](https://fdlucifer.github.io/)

## example

``` bash
D:\Go\testgofiles\oldboygo\day12\goportscan>goportscan.exe
####Authored by lucifer11####
[+]use max count 12 cpu kernels to scan host ports...
[goportscan.exe]
[+]args len is 1
usage example：goportscan.exe -host baidu.com -start-port 1 -end-port 65535
```

## help

``` bash
D:\Go\testgofiles\oldboygo\day12\goportscan>goportscan.exe -h
####Authored by lucifer11####
[+]use max count 12 cpu kernels to scan host ports...
[goportscan.exe -h]
[+]args len is 2
Usage of goportscan.exe:
  -end-port int
        the port from whitch the scanning ends (default 10000)
  -host string
        hostname or ip to test
  -start-port int
        the port on whitch the scanning starts (default 80)
  -timeout duration
        timeout (default 200ms)
```

## useage

``` bash
D:\Go\testgofiles\oldboygo\day12\goportscan>goportscan.exe -host 121.42.79.164 -start-port 1 -end-port 65535
......
[+]scanning host: 121.42.79.164, port: 64572
[+]scanning host: 121.42.79.164, port: 64584
[+]scanning host: 121.42.79.164, port: 65183
[+]scanning host: 121.42.79.164, port: 64840
[+]scanning host: 121.42.79.164, port: 64553
[+]scanning host: 121.42.79.164, port: 64755
[+]scanning host: 121.42.79.164, port: 7085
[+]scanning host: 121.42.79.164, port: 7666
[+]scanning host: 121.42.79.164, port: 8127
[+]scanning host: 121.42.79.164, port: 9716
[+]scanning host: 121.42.79.164, port: 22151
[+]scanning host: 121.42.79.164, port: 30896
[+]scanning host: 121.42.79.164, port: 49427
[+]scanning host: 121.42.79.164, port: 59080
[+]scanning host: 121.42.79.164, port: 30897
------------------------------------

A Tcp Fast Port Scanner

------------------------------------

 ____          _               _  __           _ _
| __ ) _   _  | |   _   _  ___(_)/ _| ___ _ __/ / |
|  _ \| | | | | |  | | | |/ __| | |_ / _ \ '__| | |
| |_) | |_| | | |__| |_| | (__| |  _|  __/ |  | | |
|____/ \__, | |_____\__,_|\___|_|_|  \___|_|  |_|_|
       |___/

My QQ:1185151867

Blog: https://fdlucifer.github.io/

####################################

[+]target's opening ports: [21 80 443]
[+]this scan used time: 11 seconds
-------------scan over--------------
```

### 优点特色

 - 速度极快，扫描65535个端口只需要5秒左右
 - ip和域名都能扫

### 注意
 - 如果工具打开报错请把flag.txt和go build生成的goportscan.exe放在同一目录下即可