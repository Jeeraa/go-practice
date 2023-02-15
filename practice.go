package main

import (
	"fmt"
	"golang/calculator" //Package
)

//Struct
type Product struct {
	name     string
	price    float64
	category string
	discount int
}

func main() {

	//Define Variable
	name := "Manasicha" //Inference Type
	age := 20
	var score float32 = 95 //Manual Type
	var isPass bool = true
	fmt.Println("My name is", name) //fmt.Println = Have Enter
	fmt.Println("Age =", age)
	fmt.Println("Score", score)
	fmt.Println("Status", isPass)

	//Define Constance
	const FullScore int = 100
	fmt.Println("Full Score =", FullScore)

	//Print Data By Type
	fmt.Printf("Name : %v\n", name) // v = value
	fmt.Printf("Age : %v\n", age)

	//Print Data Type
	fmt.Printf("Name's Type : %T\n", name) //fmt,Printf = No have Enter
	fmt.Printf("Age's Type : %T\n", age)

	//Math Operation
	num1, num2 := 20, 3
	fmt.Println("Sum =", num1+num2)
	fmt.Println("Diff =", num1-num2)
	fmt.Println("Multiple =", num1*num2)
	fmt.Println("Divider =", num1/num2)
	fmt.Println("Mod =", num1%num2)

	//Comparison Operation
	num3, num4 := 20, 3
	fmt.Println("Equal?", num3 == num4)
	fmt.Println("Not Equal?", num3 != num4)
	fmt.Println("num3 more num4?", num3 > num4)
	fmt.Println("num3 less num4?", num3 < num4)
	fmt.Println("num3 more and equal num4?", num3 >= num4)
	fmt.Println("num3 less and equal num4?", num3 <= num4)

	//Get Input
	var scoreExam int
	fmt.Print("Your SCORE : ") //fmt.Print = No have Enter
	fmt.Scanf("%d", &scoreExam)
	fmt.Println("TOTAL SCORE :", scoreExam)

	//IF...ELSE Condition
	if scoreExam >= 50 {
		fmt.Println("Status : PASS")
	} else {
		fmt.Println("Status : NOT PASS")
	}

	//SWITCH...CASE Condition
	switch scoreExam {
	case 50:
		fmt.Println("Status : PASS")
	case 49:
		fmt.Println("Status : NOT PASS")
	default:
		fmt.Println("Status : NO DATA")
	}

	//Array
	var arr1 [3]int
	fmt.Println(arr1) //[0,0,0]
	var data [3]int
	data[0] = 11
	data[1] = 12
	data[2] = 13
	fmt.Println(data)
	var arr3 [3]int = [3]int{11, 12, 13} //Type-1
	fmt.Println(arr3)
	day := [4]string{"MON", "TUE", "FRI", "SAT"} //Type-2
	fmt.Println(day)
	fmt.Println(day[3])
	fmt.Println(len(day))                                   //Size of array
	month := [...]string{"JAN", "FEB", "AUG", "DEC", "OCT"} //unlimit size
	fmt.Println(len(month))

	//Slice
	day2 := []string{"MON", "TUE", "WED", "THU", "FRI"}
	fmt.Println(day2)      //[MON TUE WED THU FRI]
	fmt.Println(day2[3])   //THU
	fmt.Println(len(day2)) //5
	day2[0] = "MONDAY"
	fmt.Println(day2) //[MONDAY TUE WED THU FRI]
	day2 = append(day2, "SAT")
	fmt.Println(day2)      //[MONDAY TUE WED THU FRI SAT]
	fmt.Println(day2[:])   //[MONDAY TUE WED THU FRI SAT] all index
	fmt.Println(day2[1:])  //[TUE WED THU FRI SAT] index 1->last
	fmt.Println(day2[1:3]) //[TUE WED] index 1->2

	//Map (setting index using keyword)
	book := map[string]string{"SCI": "Plants", "MATH": "Vectors"}
	book["ENG"] = "Grammar"
	fmt.Println(book)        //map[ENG:Grammar MATH:Vectors SCI:Plants]
	fmt.Println(book["SCI"]) //Plants
	fmt.Println(len(book))   //3

	valuee, isKey := book["MATH"]
	if isKey {
		fmt.Println(valuee)
	} else {
		fmt.Println("No DATA")
	}

	//LOOP
	for count := 1; count <= 10; count++ {
		if count == 5 {
			break
		}
		fmt.Println("Hello, Golang <", count, ">")
	}
	fmt.Println("END")

	for count := 1; count <= 10; count++ {
		if count == 5 {
			continue //skip count=5
		}
		fmt.Println("Hello, Golang <", count, ">")
	}
	fmt.Println("END")

	//LOOP with Slice
	numm := []int{10, 20, 30, 40, 50, 60}
	for index := 0; index < len(numm); index++ {
		fmt.Println("No. <", index, ">", numm[index])
	}

	for index, vallue := range numm {
		fmt.Println("No. <", index, ">", vallue)
	}

	for _, vallue := range numm { // _ is ignore index
		fmt.Println(vallue)
	}

	//LOOP with Map
	bookk := map[string]string{"SCI": "Plants", "MATH": "Vectors"}
	for keys, valuue := range bookk {
		fmt.Println("Subject ", keys, ", Name :", valuue)
	}

	//Function
	showWord()
	showWordAndPersonal("Python")
	totalCost(90, 50)
	deli := getDelivery()
	fmt.Println("Delivery", deli, "Baht")
	myCart := getTotalCart(8, 4)
	fmt.Println("Total Cart", myCart, "Piece")
	myCart2, check := getTotalCart2(2, 7)
	fmt.Println("Total Cart", myCart2, "Piece", "status :", check)
	myCart3 := getTotalCart3(10, 10, 9, 10)
	fmt.Println("Total Cart", myCart3, "Piece")

	//Struct
	product1 := Product{name: "IPhone", price: 30000, category: "Phone", discount: 2500}
	fmt.Println(product1)
	fmt.Println(product1.name)
	product1.discount = 1000
	fmt.Println(product1)
	fmt.Println(product1.discount)

	//Package
	fmt.Println(calculator.Add(100, 50))
}

//Function
func showWord() {
	fmt.Println("Golang,Welcome")
}

func showWordAndPersonal(name string) {
	fmt.Println("Hi", name)
}

func totalCost(cost1, cost2 int) {
	fmt.Println("Total Cost =", cost1+cost2, "Baht")
}

func getDelivery() int {
	return 50
}

func getTotalCart(cost1, cost2 int) int {
	total := cost1 + cost2
	return total
}

func getTotalCart2(cost1, cost2 int) (int, string) {
	total := cost1 + cost2
	status := ""
	if total%2 == 0 {
		status = "EVEN"
	} else {
		status = "ODD"
	}
	return total, status
}

func getTotalCart3(number ...int) int {
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}
