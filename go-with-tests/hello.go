package main

import "fmt"

const spanish = "Spanish"
const englishPrefix = "hello, "
const spanishPrefix = "hola, "

func Hello(name, language string) string {
    if name == "" {
        name = "world"
    }

    if language == spanish {
        return spanishPrefix + name
    }

    return englishPrefix + name
}

func main() {
    fmt.Println(Hello("world", ""))
}
