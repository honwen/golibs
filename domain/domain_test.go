package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitDomain001(t *testing.T) {
	rr, domain := SplitDomainToRR("a.example.com")

	assert.Equal(t, rr, "a")
	assert.Equal(t, domain, "example.com")
}

func TestSplitDomain002(t *testing.T) {
	rr, domain := SplitDomainToRR("example.com")

	assert.Equal(t, rr, "@")
	assert.Equal(t, domain, "example.com")
}

func TestSplitDomain003(t *testing.T) {
	rr, domain := SplitDomainToRR("*.example.com")

	assert.Equal(t, rr, "*")
	assert.Equal(t, domain, "example.com")
}

func TestSplitDomain004(t *testing.T) {
	rr, domain := SplitDomainToRR("*.a.example.com")

	assert.Equal(t, rr, "*.a")
	assert.Equal(t, domain, "example.com")
}

func TestSplitDomain005(t *testing.T) {
	rr, domain := SplitDomainToRR("*.b.a.example.com")

	assert.Equal(t, rr, "*.b.a")
	assert.Equal(t, domain, "example.com")
}
func TestSplitDomain006(t *testing.T) {
	rr, domain := SplitDomainToRR("a.example.co.kr")

	assert.Equal(t, rr, "a")
	assert.Equal(t, domain, "example.co.kr")
}

func TestSplitDomain007(t *testing.T) {
	rr, domain := SplitDomainToRR("*.a.example.co.kr")

	assert.Equal(t, rr, "*.a")
	assert.Equal(t, domain, "example.co.kr")
}

func TestSplitDomain008(t *testing.T) {
	rr, domain := SplitDomainToRR("example.co.kr")

	assert.Equal(t, rr, "@")
	assert.Equal(t, domain, "example.co.kr")
}

func TestIsValidHostname(t *testing.T) {
	validHosts := []string{
		"0example.com", "example.com", "ex.example.com", "ex-1ample.com.ru",
		"xn---asdasd.com", "local", "aa.ru", "a.ru", "00.11.22.33",
	}
	invalidHosts := []string{
		"-a.c", "a-.c", "a.-c", "a.c-",
		"host-", "h@st", "*.com", "ex_ample.com", "!asd.ru", "google..com",
		".google.com", "google.com.",
	}

	for _, h := range validHosts {
		assert.True(t, IsValidHostname(h), h)
	}

	for _, h := range invalidHosts {
		assert.False(t, IsValidHostname(h), h)
	}
}
