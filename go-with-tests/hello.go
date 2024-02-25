package main

import "fmt"

const englishPrefix = "hello, "

func Hello(name, language string) string {
    if name == "" {
        name = "world"
    }

    if language == "Spanish" {
        return "hola, " + name
    }

    return englishPrefix + name
}

func main() {
    fmt.Println(Hello("world", ""))
}
