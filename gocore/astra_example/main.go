package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type database struct {
	Id    string `json:"id"`
	OrgId string `json:"orgId"`
	Info  info   `json:"info"`
	//Info struct {
	//	Region string `json:"region"`
	//}
}

type info struct {
	Region string `json:"region"`
}

//type T struct {
//	PrometheusRemote struct {
//		Endpoint     string `json:"endpoint"`
//		AuthStrategy string `json:"auth_strategy"`
//		User         string `json:"user"`
//		Password     string `json:"password"`
//	} `json:"prometheus_remote"`
//}

func main() {

	url := "https://api.astra.datastax.com/v2/databases"
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", "Bearer AstraCS:idyxmCUSysBnTZLSqKEmXglD:1b2a519d7348c20097e68758b28657a27744bfaedb4d8d769bf10968f1a93234")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	//var resp APIResponse
	var resp []database
	//var resp []interface{}

	if err := json.Unmarshal(body, &resp); err != nil {
		fmt.Println(err)
		panic(err)
	}
	//fmt.Println(resp[0]["ownerId"])
	//fmt.Println(resp[0])
	//
	//data := resp[0].(map[string]interface{})
	//dataMapInfo := data["info"].(map[string]interface{})
	//
	//fmt.Println(dataMapInfo["region"])
	//fmt.Println(data["ownerId"])
	fmt.Println(resp)
	// fmt.Println(string(body))
	//for k, v := range resp {
	//	fmt.Println(k, ": ", v)
	//}

}
