package main

import (
	"docker/configs"
	"docker/routes"
)

func main (){
	configs.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}