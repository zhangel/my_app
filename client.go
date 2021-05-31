package main

import (
    "fmt"
    "log"
    "time"
    "context"
    "google.golang.org/grpc"
    "github.com/zhangel/my_app/api"
)

const (
    address = "localhost:8080"
)

func main() {
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    client := api.NewHelloClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    args:= api.HelloRequest{Name:"Tomzhang"}
    resp, err := client.SayHello(ctx, &args) 
    if err != nil {
        log.Fatalf("could not call SayHello: %v", err)
    }
    fmt.Printf("Name=%s\tresp= %+v\n", args.Name, resp) 
}
