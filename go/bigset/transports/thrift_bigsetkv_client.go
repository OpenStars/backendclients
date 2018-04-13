package transports


import (
	bs "github.com/OpenStars/backendclients/go/bigset/thrift/gen-go-0.11.0/openstars/core/bigset/generic"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	thriftpool "github.com/OpenStars/ThriftPool/src/thriftpool"	
	)


func Close(c *thriftpool.ThriftSocketClient) error {
	err := c.Socket.Close()
	//err = c.Client.(*bs.TStringBigSetKVServiceClient).Transport.Close()
	return err
}	

var (
	bsMapPool = thriftpool.NewMapPool(1000, 3600, 3600, 
		thriftpool.GetThriftClientCreatorFunc( func(t thrift.TTransport, f thrift.TProtocolFactory) (interface{}) { return  (bs.NewTStringBigSetKVServiceClientFactory(t,f)) }),
		Close)
	)

 


func init(){
	fmt.Println("init thrift bigsetkv client ");
}

//Get client by host:port
func GetBSClient(bsHost, bsPort string) *thriftpool.ThriftSocketClient{
	client, _ := bsMapPool.Get(bsHost, bsPort).Get()
	return client;
}