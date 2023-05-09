package main

import (
	"errors"
) 

func add(num1 int, num2 int) error{

	result:= num1+num2
	if result == 0{
		return errors.New("Incorrect adds to zero")
	}

}


func main(){

valu
 
}





package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func askUser(prompt string) string {
	fmt.Println(prompt)
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	err := scanner.Err()
	if err != nil {
		panic(err)

		//fiveNumbers() 
		//favouriteColour()
	}

	return scanner.Text()
}

func fiveNumbers(prompt int) int{
 number := fiveNumbers("Tell me a number first")
	fmt.Println(number)

	number1 := fiveNumbers("Tell me a number")
	fmt.Println(number1)

	number2 := fiveNumbers("Tell me a number")
	fmt.Println(number2)

	number3 := fiveNumbers("Tell me a number")
	fmt.Println(number3)

	number4 := fiveNumbers("Tell me a number last")
	fmt.Println(number4)

	n, _ := strconv.Atoi(number)
	n1, _ := strconv.Atoi(number1)
	n2, _ := strconv.Atoi(number2)
	n3, _ := strconv.Atoi(number3)
	n4, _ := strconv.Atoi(number4)

	total:=n+n1+n2+n3+n4
	
	fmt.Println(total)

	  colour:=askUser("Tell me yout favourite colour")
	if   colour=="yellow"{ 
	 fmt.Println("What a great colour")
	} else {
	    fmt.Println("Bad colour")
	
}


func main() {

	fiveNumbers()
	manyNumbers()
	favouritecolour()

	



}

 
}

















//	count:=10
//	name:=1
//	for name <=count {
//		fmt.Println("Joseph")

//		name++

//	}


//	total:=1
//	for sum:= 1; sum <= 10; sum++ {
//		total*=sum
//	}
    //fmt.Println(total)
//}

	
	
		




















// func main() {
	//a:=1.0
	//b:=7.0
	//c:=9.0
	//total:= a+b+c
	//fmt.Println("Total=",total)

	//item1:=1.95
	//item2:=4.95
	//item3:=7.99
	//totalprice:= item1+item2+item3
	//.fmt.Println("Total price=",totalprice)

	//ttotal:= totalprice+total
	//fmt.Println("Total before Vat=",ttotal)

	//totalvat:=ttotal *20 /100 
	//fmt.Println("Total with vat=",totalvat+ttotal)

	//fmt.Println("VAT 20%=",totalvat)
     

 // }


