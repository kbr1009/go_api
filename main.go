package main

// go modulesで管理すると絶対パスでのimportが必須です
import (
	"go_api/controllers"
)

func main() {
	controllers.StartWebServer()
}
