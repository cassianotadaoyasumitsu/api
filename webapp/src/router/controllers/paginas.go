package controllers

import (
	"net/http"
	"webapp/utils"
)

// Carregar tela de login
func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

// Carregar tela de cadastro de usuário
func CarregarPaginaDeCadastrosDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}
