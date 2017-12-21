package item

type Item struct {
	Price       float64
	Description string
	Url			string
}

func CreateItem (newPrice float64, desc string, url string)(item *Item){
	newItem := new(Item)
	newItem.Price = newPrice
	newItem.Description = desc
	newItem.Url = url
	return newItem
}

func (item *Item) SetPrice (newPrice float64){
	item.Price = newPrice
}