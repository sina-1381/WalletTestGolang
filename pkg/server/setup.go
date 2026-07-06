package server

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Grpc struct {
	Server   *grpc.Server
	Listener net.Listener
}

func NewGrpc(port string) (*Grpc, error) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, err
	}
	server := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     5 * time.Minute,
			MaxConnectionAge:      30 * time.Minute,
			MaxConnectionAgeGrace: 5 * time.Second,
			Time:                  1 * time.Minute,
			Timeout:               3 * time.Second,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             10 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.MaxRecvMsgSize(4*1024*1024), // 4MB
		grpc.MaxSendMsgSize(4*1024*1024), // 4MB
		grpc.MaxConcurrentStreams(100),
	)
	return &Grpc{
		Server:   server,
		Listener: listener,
	}, nil
}

func (g *Grpc) Start() error {
	return g.Server.Serve(g.Listener)
}

func (g *Grpc) GracefulClose() {
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		g.Server.GracefulStop()
	}()
}
