/*
main.go
Copyright (C) 2013-2016  Max Karpenkov

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
*/

package main

import (

	"fmt"
	"zabconf/zabbix_api"
	
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

//    f1 := make(map[string]string, 0)
//    o1 := make(map[string]string, 0)
//    s1 := make(map[string]string, 0)

//    groups := make(map[string]string, 0)
//    templates := make(map[string]string, 0)
//    interfaces := make(map[string]string, 0)

//    f1["host"] = "Template OS Mac OS X"
//    o1["h1"] = "status"
//    o1["h2"] = "available"
//    o1["h3"] = "host"
//    s1["key_"] = "net.if.in[en1]"

//    groups["groupid"] = "6"
//    templates["templateid"] = "10079"
//    interfaces["type"] = "1"
//    interfaces["main"] = "1"
//    interfaces["useip"] = "1"
//    interfaces["ip"] = "192.168.160.166"
//    interfaces["dns"] = ""
//    interfaces["port"] = "10050"

//    params["filter"] = f1
//    params["search"] = s1
//    params["host"] = "Kasenna MacBook"
//    params["hostids"] = "10089"
//    params["interfaces"] = interfaces
//    params["groups"] = groups
//    params["templates"] = templates
    params["history"] = "3"
    params["itemids"] = "23466"
    params["limit"] = "10"
    params["sortfield"] = "clock"
    params["sortorder"] = "DESC"
    params["output"] = "extend"
//    params["templated_hosts"] = 1
    
    ret, err := api.History("get", params)
    
    if err != nil {
	fmt.Println(err)
    }
    
//    fmt.Println(ret[0].Value)
    fmt.Println(ret)
}
