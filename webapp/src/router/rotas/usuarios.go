package rotas

import (
	"net/http"
	"webapp/src/router/controllers"
)

var rotasUsuarios = []Rota{
	{
		URI:              "/criar-usuario",
		Metodo:           http.MethodGet,
		Funcao:           controllers.CarregarPaginaDeCadastrosDeUsuario,
		RequerAutenticao: false,
	},
}
