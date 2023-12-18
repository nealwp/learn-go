package binarysearch

func BinarySearch(list []int, search int) int {
    low := 0
    high := len(list) - 1

    if high == -1 {
        return -1
    }
    
    if high == low && list[0] == search {
        return 0
    }

    for low <= high {
        mid := (low + high) / 2
        guess := list[mid]
        if guess == search {
            return mid
        } else if guess > search {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return -1
}
