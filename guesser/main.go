package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {     
	fiveNumbers() 
	favouriteColour()
	manyNumbers()
	expenseCheck()
	validInput()
	guessNum()	
	addingMac()
	acmeChat()
	Anthems()      
	favPet()
}

func fiveNumbers() { 

number := askUser("Tell me a number first")
fmt.Println(number)

number1 := askUser("Tell me a number")
fmt.Println(number1)

number2 := askUser("Tell me a number")
fmt.Println(number2)

number3 := askUser("Tell me a number")
fmt.Println(number3)

number4 := askUser("Tell me a number last")
fmt.Println(number4)

n, _ := strconv.Atoi(number)   // converts string to interger
n1, _ := strconv.Atoi(number1)
n2, _ := strconv.Atoi(number2)
n3, _ := strconv.Atoi(number3)
n4, _ := strconv.Atoi(number4)

total:=n+n1+n2+n3+n4

fmt.Println(total)

}

func guessNum(){

	player1:=askUser("Enter a number")
	for guessc :=0; guessc<5; guessc++ {

		player2:=askUser("Guess number you have five attempts in total")
		p1, _ := strconv.Atoi(player1)
		p2, _ := strconv.Atoi(player2)
		
		if p2==p1{
			fmt.Println("Thats correct")
			return //used to exit the function so no extra guesses can happen
			
		} else if p2>p1{		
			fmt.Println(" Try again your number was too high")
			
		} else{
			
			fmt.Println("Try again too low")
		}
		
	}
		
		fmt.Println("You ran out of guesses it was", player1)
}
	
func acmeChat(){
	fmt.Println("Welcome we only sell one product")
	e:=askUser("What would you like help with? type help or price")
	if e=="help"{
		fmt.Println("Read the manual")
	}else if e=="price"{
		fmt.Println("Product price is £99")
	}else{
		fmt.Println("Not understood please call 0845 555 5555")

	}
	}

func addingMac(){
var input string
to:=0
for input!= "99999"{
	input=askUser("Enter a whole number (Enter 99999 to stop)")

	if input== "99999"{
		break
	}
	num, err :=strconv.Atoi(input) // num stores the int vaalue err  is returned if fails. which proceeds with error message code.
	if err != nil {  // checking if an error has occured err variable stores all errors that happen during this line of code. Nil is a non existent value for 
		fmt.Println("Invalid input please enter a number")
		continue //stops the loop and starts again from begining when user enters invalid input.
	}
	to+= num
	fmt.Println(to)

}

}


func validInput(){
	
    for{
		a:=askUser("Please enter 1,2 or 3")
		if a == "1" || a=="2" || a=="3" {

			fmt.Println("Entered correctly", a)
				break
			}else{
				fmt.Println("Enter a valid number")

		}
		
	
	}
}	

	func Anthems(){
		British:= "god save our gracious king"
		American:="oh, say can you see"
		German:= "unity and justice and freedom"
		

		anthem:=askUser("Enter the first line of a national anthem:")
		anthem= strings.ToLower(anthem) //converts case to lower case
		fmt.Println("Anthem is", anthem, ".")
		if strings.HasPrefix(British, anthem){ //function that checks if what is inputed natches the first few words of the national anthems. compares them??
			fmt.Println("British anthem")
		}else if strings.HasPrefix(American, anthem) {
			fmt.Print("American anthem")
		}else if strings.HasPrefix(German, anthem) {
			fmt.Println("German anthem")
		}else{
			fmt.Println("This code isnt built to recognise any other anthem")
		}
		}
		
	
	
	











//var a string
	//a= askUser("Enter a Number 1,2 or 3:")
	//for a!= "1" || a!= "2" || a!= "3"{

	//fmt.Println("You have entered correclty",a)

//	}


	//a:=askUser("Enter a number thats either 1,2 or 3")

	//if (a=="1") || (a=="2")||(a=="3"){
	//	fmt.Println("You Have Entered correctly which is",a)
	//}else{
	//	fmt.Println("Enter a number thats either 1,2 or 3")
	//	b:= askUser("enter a number thats incorrect")
	//	if (b =="1") || (b == "2") || (b =="3"){
	//		fmt.Println(b,"has been entered")
	//	}else{
	//		fmt.Println("Thats wrong try again")
	//	}

	//	}
