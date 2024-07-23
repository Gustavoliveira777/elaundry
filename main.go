package main

import (
	"elaundry/connection"
	"elaundry/model"
	"elaundry/repository"
	"fmt"
)

func main() {
	fmt.Println("E aí, tudo bom?")
	user := model.Usuario{NomeCompleto: "Zé", Email: "zezin@tonho.com", Senha: "senha", Funcionario: false}
	userRepository := repository.UsuarioRepository{DbConnection: connection.CreateDbConn()}
	userRepository.UsuarioCreateOrUpdateTable()
	userRepository.UsuarioSave(&user)

	var users []model.Usuario
	userRepository.TodosUsuarios(&users)
	fmt.Println(users)
}
