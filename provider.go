package main 

import (
	"fmt"
	"strconv"
	"log"
	"math/rand"
	"time"
	"github.com/franela/goreq"
	)

type Provider struct {
	Rut string
	Name string
	Mail string
	Phone string
}
var provider Provider

func getRut(ix int) string {
	return strconv.Itoa(70000000+ix)
}

func getPhone() string{
	return strconv.Itoa(random(111111111,999999999))
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max -min) + min
}

func setProvider(ix int){ 
	id := strconv.Itoa(ix)
	provider.Rut 	= getRut(ix)
	provider.Name 	= "proveedor"+id
	provider.Mail 	= "mail-pro"+id+"@mail.com"
	provider.Phone	= getPhone()
}

func main() {
	for ix := 1; ix <= 45; ix++ {
		setProvider(ix)
		req := goreq.Request{
			Method: "POST",
			Uri: "http://coimco.herokuapp.com/api/providers",
			Body: provider,
		}
		req.AddHeader("Authorization","Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTk0OTIyODQsImlhdCI6MTQ5NjkwMDI4NCwibWFpbCI6ImFkbWluIn0.I_Kw0wVadDfv0RSleWg2Ernt4aKgP9v6pSq5BvtUB3Q")
		res, err := req.Do()
		checkErr(err)
		fmt.Println(res.Body.ToString())
	}

}