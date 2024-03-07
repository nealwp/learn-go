package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T){

    t.Run("compares speed of the servers, return fastest one", func(t *testing.T) {
        slowServer := makeDelayedServer(20 * time.Millisecond)
        fastServer := makeDelayedServer(0 * time.Millisecond) 

        defer slowServer.Close()
        defer fastServer.Close()

        slowUrl := slowServer.URL
        fastUrl := fastServer.URL

        want := fastUrl
        got, _ := Racer(slowUrl, fastUrl)

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })

    t.Run("returns an error if a server doesn't respond within 10 seconds", func(t *testing.T) {
        slowServer := makeDelayedServer(11 * time.Second)
        fastServer := makeDelayedServer(12 * time.Second) 

        defer slowServer.Close()
        defer fastServer.Close()

        slowUrl := slowServer.URL
        fastUrl := fastServer.URL

        _, err := Racer(slowUrl, fastUrl)

        if err == nil {
            t.Error("wanted error but got none")
        }
    })
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
    return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(delay)
        w.WriteHeader(http.StatusOK)
    }))
}
