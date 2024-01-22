include "greet.thrift"
namespace go rpc

service Dionysus {
    greet.GreetResponse greet(1: greet.GreetRequest req);
}