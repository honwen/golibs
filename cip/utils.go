package cip

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var UA_OSs = []string{
	"X11; Linux x86_64", "Linux; Android 7.0", "Macintosh; Intel Mac OS X",
	"Android 7.0; Mobile", "iPhone; CPU iPhone OS 12_1 like Mac OS X",
}

func wGet(url string, timeout time.Duration) (str string) {
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", fmt.Sprintf("Mozilla/5.0 (%s)", UA_OSs[rand.Intn(len(UA_OSs))]))
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
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	str = string(body)
	// fmt.Println(url, regexp.MustCompile(regxIP).FindString(str))
	return
}
