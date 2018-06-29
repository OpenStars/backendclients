package main
import (
	"github.com/OpenStars/backendclients/go/pprofile/thrift/gen-go/OpenStars/Platform/Profile"
	"github.com/OpenStars/backendclients/go/pprofile/transports"
	"fmt"
	"context"
)

func main(){
	aClient := transports.GetPProfileCompactClient("127.0.0.1", "8879")
	defer aClient.BackToPool();
	fmt.Println("Client: ", aClient);
	aClient.Client.(*Profile.TPlatformProfileServiceClient).PutData( context.Background(), &Profile.TPlatformProfile{ UID: 10,
		Username: "trungthanh",
		DisplayName : "Nguyen Trung Thanh",
		TrustedEmails : map[string]bool {"thanhnt@123xe.vn" : true, "trungthanhnt@gmail.com" : false} ,
		TrustedMobiles : map[string]bool {"0988903198": true, "0868887988":false},
	} )

	aClient.Client.(*Profile.TPlatformProfileServiceClient).SetExtData(context.Background(), 10, "Ability", "Coding very good")

	R,_ := aClient.Client.(*Profile.TPlatformProfileServiceClient).GetData(context.Background(), 10);
	fmt.Println("Result: ", R.Data);

}