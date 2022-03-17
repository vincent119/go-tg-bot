package main

import(

	"fmt"
	//"github.com/gin-gonic/gin"
	//"github.com/vincent119/go-tg-bot/config"
	"github.com/vincent119/go-tg-bot/middleware/config"
	
	//"./config/getconfig"
)
func init(){
	config.Init()
}

func main(){
	port := config.ServerPort()
	fmt.Println(port)








}