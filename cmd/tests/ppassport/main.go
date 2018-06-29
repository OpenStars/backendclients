package main
import (
	"github.com/OpenStars/backendclients/go/ppassport/thrift/gen-go/OpenStars/Platform/Passport"
	"github.com/OpenStars/backendclients/go/ppassport/transports"
	"fmt"
	"context"
)

func main(){
	aClient := transports.GetPassportCompactClient("127.0.0.1", "8883")
	defer aClient.BackToPool();
	fmt.Println("Client: ", aClient);

	r,_ := aClient.Client.(*Passport.TPassportServiceClient).PutData( context.Background(),
		 Passport.TKey(100), &Passport.TPassportInfo{Sha2Pwd:"MySHAhere", ExtData :map[string]string{"hello": "world" } } );

	fmt.Println("result: ",r)

	r2,_ := aClient.Client.(*Passport.TPassportServiceClient).GetData( context.Background(), Passport.TKey(100) );
	fmt.Println(r2);
}