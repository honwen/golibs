package domain

import (
	"github.com/AdguardTeam/golibs/utils"
)

func IsValidHostname(hostname string) bool {
	return utils.IsValidHostname(hostname) == nil
}
