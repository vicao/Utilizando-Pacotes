package bytespckg

import (
	"bytes"
	"fmt"
)

// IndexByte returns the index of the first instance of c in b, or -1 if c is not present in b.
func IndexByte(b []byte, c byte) int {
	fmt.Println("Aula de Pacotes em Go - Bytes")
	return bytes.IndexByte(b, c)
}

func Title() {
	fmt.Println("Função Title")
	fmt.Println("A função Title está deprecada, use a função ToTitle para converter uma string para título.")
	fmt.Println("================================")
	//fmt.Print("%s", bytes.Title([]byte("Pacote bytes")))
}
