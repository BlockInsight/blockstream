package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/blockinsight/blockstream/cached"
	"github.com/blockinsight/blockstream/types"
	"github.com/blockinsight/model.proto/golang/eth"

	leveldbcached "github.com/blockinsight/blockstream/cached/leveldb"
	"github.com/blockinsight/blockstream/probe"
	_ "github.com/libs4go/scf4go/codec"
	_ "github.com/libs4go/slf4go/backend/console"
	_ "github.com/libs4go/slf4go/filter/cached"
	"github.com/libs4go/smf4go"
	"github.com/libs4go/smf4go/app"
	"github.com/libs4go/smf4go/service/grpcservice"
	"google.golang.org/grpc"
)

func testJob(client probe.ProbeClient) {

	time.Sleep(time.Second * 4)

	job, err := client.CreateJob(context.Background(), &probe.JobRequest{
		Blockchain: types.Blockchain_ETH,
		RemoteUrl:  "https://mainnet.infura.io/v3/44ab06a5fca644df953378ac1c16d2b9",
		ChainId:    "mainnet",
		Offset:     1,
		Count:      -1,
	})

	if err != nil {
		panic(err)
	}

	for {
		resp, err := job.Recv()

		if err != nil {
			panic(err)
		}

		var block eth.Blockchain

		err = proto.Unmarshal(resp.Block, &block)

		if err != nil {
			panic(err)
		}

		printObjec(block)
	}
}

func printObjec(val interface{}) {
	buff, _ := json.MarshalIndent(val, "\t", "\t")
	println(string(buff))
}

func main() {
	grpcService := grpcservice.New("test.grpc")

	grpcService.Local("probe.server", probe.New)

	grpcService.Remote("probe.client", func(conn *grpc.ClientConn) (smf4go.Service, error) {
		client := probe.NewProbeClient(conn)

		go testJob(client)

		return client, nil
	})

	grpcService.Local("cached.leveldb.server", leveldbcached.New)

	grpcService.Remote("cached.client", func(conn *grpc.ClientConn) (smf4go.Service, error) {
		return cached.NewCachedClient(conn), nil
	})

	app.Run("blockstream-lite")
}
