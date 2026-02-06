//Aula sobre o que são Pacotes em Go
//Exemplo de pacotes : String,Io, Bytes, Os, Path/Filepath, Errors, Container, List

package stringspckg

import (
	"fmt"
	"strings"
)

func Stringsfunc() {
	
	fmt.Println("Aula de Pacotes em Go - Strings")
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println("================================")
}
