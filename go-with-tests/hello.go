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

    prefix := englishPrefix

    switch language {
    case french:
        prefix = frenchPrefix
    case spanish:
        prefix = spanishPrefix
    }

    return prefix + name
}

func main() {
    fmt.Println(Hello("world", ""))
}
