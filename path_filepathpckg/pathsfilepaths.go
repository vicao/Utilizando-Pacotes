package pathfilepathpckg

import (
	"fmt"
	"os"
	"path/filepath"
	"utilizandopacotes/ospckg"
)

func Funcwalk() {
	fmt.Println("Aula de Pacotes em Go - Path e Filepath")
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		//fmt.Println(path)
		ospckg.FuncEscreverArquivo(path)		
		return nil
	})
	fmt.Println("Arquivo salvo com sucesso!")
	fmt.Println("================================")

}

/*
f,err := os.OpenFile("AulaGo.txt",os.O_RDWR|os.O_CREATE,0666)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}


*/
