package main

import (
	"fmt"
	"log"

	"github.com/zserge/lorca"
)

func main() {
	ui, err := lorca.New("https://web.whatsapp.com", "", 800, 600)
	if err != nil {
		log.Fatal(fmt.Sprintf("Erro ao caregar o serviço: %s", err))
	}

	defer ui.Close()

	<-ui.Done()
}
