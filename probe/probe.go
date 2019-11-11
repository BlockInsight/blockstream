package probe

import (
	"github.com/blockinsight/blockstream/cached"
	types "github.com/blockinsight/blockstream/types"
	"github.com/libs4go/errors"
	"github.com/libs4go/scf4go"
	"github.com/libs4go/slf4go"
	"github.com/libs4go/smf4go/service/grpcservice"
	grpc "google.golang.org/grpc"
)

var errVendor = errors.WithVendor("blockstream-probe")

// Errors .
var (
	ErrBlockchain = errors.New("unsupport blockchain", errVendor, errors.WithCode(-1))
	ErrChainID    = errors.New("unsupport blockchain ID", errVendor, errors.WithCode(-2))
)

type probeImpl struct {
	Cached           cached.CachedClient `inject:"cached.client"`
	chainIDValidator types.ChainIDValidator
}

// New create new probe instance
func New(config scf4go.Config) (grpcservice.Service, error) {

	chainIDValidator, err := types.GetChainIDValidator(config.SubConfig("chainId"))

	if err != nil {
		return nil, errors.Wrap(err, "get valid chainIds error")
	}

	return &probeImpl{
		chainIDValidator: chainIDValidator,
	}, nil
}

func (probe *probeImpl) GrpcHandler(server *grpc.Server) error {
	RegisterProbeServer(server, probe)
	return nil
}

func (probe *probeImpl) CreateJob(request *JobRequest, server Probe_CreateJobServer) error {

	if !probe.chainIDValidator.CheckChainID(request.Blockchain, request.ChainId) {
		return errors.Wrap(ErrChainID, "invalid blockchain %s id %s", request.Blockchain, request.ChainId)
	}

	job := &jobImpl{
		Logger:     slf4go.Get("probe-job"),
		JobRequest: request,
		cached:     probe.Cached,
	}

	return job.run(server)
}
