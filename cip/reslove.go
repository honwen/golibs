package cip

import (
	"context"
	"io"
	"log/slog"
	"strings"
	"time"

	"github.com/miekg/dns"
	"github.com/mr-karan/doggo/pkg/resolvers"
)

var mResolvers = []resolvers.Resolver{}

func init() {
	upstreams := []string{
		"tls://223.5.5.5:853", "tls://223.6.6.6:853", "https://223.5.5.5/dns-query", "https://223.6.6.6/dns-query",
		// "tls://1.12.12.12:853", "https://120.53.53.53/dns-query",
		"tls://1.12.12.12:853", "tls://120.53.53.53:853", "https://1.12.12.12/dns-query", "https://120.53.53.53/dns-query",
		"https://1.15.50.48/verse", "https://106.52.218.142/verse",
	}
	opts := resolvers.Options{
		Timeout: 2000 * time.Millisecond,
		Logger:  slog.New(slog.NewJSONHandler(io.Discard, nil)),
	}
	dotOpts := resolvers.ClassicResolverOpts{
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
	ctx := context.Background()

	qFlags := resolvers.QueryFlags{}

	for i := range mResolvers {
		go func(resolver *resolvers.Resolver) {
			addr := "" // SOA
			resp, err := (*resolver).Lookup(ctx, []dns.Question{{Name: domain, Qtype: qtype, Qclass: dns.ClassINET}}, qFlags)
			if err == nil && len(resp) > 0 && len(resp[0].Answers) > 0 {
				for _, it := range resp[0].Answers {
					if it.Type == dns.TypeToString[qtype] {
						addr = it.Address
						break
					}
				}
			}
			// fmt.Println(len(resp.Answers), addr, (*resolver), err)
			cchan <- addr
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
