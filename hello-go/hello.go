package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
	frenchHelloPrefix  = "Bonjour"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return fmt.Sprintf("%s, %s", greetingPrefix(language), name)
}


func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func main() {
	fmt.Println(Hello("Go!", ""))
}
