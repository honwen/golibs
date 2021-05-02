package domain

import "regexp"

const RegxHostname = "^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])$"

// refer: https://github.com/AdguardTeam/golibs/blob/master/utils/utils.go
func IsValidHostname(hostname string) bool {
	return regexp.MustCompile(RegxHostname).MatchString(hostname)
}
