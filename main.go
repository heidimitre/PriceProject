package main

import (
	"github.com/heidimitre/priceProject/item"
	"time"
	"os"
)

func main(){

	tempURL:= os.Getenv("TARGET")

	item.CreateItem("Exploding Kittens", tempURL)
	time.Sleep(5*time.Minute)
}
