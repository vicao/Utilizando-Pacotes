package iopckg

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func Iofunc() {
	fmt.Println("Aula de Pacotes em Go - IO")
	if _, err := io.WriteString(os.Stdout, "Olá Mundo!\n"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("================================")
}

func Ioutil() {
	fmt.Println("Aula de Pacotes em Go - IOUtil")
	fmt.Println("Criando um arquivo")
	message := []byte("Hello, Gophers!")
	err := ioutil.WriteFile("teste.txt", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("================================")
}
