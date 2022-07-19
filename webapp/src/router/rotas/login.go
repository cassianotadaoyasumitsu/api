package rotas

import (
	"net/http"
	"webapp/src/router/controllers"
)

var rotasLogin = []Rota{
	{
		URI:              "/",
		Metodo:           http.MethodGet,
		Funcao:           controllers.CarregarTelaLogin,
		RequerAutenticao: false,
	},

	{
		URI:              "/login",
		Metodo:           http.MethodGet,
		Funcao:           controllers.CarregarTelaLogin,
		RequerAutenticao: false,
	},
}
