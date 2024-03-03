package countdown

import (
	"bytes"
	"reflect"
	"testing"
)

type SpySleeper struct {
    Calls int
}

type SpyCountdownOperations struct {
    Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
    s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
    s.Calls = append(s.Calls, write)
    return
}

const write = "write"
const sleep = "sleep"

func (s *SpySleeper) Sleep() {
    s.Calls++
}

func TestCountdown(t *testing.T){

    t.Run("print 3 to Go!", func(t *testing.T) {
        buffer := &bytes.Buffer{}

        Countdown(buffer, &SpyCountdownOperations{})

        got := buffer.String()
        want := "3\n2\n1\nGo!"

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })

    t.Run("sleep before every print", func(t *testing.T) {
        spySleepPrinter := &SpyCountdownOperations{}
        Countdown(spySleepPrinter, spySleepPrinter)
        
        want := []string{
            write,
            sleep,
            write,
            sleep,
            write,
            sleep,
            write,
        }

        if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
            t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
        }
    })
}
