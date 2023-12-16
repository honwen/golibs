package domain

import "github.com/AdguardTeam/golibs/netutil"

// refer: https://github.com/AdguardTeam/golibs/blob/master/netutil/addr.go#L215
func IsValidHostname(hostname string) bool {
	return netutil.ValidateDomainName(hostname) == nil
}
