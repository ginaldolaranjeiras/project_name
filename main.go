package main

import (
	"fmt"

	conta "gino.com/project_name/conta"
	pagamento "gino.com/project_name/interfaces"
	pessoa "gino.com/project_name/pessoa"
)

func main() {

	gino := pessoa.Pessoa{
		Nome: "Gino",
	}

	amanda := pessoa.Pessoa{
		Nome: "Amanda",
	}

	luis := pessoa.Pessoa{
		Nome: "Luis",
	}

	contaDoGino := conta.Conta{
		Numero:  "123456",
		Titular: gino,
	}

	contaDaAmanda := conta.Conta{
		Numero:  "654321",
		Titular: amanda,
	}

	contaDoLuis := conta.ContaPoupanca{
		Numero:  "123456",
		Titular: luis,
	}

	fmt.Println(contaDoGino, contaDaAmanda, contaDoLuis)

	contaDoGino.Depositar(100)

	contaDaAmanda.Depositar(20)

	fmt.Println(contaDoGino.ObterSaldo())

	boleto := 50.98

	pagamento.PagarBoleto(&contaDoGino, boleto)

	fmt.Println(contaDoGino.ObterSaldo())

	contaDoLuis.Depositar(800)

	fmt.Println(contaDoLuis.ObterSaldo())

	boletoLuis := 32.56

	pagamento.PagarBoleto(&contaDoLuis, boletoLuis)

	fmt.Println(contaDoLuis.ObterSaldo())

}
