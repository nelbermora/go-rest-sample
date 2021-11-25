package service

import "github.com/nelbermora/go-oracle-sample/internal/repository"

// en este pkg puede ir la logica de negocio

func DummyFunc() (string, error) {
	// para este ejemplo la funcion solo ejecutara la llamada al SP,
	// si esta se ejecuta correctamente entonces retornara un mensajem si no el error
	return repository.LlamarSP()
}
