package main
import (
	bs "github.com/OpenStars/backendclients/go/bigset/thrift/gen-go-0.11.0/openstars/core/bigset/generic"

	"github.com/OpenStars/backendclients/go/bigset/transports"
	"fmt"
	"context"
)
func main(){
	aClient := transports.GetBSClient("127.0.0.1", "18407")
	defer aClient.BackToPool();
	fmt.Println("Client: ", aClient);
	aClient.Client.(*bs.TStringBigSetKVServiceClient).BsPutItem(context.Background(), "x"  , &bs.TItem{[]byte( "hello" ), []byte("lau lam khong gap")}) 
	for i:=0 ; i<10 ; i++ {
		// fmt.Println(aClient.Client.(*KVStepCounter.KVStepCounterServiceClient).GetValue(context.Background(), "hello") )
	}
}