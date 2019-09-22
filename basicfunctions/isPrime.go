package basicfunctions

func isPrime( number int) bool{
	numberIsPrime := true
	cont:= 2 
	for cont < number/2  && numberIsPrime{
		if(number%cont == 0 ){
			numberIsPrime = false
		}
		cont++
	}
	return numberIsPrime
}