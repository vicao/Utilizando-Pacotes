package sortpckg

import (
	"fmt"
	"sort"
)

type Dados struct {
	Nome  string
	Idade int
}

type ParaNome []Dados

func (ps ParaNome) Len() int {
	return len(ps)
}
func (ps ParaNome) Less(i, j int) bool { //Item na posição i é menor que o da posição j?
	return ps[i].Nome < ps[j].Nome
}
func (ps ParaNome) Swap(i, j int) { //Troca os itens de posição
	ps[i], ps[j] = ps[j], ps[i]
}

func Funcsort() {
	criancas := []Dados{
		{"Julia", 10},
		{"João", 5},
		{"Ana", 8},
	}
	sort.Sort(ParaNome(criancas))
	fmt.Println(criancas)
}
