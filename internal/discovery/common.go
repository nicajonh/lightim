package discovery

import (
	"fmt"
	"log"
	consulapi "github.com/hashicorp/consul/api"
)

func GenerateConsulClient(consuladdr string) (*consulapi.Client,error){
	//consul client
	config:=consulapi.DefaultConfig()
	config.Address=consuladdr
	client,err:=consulapi.NewClient(config)
	if err!=nil{
		log.Println("create consul client error:", err.Error())
		return nil,err
	}
	return client,nil
}

func GenerateServiceId(name string,host string,port int) string{
	return fmt.Sprintf("%s-%s-%d",name,host,port)
}

