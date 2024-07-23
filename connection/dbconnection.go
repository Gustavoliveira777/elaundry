package connection

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDbConn() *gorm.DB {
	var dsn string = ""
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get sql.DB")
	}

	// Configurar o pool de conexões
	sqlDB.SetMaxOpenConns(25) // Número máximo de conexões abertas
	sqlDB.SetMaxIdleConns(25) // Número máximo de conexões ociosas

	// Use o banco de dados

	// Fechar a conexão quando terminar

	return db
}
