package item

import (
	"github.com/heidimitre/priceProject/price"
	"github.com/heidimitre/priceProject/updater"

	"time"
	"github.com/davecgh/go-spew/spew"
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
	go func () {
		for true {
			newItem.AddPrice()
			spew.Dump(newItem)
			time.Sleep(30*time.Second)
		}
	}()
	return newItem
}

func (item *Item) AddPrice (){
	item.Pricelist = append(item.Pricelist, updater.GetCurrentPrice(item.Url))
}