package service

import "github.com/nelbermora/go-rest-sample/internal/repository"

// en este pkg puede ir la logica de negocio

func DummyFunc() (string, error) {
	// para este ejemplo la funcion solo ejecutara la llamada al SP,
	// si esta se ejecuta correctamente entonces retornara un mensajem si no el error
	// comento esta linea para incorporar ahora al llamado del API Rest
	// return repository.LlamarSP()
	return repository.LlamarRest()
}