//}
	





	//var a string
	//for a!="1" || a!="2"|| a!="3"{
	//	askUser("Enter a Number 1,2 or 3:")
	
	//}
	//fmt.Println("You have entered correclty",a)
		

	
	
		

	

func expenseCheck(){
	for {
        r := askUser("Do you have a receipt? (Enter yes or no)")
        if r != "yes" {
            fmt.Println("Plese resubmit with a receipt.")
            continue //stops the loop and starts again from begining when user enters invalid input.
        }
        VAT := askUser("Does the receipt have registration number, date, and amount? (Enter yes or no)")
        if VAT != "yes" {
            fmt.Println("Please resubmit with a valid VAT recept.")
            continue
        }
        submit := askUser("Have you submitted the VAT receipt? (Enter yes or no)")
        if submit != "yes" {
            fmt.Println("Please upload and resubmit the VAT receipt.")
            continue
        }
        date := askUser("Is the date on the VAT receipt used on the claim correctly? (Enter yes or no)")
        if date != "yes" {
            fmt.Println("Please correct the date on the claim or resubmit.")
            continue
        }
        amount := askUser("Is the amount on the VAT receipt used on the claim form correct? (Enter yes or no)")
        if amount != "yes" {
            fmt.Println("Please correct the amount on the claim form and resubmit.")
            continue
        }
          attend := askUser("Are the attendents field filled out with at least one name? (Enter yes or no)")
        if attend != "yes" {
            fmt.Println("Please fill out the attendees field or resubmit.")
            continue
        }
         Duplicate := askUser("Are you claiming for something that has already been provided by BJSS e.g breakfeast? (Enter yes or no)")
        if Duplicate == "yes" {
            fmt.Println("Please remove the item from your claim.")
            continue
        }
        more:= askUser("Do you have more expenses to claim? (Enter yes or no)")
        if more == "yes" {
            fmt.Println("Use Add Expense to add the expenses to the same claim form.")
        } else {
            fmt.Println("Your claim and VAT receipt has been submitted..")
            break
        }
    }
}
func manyNumbers(){
	t:=0
		for {
		manynum:=askUser("Enter a number?") 
		if manynum == "-1"{
			fmt.Println("Total:",t)
			break
		} else{
			num, err := strconv.Atoi(manynum) // num stores the int vaalue converted here. err  is returned if fails. which proceeds with error message code. if there was no error error variable would be nil(non existent)
			if err != nil{   //  checking if an error has occured err variable stores all errors that happen during this line of code. Nil is a non existent value. 
				fmt.Println("Incorrect  enter again")
				continue //stops the loop and starts again from begining when user enters invalid input.
			}
			t += num  // would be executed if  there was no error and continue
		}
		
		}
}
	

func favouriteColour() {
	colour:=askUser("Tell me yout favourite colour")
	if   colour=="yellow"{ 
	 fmt.Println("What a great colour")
	} else {
	    fmt.Println("Bad colour")
	
	
}

}


func favPet(){
	res:=askUser("Do you like meows, barks or bubbles?")


	if res == "bubbles" {
        fmt.Println("your favourite pet is a fish.")
    } else if res == "meows" {
        t:=askUser("do you like short or long tails?")
        if t == "short" {
            fmt.Println("your favourite is a Manx cat.")
        } else if t == "long" {
            c:=askUser("do you like solid colour fur or mixed colours?")
            if c == "mixed colours" {
                fmt.Println("you like tabby cats.")
            } else if c == "solid" {
                i:=askUser("My guess is black cats")
                if i == "yes" {
                    fmt.Println("The program has ended")
                } else {
                    fmt.Println("My guess is White cat")
                }
            }
        }
    } else if res == "barks" {
        g:=askUser("Do you like tiny or big dogs?")
        if g == "tiny" {
            fmt.Println("You like chihuahuas")
        } else if g == "big" {
            fmt.Println("There are lots of dogs you might like")
        }
    }
}



	


func askUser(prompt string) string {
	fmt.Println(prompt)
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	err := scanner.Err()
	if err != nil {
		panic(err)
	}

	return scanner.Text()
}




	
	
	
   
	


 
	







