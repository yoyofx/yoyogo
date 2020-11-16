package LB

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/yoyofx/yoyogo/Abstractions/ServiceDiscovery"
	"log"
	"net"
	"sort"
	"time"
)

type serviceDuration struct {
	service  ServiceDiscovery.ServiceInstance
	duration int64
}

type WeightedResponseTime struct {
	s ServiceDiscovery.IServiceDiscovery
}

func (w *WeightedResponseTime) Next(serviceName string) (ServiceDiscovery.ServiceInstance, error) {
	//获取服务节点
	endpoints := w.s.GetAllInstances(serviceName)
	//初始化服务信息切片
	serviceDurationSlice := make([]serviceDuration, len(endpoints))
	for _, v := range endpoints {
		host := v.GetHost()
		destAddress := net.IPAddr{IP: net.ParseIP(host)}
		//初始化ICMP协议头
		icmp := ICMP{
			Type:        8,
			Code:        0,
			CheckSum:    0,
			Identifier:  0,
			SequenceNum: 0,
		}
		var buffer bytes.Buffer
		binary.Write(&buffer, binary.BigEndian, icmp)
		icmp.CheckSum = checkSum(buffer.Bytes())
		buffer.Reset()
		conn, err := net.DialIP("ip4:icmp", nil, &destAddress)
		if err != nil {
			fmt.Printf("Fail to connect to remote host: %s\n", err)
		}
		binary.Write(&buffer, binary.BigEndian, icmp)
		if _, err := conn.Write(buffer.Bytes()); err != nil {
			log.Fatal(err)
		}
		//请求时间计算
		startTime := time.Now()
		conn.SetReadDeadline(time.Now().Add(time.Second * 2))
		recv := make([]byte, 1024)
		_, err = conn.Read(recv)
		if err != nil {
			log.Fatal(err)
		}
		endTime := time.Now()
		duration := endTime.Sub(startTime).Nanoseconds() / 1e6
		serviceDurationSlice = append(serviceDurationSlice, serviceDuration{service: v, duration: duration})
		conn.Close()
	}
	sort.SliceStable(serviceDurationSlice, func(i, j int) bool {
		return serviceDurationSlice[i].duration < serviceDurationSlice[j].duration
	})
	return serviceDurationSlice[0].service, _

}

//定义ICMP协议结构体
type ICMP struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16
	Identifier  uint16
	SequenceNum uint16
}

func checkSum(data []byte) uint16 {
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)
	return uint16(^sum)
}
