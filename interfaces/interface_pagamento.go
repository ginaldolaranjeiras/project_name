package interfaces

func PagarBoleto(conta pagamento, valorDoTitulo float64) {
	conta.Sacar(valorDoTitulo)
}

type pagamento interface {
	Sacar(valor float64)
}
