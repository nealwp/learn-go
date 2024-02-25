package arrayslice

func Sum(numbers []int) (sum int) {
    sum = 0
    for _, number := range numbers {
        sum += number
    }
    return
}

func SumAll(numbersToSum ...[]int) []int {
    var sums []int
    for _, numbers := range numbersToSum {
        sums = append(sums, Sum(numbers))
    }

    return sums
}
