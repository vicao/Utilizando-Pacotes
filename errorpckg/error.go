package errorpckg

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	fmt.Println("Aula de Pacotes em Go - Error")
	return fmt.Sprintf("Erro em %v: %v", e.When, e.What)
}

func Funcoops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"O arquivo do sistema desapareceu",
	}
}
