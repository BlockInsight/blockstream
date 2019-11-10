package probe

import (
	"time"
)

type jobImpl struct {
	*JobRequest
}

func (job *jobImpl) run(server Probe_CreateJobServer) error {
	for {
		time.Sleep(time.Second)

		response := &JobResponse{
			Block: []byte("hello"),
		}

		if err := server.Send(response); err != nil {
			return err
		}
	}
}
