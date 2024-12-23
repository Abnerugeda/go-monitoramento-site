package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramento int = 3
const delay = 3

func main() {
	exibeInstrucao()
	for {
		exibeMenu()

		var comando int
		fmt.Scan(&comando)
		execCommand(comando)
	}

}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair do programa")
}

func exibeInstrucao() {
	var nome string = "Abner"
	var versao float32 = 1.1
	fmt.Println("olá, sr", nome)
	fmt.Println("Este programa está na versão", versao)
}

func execCommand(comando int) {
	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço essa instrução")
		os.Exit(-1)
	}
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSiteDoArquivo()

	for i := 0; i < monitoramento; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		message := fmt.Sprintf("Teste %d/%d", i+1, monitoramento)
		fmt.Println(message)
		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status code: ", resp.StatusCode)
	}
}

func leSiteDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("./sites.txt")

	if err != nil {
		fmt.Println("Ocorreu algum erro:", err)
		return nil
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, lErr := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if lErr == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}
