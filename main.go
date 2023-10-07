package main

import (
	"Praktikum/configs"
	"Praktikum/routes"
)

func main (){
	configs.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}