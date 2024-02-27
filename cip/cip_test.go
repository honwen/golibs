package cip

import (
	"regexp"
	"testing"

	"github.com/ysmood/got"
)

func TestGetIPv4(t *testing.T) {
	ip := MyIPv4()
	// log.Println(ip)
	got.T(t).True(regexp.MustCompile(RegxIPv4).MatchString(ip) || len(ip) == 0)
}

func TestGetIPv6(t *testing.T) {
	ip := MyIPv6()
	// log.Println(ip)
	got.T(t).True(regexp.MustCompile(RegxIPv6).MatchString(ip) || len(ip) == 0)
}

func TestResloveIPv4(t *testing.T) {
	got.T(t).Has([]string{"8.8.8.8", "8.8.4.4"}, ResloveIPv4("dns.google"))
	got.T(t).Has([]string{"223.6.6.6", "223.5.5.5"}, ResloveIPv4("dns.alidns.com"))

	got.T(t).True(regexp.MustCompile(RegxIPv4).MatchString(ResloveIPv4("www.qq.com")))
}

func TestResloveIPv6(t *testing.T) {
	got.T(t).
		Has([]string{"2001:4860:4860::8844", "2001:4860:4860::8888"}, ResloveIPv6("dns.google"))
	got.T(t).Has([]string{"2400:3200::1", "2400:3200:baba::1"}, ResloveIPv6("dns.alidns.com"))

	got.T(t).True(regexp.MustCompile(RegxIPv6).MatchString(ResloveIPv6("www.qq.com")))
}

func TestIPsOfASN(t *testing.T) {
	// skip tests
	// got.T(t).Gt(len(IPsOfASN(ASN_TENCENT_CN)), 100)
	// got.T(t).Gt(len(IPsOfASN(ASN_ALIBABA_CN)), 100)
	// got.T(t).Gt(len(IPsOfASN(ASN_HWCLOUDS)), 100)
	// got.T(t).Gt(len(IPsOfASN(ASN_HINET)), 100)
	// got.T(t).Gt(len(IPsOfASN(ASN_NEWTT)), 100)
	// got.T(t).Gt(len(IPsOfASN(ASN_HKBN)), 100)
	// got.T(t).Gt(len(IPsOfASN(ASN_UHGL)), 100)
}
