package main

//go get github.com/franela/goreq

//import "github.com/franela/goreq"
import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
	"github.com/franela/goreq"
	"log"
	)

type Customer struct {
	Rut string
	Name string
	Mail string
	Phone string
}

var customer Customer

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func getRut(ix int) string {
	return strconv.Itoa(25000000+ix)
}

func getPhone() string{
	return strconv.Itoa(random(111111111,999999999))
}

func setCustomer(ix int){ 
	id := strconv.Itoa(ix)
	customer.Rut 	= getRut(ix)
	customer.Name 	= "cliente"+id
	customer.Mail 	= "mail-cli"+id+"@mail.com"
	customer.Phone	= getPhone()
}

func main() {
	for ix := 1; ix <= 200; ix++ {
		setCustomer(ix)
		req := goreq.Request{
			Method: "POST",
			Uri: "http://coimco.herokuapp.com/api/customers",
			Body: customer,
		}
		req.AddHeader("Authorization","Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTk0OTIyODQsImlhdCI6MTQ5NjkwMDI4NCwibWFpbCI6ImFkbWluIn0.I_Kw0wVadDfv0RSleWg2Ernt4aKgP9v6pSq5BvtUB3Q")
		res, err := req.Do()
		checkErr(err)
		fmt.Println(res.Body.ToString())
	}
}
