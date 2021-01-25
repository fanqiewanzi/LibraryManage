package main

import (
	"connectsql"
	"login"
)


func main(){
	connectsql.Conectsql()
	login.Login()
	connectsql.DB.Close()
}




