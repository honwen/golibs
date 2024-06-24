package cip

import (
	"regexp"
	"time"
)

const minTimeout = 2000 * time.Millisecond

const RegxIPv4 = `(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)`

const RegxIPv6 = `([0-9A-Fa-f]{0,4}:){2,7}([0-9A-Fa-f]{1,4}$|((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`

var ApiIPv4 = []string{
	"http://www.net.cn/static/customercare/yourip.asp", "http://ddns.oray.com/checkip",
	"https://test.ustc.edu.cn/backend/getIP.php", "https://wsus.sjtu.edu.cn/speedtest/backend/getIP.php",
	"http://ipv4.ddnspod.com", "http://speedtest.zju.edu.cn/getIP.php",
	"https://www.lib.whu.edu.cn/speedtest/backend/getIP.php",
	// "https://v6r.ipip.net",
	"http://members.3322.org/dyndns/getip", "https://ip.3322.net",
	"http://whois.pconline.com.cn/ipJson.jsp", "http://ifconfig.cc", "http://cip.cc",
	"http://v4.myip.la", "https://api.ipify.org", "http://ip-api.com/json", "http://whatismyip.akamai.com",
	"https://www.cloudflare-cn.com/cdn-cgi/trace", "https://www.cloudflare.com/cdn-cgi/trace",

	// refer: https://github.com/honwen/aliyun-ddns-cli/issues/85
	// "http://pv.sohu.com/cityjson?ie=utf-8",
}

var ApiIPv6 = []string{
	"http://ipv6.testipv6.cn/ip/", "http://ipv6.ddnspod.com",
	"http://v6.myip.la", "https://api64.ipify.org",
	"https://test6.ustc.edu.cn/backend/getIP.php",
	"https://www.lib.whu.edu.cn/speedtest/backend/getIP.php",
	"http://ip6only.me/api/", "http://v6.ipv6-test.com/api/myip.php", "https://v6.ident.me",
	"https://www.cloudflare-cn.com/cdn-cgi/trace", "https://www.cloudflare.com/cdn-cgi/trace",
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
		if len(v) == 0 {
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
	if len(ip) == 0 {
		ip = vailder(wGet(ipAPI[0], 5*minTimeout))
	}
	return
}
