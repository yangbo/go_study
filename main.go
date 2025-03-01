package main

import (
	"fmt"
	"rsc.io/quote"
)

func main(){
	fmt.Println("你好 Go!");
	fmt.Print(quote.Opt());
}
