package main

import (
		"net/http"
		"io/ioutil"
		"os"
		"fmt"
)

func main() {
	
	user :=os.Getenv("DNS_USERNAME")
	password :=os.Getenv("DNS_PASSWORD")
	hostname :=os.Getenv("HOSTNAME")
	res, _ := http.Get("https://api.ipify.org")
	ip, _ := ioutil.ReadAll(res.Body)

	client := &http.Client{}
	URL := "https://domains.google.com/nic/update?hostname="+hostname+"&"+string(ip)
	
	fmt.Println(URL)
    	
	//pass the values to the request's body
    	req, _ := http.NewRequest("POST", URL, nil)
    	req.SetBasicAuth(user, password)
	
	resp, err:=client.Do(req)
	
	fmt.Println(resp)
	fmt.Println(err)
 
}
		
