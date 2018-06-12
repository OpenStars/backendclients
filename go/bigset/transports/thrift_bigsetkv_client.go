package transports


import (
	bs "github.com/OpenStars/backendclients/go/bigset/thrift/gen-go-0.11.0/openstars/core/bigset/generic"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	thriftpool "github.com/OpenStars/thriftpool"	
	)


var (
	bsMapPool = thriftpool.NewMapPool(1000, 3600, 3600, 
		thriftpool.GetThriftClientCreatorFunc( func(c thrift.TClient) (interface{}) { return  (bs.NewTStringBigSetKVServiceClient(c)) }),
		thriftpool.DefaultClose)
	)

 


func init(){
	fmt.Println("init thrift bigsetkv client ");
}

//Get client by host:port
func GetBSClient(bsHost, bsPort string) *thriftpool.ThriftSocketClient{
	client, _ := bsMapPool.Get(bsHost, bsPort).Get()
	return client;
}