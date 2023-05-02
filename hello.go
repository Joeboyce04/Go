package main

import (
	"fmt"
) 

func main() {
	a:=1.0
	b:=7.0
	c:=9.0
	total:= a+b+c
	fmt.Println("Total",total)

	item1:=1.95
	item2:=4.95
	item3:=7.99
	totalprice:= item1+item2+item3
	fmt.Println("Total price",totalprice)

	ttotal:= totalprice+total
	fmt.Println("Total before Vat",ttotal)

	totalvat:=ttotal / 100 * 20
	fmt.Println("Total with vat",totalvat+ttotal)

	fmt.Println("VAT 20%",totalvat)
     

}
