package main

import (
	"fmt"
	"math/rand"
	"time"
	"log"
	"github.com/franela/goreq"
	"strconv"
)	

type CustomerTag struct{
	Id_tag int 
	Id_customer string
}

var customerTag CustomerTag

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func setCustomerTag(rut int){
	customerTag.Id_tag 		= 1 + random(1,1000)%10 
	customerTag.Id_customer = strconv.Itoa(rut)
}

func main() {
	for ix := 25000001; ix <= 25000200 ; ix++ {
		setCustomerTag(ix);
		req := goreq.Request{
			Method: "POST",
			Uri: "http://coimco.herokuapp.com/api/tags_customer",
			Body: customerTag,
		}
		req.AddHeader("Authorization","Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTk0OTIyODQsImlhdCI6MTQ5NjkwMDI4NCwibWFpbCI6ImFkbWluIn0.I_Kw0wVadDfv0RSleWg2Ernt4aKgP9v6pSq5BvtUB3Q")
		res, err := req.Do()
		checkErr(err)
		fmt.Println(res.Body.ToString())
	}
}
