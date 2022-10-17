package conta

import (
	"fmt"

	pessoa "gino.com/project_name/pessoa"
)

type ContaPoupanca struct {
	Titular pessoa.Pessoa
	Numero  string
	saldo   float64
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}

func (c *ContaPoupanca) Depositar(valor float64) {
	if valor > 0 {
		c.saldo += valor
	} else {
		fmt.Println("Erro: valor inválido")
	}
}

func (c *ContaPoupanca) Sacar(valor float64) {
	if valor > 0 && valor <= c.saldo {
		c.saldo -= valor
	} else {
		fmt.Println("Erro: valor inválido")
	}
}
