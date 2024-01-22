include "greet.thrift"

namespace go api

service Dionysus {
    greet.GreetResponse greet(1: greet.GreetRequest req) (api.get="greet");
}