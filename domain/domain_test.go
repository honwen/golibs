package domain

import (
	"testing"

	"github.com/ysmood/got"
)

func TestSplitDomain001(t *testing.T) {
	rr, domain := SplitDomainToRR("a.example.com")

	got.T(t).Eq(rr, "a")
	got.T(t).Eq(domain, "example.com")
}

func TestSplitDomain002(t *testing.T) {
	rr, domain := SplitDomainToRR("example.com")

	got.T(t).Eq(rr, "@")
	got.T(t).Eq(domain, "example.com")
}

func TestSplitDomain003(t *testing.T) {
	rr, domain := SplitDomainToRR("*.example.com")

	got.T(t).Eq(rr, "*")
	got.T(t).Eq(domain, "example.com")
}

func TestSplitDomain004(t *testing.T) {
	rr, domain := SplitDomainToRR("*.a.example.com")

	got.T(t).Eq(rr, "*.a")
	got.T(t).Eq(domain, "example.com")
}

func TestSplitDomain005(t *testing.T) {
	rr, domain := SplitDomainToRR("*.b.a.example.com")

	got.T(t).Eq(rr, "*.b.a")
	got.T(t).Eq(domain, "example.com")
}
func TestSplitDomain006(t *testing.T) {
	rr, domain := SplitDomainToRR("a.example.co.kr")

	got.T(t).Eq(rr, "a")
	got.T(t).Eq(domain, "example.co.kr")
}

func TestSplitDomain007(t *testing.T) {
	rr, domain := SplitDomainToRR("*.a.example.co.kr")

	got.T(t).Eq(rr, "*.a")
	got.T(t).Eq(domain, "example.co.kr")
}

func TestSplitDomain008(t *testing.T) {
	rr, domain := SplitDomainToRR("example.co.kr")

	got.T(t).Eq(rr, "@")
	got.T(t).Eq(domain, "example.co.kr")
}

func TestIsValidHostname(t *testing.T) {
	validHosts := []string{
		"0example.com", "example.com", "ex.example.com", "ex-1ample.com.ru", "local", "xn--oiq.cc", "yandex.ru",
	}
	invalidHosts := []string{
		"-a.c", "a-.c", "a.-c", "a.c-",
		"host-", "h@st", "*.com", "ex_ample.com", "!asd.ru", "google..com",
		".google.com", "google.com.", "yandex.*",
	}

	for _, h := range validHosts {
		got.T(t).True(IsValidHostname(h))
	}

	for _, h := range invalidHosts {
		got.T(t).False(IsValidHostname(h))
	}
}

func TestSorted(t *testing.T) {
	list := []string{"www.qq.com", "qq.com", "google.com", "google.com.hk", "www.google.com", "web.dev", "nasa.gov", "im.tv", "b23.tv", "qq.com"}
	golden := []string{"google.com.hk", "google.com", "qq.com", "web.dev", "nasa.gov", "b23.tv", "im.tv"}

	got.T(t).Eq(Sort(list), golden)
}
