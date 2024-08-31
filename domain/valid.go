package domain

import "github.com/AdguardTeam/golibs/netutil"

// refer: https://github.com/AdguardTeam/golibs/blob/v0.27.0/netutil/addr.go#L222
func IsValidHostname(hostname string) bool {
	return netutil.ValidateDomainName(hostname) == nil
}
