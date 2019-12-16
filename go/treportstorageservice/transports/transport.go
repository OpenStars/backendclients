package transports

import (
	"github.com/OpenStars/backendclients/go/treportstorageservice/thrift/gen-go/OpenStars/Common/TReportStorageService"
	"log"

	"github.com/OpenStars/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mTReportStorageServiceBinaryMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (TReportStorageService.NewTReportStorageServiceClient(c)) }),
		thriftpoolv2.DefaultClose)

	mTReportStorageServiceCommpactMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (TReportStorageService.NewTReportStorageServiceClient(c)) }),
		thriftpoolv2.DefaultClose)
)

func init() {
	log.Println("init thrift TReportStorageService client ")
}

//GetTReportStorageServiceBinaryClient client by host:port
func GetTReportStorageServiceBinaryClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mTReportStorageServiceBinaryMapPool.Get(aHost, aPort).Get()
	return client
}

//GetTReportStorageServiceCompactClient get compact client by host:port
func GetTReportStorageServiceCompactClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mTReportStorageServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
