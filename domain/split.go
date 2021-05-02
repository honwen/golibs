package domain

import (
	"log"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func SplitDomainToRR(fulldomain string) (rr, mainDomain string) {
	wildCard := false
	if strings.HasPrefix(fulldomain, `*.`) {
		wildCard = true
		fulldomain = fulldomain[2:]
	}

	for len(fulldomain) > 0 && strings.HasSuffix(fulldomain, `.`) {
		fulldomain = fulldomain[:len(fulldomain)-1]
	}

	eTLD, _ := publicsuffix.PublicSuffix(fulldomain)
	mainDomain, _ = publicsuffix.EffectiveTLDPlusOne(fulldomain)
	if !IsValidHostname(fulldomain) || len(eTLD) == 0 || len(mainDomain) <= len(eTLD) {
		log.Fatal("Not a Vaild Domain")
		return
	}

	if fulldomain != mainDomain {
		rr = fulldomain[:len(fulldomain)-len(mainDomain)-1]
	}
	if wildCard {
		if len(rr) == 0 {
			rr = `*`
		} else {
			rr = `*.` + rr
		}
	}

	if len(rr) == 0 {
		rr = `@`
	}

	// fmt.Println(rr, domain)
	return
}
