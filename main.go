package main

import(

	"fmt"
	//"github.com/gin-gonic/gin"
	//"github.com/vincent119/go-tg-bot/config"
	"github.com/vincent119/go-tg-bot/middleware/configs"
	//"./config/getconfig"
)
func init(){
	configs.Init()
}

func main(){
	port := configs.ServerPort()
	fmt.Println(port)








}