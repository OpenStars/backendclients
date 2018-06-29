package main
import (
	"github.com/OpenStars/backendclients/go/s2i64kv/thrift/gen-go/OpenStars/Common/S2I64KV"
	"github.com/OpenStars/backendclients/go/s2i64kv/transports"
	"fmt"
	"context"
)

func main(){
	aClient := transports.GetS2I64CompactClient("127.0.0.1", "8883")
	defer aClient.BackToPool();
	fmt.Println("Client: ", aClient);

	rcas,_ := aClient.Client.(*S2I64KV.TString2I64KVServiceClient).CasData(context.Background(), (S2I64KV.TKey)(10), &S2I64KV.TCasValue{0, 100});
	fmt.Println("rcas:", rcas);
	res, _ := aClient.Client.(*S2I64KV.TString2I64KVServiceClient).GetData(context.Background(), (S2I64KV.TKey)(10) )
	fmt.Println("get result: ",res)

	aClient.Client.(*S2I64KV.TString2I64KVServiceClient).PutData(context.Background(), (S2I64KV.TKey)(10), &S2I64KV.TI64Value{101});
	res, _ = aClient.Client.(*S2I64KV.TString2I64KVServiceClient).GetData(context.Background(), (S2I64KV.TKey)(10) )
	fmt.Println("get after put :", res)

}