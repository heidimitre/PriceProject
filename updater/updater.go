package updater

import (
	"net/http"
	"io/ioutil"
	"regexp"
	"strings"
	"strconv"
	"github.com/heidimitre/priceProject/price"
	"time"
)

func GetCurrentPrice(url string) (newPrice *price.Price){
	req,_ := http.NewRequest("GET", url, nil)
	client:= http.Client{}
	response,_ := client.Do(req)
	body,_ := ioutil.ReadAll(response.Body)
	bodyString:= string (body)
	re := regexp.MustCompile("\"buyboxPrice\"\\:\"\\$([0-9]+).([0-9]+)")
	priceString:= re.FindAllString(bodyString, -1)[0]
	priceString = priceString[strings.Index(priceString,"$")+1 :len(priceString)]
	value,_ := strconv.ParseFloat(priceString, 64)
	price:= price.CreatePrice(value, time.Now())
	return price
}

