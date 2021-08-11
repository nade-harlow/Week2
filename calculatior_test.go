package main

import (
	"reflect"
	"testing"
)
func TestAddition(t *testing.T) {
	values := []struct {
	Numbers []float64
	Expected float64
	message string
	}{
		{
			Numbers: []float64{2,4,2},
			Expected: 8,
			message: "Correct",
		},
		{
			Numbers: []float64{5,3,1},
			Expected: 9,
			message: "Correct",
		},
		{
			Numbers: []float64{11,2,7,2},
			Expected: 22,
			message: "Wrong",
		},
		{
			Numbers: []float64{43,67,1},
			Expected: 111,
			message: "Wrong",
		},
	}

	for i, v := range values {
		f := addition(v.Numbers)
		if f != v.Expected {
			t.Fatalf("Index: %v %v\tMessage: %s\tExpecting: %v\n", i, v.Numbers, v.message, f)
		}
	}
}


func TestSubtraction(t *testing.T) {
	values := []struct {
		Numbers []float64
		Expected float64
		message string
	}{
		{
			Numbers: []float64{2,4,2},
			Expected: -4,
			message: "Correct",
		},
		{
			Numbers: []float64{5,3,1},
			Expected: 1,
			message: "Correct",
		},
		{
			Numbers: []float64{11,2,7,2},
			Expected: 0,
			message: "Correct",
		},
		{
			Numbers: []float64{67,43,1},
			Expected: 23,
			message: "Correct",
		},
	}

	for i, v := range values {
		f := subtraction(v.Numbers)
		if f != v.Expected {
			t.Fatalf("Index: %v %v\tMessage: %s\tExpecting: %v\n", i, v.Numbers, v.message, f)
		}
	}
}

func TestMultiplication(t *testing.T) {
	values := []struct {
		Numbers []float64
		Expected float64
		message string
	}{
		{
			Numbers: []float64{2,4,2},
			Expected: 16,
			message: "Correct",
		},
		{
			Numbers: []float64{2,2,4},
			Expected: 16,
			message: "Correct",
		},
		{
			Numbers: []float64{11,2,7,2},
			Expected: 308,
			message: "Wrong",
		},
		{
			Numbers: []float64{43,67,1},
			Expected: 2881,
			message: "Correct",
		},
	}

	for i, v := range values {
		f := multiplication(v.Numbers)
		if f != v.Expected {
			t.Fatalf("Index: %v %v\tMessage: %s  Expecting: %v\n", i, v.Numbers, v.message, f)
		}
	}
}

func TestDivision(t *testing.T) {
	values := []struct {
		Numbers []float64
		Expected float64
		message string
	}{
		{
			Numbers: []float64{21,7,3},
			Expected: 1,
			message: "Correct",
		},
		{
			Numbers: []float64{6,3,1},
			Expected: 2,
			message: "Correct",
		},
		{
			Numbers: []float64{12,2,2,2},
			Expected: 1.5,
			message: "Correct",
		},
		{
			Numbers: []float64{100,2,5,5},
			Expected: 2,
			message: "Correct",
		},
	}

	for i, v := range values {
		f := division(v.Numbers)
		if f != v.Expected {
			t.Fatalf("Index: %v %v\tMessage: %s\tExpecting: %v\n", i, v.Numbers, v.message, f)
		}
	}
}

func TestCalculate(t *testing.T)  {
	values := []struct {
		Strings []string
		Expected []float64
		message string
	}{
		{
			Strings: []string{"2/2","2+2","6-2","4*2"},
			Expected: []float64{1,4,4,8},
		},
	}
	for i, v := range values {
		f := calculate(v.Strings...)
		if !reflect.DeepEqual(f,v.Expected){
			t.Fatalf("Index: %v %v\tMessage: %s\tExpecting: %v\n", i, v.Strings, v.message, f)
		}
	}
}



//func TestAddition(t *testing.T) {
//addi := addition([]float64{2, 2, 0})
//if addi != 3 { t.Errorf("inaccurate result")}
//
//}
