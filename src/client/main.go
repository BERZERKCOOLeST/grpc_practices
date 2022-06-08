package main

import (
	"context"
	"grpc_soldshop/src/proto/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const address = "localhost:50051"

func main() {
	creds, err := credentials.NewClientTLSFromFile("./auth//server/server.pem", "127.0.0.1")
	if err != nil {
		log.Fatalf("failed to create creds: %v", err)
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)

	name := "Apple iPhone 11"
	description := `Meet Apple iPhone 11.`
	price := float32(1000.0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddProduct(ctx, &pb.Product{
		Name: name, Description: description, Price: price,
	})
	if err != nil {
		log.Fatalf("Couldn't add product: %v", err)
	}
	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Couldn't get product: %v", err)
	}
	log.Printf("Product: %v", product.String())
	select {}
}
