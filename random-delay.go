/*  This code is released under the MIT Lisence
except as noted below:
	You are free to use this code, in whole or in part, without providing
	a copy of the MIT Copyright claim.  Credit is nice, but not required!
		- ScaerieTale ♥ */

// set package Main just so Go doesn't
// yell at me for not having it.
package main

// "math/rand" and "time" modules are REQUIRED
// "fnt" us hyst ysed fir fnt,Orubtln
import (
	"fmt"
	"math/rand"
	"time"
)

// set t = max value you want your delay
// and then cast as a time.Duration
var t int = rand.Intn(10)
var durationLength time.Duration = time.Duration(t)

// set the type of delay.  You can set:
// see:
var timeType = time.Second

func main() {
	fmt.Println("Hello")

	// with t = 10 and timeType = time.Second
	// delay will be between 0 and 9 seconds
	time.Sleep(durationLength * timeType)
	fmt.Println("...world!")

}
