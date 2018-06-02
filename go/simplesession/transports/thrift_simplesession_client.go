package transports


import (
	ss "github.com/OpenStars/backendclients/go/simplesession/thrift/gen-go/simplesession"
	"git.apache.org/thrift.git/lib/go/thrift"
	thriftpool "github.com/OpenStars/thriftpool"	
	)



var (
	bsMapPool = thriftpool.NewMapPool(1000, 3600, 3600, 
		thriftpool.GetThriftClientCreatorFunc( func(t thrift.TTransport, f thrift.TProtocolFactory) (interface{}) { return  (ss.NewTSimpleSessionServiceClientFactory(t,f)) }),
		thriftpool.DefaultClose)
	)

 


func init(){
	
}

//Get client by host:port
func GetSimpleSessionClient(bsHost, bsPort string) *thriftpool.ThriftSocketClient{
	client, _ := bsMapPool.Get(bsHost, bsPort).Get()
	return client;
}