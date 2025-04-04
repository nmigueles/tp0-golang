package main

import (
	"client/globals"
	"client/utils"
)

func main() {
	utils.ConfigurarLogger()
	// Inicializamos la configuración
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	// Enviamos el mensaje de la config al servidor
	utils.EnviarMensaje(
		globals.ClientConfig.Ip,
		globals.ClientConfig.Puerto,
		globals.ClientConfig.Mensaje,
	)

	// Generamos un paquete con los mensajes y lo enviamos al servidor
	utils.GenerarYEnviarPaquete()
}
