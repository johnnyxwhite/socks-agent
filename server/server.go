package server

import (
	"github.com/armon/go-socks5"
	"golang.org/x/net/context"
	"net"
	"time"
)

type SocksServer struct {
	server *socks5.Server
}

func NewSocksServer6(pool ipool, whiteIp socks5.RuleSet) *SocksServer {
	conf := &socks5.Config{
		Dial: func(ctx context.Context, netw, addr string) (net.Conn, error) {
			netw = "tcp6"
			//本地地址  ipaddr是本地外网IP
			lAddr, err := net.ResolveTCPAddr(netw, pool.GetIp())
			if err != nil {
				return nil, err
			}
			//被请求的地址
			rAddr, err := net.ResolveTCPAddr(netw, addr)
			if err != nil {
				return nil, err
			}
			conn, err := net.DialTCP(netw, lAddr, rAddr)
			if err != nil {
				return nil, err
			}
			deadline := time.Now().Add(2 * time.Second)
			conn.SetDeadline(deadline)
			return conn, nil
		},
		Rules: whiteIp,
	}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}
	return &SocksServer{server: server}
}

func NewSocksServer4(whiteIp socks5.RuleSet) *SocksServer {
	conf := &socks5.Config{
		Rules: whiteIp,
	}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}
	return &SocksServer{server: server}
}

func (s *SocksServer) ListenAndServe(bind string) error {
	return s.server.ListenAndServe("tcp", bind)
}
