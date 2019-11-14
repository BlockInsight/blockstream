package worker

import (
	"github.com/libs4go/scf4go"
	"github.com/libs4go/smf4go/service/grpcservice"
)

type workerImpl struct {
}

// New create new worker instance
func New(config scf4go.Config) (grpcservice.Service, error) {
	return nil, nil
}
