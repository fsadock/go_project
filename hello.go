package main

import "fmt"

const (
	spanish = "Spanish"
	french = "French"
	portuguese = "Portuguese"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	ptHelloPrefix = "Olá, "
)
func Hello(name, language string) string{
	if name == ""{
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string){
	switch language{
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = ptHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main(){
	fmt.Println(Hello("world", ""))
}
