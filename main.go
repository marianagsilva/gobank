package main

import (
	"fmt"

	"github.com/banco/contas"
)

func main() {

	var contaGuilherme ContaCorrente = ContaCorrente{ //contaGuilherme := ContaCorrente{}
		Titular:       "Guilherme",
		NumeroAgencia: 589,
		NumeroConta:   12345,
		Saldo:         125.5,
	}

	var contaGuilherme2 ContaCorrente = ContaCorrente{ //contaGuilherme := ContaCorrente{}
		Titular:       "Guilherme",
		NumeroAgencia: 580,
		NumeroConta:   12345,
		Saldo:         125.5,
	}
	fmt.Println(contaGuilherme == contaGuilherme2)

	contaBruna := contas.ContaCorrente{"Bruna", 222, 1112223, 200} // mesma coisa que acima, declaração curta
	contaBruna2 := contas.ContaCorrente{"Bruna", 222, 1112223, 200}

	fmt.Println(contaBruna == contaBruna2)

	var contaCris *ContaCorrente
	contaCris = new(ContaCorrente)
	contaCris.Titular = "Cris"
	contaCris.Saldo = 500

	var contaCris2 *ContaCorrente
	contaCris2 = new(ContaCorrente)
	contaCris2.Titular = "Cris"
	contaCris2.Saldo = 500

	fmt.Println(*contaCris == *contaCris2)
	fmt.Println(contaCris2)

	contaSilvia := ContaCorrente{}
	contaSilvia.titular = "Silvia"
	contaSilvia.Saldo = 500

	fmt.Println(contaSilvia.saldo)
	fmt.Println(contaSilvia.Depositar(2500))
	//fmt.Println(contaSilvia.Sacar(300))

	contaSilvia = ContaCorrente{titular: "Silvia", saldo: 300}
	contaGustavo := ContaCorrente{titular: "Gustavo", saldo: 100}

	status := contaSilvia.Transferir(200, &contaGustavo)

	fmt.Println(" ")
	fmt.Println("Transferência")
	fmt.Println(status)
	fmt.Println(contaSilvia)
	fmt.Println(contaGustavo)

}
