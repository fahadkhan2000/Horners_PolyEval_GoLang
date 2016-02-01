package main

import "fmt"
import s "strings"
import "math"
import "strconv"

var print = fmt.Println

func convertPolyToStandardFormat(pol string) []string {
	polWithoutSpaces := s.Replace(pol , " ", "", -1)
	polWithPlusSigns := s.Replace(polWithoutSpaces, "-", "+-", -1)
	monomialsArray := createTermsBySplitingPoly(polWithPlusSigns)
	
	print("Array of monomials = " , monomialsArray)
	return monomialsArray
}

func createTermsBySplitingPoly(polWithPlusSigns string) []string {
	monomialsArray := s.Split(polWithPlusSigns , "+")

	if monomialsArray[0] == "" {
		monomialsArray = append(monomialsArray[:0], monomialsArray[1:]...)
	}
	return monomialsArray
}

func evaluateMonomial(singleTerm string , val float64) float64 {
	splittedCoeffAndPower := determineTypeOfMonomialForSplitting(singleTerm)
	coeff, exp := convertMonomialFromStringToDouble(splittedCoeffAndPower)
	return coeff * (math.Pow(val , exp))
}

func determineTypeOfMonomialForSplitting(singleTerm string) []string {
	splittedCoeffAndPower := []string {"0", "0"}

	if s.Contains(singleTerm, "^") == true {
		if s.HasPrefix(singleTerm, "x") == true || s.HasPrefix(singleTerm, "-x") == true || s.HasPrefix(singleTerm, "+x") == true {
			singleTerm = s.Replace(singleTerm, "x", "1x", -1)
			splittedCoeffAndPower = splitTermIntoCoeffAndPower(singleTerm)
			} else {
			splittedCoeffAndPower = splitTermIntoCoeffAndPower(singleTerm)
			}
	} else if s.Contains(singleTerm, "^") == false {
			if s.Contains(singleTerm , "x") == true {
			singleTerm = s.Replace(singleTerm, "x", "x^1", -1)
			splittedCoeffAndPower = splitTermIntoCoeffAndPower(singleTerm)
		} else {
			var appender string = "x^0"
			singleTerm = s.Join([]string{singleTerm, appender} , "")
			splittedCoeffAndPower = splitTermIntoCoeffAndPower(singleTerm)
		}
	} 
	return splittedCoeffAndPower
}

func splitTermIntoCoeffAndPower(singleTerm string) []string {
	var splittedCoeffAndPower = []string{}

	if singleTerm == "x^1" || singleTerm == "-x^1" {
		singleTerm = s.Replace(singleTerm, "x", "1", -1)
		splittedCoeffAndPower = s.Split(singleTerm, "^")
	} else {
			splittedCoeffAndPower = s.Split(singleTerm, "^")
	        tempString := s.Join(splittedCoeffAndPower, "")
	        splittedCoeffAndPower = splittedCoeffAndPower[:0]
	        splittedCoeffAndPower = s.Split(tempString, "x")
	}
	return splittedCoeffAndPower
}

func convertMonomialFromStringToDouble(splittedCoeffAndPower []string) (coeff, exp float64) {
	var coeffAndExpArray = []float64{}
	for _,i := range splittedCoeffAndPower {
		flt, err := strconv.ParseFloat(i , 64)
		if err != nil {
			panic(err)
		}
		coeffAndExpArray = append(coeffAndExpArray, flt)
	}
	coeff = coeffAndExpArray[0]
	exp = coeffAndExpArray[1]
	return coeff, exp
}

func calculateFinalResult(monomialsArray []string, val float64) {
	var finalRes float64 = 0.0

	for i := 0; i < len(monomialsArray); i++ {
		finalRes = finalRes + evaluateMonomial(monomialsArray[i] , val) 
	}
	print("final result = ", finalRes)
}

func main() {
    pol := "50x^5 - 40x^4 + 30x^3 -20x^2 +10x+81.1"
    val := 1.0

	fmt.Println("Polynomial String = ", pol)
	fmt.Println("Float type value = ", val)
    monomialsArray := convertPolyToStandardFormat(pol)
    calculateFinalResult(monomialsArray , val)
}