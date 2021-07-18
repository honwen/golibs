package cip

import (
	"regexp"
	"time"
)

const minTimeout = 2000 * time.Millisecond
const RegxIPv4 = `(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)`
const RegxIPv6 = `([0-9A-Fa-f]{0,4}:){2,7}([0-9A-Fa-f]{1,4}$|((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`

var ApiIPv4 = []string{
	"http://www.net.cn/static/customercare/yourip.asp", "http://ddns.oray.com/checkip", "http://speedtest.ecnu.edu.cn/getIP.php",
	"http://members.3322.org/dyndns/getip", "http://ns1.dnspod.net:6666", "http://ip.bczs.net",
	"http://pv.sohu.com/cityjson?ie=utf-8", "http://whois.pconline.com.cn/ipJson.jsp",
	"http://ipba.cc", "http://v4.myip.la", "https://api.ipify.org", "http://ip-api.com", "http://whatismyip.akamai.com",
}

var ApiIPv6 = []string{
	"http://speed.neu6.edu.cn/getIP.php", "http://v6.myip.la", "https://api64.ipify.org", "http://speedtest6.ecnu.edu.cn/getIP.php",
	"http://ip6only.me/api/", "http://v6.ipv6-test.com/api/myip.php", "https://v6.ident.me",
}

var curlVer = []string{
	"7.76.1", "7.76.0", "7.75.0", "7.74.0", "7.73.0", "7.72.0", "7.71.1", "7.71.0", "7.70.0",
	"7.69.1", "7.69.0", "7.68.0", "7.67.0", "7.66.0", "7.65.3", "7.65.2", "7.65.1", "7.65.0",
	"7.64.1", "7.64.0", "7.63.0", "7.62.0", "7.61.1", "7.61.0", "7.60.0", "7.59.0", "7.58.0",
}

func MyIPv4() (ip string) {
	regx := regexp.MustCompile(RegxIPv4)
	return FastWGetWithVailder(ApiIPv4, func(s string) string {
		return regx.FindString((s))
	})
}

func MyIPv6() (ip string) {
	regx := regexp.MustCompile(RegxIPv6)
	return FastWGetWithVailder(ApiIPv6, func(s string) string {
		return regx.FindString((s))
	})
}

func FastWGetWithVailder(ipAPI []string, vailder func(string) string) (ip string) {
	var (
		length   = len(ipAPI)
		ipMap    = make(map[string]int, length/5)
		cchan    = make(chan string, length/2)
		maxCount = -1
	)
	for _, url := range ipAPI {
		go func(url string) {
			cchan <- vailder(wGet(url, minTimeout))
		}(url)
	}
	for i := 0; i < length; i++ {
		v := <-cchan
		if 0 == len(v) {
			continue
		}
		if ipMap[v]++; ipMap[v] >= length/2 {
			return v
		}
	}
	for k, v := range ipMap {
		if v > maxCount {
			maxCount = v
			ip = k
		}
	}

	// Use First ipAPI as failsafe
	if 0 == len(ip) {
		ip = vailder(wGet(ipAPI[0], 5*minTimeout))
	}
	return
}
