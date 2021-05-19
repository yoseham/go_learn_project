package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0)
	slowURL := slowServer.URL
	fastURL := fastServer.URL

	defer slowServer.Close()
	defer fastServer.Close()

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestSelectRacer(t *testing.T) {
	t.Run("compare speeds of server", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		slowURL := slowServer.URL
		fastURL := fastServer.URL
		defer slowServer.Close()
		defer fastServer.Close()
		want := fastURL
		got, err := SelectRacer(slowURL, fastURL)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("returns an error if a server doesn't response within 10s", func(t *testing.T) {
		server := makeDelayedServer(21 * time.Millisecond)
		defer server.Close()
		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)
		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})

}
