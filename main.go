package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Função que faz a solicitação HTTP e retorna a resposta
func getResponseBody(url string) ([]byte, error) {
	// Faz a solicitação GET à URL fornecida
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return nil, err
	}

	return body, nil
}
func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	// Recupera os argumentos da linha de comando
	args := os.Args

	// Verifica se há argumentos suficientes
	if len(args) < 2 {
		fmt.Println("Passe o CEP no formato 99999999 como argumento na chamada. Exemplo: go run main.go 01153000 ")
		return
	}

	// Recupera o primeiro argumento (o segundo item em args, porque o primeiro é o nome do programa)
	cep := args[1]

	go func() {
		//  atrasando a chamada dessa API
		//	time.Sleep(time.Second)

		// Chama a função getResponseBody e obtém a resposta
		body, err := getResponseBody("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Println("Erro ao obter o corpo da resposta:", err)
			return
		}

		// Imprime o corpo da resposta
		fmt.Println(string(body))
		c1 <- "Via CEP"
	}()

	go func() {
		// Chama a função getResponseBody e obtém a resposta
		body, err := getResponseBody("https://brasilapi.com.br/api/cep/v1/" + cep + " + cep/")
		if err != nil {
			fmt.Println("Erro ao obter o corpo da resposta:", err)
			return
		}

		// Imprime o corpo da resposta
		fmt.Println(string(body))

		c2 <- "Brasil API"
	}()

	select {
	case msg1 := <-c1:
		println("Processado pela função: ", msg1)
	case msg2 := <-c2:
		println("Processado pela função: ", msg2)

	case <-time.After(time.Second):
		println("timeout ")
	}
}
