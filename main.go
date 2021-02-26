package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {

	var contaGuilherme ContaCorrente = ContaCorrente{ //contaGuilherme := ContaCorrente{}
		titular:       "Guilherme",
		numeroAgencia: 589,
		numeroConta:   12345,
		saldo:         125.5,
	}

	var contaGuilherme2 ContaCorrente = ContaCorrente{ //contaGuilherme := ContaCorrente{}
		titular:       "Guilherme",
		numeroAgencia: 580,
		numeroConta:   12345,
		saldo:         125.5,
	}
	fmt.Println(contaGuilherme == contaGuilherme2)

	contaBruna := ContaCorrente{"Bruna", 222, 1112223, 200} // mesma coisa que acima, declaração curta
	contaBruna2 := ContaCorrente{"Bruna", 222, 1112223, 200}

	fmt.Println(contaBruna == contaBruna2)

	var contaCris *ContaCorrente
	contaCris = new(ContaCorrente)
	contaCris.titular = "Cris"
	contaCris.saldo = 500

	var contaCris2 *ContaCorrente
	contaCris2 = new(ContaCorrente)
	contaCris2.titular = "Cris"
	contaCris2.saldo = 500

	fmt.Println(*contaCris == *contaCris2)
	fmt.Println(contaCris2)

	contaSilvia := ContaCorrente{}
	contaSilvia.titular = "Silvia"
	contaSilvia.saldo = 500

	fmt.Println(contaSilvia.saldo)
	fmt.Println(contaSilvia.Sacar(300))

}
