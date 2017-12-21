package item

type Item struct {
	Price       float64
	Description string
}

func CreateItem (newPrice float64, desc string)(item *Item){
	newItem := new(Item)
	newItem.Price = newPrice
	newItem.Description = desc
	return newItem
}

func (item *Item) SetPrice (newPrice float64){
	item.Price = newPrice
}