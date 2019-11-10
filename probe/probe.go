package probe

import (
	"github.com/libs4go/scf4go"
	"github.com/libs4go/smf4go/service/grpcservice"
	grpc "google.golang.org/grpc"
)

type probeImpl struct {
}

// New create new probe instance
func New(config scf4go.Config) (grpcservice.Service, error) {
	return &probeImpl{}, nil
}

func (probe *probeImpl) GrpcHandler(server *grpc.Server) error {
	RegisterProbeServer(server, probe)
	return nil
}

func (probe *probeImpl) CreateJob(request *JobRequest, server Probe_CreateJobServer) error {
	job := &jobImpl{request}

	return job.run(server)
}
