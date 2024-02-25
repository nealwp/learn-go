package arrayslice

func Sum(numbers [5]int) (sum int) {
    sum = 0
    for i := range 5 {
        sum += numbers[i]
    }
    return
}
