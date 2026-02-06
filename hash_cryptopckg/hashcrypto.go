package hashcryptopckg

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
)

func FuncHas() {
	fmt.Println("Aula de Pacotes em Go - Hash/Crc32")
	h := crc32.NewIEEE()
	h.Write([]byte("Código com pacote hash/crc32"))
	v := h.Sum32()
	fmt.Printf("Valor CRC32: %v\n", v)
}

func FuncCripto() {

	fmt.Println("Aula de Pacotes em Go - Criptografia")
	h := sha1.New()
	h.Write([]byte("Código com pacote hash/crc32"))
	bs := h.Sum([]byte{})
	fmt.Printf("Valor SHA1: %x\n", bs)
}
