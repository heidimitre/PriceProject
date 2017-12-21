package main

import (
	//"net/http"
	//"io/ioutil"
	"fmt"
	"github.com/heidimitre/priceProject/item"
	//"regexp"
	//"os"
)

func main(){
	//req,_ := http.NewRequest("GET", os.Getenv("TARGET"), nil)
	//client:= http.Client{}
	//response,_ := client.Do(req)
	//body,_ := ioutil.ReadAll(response.Body)
	//bodyString:= string (body)
	//re := regexp.MustCompile("\"buyboxPrice\"\\:\"\\$([0-9]+).([0-9]+)")
	item:= item.CreateItem(19.99, "Exploding Kittens")
	fmt.Println(item.Description, item.Price)
	//fmt.Println(re.FindAllString(bodyString, -1))
}
