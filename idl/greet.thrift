namespace go greet

struct GreetRequest {
    1: string Greet;
    2: string Name;
    3: string Gender;
}

struct GreetResponse {
    1: bool Ok;
    2: string Msg;
}