package main

import(
	"fmt"
	"github.com/franela/goreq"
	"log"
	"strconv"
)

type User struct{
	Mail string
	Name string
	LastName string
	Rut string
	Pass string
	Role int8
	Active bool
}

var admin User
var manager User
var seller User

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func setAdmin(){
	admin.Role 		= 2
	admin.Mail		= "admin@mail.com"
	admin.Name 		= "admin"
	admin.LastName 	= "coimco"
	admin.Pass 		= "admin"
	admin.Rut  		= "12000000"
	admin.Active 	= true
}

func setManager(){
	manager.Role 		= 1
	manager.Mail 		= "manager@mail.com"
	manager.Name 		= "manager"
	manager.LastName 	= "coimco"
	manager.Pass 		= "manager"
	manager.Rut  		= "11000000"
	manager.Active 	= true
}

func setSeller(ix int){
	id := strconv.Itoa(ix)
	seller.Role 		= 0
	seller.Mail		    = "seller"+id+"@mail.com"
	seller.Name 		= "seller"+id
	seller.LastName 	= "coimco"
	seller.Pass 		= "seller"
	seller.Rut  		= "10000000"	
	seller.Active 	    = true
}
func main(){
	setAdmin()
	setManager()
	var user User
	for ix := 1; ix <= 12; ix++ {
		if ix==11 {
			user = admin
		}else if ix ==12 {
			user = manager
		}else{
			setSeller(ix)
			user = seller
		}
		req := goreq.Request{
			Method: "POST",
			Uri: "http://coimco.herokuapp.com/api/accounts",
			Body: user,
		}
		req.AddHeader("Authorization","Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTk0OTIyODQsImlhdCI6MTQ5NjkwMDI4NCwibWFpbCI6ImFkbWluIn0.I_Kw0wVadDfv0RSleWg2Ernt4aKgP9v6pSq5BvtUB3Q")
		res, err := req.Do()
		checkErr(err)
		fmt.Println(res.Body.ToString())
	}

}