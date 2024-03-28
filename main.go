package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func CEP1() {
	// Faça uma solicitação GET a uma URL
	resp, err := http.Get("http://viacep.com.br/ws/01153000/json/")
	if err != nil {
		fmt.Println("Erro ao fazer a solicitação:", err)
		return
	}
	defer resp.Body.Close()

	// Leia o corpo da resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
		return
	}

	// Imprima o corpo da resposta

	fmt.Println(string(body))
}
func CEP2() {
	// Faça uma solicitação GET a uma URL
	resp, err := http.Get("https://brasilapi.com.br/api/cep/v1/01153000 + cep/")
	if err != nil {
		fmt.Println("Erro ao fazer a solicitação:", err)
		return
	}
	defer resp.Body.Close()

	// Leia o corpo da resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
		return
	}

	// Imprima o corpo da resposta
	fmt.Println(" ")
	fmt.Println(string(body))
	fmt.Println(" ")

}
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		//time.Sleep(time.Second)

		CEP1()
		c1 <- "Via CEP"

	}()

	go func() {

		CEP2()
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
