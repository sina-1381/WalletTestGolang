package server

import (
	"WalletTest/internal"
	"WalletTest/proto"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func NewWallet(port string) (*internal.Wallet, error) {
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
	proto.RegisterWalletServiceServer(server, &internal.Wallet{})
	return &internal.Wallet{
		Server:   server,
		Listener: listener,
	}, nil
}

func StartWallet(server *grpc.Server, listener net.Listener) error {
	return server.Serve(listener)
}

func GracefulCloseWallet(server *grpc.Server) {
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		server.GracefulStop()
	}()
}
