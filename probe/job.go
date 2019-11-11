package probe

import (
	context "context"

	"github.com/golang/protobuf/proto"

	"github.com/blockinsight/blockstream/cached"
	types "github.com/blockinsight/blockstream/types"
	"github.com/blockinsight/model.proto/golang/eth"
	"github.com/blockinsight/rpc4go"
	_ "github.com/blockinsight/rpc4go/driver/eth" //
	"github.com/libs4go/errors"
	"github.com/libs4go/slf4go"
)

type jobImpl struct {
	slf4go.Logger
	*JobRequest
	cached cached.CachedClient
}

func (job *jobImpl) createRPCClient() (rpc4go.Client, error) {
	switch job.Blockchain {
	case types.Blockchain_ETH:
		return rpc4go.New("eth", job.RemoteUrl)
	default:
		return nil, errors.Wrap(ErrBlockchain, "unsupport blockchain rpc client %s", job.Blockchain)
	}
}

func (job *jobImpl) fetchAndCacheBlock(client rpc4go.Client, offset uint64) ([]byte, error) {

	switch job.Blockchain {
	case types.Blockchain_ETH:
		var block *eth.Blockchain
		err := client.BlockByNumber(int64(offset), &block)

		if err != nil {
			return nil, errors.Wrap(err, "fetch blockchain from %s error", job.RemoteUrl)
		}

		buff, err := proto.Marshal(block)

		if err != nil {
			return nil, errors.Wrap(err, "marshal blockchain from %s error", job.RemoteUrl)
		}

		job.I("put block {@blockchain} chainId {@chainId} offset {@offset}", job.Blockchain, job.ChainId, offset)
		_, err = job.cached.Put(context.Background(), &cached.PutRequest{
			Blockchain: job.Blockchain,
			ChainId:    job.ChainId,
			Block:      buff,
			Offset:     offset,
		})

		if err != nil {
			job.E("put block {@blockchain} chainId {@chainId} offset {@offset} error: {@error}", job.Blockchain, job.ChainId, offset, err)
		}

		return buff, nil

	default:
		return nil, errors.Wrap(ErrBlockchain, "unsupport blockchain rpc client %s", job.Blockchain)
	}

}

func (job *jobImpl) run(server Probe_CreateJobServer) error {

	client, err := job.createRPCClient()

	if err != nil {
		return err
	}

	for i := job.Offset; ; i++ {

		block, err := job.getCached(i)

		if err != nil {
			job.E("fetch cached block error: {@error}", err)
		}

		if block == nil {
			block, err = job.fetchAndCacheBlock(client, i)

			if err != nil {
				return err
			}
		}

		if err := server.Send(&JobResponse{Block: block}); err != nil {
			return err
		}
	}
}

func (job *jobImpl) getCached(offset uint64) ([]byte, error) {
	resp, err := job.cached.Get(context.Background(), &cached.GetRequest{
		Blockchain: job.Blockchain,
		ChainId:    job.ChainId,
		Offset:     offset,
	})

	if err != nil {
		return nil, errors.Wrap(err, "get cached block error")
	}

	return resp.Block, nil
}
