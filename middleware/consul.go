package middleware

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
)

//注册服务到consul
func ConsulRegister() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = "192.168.1.2:8500"
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	// 创建注册到consul的服务到
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "wallet-service"
	registration.Name = "wallet service"
	registration.Port = 9002
	registration.Tags = []string{"wallet-service"}
	registration.Address = "192.168.1.2"

	// 增加consul健康检查回调函数
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d", registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
}

// 取消consul注册的服务
func ConsulDeRegister() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = "192.168.1.2:8500"
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}
	err = client.Agent().ServiceDeregister("111")
}

// 从consul中发现服务
func ConsulFindServer() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = "192.168.1.2:8500"
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	// 获取所有service
	services, _ := client.Agent().Services()
	for _, value := range services {
		fmt.Println(value.Address)
		fmt.Println(value.Port)
	}
	// 获取指定service
	service, _, err := client.Agent().Service("wallet-service", nil)
	if err == nil {
		fmt.Println(service.Address)
		fmt.Println(service.Port)
	}

}
