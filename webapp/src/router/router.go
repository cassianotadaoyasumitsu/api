package router

import (
	"webapp/src/router/rotas"

	"github.com/gorilla/mux"
)

// Retorna rotas do sistema
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)
}
