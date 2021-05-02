package cip

import (
	"strings"
	"time"

	"github.com/miekg/dns"
	"github.com/mr-karan/doggo/pkg/resolvers"
	"github.com/mr-karan/doggo/pkg/utils"
)

var mResolvers = []resolvers.Resolver{}

func init() {
	var upstreams = []string{
		"tls://223.5.5.5:853", "tls://223.6.6.6:853", "https://223.5.5.5/dns-query", "https://223.6.6.6/dns-query",
	}
	var opts = resolvers.Options{
		Timeout: 2000 * time.Millisecond,
		Logger:  utils.InitLogger(),
	}
	var dotOpts = resolvers.ClassicResolverOpts{
		UseTLS: true,
		UseTCP: true,
	}

	for _, upstream := range upstreams {
		var ns resolvers.Resolver
		switch {
		case strings.HasPrefix(upstream, "https://"):
			ns, _ = resolvers.NewDOHResolver(upstream, opts)
		case strings.HasPrefix(upstream, "tls://"):
			ns, _ = resolvers.NewClassicResolver(upstream[6:], dotOpts, opts)
		default:
			continue
		}
		mResolvers = append(mResolvers, ns)
	}
}

func ResloveIPv4(domain string) (ip string) {
	return reslove(domain, dns.TypeA)
}

func ResloveIPv6(domain string) (ip string) {
	return reslove(domain, dns.TypeAAAA)
}

func reslove(domain string, qtype uint16) (ip string) {
	var (
		dnsMap   = make(map[string]int, len(mResolvers))
		cchan    = make(chan string, len(mResolvers))
		maxCount = -1
	)

	for i := range mResolvers {
		go func(resolver *resolvers.Resolver) {
			resp, err := (*resolver).Lookup(dns.Question{Name: domain, Qtype: qtype})
			if err == nil && len(resp.Answers) > 0 {
				for idx := range resp.Answers {
					if resp.Answers[idx].Type == dns.TypeToString[qtype] {
						cchan <- resp.Answers[idx].Address
					}
				}
			}
			cchan <- "" // SOA
		}(&mResolvers[i])
	}

	for i := 0; i < len(mResolvers); i++ {
		v := <-cchan
		if len(v) == 0 {
			continue
		}
		if dnsMap[v] >= len(mResolvers)/2 {
			return v
		}
		dnsMap[v]++
	}

	for k, v := range dnsMap {
		if v > maxCount {
			maxCount = v
			ip = k
		}
	}
	return
}
