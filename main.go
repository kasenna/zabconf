package main

import (

	"fmt"
	"github.com/zabbix"
	
)

func main() {

    api, err := zabbix.NewAPI("192.168.160.171", "Admin", "qweasd123")
    if err != nil {
	fmt.Println(err)
	return
    }

    fmt.Println("Connected to API")
    _, err = api.Login()
    if err != nil {
	fmt.Println(err)
	return
    }

    versionresult, err := api.Version()
    if err != nil {
	fmt.Println(err)
    }

    fmt.Println(versionresult)

    params := make(map[string]interface{}, 0)
    f1 := make(map[string]string, 0)
    o1 := make(map[string]string, 0)
    f1["host"] = "Kasenna MacBook"
    o1["h1"] = "status"
    o1["h2"] = "available"
    o1["h3"] = "host"
//    params["filter"] = f1
    params["output"] = o1
//    params["templated_hosts"] = 1
    
    ret, err := api.Host("get", params)
    
    if err != nil {
	fmt.Println(err)
    }
    
    fmt.Println(ret)

}
