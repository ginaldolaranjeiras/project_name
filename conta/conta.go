package conta

import pessoa "gino.com/project_name/pessoa"

type Conta struct {
	Titular pessoa.Pessoa
	Numero  string
	Saldo   float64
}
