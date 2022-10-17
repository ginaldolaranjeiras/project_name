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
		Saldo:   1000,
		Titular: gino,
	}

	contaDaAmanda := conta.Conta{
		Numero:  "654321",
		Saldo:   2000,
		Titular: amanda,
	}

	fmt.Println(contaDoGino, contaDaAmanda)
}
