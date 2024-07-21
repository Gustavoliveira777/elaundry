package main

import (
	"elaundry/model"
	"fmt"
)

func main() {
	fmt.Println("E aí, tudo bom?")
	pessoa := model.UsuarioFactory(10, "José Alencar", "ze@zemail.com", "12312312", false)
	// pessoa.SetNomeCompleto("José")
	fmt.Printf("Nome completo: %s\n Id: %d", pessoa.GetNomeCompleto(), pessoa.GetUsuarioId())
}
