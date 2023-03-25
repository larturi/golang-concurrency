package main

import (
	"pkgs/labs"
)

func main() {

	/*
	* Consume una API de manera concurrente
	 */
	// labs.StartTodos()

	/*
	* Genera un txt con la lista de numeros hexadecimales con concurrencia
	 */
	// labs.StartHexa()

	/*
	* Sincronización de una cuenta bancaria con depositos y extractos
	* Warning de DATA RACE
	* > go build --race main.go
	* > ./main
	* Se resuelve usando Lock() y Unlock()
	 */
	labs.StartSyncCuentaBancaria()
}
