package main

import (
	"fmt"
	"log"
	"github.com/franela/goreq"
	)
var tagname = []string{
	"Educación",
	"Salud",
	"ONG",
	"Salud",
	"Retail",
	"Organizacion",
	"Mantenimiento",
	"Mineria",
	"Construcción",
	"Almacenamiento y Comunicaciones"}

type Tag struct{
	Name string
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

var tag Tag

func main() {
	for ix := 0; ix < 10; ix++ {
		tag.Name = tagname[ix]
		req := goreq.Request{
			Method: "POST",
			Uri:    "http://coimco.herokuapp.com/api/tags",
			Body:   tag,
		}
		req.AddHeader("Authorization","Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTk0OTIyODQsImlhdCI6MTQ5NjkwMDI4NCwibWFpbCI6ImFkbWluIn0.I_Kw0wVadDfv0RSleWg2Ernt4aKgP9v6pSq5BvtUB3Q")

		res, err := req.Do()
		//fmt.Println(res.Body.ToString())
		checkErr(err)
		fmt.Println(tag)
		fmt.Println(res.Body.ToString())
	}	
}

