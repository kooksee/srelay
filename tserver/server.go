package tserver

import (
	"strings"
	"time"

	knet "SRelay/utils/net"
)

type KcpServer struct {
	clients map[string]knet.Conn
	l       *knet.KcpListener

	host string
	port int
}

func NewKcpServer(host string, port int) *KcpManager {
	return &KcpManager{host: host, port: port}
}

func (km *KcpServer) Listen() (err error) {
	km.l, err = knet.ListenKcp(km.host, km.port)
	return
}

func (km *KcpServer) Start() {

	for {
		c, err := km.l.Accept()
		c.SetReadDeadline(time.Now().Add(connReadTimeout))

		addr := strings.Split(c.RemoteAddr().String(), ":")
		if err != nil {
			delete(km.clients, addr[1])
			continue
		}
		km.clients[addr[1]] = c



		// go io.Copy()
	}
}