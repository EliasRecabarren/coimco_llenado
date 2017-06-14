package main

import (
	"fmt"
	"github.com/franela/goreq"
	)

type Login struct {
	Mail string
	Pass string
}

func init() {
	login.Mail = "admin@mail.com"
	login.Pass = "admin"
}

var login Login

func main() {
	fmt.Println(login)

	res, _ := goreq.Request{
		Method: "POST",
		Uri:    "http://coimco.herokuapp.com/login",
		Body:   login,
	}.Do()
	fmt.Println(res.Body.ToString())
	
}
