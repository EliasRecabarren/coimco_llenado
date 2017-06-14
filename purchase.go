package main 

import(
	"fmt"
	"github.com/franela/goreq"
	"math/rand"
	"time"
	"log"
	"strconv"
	)

type Purchase struct{
 	Id_provider string
 	Date string
 	Shiptime string
}

type PurchaseDetail struct{
	Purchase_id int
	Product_id int
	Price int
	Quantity int
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
	yy := strconv.Itoa(random(1,5))
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
func getShiptime() string {
	dd := strconv.Itoa(random(1,5))
	return "2014-01-0"+dd+"T10:00:00.000Z"
}

var purchase Purchase
var purchasedetail PurchaseDetail

func setPurchase() {
	purchase.Id_provider = strconv.Itoa(random(1,45)+70000000)
	purchase.Date 		 = getDate()
	purchase.Shiptime 	 = getShiptime()

	fmt.Println(purchase.Id_provider, purchase.Date, purchase.Shiptime)

}

func getProduct()(int,int,int){
	id := random(0,189)
	var indx,p,q int
	indx = int(id/10)
	min := price[indx]*float64(0.8)
	max := price[indx]*float64(1.2)
	p = random(int(min), int(max))
	q = random(100,150)
	return id+1,p,q

}

func setPurchaseDetail(id int){
	p_id, price, q := getProduct() 
	purchasedetail.Purchase_id = id
	purchasedetail.Product_id  = p_id
	purchasedetail.Price	   = price
	purchasedetail.Quantity    = q

	//fmt.Println(id,p_id,price,q)

}

func main() {
	var details int
	for ix := 1 ; ix <= 1000 ; ix++ {
		setPurchase()
		req := goreq.Request{
			Method: "POST",
			Uri: "https://coimco.herokuapp.com/api/purchases",
			//Uri: "https://9626204b.ngrok.io/api/purchases",
			Body: purchase,
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
					setPurchaseDetail(ix)
					req := goreq.Request{
					Method: "POST",
					Uri: "https://coimco.herokuapp.com/api/purchase_detail",
					//Uri: "https://9626204b.ngrok.io/api/purchase_detail",
					Body: purchasedetail,
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