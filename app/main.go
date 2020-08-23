package main

import (
	"github.com/SamuelBagattin/cesi-projet-apero/config"
)

func main() {
	config.DatabaseInit()
	InitalizeRouter()
}
