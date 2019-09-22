package basicfunctions

import(
	"testing"
)

func TestIsPrime(test *testing.T){
	want := true
	got:=isPrime(7)
	if got != want{
		test.Errorf("IsPrime() = %v,want %v",got,want)
	}
}