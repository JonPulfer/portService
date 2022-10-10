package adapter

import (
	"context"

	"github.com/JonPulfer/portsService/internal/portService/service"
	portPb "github.com/JonPulfer/portsService/pb/message/port"
)

// One type of input source could be a grpc client stream server. I have generated
// the code for this and shown how I would fit it into the design by injecting the
// portSvc.

type grpcServer struct {
	portSvc service.Port
	portPb.PortService_StorePortServer
}

func (g *grpcServer) StorePort(server portPb.PortService_StorePortServer) error {
	// TODO implement me
	panic("implement me")
}

func (g *grpcServer) FindPort(ctx context.Context, request *portPb.FindPortRequest) (*portPb.SerialisedPort, error) {
	// TODO implement me
	panic("implement me")
}

func NewGrpcStreamServer(portSvc service.Port) portPb.PortServiceServer {
	return &grpcServer{portSvc: portSvc}
}
