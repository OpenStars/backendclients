package main
import (
	"github.com/OpenStars/backendclients/go/bigset/thrift/gen-go/openstars/core/bigset/generic"

	"github.com/OpenStars/backendclients/go/bigset/transports"
	"fmt"
	"context"
	"os"
	"strconv"
	"bufio"
)

func testPut( s string ) {
	for x := 0; x < 10 ; x ++{
		aClient := transports.GetBsGenericClient("127.0.0.1", "18407")
		if aClient != nil {
			
			fmt.Println("Client: ", aClient);
			for i := 0; i < 100; i ++ {
				aClient.Client.(*generic.TStringBigSetKVServiceClient).BsPutItem(context.Background(), "x"  , &generic.TItem{[]byte( "hello "+s ), []byte("lau lam khong gap")})
			}
			fmt.Println("test ", s);
			aClient.BackToPool();
		} 	else {
			fmt.Println("test error ", s);
		}
	}
}

func main(){
	for i:=0 ; i<1000 ; i++ {
		go testPut( strconv.Itoa(i) )
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}