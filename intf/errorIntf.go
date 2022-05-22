package intf

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("f < 0 is error")
	}
	return 11.0, nil
}

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	str := ""
	return fmt.Sprintln(str, de.dividee)
}

func Divide(varDividee int,varDivider int) (result int,errMsg string){
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider , ""
	}
}

func Invoke() {
	

	if result,errorMsg := Divide(10,5); errorMsg ==  "" {
		fmt.Println("10/5 = ",result)
	}

	if _, errorMsg := Divide(100, 0 );errorMsg != "" {
		fmt.Println("errorMsg is:", errorMsg)
	}

}
