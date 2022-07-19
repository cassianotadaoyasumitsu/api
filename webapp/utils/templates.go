package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// CarregarTemplates carrega todos os templates do projeto
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// Executar Template renderiza um template
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
