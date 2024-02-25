package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("say hello to person", func (t *testing.T) {
        got := Hello("nealwp")
        want := "hello, nealwp"

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })

    t.Run("say hello world when empty string", func (t *testing.T) {
        got := Hello("")
        want := "hello, world"

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })
}
