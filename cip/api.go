package cip

import (
	"regexp"
	"time"
)

const minTimeout = 2000 * time.Millisecond

const RegxIPv4 = `(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)`

const RegxIPv6 = `([0-9A-Fa-f]{0,4}:){2,7}([0-9A-Fa-f]{1,4}$|((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`

var ApiIPv4 = []string{
	// cn
	"https://itango.tencent.com/out/itango/myip",       // verified 2026-07-27
	"https://ip9.com.cn/get",                           // verified 2026-07-27
	"http://www.net.cn/static/customercare/yourip.asp", // verified 2026-07-27
	"http://ddns.oray.com/checkip",                     // verified 2026-07-27
	"https://ifcfg.cn",                                 // verified 2026-07-27
	"http://ipv4.ddnspod.com",                          // verified 2026-07-27
	"https://ipv4.ddnsip.cn/",                          // verified 2026-07-27
	"http://members.3322.org/dyndns/getip",             // verified 2026-07-27
	"https://ip.3322.net",                              // verified 2026-07-27

	// global
	"http://ifconfig.cc",                          // verified 2026-07-27
	"http://cip.cc",                               // verified 2026-07-27
	"http://v4.myip.la",                           // verified 2026-07-27
	"http://ip-api.com/json",                      // verified 2026-07-27
	"http://whatismyip.akamai.com",                // verified 2026-07-27
	"https://www.cloudflare-cn.com/cdn-cgi/trace", // verified 2026-07-27
	"https://www.cloudflare.com/cdn-cgi/trace",    // verified 2026-07-27

	// refer: https://github.com/honwen/aliyun-ddns-cli/issues/85
	// "http://pv.sohu.com/cityjson?ie=utf-8",
	// "https://v6r.ipip.net",
}

var ApiIPv6 = []string{
	// cn
	"http://ipv6.ddnspod.com", // verified 2026-07-27
	"https://ipv6.ddnsip.cn/", // verified 2026-07-27

	// global
	"http://v6.myip.la",                           // verified 2026-07-27
	"https://api64.ipify.org",                     // verified 2026-07-27
	"http://ip6only.me/api/",                      // verified 2026-07-27
	"https://v6.ident.me",                         // verified 2026-07-27
	"https://www.cloudflare-cn.com/cdn-cgi/trace", // verified 2026-07-27
	"https://www.cloudflare.com/cdn-cgi/trace",    // verified 2026-07-27
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
