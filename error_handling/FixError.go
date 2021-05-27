// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/pkg/errors"
	"log"
)

// ---------------------------------------------------------
// CHALLENGE
//  Add error handling to the feet to meters program.
//
// EXPECTED OUTPUT
//  go run main.go hello
//    error: 'hello' is not a number.
//
//  go run main.go what
//    error: 'what' is not a number.
//
//  go run main.go 100
//    100 feet is 30.48 meters.
// ---------------------------------------------------------

const usage = `
Feet to Meters
--------------
This program converts feet into meters.
Usage:
feet [feetsToConvert]`

func parseFloat(arg string) (float64, error) {

	feet, err := strconv.ParseFloat(arg, 64)

	if err != nil {
		msg:=fmt.Sprintf("error: '%v' is not a number.", arg)
        return 0, errors.Wrap(err, msg)
	}

	return feet, err
}

func convertToMeters(feet float64) {
	meters := feet * 0.3048
	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}


func main() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	arg := os.Args[1]	

	feet, err := parseFloat(arg)

	if err != nil {
       log.Println(err)
	   return 
	}
	
    convertToMeters(feet)	
}