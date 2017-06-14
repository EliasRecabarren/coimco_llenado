package main 

import(
	"fmt"
	"github.com/franela/goreq"
	"math/rand"
	"time"
	"log"
	"strconv"
	)

type Sale struct{
 	Id_user string
 	Id_customer string
 	Date string
}

type SaleDetail struct{
	Sale_id int
	Product_id int
	Price int
	Quantity int
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func getDate() string {
	yy := strconv.Itoa(random(1,6))
	mm := random(1,12)
	dd := random(1,28)
	var d string
	var m string
	if mm < 10 {
		//m := strconv.Itoa(mm)
		m = "0"+strconv.Itoa(mm)
	}else {
		m = strconv.Itoa(mm)
	}
	if dd < 10 {
		//m := strconv.Itoa(mm)
		d = "0"+strconv.Itoa(dd)
	}else {
		d = strconv.Itoa(dd)
	}
	return "201"+yy+"-"+m+"-"+d+"T10:00:00.000Z"

}
var price = []float64{
	15000,
	15000,	
	20000,
	13000,
	24000,
	40000,
	45000,
	7000,	
	8000,
	30000,
	28000,
	45000,		
	200000,
	80000,
	250000,
	30000, 	
	30000,
	100000,
	350000}

var sale Sale
var saledetail SaleDetail

func setSale() {
	id := strconv.Itoa(random(1,10))
	sale.Id_user 	 = "seller"+id+"@mail.com"
	sale.Id_customer = strconv.Itoa(random(1,200)+25000000)
	sale.Date 		 = getDate()

}
func getProduct()(int,int,int){
	id := random(0,189)
	var indx,p,q int
	indx = int(id/10)
	min := price[indx]*float64(1.15)
	max := price[indx]*float64(1.33)
	p = random(int(min), int(max))
	q = random(80,140)
	return id+1,p,q

}
func setSaleDetail(id int){
	p_id, price, q:= getProduct() 
	saledetail.Sale_id = id
	saledetail.Product_id  = p_id
	saledetail.Price	   = price
	saledetail.Quantity    = q

}


func main() {
	var details int
	for ix := 1 ; ix <= 1000 ; ix++ {
		setSale()
		req := goreq.Request{
			Method: "POST",
			Uri: "https://coimco.herokuapp.com/api/sales",
			//Uri: "https://9626204b.ngrok.io/api/purchases",
			Body: sale,
		}
		req.AddHeader("Authorization","Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTk0OTIyODQsImlhdCI6MTQ5NjkwMDI4NCwibWFpbCI6ImFkbWluIn0.I_Kw0wVadDfv0RSleWg2Ernt4aKgP9v6pSq5BvtUB3Q")
		res,err := req.Do()
		checkErr(err)
		if res == nil {
			fmt.Println("ADIOS")}
			fmt.Println(err)

		if err == nil {
				fmt.Println(res.Body.ToString())
				details = random(1,5)

				for iy := 1; iy <= details ; iy++ {
					setSaleDetail(ix)
					req := goreq.Request{
					Method: "POST",
					Uri: "https://coimco.herokuapp.com/api/sale_detail",
					//Uri: "https://9626204b.ngrok.io/api/purchase_detail",
					Body: saledetail,
					}
					req.AddHeader("Authorization","Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTk0OTIyODQsImlhdCI6MTQ5NjkwMDI4NCwibWFpbCI6ImFkbWluIn0.I_Kw0wVadDfv0RSleWg2Ernt4aKgP9v6pSq5BvtUB3Q")
					res_,err_ := req.Do()
					
					checkErr(err_)
					if res_ != nil {
						fmt.Println(res_.Body.ToString())
					}else{
						iy--
					}
				}
			
		}else{
			ix--
			fmt.Println("error ")
		}	
	
	}
	
}