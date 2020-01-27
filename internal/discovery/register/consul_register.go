package register

import (
	"fmt"
	"lightim/internal/discovery"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type ConsulRegister struct{
	Target string
	Ttl int
}

func NewConsulRegister(target string,ttl int) *ConsulRegister{
	return &ConsulRegister{target,ttl}
}

func(self *ConsulRegister) Register(info *discovery.RegisterInfo) error{
	//initail consul client
	client,err := discovery.GenerateConsulClient(self.Target)
	if err!=nil{
		log.Println("Create consul client error:", err.Error())
	}
	serviceId:=discovery.GenerateServiceId(info.ServiceName,info.Host,info.Port)
	reg := &consulapi.AgentServiceRegistration{
		ID:      serviceId,
		Name:    info.ServiceName,
		Tags:    []string{info.ServiceName},
		Port:    info.Port,
		Address: info.Host,
	}
	if err = client.Agent().ServiceRegister(reg);err!=nil{
		panic(err)
	}
	//initial register service check
	check :=consulapi.AgentServiceCheck{TTL: fmt.Sprintf("%ds", self.Ttl), Status: consulapi.HealthPassing}
	agentcheckInfo:=consulapi.AgentCheckRegistration{
		ID:                serviceId,
		Name:              info.ServiceName,
		ServiceID:         serviceId,
		AgentServiceCheck: check,
	}
	if err=client.Agent().CheckRegister(&agentcheckInfo);err!=nil{
		return fmt.Errorf("ConsulGrpc: initial register service check to consul error: %s", err.Error())
	}
	go func(){
		ch:=make(chan os.Signal,1)
		signal.Notify(ch,syscall.SIGTERM,syscall.SIGINT,syscall.SIGKILL,syscall.SIGHUP,syscall.SIGQUIT)
		x:=<-ch
		log.Println("Grpc: receive signal: ", x)
		// deregister service
		self.DeRegister(info)
		s, _ := strconv.Atoi(fmt.Sprintf("%d", x))
		os.Exit(s)
	}()
	return nil
}

func(self *ConsulRegister) DeRegister(info *discovery.RegisterInfo) error{
	serviceId := discovery.GenerateServiceId(info.ServiceName, info.Host, info.Port)
	client,err:=discovery.GenerateConsulClient(self.Target)
	if err!=nil{
		log.Println("Create consul client error:", err.Error())
	}
	err=client.Agent().ServiceDeregister(serviceId)
	if err != nil {
		log.Println("ConsulGrpc: deregister service error: ", err.Error())
	} else {
		log.Println("ConsulGrpc: deregistered service from consul server.")
	}

	err = client.Agent().CheckDeregister(serviceId)
	if err != nil {
		log.Println("ConsulGrpc: deregister check error: ", err.Error())
	}

	return nil


}
