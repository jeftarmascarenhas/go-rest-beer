package main

import "fmt"

func main() {
	ok, err := say("")
	if err != nil {
		panic(err.Error())
	}
	switch ok {
	case true:
		fmt.Println("Deu certo")
	default:
		fmt.Println("deu errado")
	}
}

func say(what string) (bool, error) {
	if what == "" {
		return false, fmt.Errorf("Empty string")
	}
	fmt.Println(what)
	return true, nil
}
