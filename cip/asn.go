package cip

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var apis = []string{
	`http://ip.bczs.net/AS`,
	`https://api.hackertarget.com/aslookup/?q=AS`,
}

const (
	ASN_TENCENT_CN = 132203
	ASN_ALIBABA_CN = 45102
	ASN_HWCLOUDS   = 136907
	ASN_HINET      = 3462
	ASN_NEWTT      = 9381
	ASN_HKBN       = 9269
	ASN_UHGL       = 135377
	ASN_MICROSOFT  = 8075
)

func scanCIDR(content string) (ips []string) {
	regx := regexp.MustCompile(RegxIPv4)

	if strings.Contains(content, "<html") && strings.Contains(content, "<code") {
		doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(content)))
		if err != nil {
			log.Fatal(err)
		}

		doc.Find("code").Each(func(i int, s *goquery.Selection) {
			text := strings.TrimSpace(s.Text())
			if regx.MatchString(strings.Split(text, `/`)[0]) {
				ips = append(ips, text)
			}
		})
	} else {
		sc := bufio.NewScanner(strings.NewReader(content))
		for sc.Scan() {
			text := sc.Text()
			if regx.MatchString(strings.Split(text, `/`)[0]) {
				ips = append(ips, text)
			}
		}
	}
	return ips
}

func IPsOfASN(id int) (ips []string) {
	var (
		length = len(apis)
		cchan  = make(chan []string, length)
	)
	for _, api := range apis {
		go func(api string) {
			body := wGet(api+fmt.Sprint(id), 10*time.Second)
			cchan <- scanCIDR(body)
		}(api)
	}

	ipMap := map[string]bool{}
	for i := 0; i < length; i++ {
		for _, v := range <-cchan {
			ipMap[v] = true
		}
	}

	ips = make([]string, len(ipMap))
	i := 0
	for k := range ipMap {
		ips[i] = k
		i++
	}

	return
}
