package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntro()
	for {
		exibeMenu()

		comando := leComando()

		fmt.Println("Comando selecionado:", comando)

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 3:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido.")
			os.Exit(-1)
		}
	}
}

func exibeIntro() {
	nome := "bruno"
	versao := "v0.1.0"
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair")
}

func leComando() int {
	var comando int
	fmt.Print("Digite o comando: ")
	fmt.Scan(&comando)
	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// site com URL inexistente
	site := "https://httpbin.org/status/404" // ou 200
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
