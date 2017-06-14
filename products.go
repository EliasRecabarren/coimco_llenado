package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"github.com/franela/goreq"
	"log"
	"time"
)

type Product struct {
	Name string
	Details string
	Brand string
	Category string
}


var cat = []string{
	"Accesorios",
	"Accesorios",
	"Accesorios",
	"Accesorios",
	"Accesorios",
	"Almacenamiento",
	"Almacenamiento",
	"Almacenamiento",
	"Almacenamiento",
	"Conectividad",
	"Conectividad",
	"Conectividad",
	"Computadores",
	"Computadores",
	"Computadores",
	"Gabinetes",
	"Gabinetes",
	"Racks",
	"Servidores"}

var name = []string{
	"teclado",
	"mouse",
	"bolso",
	"funda",
	"parlantes",
	"disco-duro",
	"disco-externo",
	"memoria-flash",
	"pendrive",
	"router",
	"switch",
	"patch-panel",
	"computador",
	"tablet",
	"notebook",
	"fuente-de-poder",
	"gabinetes",
	"rack",
	"servidor"}

var brand = []string{

	"Proin",
	"iaculis",
	"ipsum",
	"etk",
	"luctus",
	"faucibus",
	"justo",
	"turpis",
	"varius",
	"felis",
	"atd",
	"gravida",
	"leo",
	"nequ",
	"eul",
	"eni",
	"Pellentesque",
	"quis",
	"tempus",
	"elit",
	"Vivamus",
	"turpis",
	"mip",
	"hendrerit"}

var product Product

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return (rand.Intn(max-min) + min)%24
}

func setProduct(ix,n int){
	id := strconv.Itoa(n)
	product.Name 	 =  name[ix]+"#"+id
	product.Details  = "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
	product.Brand	 = brand[random(1,99999)]
	product.Category = cat[ix]
}

func main() {
	
	for ix := 0; ix < 19; ix++ {
		for iy := 1; iy <= 10; iy++ {
				setProduct(ix,iy)
				req := goreq.Request{
				Method: "POST",
				Uri: "http://coimco.herokuapp.com/api/products",
				Body: product,
			}
			req.AddHeader("Authorization","Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTk0OTIyODQsImlhdCI6MTQ5NjkwMDI4NCwibWFpbCI6ImFkbWluIn0.I_Kw0wVadDfv0RSleWg2Ernt4aKgP9v6pSq5BvtUB3Q")
			res, err := req.Do()
			checkErr(err)
			fmt.Println(res.Body.ToString())
		}	
	}
}