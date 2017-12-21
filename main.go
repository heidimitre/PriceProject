package main

import (
	"fmt"

	"github.com/heidimitre/priceProject/item"
)

func main(){

	tempURL:= ""

	item:= item.CreateItem("Exploding Kittens", tempURL)
	item.AddPrice()
	fmt.Println(item.Description, item.Pricelist[0].Value)
}
