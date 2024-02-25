package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishPrefix = "hello, "
const spanishPrefix = "hola, "
const frenchPrefix = "bonjour, "

func Hello(name, language string) string {
    if name == "" {
        name = "world"
    }

    if language == spanish {
        return spanishPrefix + name
    }

    if language == french {
        return frenchPrefix + name
    }

    return englishPrefix + name
}

func main() {
    fmt.Println(Hello("world", ""))
}
