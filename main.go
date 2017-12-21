package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
	"strings"
	"strconv"
	"github.com/heidimitre/priceProject/item"
)

func main(){

	tempURL:= ""


	req,_ := http.NewRequest("GET", tempURL, nil)
	client:= http.Client{}
	response,_ := client.Do(req)
	body,_ := ioutil.ReadAll(response.Body)
	bodyString:= string (body)
	re := regexp.MustCompile("\"buyboxPrice\"\\:\"\\$([0-9]+).([0-9]+)")
	priceString:= re.FindAllString(bodyString, -1)[0]
	priceString = priceString[strings.Index(priceString,"$")+1 :len(priceString)]
	price,_ := strconv.ParseFloat(priceString, 64)

	item:= item.CreateItem(price, "Exploding Kittens", tempURL)
	fmt.Println(item.Description, item.Price)
}
