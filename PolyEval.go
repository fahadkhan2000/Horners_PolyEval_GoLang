package main

import "fmt"
import s "strings"
import "math"
import "strconv"

var print = fmt.Println

func main() {

    fmt.Println("Polynomial Evaluation")

    pol := "50x^5 - 40x^4 + 30x^3 -20x^2 +10x+1000.1"
    val := 1.0

    printInputData(pol, val)

    monomialsArray := convertPolyToStandardFormat(pol)

    //evaluateMonomial("-x" , 2)

    print("..............." , monomialsArray[0])

    calculateFinalResult(monomialsArray , val)
}

func printInputData(pol string , val float64) {

	fmt.Println("Polynomial String = ", pol)

	fmt.Println("Float type value = ", val)
}

func convertPolyToStandardFormat(pol string) []string {

	polWithoutSpaces := removeAllSpacesFromPoly(pol)
	print("Polynomial without spaces = " , polWithoutSpaces)

	polWithPlusSigns := addPlusSignstoPoly(polWithoutSpaces)
	print("Polynomial with added plus signs = " , polWithPlusSigns)

	monomialsArray := createTermsBySplitingPoly(polWithPlusSigns)
	print("Array of monomials = " , monomialsArray)


	return monomialsArray
}

func removeAllSpacesFromPoly(pol string) string {

	polWithoutSpaces := s.Replace(pol , " ", "", -1)

	return polWithoutSpaces
}

func addPlusSignstoPoly(polWithoutSpaces string) string {

    polWithPlusSigns := s.Replace(polWithoutSpaces, "-", "+-", -1)

	return polWithPlusSigns
}

func createTermsBySplitingPoly(polWithPlusSigns string) []string {

	monomialsArray := s.Split(polWithPlusSigns , "+")

	if monomialsArray[0] == "" {

		monomialsArray = append(monomialsArray[:0], monomialsArray[1:]...)
	}

	return monomialsArray
}

/*
Evaluation of a single monomial 
*/

func evaluateMonomial(singleTerm string , val float64) float64 {

	splittedCoeffAndPower := determineTypeOfMonomialForSplitting(singleTerm)

	floatArrayOfMonomials := convertMonomialFromStringToDouble(splittedCoeffAndPower)

	monoRes := calculateMonomial(floatArrayOfMonomials , val)

	return monoRes

}

func determineTypeOfMonomialForSplitting(singleTerm string) []string {

	splittedCoeffAndPower := []string {"0", "0"}

if s.Contains(singleTerm, "^") == true {

	print("____________________________term contains ^ ")

	if s.HasPrefix(singleTerm, "x") == true || s.HasPrefix(singleTerm, "-x") == true || s.HasPrefix(singleTerm, "+x") == true {

		print("____________________________term hasPrefix +x or -x ")

		singleTerm = s.Replace(singleTerm, "x", "1x", -1)
		splittedCoeffAndPower = splitTermIntoCoeffAndPower(singleTerm)
		print("string type array of monomial's coeff and power" , splittedCoeffAndPower)

		} else {

		print("____________________________term hasPrefix +x or -x ")

		splittedCoeffAndPower = splitTermIntoCoeffAndPower(singleTerm)
        print("string type array of monomial's coeff and power" , splittedCoeffAndPower)


		}

} else if s.Contains(singleTerm, "^") == false {

		if s.Contains(singleTerm , "x") == true {

		print("____________________________term is x or -x ")

		singleTerm = s.Replace(singleTerm, "x", "x^1", -1)
		splittedCoeffAndPower = splitTermIntoCoeffAndPower(singleTerm)
		print("string type array of monomial's coeff and power" , splittedCoeffAndPower)

	} else {

		print("____________________________term is constant ")

		var appender string = "x^0"

		singleTerm = s.Join([]string{singleTerm, appender} , "")
		splittedCoeffAndPower = splitTermIntoCoeffAndPower(singleTerm)
		print("string type array of monomial's coeff and power" , splittedCoeffAndPower)
	}

} 
return splittedCoeffAndPower

}


func splitTermIntoCoeffAndPower(singleTerm string) []string {

	var splittedCoeffAndPower = []string{}

	if singleTerm == "x^1" || singleTerm == "-x^1" {

		singleTerm = s.Replace(singleTerm, "x", "1", -1)

		splittedCoeffAndPower = s.Split(singleTerm, "^")

	} else{

			splittedCoeffAndPower = s.Split(singleTerm, "^")
	        tempString := s.Join(splittedCoeffAndPower, "")
	        splittedCoeffAndPower = splittedCoeffAndPower[:0]
	        splittedCoeffAndPower = s.Split(tempString, "x")

	}

	return splittedCoeffAndPower
}

func convertMonomialFromStringToDouble(splittedCoeffAndPower []string) [] float64 {

	var floatArrayOfMonomials = []float64{}

	for _,i := range splittedCoeffAndPower {

		flt, err := strconv.ParseFloat(i , 64)

		if err != nil {
			panic(err)
		}

		floatArrayOfMonomials = append(floatArrayOfMonomials, flt)
	}
	print("float type array of monomial coeff and power = " , floatArrayOfMonomials)
	return floatArrayOfMonomials
}

func calculateMonomial(floatArrayOfMonomials []float64 , val float64) float64 {

	 var monoRes float64 = 0.0

	for i := 0; i < len(floatArrayOfMonomials); i++ {

		monoRes = floatArrayOfMonomials[0] * (math.Pow(val , floatArrayOfMonomials[1]))

		print("value of monomial term = " , monoRes)
	}
	return monoRes
}

func calculateFinalResult(monomialsArray []string, val float64) {

	var finalRes float64 = 0.0

	for i := 0; i < len(monomialsArray); i++ {

		finalRes = finalRes + evaluateMonomial(monomialsArray[i] , val) 
	}

	print("final result = ", finalRes)
}





