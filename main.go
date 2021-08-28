package main

import (
	"os"
	"url_shortner/logic"
	"url_shortner/server"
	"fmt"
)

func main() {

	longURL:=os.Args[len(os.Args)-2]
	x:= os.Args[len(os.Args)-1]

	switch x {

	case "1":logic.API(longURL)
	case "2":server.Run()
	default:
				fmt.Println("Error")

	}



}