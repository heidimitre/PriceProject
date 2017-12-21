package main

import (
	"fmt"

	"github.com/heidimitre/priceProject/item"
	"github.com/heidimitre/priceProject/updater"
)

func main(){

	tempURL:= ""

	price:= updater.GetCurrentPrice(tempURL)
	item:= item.CreateItem(price, "Exploding Kittens", tempURL)
	fmt.Println(item.Description, item.Price)
}
