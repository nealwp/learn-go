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
        got, err := Racer(slowUrl, fastUrl, 10 * time.Second)

        if err != nil {
            t.Fatalf("did not expect error but got %v", err)
        }

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })

    t.Run("returns an error if a server doesn't respond within 10 seconds", func(t *testing.T) {
        server := makeDelayedServer(25 * time.Millisecond)

        defer server.Close()

        _, err := Racer(server.URL, server.URL, 20 * time.Millisecond)

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
