package httpClent


import (
	// "bytes"
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)


func Post(url string,postData map[string]string,Header map[string]string) string {


	var str string
	for i, v := range postData {
		str += i+"="+v+"&"
	}

    // byte, _ := json.Marshal(postData)
    // fmt.Println(string(byte))
    //req, _ := http.NewRequest("POST", url,strings.NewReader(string(byte)))
    req, _ := http.NewRequest("POST", url,strings.NewReader(str))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    for i, v := range Header {
		req.Header.Set(i, v)    	
    }
    

    resp, err := (&http.Client{}).Do(req)
    if err != nil {
        fmt.Println("save topic failed", err.Error())
    }
    defer resp.Body.Close()
    
    //var serviceResp ServiceResp
    respByte, _ := ioutil.ReadAll(resp.Body)
	return string(respByte)
}




func Get(url string) string {
	req, _ := http.NewRequest("GET", url, nil)
	//req.Header.Set("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC9hcHAudjQuYXBpLm5sc2dhcHAuY29tXC9hcGlcL3Y0XC9hdXRoXC9sb2dpbiIsImlhdCI6MTYyODA2MDcyOCwiZXhwIjoxNjI5MzU2NzI4LCJuYmYiOjE2MjgwNjA3MjgsImp0aSI6InJKTVVjMGNPRjBjMjFhRmIiLCJzdWIiOjc1Njc0NCwicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyJ9.m5xV9EWf6jyuXYcG8XGvqnlP98l3lcmddv3ZCblzelY")
	
	resp, err := (&http.Client{}).Do(req)
	//resp, err := http.Get(serviceUrl + "/topic/query/false/lsj")
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	defer resp.Body.Close()

	//var serviceResp ServiceResp
	respByte, _ := ioutil.ReadAll(resp.Body)
    return string(respByte)
}


