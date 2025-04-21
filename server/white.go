package server

import (
	"github.com/armon/go-socks5"
	"golang.org/x/net/context"
	"net"
)

type WhiteIP struct {
	ips []string
}

func (w *WhiteIP) Refresh(ip []string) {
	w.ips = make([]string, len(ip))
	copy(w.ips, ip)
}

func (w *WhiteIP) Allow(ctx context.Context, req *socks5.Request) (context.Context, bool) {
	clientIP, _, _ := net.SplitHostPort(req.RemoteAddr.String())
	for _, ip := range w.ips {
		if ip == clientIP {
			return ctx, true
		}
	}
	return ctx, false
}
