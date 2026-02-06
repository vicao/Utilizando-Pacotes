package main

import (
	"fmt"
	"utilizandopacotes/bytespckg"
	"utilizandopacotes/iopckg"
	"utilizandopacotes/ospckg"
	"utilizandopacotes/stringpckg"
	"utilizandopacotes/path_filepathpckg"
	"utilizandopacotes/errorpckg"
	"utilizandopacotes/container_listpckg"
	"utilizandopacotes/sortpckg"
	"utilizandopacotes/hash_cryptopckg"
)

func main() {
	iopckg.Iofunc()
	iopckg.Ioutil()
	stringspckg.Stringsfunc()
	fmt.Println(bytespckg.IndexByte([]byte("hello"), 'e'))
	bytespckg.Title()
	ospckg.Funcos()
	pathfilepathpckg.Funcwalk()
		if err:= errorpckg.Funcoops(); err != nil {
		fmt.Println(err)

	}
	containerlistpckg.FuncList()
	sortpckg.Funcsort()
	hashcryptopckg.FuncHas()
	hashcryptopckg.FuncCripto()
}
