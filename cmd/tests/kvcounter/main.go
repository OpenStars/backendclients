package main
import (
	"github.com/OpenStars/backendclients/go/kvcounter/thrift/gen-go/OpenStars/Counters/KVStepCounter"
	"github.com/OpenStars/backendclients/go/kvcounter/transports"
	"fmt"
	"context"
)
func main(){
	aClient := transports.GetKVCounterBinaryClient("127.0.0.1", "8881")
	defer aClient.BackToPool();
	fmt.Println("Client: ", aClient);
	aClient.Client.(*KVStepCounter.KVStepCounterServiceClient).CreateGenerator(context.Background(),"hello")
	for i:=0 ; i<10 ; i++ {
		fmt.Println(aClient.Client.(*KVStepCounter.KVStepCounterServiceClient).GetValue(context.Background(), "hello") )
	}
}