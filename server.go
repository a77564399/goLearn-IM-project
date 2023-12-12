package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// 创建一个返回Server接口方法
func NewServer(ip string, port int) *Server {
	server := &Server{Ip: ip, Port: port}
	return server
}

func (this *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("Listen err:", err)
	}
	defer listener.Close()

	//死循环将服务一直保持在前台
	for {
		//	accept，监听端口，当有其他服务连接到此端口的时候，执行handle操作
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}
		//	do Handle
		go this.Handle(conn)
	}
}

// 建立链接之后的操作
func (this *Server) Handle(conn net.Conn) {
	fmt.Println("端口链接成功")
}
