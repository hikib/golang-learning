package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const danish = "Danish"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const danishHelloPrefix = "Hej, "
const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case danish:
		prefix = danishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
