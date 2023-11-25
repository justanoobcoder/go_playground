package website_racer

import (
	"fmt"
	"net/http"
	"time"
)

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func WebsiteRacer(url1, url2 string) (string, error) {
	return ConfigurableWebsiteRacer(url1, url2, 10*time.Second)
}

func ConfigurableWebsiteRacer(url1, url2 string, delay time.Duration) (string, error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(delay):
		return "", fmt.Errorf("timed out")
	}
}
