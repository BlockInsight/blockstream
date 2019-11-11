package cached

import (
	"context"
	"encoding/binary"
	"fmt"
	"path/filepath"
	"sync"

	"github.com/blockinsight/blockstream/cached"
	"github.com/blockinsight/blockstream/types"
	"github.com/libs4go/errors"
	"github.com/libs4go/scf4go"
	"github.com/libs4go/slf4go"
	"github.com/libs4go/smf4go/service/grpcservice"
	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/grpc"
)

var errVendor = errors.WithVendor("blockstream-cached-leveldb")

// Errors .
var (
	ErrChainIDs = errors.New("valid chainIDs list not provide", errVendor, errors.WithCode(-1))
	ErrChainID  = errors.New("unsupport chainID", errVendor, errors.WithCode(-2))
)

// block cached with leveldb
type cachedImpl struct {
	slf4go.Logger
	sync.RWMutex
	rootpath         string
	databases        map[string]*leveldb.DB
	chainIDValidator types.ChainIDValidator
}

// New create leveldb cached for blockchain data
func New(config scf4go.Config) (grpcservice.Service, error) {

	rootpath := config.Get("root").String("./cached")

	chainIDValidator, err := types.GetChainIDValidator(config.SubConfig("chainId"))

	if err != nil {
		return nil, errors.Wrap(err, "get valid chainIds error")
	}

	return &cachedImpl{
		Logger:           slf4go.Get("cached-leveldb"),
		rootpath:         rootpath,
		databases:        make(map[string]*leveldb.DB),
		chainIDValidator: chainIDValidator,
	}, nil
}

func (impl *cachedImpl) GrpcHandler(server *grpc.Server) error {
	cached.RegisterCachedServer(server, impl)
	return nil
}

func (impl *cachedImpl) createCache(blockchain types.Blockchain, chainID string) (*leveldb.DB, error) {
	impl.Lock()
	defer impl.Unlock()

	key := fmt.Sprintf("%s%s", blockchain, chainID)

	db, ok := impl.databases[key]

	if ok {
		return db, nil
	}

	var err error

	path := filepath.Join(impl.rootpath, blockchain.String(), chainID)

	db, err = leveldb.OpenFile(path, nil)

	if err != nil {
		return nil, errors.Wrap(err, "create leveldb at %s error", path)
	}

	impl.databases[key] = db

	return db, nil
}

func (impl *cachedImpl) getOrCreateCache(blockchain types.Blockchain, chainID string) (*leveldb.DB, error) {
	impl.RLock()

	if !impl.chainIDValidator.CheckChainID(blockchain, chainID) {
		impl.RUnlock()
		return nil, errors.Wrap(ErrChainID, "unsupport blockchain %s id %s", blockchain, chainID)
	}

	key := fmt.Sprintf("%s%s", blockchain, chainID)

	db, ok := impl.databases[key]

	impl.RUnlock()

	if !ok {
		var err error
		db, err = impl.createCache(blockchain, chainID)

		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func blockNumberKey(number uint64) []byte {
	var result [8]byte
	binary.BigEndian.PutUint64(result[:], number)

	return result[:]
}

func (impl *cachedImpl) Put(ctx context.Context, request *cached.PutRequest) (*cached.PutResponse, error) {
	cache, err := impl.getOrCreateCache(request.Blockchain, request.ChainId)

	if err != nil {
		return nil, err
	}

	err = cache.Put(blockNumberKey(request.Offset), request.Block, nil)

	if err != nil {
		return nil, errors.Wrap(err, "write blockchain %s with chainId %s block %d error", request.Blockchain, request.ChainId, request.Offset)
	}

	impl.D("write blockchain {@blockchain} with chainId {@chainId} block {@offset} -- success", request.Blockchain, request.ChainId, request.Offset)

	return &cached.PutResponse{}, nil
}

func (impl *cachedImpl) Get(ctx context.Context, request *cached.GetRequest) (*cached.GetResponse, error) {
	cache, err := impl.getOrCreateCache(request.Blockchain, request.ChainId)

	if err != nil {
		return nil, err
	}

	block, err := cache.Get(blockNumberKey(request.Offset), nil)

	if err != nil {
		if err == leveldb.ErrNotFound {
			impl.D("get blockchain {@blockchain} with chainId {@chainId} block {@offset} -- not found", request.Blockchain, request.ChainId, request.Offset)
			return &cached.GetResponse{}, nil
		}

		return nil, errors.Wrap(err, "read blockchain %s with chainId %s block %d error", request.Blockchain, request.ChainId, request.Offset)
	}

	if block != nil {
		impl.D("get blockchain {@blockchain} with chainId {@chainId} block {@offset} -- success", request.Blockchain, request.ChainId, request.Offset)
	} else {
		impl.D("get blockchain {@blockchain} with chainId {@chainId} block {@offset} -- not found", request.Blockchain, request.ChainId, request.Offset)
	}

	return &cached.GetResponse{
		Block: block,
	}, nil
}
