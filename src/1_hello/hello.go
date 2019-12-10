package main

import "fmt"

const sapaan = "Hello, "
const sapaanIndo = "Hallo, "
const sapaanFrench = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	return sapaanPrefix(lang) + name
}

func sapaanPrefix(lang string) (prefix string) {
	switch lang {
	case "French":
		prefix = sapaanFrench
	case "Indo":
		prefix = sapaanIndo
	default:
		prefix = sapaan
	}
	return
	// amazing, dengan bikin (prefix string) buat return value-nya gak perlu di tulis lagi di returnya karena otomatis ngebuat variable 'prefix'

}

func main() {
	fmt.Println(Hello("Tian"))
}
