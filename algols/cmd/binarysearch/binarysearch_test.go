package binarysearch

import "testing"

func TestBinarySearchRetunsNegativeOneWhenNotFound(t *testing.T) {
    got := binarysearch()
    want := -1

    if want != got {
        t.Fatalf("wanted %v, got %v", want, got)
    }
}
