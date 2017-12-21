package item

import (
	"github.com/heidimitre/priceProject/price"
	"github.com/heidimitre/priceProject/updater"
)

type Item struct {
	Pricelist   []*price.Price
	Description string
	Url			string
}

func CreateItem (desc string, url string)(item *Item){
	newItem := new(Item)
	newItem.Pricelist = []*price.Price{}
	newItem.Description = desc
	newItem.Url = url
	return newItem
}

func (item *Item) AddPrice (){
	item.Pricelist = append(item.Pricelist, updater.GetCurrentPrice(item.Url))
}