package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	var contaGuilherme ContaCorrente = ContaCorrente{ //contaGuilherme := ContaCorrente{}
		titular:       "Guilherme",
		numeroAgencia: 589,
		numeroConta:   12345,
		saldo:         125.5,
	}
	contaBruna := ContaCorrente{"Bruna", 222, 1112223, 200} // mesma coisa que acima, declaração curta

	fmt.Println(contaGuilherme)
	fmt.Println(contaBruna)
}
