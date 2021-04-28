package hello
import "fmt"
const englishPrefix = "Hello,"
const spanishPrefix = "Hola,"
const frenchPrefix = "Bonjour"
const spanish = "Spanish"
const french = "French"
func greetingPrefix(language string) (prefix string){
	switch language{
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}
func Hello(name string, language string) string{
	if name == ""{
		name = "World"
	}
	return greetingPrefix(language)+name
}

func main(){
	fmt.Println(Hello("","English"))
}