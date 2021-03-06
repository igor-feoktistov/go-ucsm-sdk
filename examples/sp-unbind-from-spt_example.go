package main

import (
	"fmt"

	"go-ucsm-sdk/util"
)

func main() {
	endPoint := "https://ucsm01.example.com/"
	username := "admin"
	password := "secret"
	dn := "org-root/org-Linux/ls-phygymdev01-lab"
	client, err := util.AaaLogin(endPoint, username, password)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer client.AaaLogout()
	if lsServer, err := util.SpUnbindFromSpt(client, dn); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Dn: %s\n", lsServer.Dn)
		fmt.Printf("\tConfiguration State: %s\n", lsServer.ConfigState)
		fmt.Printf("\tMaintenance Policy: %s\n", lsServer.MaintPolicyName)
		fmt.Printf("\tSource Template: %s\n", lsServer.SrcTemplName)
	}
}
