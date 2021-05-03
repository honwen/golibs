package domain

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"net/url"
	"regexp"
	"strings"

	"github.com/Workiva/go-datastructures/set"
)

func trimSite(s string) string {
	site := strings.TrimSpace(s)
	switch {
	case strings.HasSuffix(site, `USA`):
		site = site[0:strings.Index(site, `USA`)]
	case strings.HasSuffix(site, `--`):
		site = site[0:strings.Index(site, `--`)]
	}
	return site
}

func ExtractFromBytes(data []byte) (domains []string) {
	// ---------------------------------gfwlist---------------------------------
	tmpList := set.New()
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())

		if s == "" ||
			strings.HasPrefix(s, "[") ||
			strings.HasPrefix(s, "!") ||
			strings.HasPrefix(s, "||!") ||
			strings.HasPrefix(s, "@@") {
			continue
		}

		if unURL, err := url.QueryUnescape(s); err == nil {
			s = unURL
		}

		switch {
		case strings.HasPrefix(s, "||"):
			site := strings.Split(s[2:], "/")[0]
			switch {
			case strings.Contains(site, "*."):
				parts := strings.Split(site, "*.")
				site = parts[len(parts)-1]
			case strings.HasPrefix(site, "*"):
				parts := strings.SplitN(site, ".", 2)
				site = parts[len(parts)-1]
			}
			tmpList.Add(trimSite(site))
		case strings.HasPrefix(s, "|https://"):
			fallthrough
		case strings.HasPrefix(s, "|http://"):
			if u, err := url.Parse(s[1:]); err == nil {
				site := u.Host
				switch {
				case strings.Contains(site, "*."):
					parts := strings.Split(site, "*.")
					site = parts[len(parts)-1]
				case strings.HasPrefix(site, "*"):
					parts := strings.SplitN(site, ".", 2)
					site = parts[len(parts)-1]
				}
				tmpList.Add(trimSite(site))
			}
		case strings.HasPrefix(s, "."):
			site := strings.Split(strings.Split(s[1:], "/")[0], "*")[0]
			if !strings.HasSuffix(site, ".") {
				tmpList.Add(trimSite(site))
			}
		case !strings.ContainsAny(s, "*"):
			site := strings.Split(s, "/")[0]
			if regexp.MustCompile(`^[a-zA-Z0-9\.\_\-]+$`).MatchString(site) {
				tmpList.Add(trimSite(site))
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return
	}
	// ---------------------------------gfwlist---------------------------------
	for _, site := range tmpList.Flatten() {
		if strings.ContainsAny(site.(string), ".") {
			domains = append(domains, site.(string))
		}
	}
	return
}

func ExtractFromB64(base64Str string) (domains []string) {
	if d, err := base64.StdEncoding.DecodeString(base64Str); err != nil {
		return nil
	} else {
		return ExtractFromBytes(d)
	}
}
