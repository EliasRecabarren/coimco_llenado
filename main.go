package main

//go get github.com/franela/goreq

//import "github.com/franela/goreq"
import (
	"fmt"
	"github.com/franela/goreq"
	)

func init() {
	login.Mail = "admin"
	login.Pass = "admin"
}
var login Login
var customer Customer
var provider Provider
//var product Product

func main() {
	fmt.Println(login)

	for ix := 0; ix < 10; ix++ {
		res, _ := goreq.Request{
			Method: "POST",
			Uri:    "http://coimco.herokuapp.com/login",
			Body:   login,
		}.Do()
		fmt.Println(res.Body.ToString())
	}
	/*for ix := 1; ix <= 100; ix++{
		customer_init(ix);
	}*/


}
