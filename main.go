package main

import (
	"flag"
	"fmt"

	"github.com/coderconvoy/passmaker/pswd"
)

func main() {
	p := flag.String("p", "", "password")
	c := flag.String("c", "", "check against")
	flag.Parse()

	if *p == "" {
		fmt.Println("No Password")
		return
	}

	if *c == "" {
		//Make Password

		pass, err := pswd.New(*p)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf(pass.String())
		return
	}

	//Check Password

	pass, err := pswd.Parse(*c)
	if err != nil {
		fmt.Println(err)
		return
	}

	if pass.Check(*p) {
		fmt.Println("Password OK")
		return
	}
	fmt.Println("Password Fail ,, np == ", pass)

}
