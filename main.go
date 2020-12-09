package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	exibeIntro()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido")
			os.Exit(-1)
		}
	}
}

func exibeIntro() {

	var nomeLido string
	fmt.Println("Qual o seu nome?: ")
	fmt.Scan(&nomeLido)
	fmt.Println("Olá sr(a), ", nomeLido)
	fmt.Println("Escolha uma das opções abaixo")
}

func exibeMenu() {
	fmt.Println("1 - Monitoramento")
	fmt.Println("2 - Sair do Programa")
}

func leComando() int {
	var comando int

	fmt.Scan(&comando)
	fmt.Println("O comando foi escolhido foi ", comando)

	return comando
}

func siteMonitorar() string {
	var site string

	fmt.Println("Qual site você deseja monitorar agora?")
	fmt.Scan(&site)

	return site
}

func monitoramentos() int {
	var monitoramentos int
	fmt.Println("Quantas vezes você deseja executar o monitoramento?")
	fmt.Scan(&monitoramentos)

	return monitoramentos
}

func testaSite(site string) {
	resposta, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resposta.StatusCode == 200 {
		fmt.Println("O site ", site, " está online")
		fmt.Println(" ")
	} else {
		fmt.Println("O site ", site, " está com problemas, StatusCode: ", resposta.StatusCode)
		fmt.Println(" ")
	}
}

func iniciarMonitoramento() {
	site := siteMonitorar()
	siteAtual := []string{site}
	const delay = 5
	monitoramentos := monitoramentos()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range siteAtual {
			fmt.Println("Testando o site: ", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}
}
