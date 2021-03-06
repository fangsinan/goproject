package httpClent


import (
	// "bytes"
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//POST 请求
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



//Get 请求
func Get(url string) string {
	req, _ := http.NewRequest("GET", url, nil)
	//req.Header.Set("authorization", "as*21^31no&")
	
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



// JsonToMap 
func JsonToMap(jsonStr string) (map[string]interface{}, error) {
     //创建一个map
    m := make(map[string]interface{}, 4)
 
    err := json.Unmarshal([]byte(jsonStr), &m) //第二个参数要地址传递
    if err != nil {
        fmt.Println("err = ", err)
        return nil, err
    }
 
    return m, nil
}
