package main

import "fmt"

func main() {
	fmt.Println(Hello("world", ""))
}

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix  = "Hello, "
	spanishhHelloPrefix = "Hola, "
	frenchHelloPrefix   = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language, name)
}

func greetingPrefix(language, name string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishhHelloPrefix + name
	case french:
		prefix = frenchHelloPrefix + name
	default:
		prefix = englishHelloPrefix + name
	}
	return
}
