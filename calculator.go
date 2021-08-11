package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {

	var s = []float64{3,3,7}

	fmt.Println(addition(s))
	fmt.Println(subtraction(s))
	fmt.Println(multiplication(s))
	fmt.Println(division(s))
	//fmt.Println(calculate("25/5/2","2*3*4*5", "2+3+4+5.9+6+7.8", "20.54-7.65-2.897" ))

}

//------------ADDITION FUNCTION------------------

func addition(n []float64)float64  {
	var sum float64 =0
	for _,v:= range n{
		sum += v

	}
	return sum
}
//-----------END OF FUNCTION----------------------



//---------------SUBTRACTION FUNCTION-------------

func subtraction (sub []float64)float64 {
	var s float64= sub[0]*2
	for _,v:= range sub{
		s -= v
	}
	return s
}
//-----------END OF FUNCTION----------------------



//-----------MULTIPLICATION FUNCTION---------------

func multiplication(m []float64)float64 {
	var diff float64=1
	for i:=0; i<len(m); i++{
		diff *= m[i]
	}
	return diff
}
//-----------END OF FUNCTION----------------------



//--------------DIVISION FUNCTION-----------------

func division (d []float64)float64 {
	var div float64= d[0]*d[0]
	for _,v:=range d{
		div = div/v
	}
	return div
}
//-----------END OF FUNCTION----------------------


//--------------CALCULATE FUNCTION-----------------

func calculate(str ...string) []float64 {
	var final []float64
	var converted []float64
	for _,v:=range str{
		if strings.ContainsAny(v, "/"){ // Checks if the string has a "/" character in it
			st:= fmt.Sprint(strings.ReplaceAll(v, "/"," "))
			arr := strings.Split(st, " ")
			for _,j:=range arr {
				if num, err := strconv.ParseFloat(j, 64); err == nil { // Here converts the strings into float
					converted = append(converted, num)
				}
			}
			final= append(final, division(converted)) // Calls the division function and appends the result into a slice
		}
//-------------------------------------------------------------------------------------------------------------------------

		if strings.ContainsAny(v, "*") { // Checks if the string has a "*" character in it
			var covet []float64
			st := fmt.Sprint(strings.ReplaceAll(v, "*", " "))
			arr := strings.Split(st, " ")
			for _, j := range arr {
				if num, err := strconv.ParseFloat(j, 64); err == nil { // Here converts the strings into float

					covet = append(covet, num)
				}
			}
			final = append(final, multiplication(covet)) // Calls the multiplication function and appends the result into a slice
		}
//-------------------------------------------------------------------------------------------------------------------------

		if strings.ContainsAny(v, "+"){ // Checks if the string has a "+" character in it
		var change []float64
			st:= fmt.Sprint(strings.ReplaceAll(v, "+"," "))
			arr := strings.Split(st, " ")
			for _,j:=range arr {
				if num, err := strconv.ParseFloat(j, 64); err == nil { // Here converts the strings into float
					change = append(change, num)
				}
			}
			final= append(final, addition(change)) // Calls the addition function and appends the result into a slice
		}
//-------------------------------------------------------------------------------------------------------------------------

		if strings.ContainsAny(v, "-"){ // Checks if the string has a "-" character in it
			var shed []float64
			st:= fmt.Sprint(strings.ReplaceAll(v, "-"," "))
			arr := strings.Split(st, " ")
			for _,j:=range arr {
				if num, err := strconv.ParseFloat(j, 64); err == nil { // Here converts the strings into float
					shed = append(shed, num)
				}
			}
			final= append(final, subtraction(shed)) // Calls the subtraction function and appends the result into a slice
		}
	}

	return final
}
//-----------END OF FUNCTION----------------------
