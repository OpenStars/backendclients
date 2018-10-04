package transports

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/OpenStars/backendclients/go/i64listi64/thrift/gen-go/OpenStars/Common/list/i64listi64"
	thriftpool "github.com/OpenStars/thriftpool"
)

var (
	i64listi64BinMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (i64listi64.NewI64ListI64ServiceClient(c)) }),
		thriftpool.DefaultClose)
)

// GetI64ListI64BinaryClient get i64listi64 thrift client
func GetI64ListI64BinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := i64listi64BinMapPool.Get(aHost, aPort).Get()
	return client
}
