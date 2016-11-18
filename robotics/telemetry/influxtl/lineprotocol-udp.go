package influxtl

import (
	"log"
	"net"
	"fmt"
	"github.com/fuzzycow/ev32go/robotics/telemetry"
)

type Client struct {
	conn *net.UDPConn
	raddr string
	net string
	cont bool
	clientId string
}

func NewClient(net, raddr, clientId string) *Client {
	tl := &Client{
		raddr: raddr,
		net: net,
		clientId: clientId,
	}
	return tl
}

func (tl *Client) Open() error {
	log.Printf("opening udp conn to InfluxDB UDP listener at %s",tl.raddr)
	udpRemoteAddr,err := net.ResolveUDPAddr(tl.net,tl.raddr)
	if err != nil {
		return err
	}
	conn,err := net.DialUDP(tl.net,nil,udpRemoteAddr)
	if err != nil {
		return err
	}
	tl.conn = conn
	return nil
}

func (tl *Client) Write(b []byte) (int,error) {
	if tl.conn == nil {
		return 0,fmt.Errorf("Connection is not open")
	}
	// log.Printf("FIXME99: %s",string(b))
	n, err := tl.conn.Write(b)
	/*if err != nil {
		log.Fatalf("FIXME100: %v",err)
	} */
	return n,err
}

func (tl *Client) Close() error {
	if tl.conn == nil {
		return nil
	}
	return tl.conn.Close()
}

func (tl *Client) NewMessage(subject ...string) telemetry.Message {
	m := newClientMessage(tl, subject...)
	return m
}

