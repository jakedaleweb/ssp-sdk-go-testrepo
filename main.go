package main

import (
	"fmt"

	"github.com/silverstripeltd/ssp-sdk-go/ssp"
)

func main() {
	client, _ := ssp.NewClient(nil)
	s := &ssp.UpdateInstanceType{
		InstanceType: "t2.small",
	}
	if err := client.UpdateInstanceType("cbmicro", "uat", s); err != nil {
		fmt.Printf("%s", err)
	}
}
