package test

import (
	"context"
	"github.com/nebulaim/telegramd/baselib/grpc_util"
	"github.com/nebulaim/telegramd/baselib/grpc_util/service_discovery"
	"github.com/nebulaim/telegramd/proto/mtproto"
	"testing"
)

func newAuthSessionRPCClient() mtproto.RPCSessionClient {
	cfg := service_discovery.ServiceDiscoveryClientConfig{
		ServiceName: "auth_session",
		EtcdAddrs:   []string{"http://127.0.0.1:2379"},
		Balancer:    "round_robin",
	}
	c, _ := grpc_util.NewRPCClient(&cfg)
	return mtproto.NewRPCSessionClient(c.GetClientConn())
}

func TestSetAuthKey(t *testing.T) {
	rpcCli := newAuthSessionRPCClient()
	rpcCli.SessionSetAuthKey(context.Background(), &mtproto.TLSessionSetAuthKey{
		AuthKey: nil,
	})
}
