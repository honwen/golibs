package cip

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const api = `http://ip.bczs.net/AS`

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

func IPsOfASN(id int) (ips []string) {
	body := wGet(api+fmt.Sprint(id), 30*time.Second)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(body)))
	if err != nil {
		log.Fatal(err)
	}

	regx := regexp.MustCompile(RegxIPv4)
	doc.Find("code").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if regx.MatchString(strings.Split(text, `/`)[0]) {
			ips = append(ips, text)
		}
	})
	return
}
