package main

import (
	"github.com/go-programming-tour-book/tour/internal/sql2struct"
	"log"

	"github.com/go-programming-tour-book/tour/cmd"
)


func main(){
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}

	sql2.TplParse()
}