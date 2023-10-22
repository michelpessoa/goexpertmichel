package main

import "github.com/michelpessoa/goexpertmichel/9-Apis/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)

	// Forma encapsulada
	// config := configs.NewConf()
	// print(config.GetDbDriver())
}
