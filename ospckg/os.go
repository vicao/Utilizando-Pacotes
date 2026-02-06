package ospckg

import (
	"fmt"
	"log"
	"os"
)

func Funcos() {
	fmt.Println("Aula de Pacotes em Go - Os")
	// Criando um diretório
	err := os.Mkdir("diretorio_teste", 0755)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Diretório criado com sucesso!")

	// Listando arquivos no diretório atual
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Arquivos no diretório atual:")
	for _, file := range files {
		fmt.Println(file.Name())
	}
	fmt.Println("================================")
}

func FuncEscreverArquivo(texto string) {
	// Criando um arquivo e escrevendo nele
	f, err := os.Create("teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(texto + "\n")
	if err != nil {
		log.Fatal(err)
	}
}
