package website_racer

import (
	"go_playground/src/website_racer"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func createMockServerWithDelay(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}

func TestWebsiteRacer(t *testing.T) {
	t.Run("race 2 urls return the fastest response url", func(t *testing.T) {
		slowServer := createMockServerWithDelay(20 * time.Millisecond)
		fastServer := createMockServerWithDelay(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		expected := fastUrl
		actual, _ := website_racer.WebsiteRacer(fastUrl, slowUrl)

		if actual != expected {
			t.Errorf("actual %q, expected %q", actual, expected)
		}
	})

	t.Run("when a server doesn't response within an amount of time should return error", func(t *testing.T) {
		server1 := createMockServerWithDelay(250 * time.Millisecond)
		server2 := createMockServerWithDelay(300 * time.Millisecond)
		defer server1.Close()
		defer server2.Close()

		_, err := website_racer.ConfigurableWebsiteRacer(server1.URL, server2.URL, 200*time.Millisecond)

		if err == nil {
			t.Errorf("expected error but didn't get one")
		}
	})
}
