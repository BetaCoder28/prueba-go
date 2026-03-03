package main

// Librería para inyección de dependencias
import (
	"github.com/BetaCoder28/prueba-go/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		// Todas las funcionas que devuelvan un struct
		fx.Provide(
			settings.New,
		),
		//Comando que necesitemos ejecutar justo antes que la aplicación corra
		fx.Invoke(),
	)

	app.Run()
}
