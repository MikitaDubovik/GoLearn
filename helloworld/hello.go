package helloworld

const (
	spanish = "Spanish"
	french  = "French"

	spanishHelloPrefix = "Hola "
	frenchHelloPrefix  = "Bonjour "
	englishHelloPrefix = "Hello "
)

func main() {
	print(Hello("World", ""))
}

func Hello(name, language string) string {
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
	default:
		prefix = englishHelloPrefix
	}
	return
}
