package main

import (
	"fmt"

	conta "gino.com/project_name/conta"
	pessoa "gino.com/project_name/pessoa"
)

func main() {

	gino := pessoa.Pessoa{
		Nome: "Gino",
	}

	amanda := pessoa.Pessoa{
		Nome: "Amanda",
	}

	contaDoGino := conta.Conta{
		Numero:  "123456",
		Titular: gino,
	}

	contaDaAmanda := conta.Conta{
		Numero:  "654321",
		Titular: amanda,
	}

	fmt.Println(contaDoGino, contaDaAmanda)

	contaDoGino.Depositar(100)
	contaDaAmanda.Depositar(20)

	fmt.Println(contaDoGino.ObterSaldo(), contaDaAmanda.ObterSaldo())
}
