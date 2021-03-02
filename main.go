package main

import (
	"fmt"

	"github.com/banco/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDenis := contas.ContaPoupanca{}
	contaDenis.Depositar(100)
	contaDenis.Sacar(5000)

	PagarBoleto(&contaDenis, 60)

	contaLuisa := contas.ContaCorrente{}
	contaLuisa.Depositar(500)
	PagarBoleto(&contaLuisa, -400)

	fmt.Println(contaDenis)
	fmt.Println(contaLuisa)

}
