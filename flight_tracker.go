package main

import "fmt"
import "time"
import "math/rand"

var totalcalls int

func fetchprice(Airline string){
	totalcalls++
	delay:=time.Duration(rand.Intn(300))*time.Millisecond
	time.Sleep(delay)
	price :=(rand.Intn(5000)+8000)
	fmt.Printf("%s->%d time:%v\n",Airline,price,delay)
}
func main() {
   Airlines:=[]string{"Indigo","Air India","Vistara","SpiceJet","Akasa Air"}
   start:=time.Now()
   for _, Airline:= range Airlines{
	fetchprice(Airline)
   }
   fmt.Printf("Time Taken:%v\n",time.Since(start))
   fmt.Printf("Calls Taken:%d\n",totalcalls)
}