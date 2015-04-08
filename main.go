// Copyright 2015 elliott@tkwcafe.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"

	"go-password/password"
)

var length, cnt int
var incSpecials bool

func init() {
	flag.IntVar(&length, "c", password.CHAR_MIN, "     character count per password")
	flag.IntVar(&cnt, "n", password.OUT_MIN, "      total passwords to generate")
	flag.BoolVar(&incSpecials, "s", false, "  include special characters (e.g. !@#$%^&*)")
}

func main() {
	flag.Parse()

	for i := 0; i < cnt; i++ {
		pwd, err := password.New(length, incSpecials)
		if err != nil {
			fmt.Errorf("unable to read from random source: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(pwd)
	}
}
