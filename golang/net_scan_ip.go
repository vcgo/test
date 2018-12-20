package main

import (
	"fmt"
	"log"
	"math"
	"net"
	"strings"
	"sync"
	"time"
)

// 单网卡模式
func main() {
	// 获取网卡信息
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal("无法获取本地网络信息:", err)
	}
	// 获取ip
	var myIP *net.IPNet
	for _, a := range addrs {
		if getIP, ok := a.(*net.IPNet); ok && !getIP.IP.IsLoopback() {
			if getIP.IP.To4() != nil {
				myIP = getIP
				break
			}
		}
	}
	// WaitGroup
	wg := sync.WaitGroup{}

	// 扫描列表
	ch := make(chan net.IP)
	for _, ip := range Table(myIP) {
		go func(ip net.IP) {
			wg.Add(1)
			defer wg.Done()
			if scanPort(ip.String(), 80, 500*time.Millisecond) {
				ch <- ip
			}
		}(ip)
	}

	ch2 := make(chan bool)
	go func() {
		wg.Wait()
		ch2 <- true
	}()

	for {
		brk := false
		select {
		case newIP := <-ch:
			fmt.Println("ok", newIP)
		case <-ch2:
			brk = true
			break
		}
		if brk {
			break
		}
	}

	// for {
	// 	fmt.Println("Do other thing.")
	// 	time.Sleep(time.Duration(9999) * time.Millisecond)
	// }
}

// Table 根据IP和mask换算内网IP范围
func Table(ipNet *net.IPNet) []net.IP {
	ip := ipNet.IP.To4()
	var min, max uint32
	var data []net.IP
	for i := 0; i < 4; i++ {
		b := uint32(ip[i] & ipNet.Mask[i])
		min += b << ((3 - uint(i)) * 8)
	}
	one, _ := ipNet.Mask.Size()
	max = min | uint32(math.Pow(2, float64(32-one))-1)
	for i := min; i < max; i++ {
		if i&0x000000ff == 0 {
			continue
		}
		data = append(data, uint32ToIP(i))
	}
	return data
}

func uint32ToIP(intIP uint32) net.IP {
	var bytes [4]byte
	bytes[0] = byte(intIP & 0xFF)
	bytes[1] = byte((intIP >> 8) & 0xFF)
	bytes[2] = byte((intIP >> 16) & 0xFF)
	bytes[3] = byte((intIP >> 24) & 0xFF)
	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

func scanPort(ip string, port int, timeout time.Duration) bool {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			return scanPort(ip, port, timeout)
		}
		return false
	}
	conn.Close()
	return true
}
