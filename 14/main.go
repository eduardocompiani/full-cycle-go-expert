package main

import "fmt"

type Cliente struct {
	nome string
}

type Conta struct {
	saldo int
}

func (c Conta) simularDeposito(valor int) int {
	c.saldo += valor
	return c.saldo
}

func (c *Conta) efetivarDeposito(valor int) int {
	c.saldo += valor
	return c.saldo
}

func (c Cliente) andou() {
	c.nome = "Eduardo Compiani"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func main() {
	// exemplo passando copia
	compiani := Cliente{
		nome: "Compiani",
	}

	compiani.andou()
	fmt.Printf("O valor da struct com nome %v\n", compiani.nome)

	conta := Conta{saldo: 100}
	valor := conta.simularDeposito(200)
	fmt.Printf("O saldo da conta simulado é %v\n", valor)
	fmt.Printf("O saldo da conta é %v\n", conta.saldo)
	conta.efetivarDeposito(200)
	println(conta.saldo)
}
