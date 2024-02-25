package helloworld

import "fmt"

const (
    spanish = "Spanish"
    french = "French"
    englishPrefix = "hello, "
    spanishPrefix = "hola, "
    frenchPrefix = "bonjour, "
)

func Hello(name, language string) string {
    if name == "" {
        name = "world"
    }

    return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
    switch language {
    case french:
        prefix = frenchPrefix
    case spanish:
        prefix = spanishPrefix
    default:
        prefix = englishPrefix
    }

    return
}

func main() {
    fmt.Println(Hello("world", ""))
}
