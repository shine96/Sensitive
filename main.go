package main

import (
	"Sensitive/app/common"
	"Sensitive/app/routes"
)

func main() {
	if common.IsFileExist("./dict.txt") == false {
		common.DownloadFile("https://media.bnickolas.com/dict_1640512181457.txt", "./dict.txt")
	}
	serve := routes.InitRoute()
	serve.Run(":9000")
}
