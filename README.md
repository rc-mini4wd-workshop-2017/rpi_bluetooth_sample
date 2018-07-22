# rpi_bluetooth_sample

## prepare

```
$ go get github.com/tarm/serial
```

## how to use

console 1

```
$ hcitool scan
Scanning ...
        xx:xx:xx:xx:xx:xx       Quattro Ace
$ sudo rfcomm connect hci0 xx:xx:xx:xx:xx:xx 2
Connected /dev/rfcomm0 to xx:xx:xx:xx:xx:xx on channel 2
Press CTRL-C for hangup
```

console 2

```
$ go run main.go
2018/07/22 13:08:40 > info
2018/07/22 13:08:40 info
version=0.0.1

result: 0
```