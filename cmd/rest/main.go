package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	hellov1 "github.com/techinnovai/hello-go-api/gen/hello/v1"
	"google.golang.org/grpc"
)

func main() {
	mux := runtime.NewServeMux()

	err := hellov1.RegisterHelloServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		"localhost:9090",
		[]grpc.DialOption{grpc.WithInsecure()},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("REST server running on :8080")
	http.ListenAndServe(":8080", mux)
}
