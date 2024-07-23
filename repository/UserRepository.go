package repository

import (
	"elaundry/model"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type UsuarioRepository struct {
	DbConnection *gorm.DB
}

func (u *UsuarioRepository) UsuarioCreateOrUpdateTable() {
	u.DbConnection.AutoMigrate(&model.Usuario{})
}

func (u *UsuarioRepository) UsuarioSave(usuario *model.Usuario) {
	u.DbConnection.Create(usuario)
}

func (u *UsuarioRepository) TodosUsuarios(usuarios *[]model.Usuario) {
	u.DbConnection.Find(usuarios)
}

func (u *UsuarioRepository) FecharConexao() {
	sqlDB, erroDb := u.DbConnection.DB()
	if erroDb != nil {
		log.Default().Println(erroDb.Error())
	}

	defer func() {
		if err := sqlDB.Close(); err != nil {
			fmt.Printf("failed to close database connection: %v\n", err)
		} else {
			fmt.Println("database connection closed")
		}
	}()
}
