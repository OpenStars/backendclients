package transports

import (
	"github.com/OpenStars/backendclients/go/mapnoitfycall/thrift/gen-go/OpenStars/Common/MapNotifyCallKV"
	// "github.com/OpenStars/backendclients/go//gen-go/OpenStars/Common/MapPhoneNumberPubkeyKV" //Todo: Fix this
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/OpenStars/thriftpool"
)

var (
	mTMapNotifyCallKVServiceBinaryMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} {
			return (MapNotifyCallKV.NewTMapNotifyKVServiceClient(c))
		}),
		thriftpool.DefaultClose)

	mTMapNotifyCallKVServiceCommpactMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} {
			return (MapNotifyCallKV.NewTMapNotifyKVServiceClient(c))
		}),
		thriftpool.DefaultClose)
)

func init() {
	fmt.Println("init thrift TMapPhoneNumberPubkeyKVService client ")
}

//GetTMapPhoneNumberPubkeyKVServiceBinaryClient client by host:port
func GetTMapNotifyCallKVServiceBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := mTMapNotifyCallKVServiceBinaryMapPool.Get(aHost, aPort).Get()
	return client
}

//GetTMapPhoneNumberPubkeyKVServiceCompactClient get compact client by host:port
func GetTMapNotifyCallKVServiceCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := mTMapNotifyCallKVServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
