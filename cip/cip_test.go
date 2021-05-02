package cip

import (
	"log"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIPv4(t *testing.T) {
	ip := MyIPv4()
	log.Println(ip)
	assert.True(t, regexp.MustCompile(regxIPv4).MatchString(ip) || len(ip) == 0)
}

func TestGetIPv6(t *testing.T) {
	ip := MyIPv6()
	log.Println(ip)
	assert.True(t, regexp.MustCompile(regxIPv6).MatchString(ip) || len(ip) == 0)
}

func TestGetDNSv4(t *testing.T) {
	assert.Contains(t, []string{"8.8.8.8", "8.8.4.4"}, ResloveIPv4("dns.google"))
	assert.Contains(t, []string{"223.6.6.6", "223.5.5.5"}, ResloveIPv4("dns.alidns.com"))
}

func TestGetDNSv6(t *testing.T) {
	assert.Contains(t, []string{"2001:4860:4860::8844", "2001:4860:4860::8888"}, ResloveIPv6("dns.google"))
	assert.Contains(t, []string{"2400:3200::1", "2400:3200:baba::1"}, ResloveIPv6("dns.alidns.com"))
}
