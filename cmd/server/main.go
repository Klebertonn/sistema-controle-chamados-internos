package main

import (
	"log"
	"sistemas-controle-chamdos-internos/config"
	"sistemas-controle-chamdos-internos/database/migrations"
	"sistemas-controle-chamdos-internos/internal/handlers"
	"sistemas-controle-chamdos-internos/internal/repositories"
	"sistemas-controle-chamdos-internos/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	// Conexão com banco
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := migrations.Migrate(db); err != nil {
		log.Fatal(err)
	}

	//Repositories
	chamadoRepo := repositories.NewChamadoRepository(db)

	//Service
	chamadoService := service.NewChamadoService(chamadoRepo)

	//Handler

	chamadohandler := handlers.NewChamadoHandler(chamadoService)

	//Router
	r := gin.Default()

	r.POST("/chamados", chamadohandler.CriarChamados)

	r.GET("/chamados", chamadohandler.ListarChamados)

	r.GET("/chamados/:id", chamadohandler.BuscarOrdemID)

	r.PUT("/chamados/:id", chamadohandler.AtualizarChamado)

	log.Println("Servidor Iniciado na porta 8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
