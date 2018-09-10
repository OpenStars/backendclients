package main
import (
	"github.com/OpenStars/backendclients/go/bigset/thrift/gen-go/openstars/core/bigset/generic"

	"github.com/OpenStars/backendclients/go/bigset/transports"
	"fmt"
	"context"
)
func main(){
	aClient := transports.GetBsGenericClient("127.0.0.1", "18407")
	defer aClient.BackToPool();
	fmt.Println("Client: ", aClient);
	aClient.Client.(*generic.TStringBigSetKVServiceClient).BsPutItem(context.Background(), "x"  , &generic.TItem{[]byte( "hello" ), []byte("lau lam khong gap")})
	for i:=0 ; i<10 ; i++ {

	}
}