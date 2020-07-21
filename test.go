package main

import (
	"fmt"

	"github.com/lithammer/shortuuid"
)

func mainB() {
	ID := shortuuid.New()
	fmt.Println(ID)
}
