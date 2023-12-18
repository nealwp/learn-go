package binarysearch

import "testing"

type testCase struct {
    list []int
    search_for int
    expected int
}

var testCases = []testCase {
    {[]int{}, 1, -1},
    {[]int{1}, 1, 0},
    {[]int{1,2,3}, 2, 1},
    {[]int{1,2,3,4,5,6}, 6, 5},
    {[]int{1,2,3,4,5,6}, 9, -1},
}

func TestBinarySearch(t *testing.T) {
    for _, test := range testCases {
        if output := BinarySearch(test.list, test.search_for); output != test.expected {
            t.Errorf("Output %v not equal to expected %v", output, test.expected)
        }
    }
}
