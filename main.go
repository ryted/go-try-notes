package main

import "fmt"



func main (){
	// this program hello word
	fmt.Println("Hello World")
	// variable
	// var 
	myName := "Gede Ages"
	// const 
	 const  iname = "gede ages"
	// print variable myName
	fmt.Println(myName)
	// print variable iName
	fmt.Println(iname)
	
	// type data
	// string
	myString := "learn programing and english  "
	// int
	myInt := 5
	// float
	myFloat := 3.14
	// boolean
	myBool := true
	// print data
	fmt.Println(myString)
	fmt.Println(myInt)
	fmt.Println(myFloat)
	fmt.Println(myBool)
	// operator
	myNu := 10 
	mybb := 5
	// +
	plus := myNu + mybb
	// -
	minus := myNu - mybb
	// *
	multi := myNu * mybb
	// /
	div := myNu / mybb
	// %
	mod := myNu % mybb
	// ^
	exp := myNu ^ mybb
	// print operator
	fmt.Println(plus)
	fmt.Println(minus)
	fmt.Println(multi)
	fmt.Println(div)
	fmt.Println(mod)
	fmt.Println(exp)
	// if
	if myBool {
		fmt.Println("true")
		}
	// if else
	if myBool{
		fmt.Println("true")
	}else {
		fmt.Println("false")
	}
	// switch
	switch myInt {
		case 5:
			fmt.Println("5")
		case 6:
			fmt.Println("6")
		default:
			fmt.Println("default")
		}
	// falltrough
	var point = 30

	switch {
	case point == 50 :
		fmt.Println("the king ")
	case (point < 40 ) && (point > 29) :
		fmt.Println("warrior human")
	default :
	fmt.Println("diel")
	}
	// for 
	for i := 0 ; i < 6 ; i++ {
		fmt.Println("number ", i)
	}
	// foreach

	// 
	for i := 1; i <= 10; i++ { 
		if i % 2 == 1 {
			continue
		}
		if i > 8{
			break
		}
		fmt.Println("angka ",i)
	}
	// array
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr)
	fmt.Println(arr[2])
	// slice
	var sli = []int{1,2,3,4,5}
	fmt.Println(sli)
	fmt.Println(sli[:2])
	//map
		var m = map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
			}
			fmt.Println(m)
			fmt.Println(m["a"])
	// map literal
	var m2 = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		}
		fmt.Println(m2)
		fmt.Println(m2["a"])

	 

}