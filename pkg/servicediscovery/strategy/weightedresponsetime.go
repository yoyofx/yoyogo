package strategy

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"log"
	"net"
	"sort"
	"time"
)

type serviceDuration struct {
	service  servicediscovery.ServiceInstance
	duration int64
}

type WeightedResponseTime struct {
	s servicediscovery.IServiceDiscovery
}

func ping(host servicediscovery.ServiceInstance, channel chan serviceDuration, curIndex int, endIndex int) {
	destAddress := net.IPAddr{IP: net.ParseIP(host.GetHost())}
	//初始化ICMP协议头
	icmp := ICMP{
		Type:        8,
		Code:        0,
		CheckSum:    0,
		Identifier:  0,
		SequenceNum: 0,
	}
	var buffer bytes.Buffer
	_ = binary.Write(&buffer, binary.BigEndian, icmp)
	icmp.CheckSum = checkSum(buffer.Bytes())
	buffer.Reset()
	conn, err := net.DialIP("ip4:icmp", nil, &destAddress)
	defer conn.Close()
	if err != nil {
		msg := fmt.Sprintf("Fail to connect to remote host: %s\n", err)
		log.Println(msg)
	}
	_ = binary.Write(&buffer, binary.BigEndian, icmp)
	if _, err := conn.Write(buffer.Bytes()); err != nil {

		log.Fatal(err)
	}
	//请求时间计算
	startTime := time.Now()
	_ = conn.SetReadDeadline(time.Now().Add(time.Second * 2))
	recv := make([]byte, 1024)
	_, err = conn.Read(recv)
	if err != nil {
		log.Fatal(err)
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime).Nanoseconds() / 1e6
	channel <- serviceDuration{duration: duration, service: host}
	if curIndex == endIndex {
		close(channel)
	}

}

func (w *WeightedResponseTime) Next(serviceName string) (servicediscovery.ServiceInstance, error) {
	//获取服务节点
	var errorMsg error = nil
	endpoints := w.s.GetAllInstances(serviceName)
	//初始化服务信息切片
	serviceDurationSlice := make([]serviceDuration, len(endpoints))
	channel := make(chan serviceDuration)
	//多线程ping
	for i, v := range endpoints {
		go ping(v, channel, i, len(endpoints)-1)
	}
	for v := range channel {
		if v.duration == 0 || v.service == nil {
			errorMsg = errors.New("get service response error")
		}
		serviceDurationSlice = append(serviceDurationSlice, v)
	}
	sort.SliceStable(serviceDurationSlice, func(i, j int) bool {
		return serviceDurationSlice[i].duration < serviceDurationSlice[j].duration
	})
	return serviceDurationSlice[0].service, errorMsg
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
