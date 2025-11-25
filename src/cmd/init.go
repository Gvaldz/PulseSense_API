package cmd

import (
	"log"
	"pulse_sense/src/app"
	"pulse_sense/src/core"
)

func Init() {
	core.LoadConfig()


	app, err := app.NewApplication()
	if err != nil {
		log.Fatal("Error al inicializar la aplicación:", err)
	}
	defer app.Close()

	if err := app.Start(); err != nil {
		log.Fatal("Error al iniciar la aplicación:", err)
	}
}
