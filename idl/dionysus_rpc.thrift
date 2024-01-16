namespace go dionysus.rpc

struct SayHelloRequest {
    1: string Say;
    2: string Name;
    3: string Gender;
}

struct SayHelloResponse {
    1: bool Ok;
    2: string Msg;
}
service DionysusRPC {
    SayHelloResponse sayHello(1: SayHelloRequest req);
}