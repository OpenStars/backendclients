// This autogenerated skeleton file illustrates how to build a server.
// You should copy it to another filename to avoid overwriting it.

#include "TDataService.h"
#include <thrift/protocol/TBinaryProtocol.h>
#include <thrift/server/TSimpleServer.h>
#include <thrift/transport/TServerSocket.h>
#include <thrift/transport/TBufferTransports.h>

using namespace ::apache::thrift;
using namespace ::apache::thrift::protocol;
using namespace ::apache::thrift::transport;
using namespace ::apache::thrift::server;

using namespace  ::OpenStars::Common::MapPhoneNumberPubkeyKV;

class TDataServiceHandler : virtual public TDataServiceIf {
 public:
  TDataServiceHandler() {
    // Your initialization goes here
  }

  void getTokenByPubkey(TDataResult& _return, const std::string& pubkey) {
    // Your implementation goes here
    printf("getTokenByPubkey\n");
  }

  void getPubkeyByToken(TDataResult& _return, const std::string& token) {
    // Your implementation goes here
    printf("getPubkeyByToken\n");
  }

  TErrorCode::type putData(const std::string& pubkey, const std::string& token) {
    // Your implementation goes here
    printf("putData\n");
  }

};

int main(int argc, char **argv) {
  int port = 9090;
  ::apache::thrift::stdcxx::shared_ptr<TDataServiceHandler> handler(new TDataServiceHandler());
  ::apache::thrift::stdcxx::shared_ptr<TProcessor> processor(new TDataServiceProcessor(handler));
  ::apache::thrift::stdcxx::shared_ptr<TServerTransport> serverTransport(new TServerSocket(port));
  ::apache::thrift::stdcxx::shared_ptr<TTransportFactory> transportFactory(new TBufferedTransportFactory());
  ::apache::thrift::stdcxx::shared_ptr<TProtocolFactory> protocolFactory(new TBinaryProtocolFactory());

  TSimpleServer server(processor, serverTransport, transportFactory, protocolFactory);
  server.serve();
  return 0;
}

