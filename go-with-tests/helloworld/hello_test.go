package helloworld

import "testing"

func TestHello(t *testing.T) {
    t.Run("say hello to person", func (t *testing.T) {
        got := Hello("nealwp", "")
        want := "hello, nealwp"

        assertCorrectMessage(t, got, want)
    })

    t.Run("say hello world when empty string", func (t *testing.T) {
        got := Hello("", "")
        want := "hello, world"

        assertCorrectMessage(t, got, want)
    })

    t.Run("say hello in spanish", func (t *testing.T) {
        got := Hello("nealwp", "Spanish")
        want := "hola, nealwp"

        assertCorrectMessage(t, got, want)
    })

    t.Run("say hello in french", func (t *testing.T) {
        got := Hello("nealwp", "French")
        want := "bonjour, nealwp"

        assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
