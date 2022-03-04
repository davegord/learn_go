package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
		lat := "50.2289 N"
		long := "99.4656 W"
		e := errors.New("Dummy, you can't take a sqrt of a negative number")
		return 0, sqrtError{lat, long, e}
	}
	return 42, nil
}
