package main

import (
	//  net/http é responsável pelos métodos de requisições
	"net/http"
)

// handler recebe dois parâmetros (requisição e resposta)
// um do tipo http.ResponseWriter e outro do tipo *http.Request
func handler(w http.ResponseWriter, r *http.Request) {

	// http.ServeFile é responsável por disponibilizar o arquivo front-end
	http.ServeFile(w, r, "./static/index.html")
}

// função principal responsável pela execução
func main() {
	// HandleFunc recebe dois parâmetros,
	// o primeiro informa que a URL http://localhost:8080 deverá executar determinada ação.
	// O segundo parâmetro indica que chamaremos a função handler ao acessar /
	http.HandleFunc("/", handler)

	// ListenAndServe especifica a porta e config do servidor (nil = padrão)
	http.ListenAndServe(":8080", nil)
}

// no terminal: go run main.go
// no navegador: http://localhost:8080/
