package main

import "fmt"

const englishPrefix = "hello, "

func Hello(name string) string {
    return englishPrefix + name
}

func main() {
    fmt.Println(Hello("world"))
}
