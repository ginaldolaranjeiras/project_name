package conta

import (
	"fmt"

	pessoa "gino.com/project_name/pessoa"
)

type Conta struct {
	Titular pessoa.Pessoa
	Numero  string
	saldo   float64
}

func (c *Conta) ObterSaldo() float64 {
	return c.saldo
}

func (c *Conta) Depositar(valor float64) {
	if valor > 0 {
		c.saldo += valor
	} else {
		fmt.Println("Erro: valor inválido")
	}
}

func (c *Conta) Sacar(valor float64) {
	if valor > 0 && valor <= c.saldo {
		c.saldo -= valor
	} else {
		fmt.Println("Erro: valor inválido")
	}
}
