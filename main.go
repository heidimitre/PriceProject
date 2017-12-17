package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"regexp"
	"os"
)

func main(){
	req,_ := http.NewRequest("GET", os.Getenv("TARGET"), nil)
	client:= http.Client{}
	response,_ := client.Do(req)
	body,_ := ioutil.ReadAll(response.Body)
	bodyString:= string (body)
	re := regexp.MustCompile("\"buyboxPrice\"\\:\"\\$([0-9]+).([0-9]+)")
	fmt.Println(re.FindAllString(bodyString, -1))

}
