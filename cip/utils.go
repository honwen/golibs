package cip

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

// https://www.useragents.me/
var UA_OSs = []string{
	"X11; Linux x86_64", "X11; Ubuntu; Linux x86_64",
	"Linux; Android 7.0", "Android 7.0; Mobile", "Linux; Android 10", "Android 14; Mobile",
	"Macintosh; Intel Mac OS X", "MacBook Air; M1 Mac OS X",
	"iPhone; CPU iPhone OS 14_1 like Mac OS X", "iPhone; CPU iPhone OS 17_1 like Mac OS X",
	"iPad; CPU iPhone OS 14_1 like Mac OS X", "iPad; CPU iPhone OS 17_1 like Mac OS X",
}

func wGet(url string, timeout time.Duration) (str string) {
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set(
		"User-Agent",
		fmt.Sprintf("Mozilla/5.0 (%s)", UA_OSs[rand.Intn(len(UA_OSs))]),
	)
	if err != nil {
		return
	}
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(request)
	if err != nil {
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resp.Body.Close()
	str = string(body)
	// log.Println(str)
	// log.Println(url, regexp.MustCompile(RegxIPv4).FindString(str))
	return
}
