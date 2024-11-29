package main

import (
	"context"
	"log"

	"github.com/wizard-corp/api-gateway/grpc/build"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	client := build.NewTicketClient(conn)

	response, err := client.CreateTicket(ctx, &build.CreateTicketRequest{
		WriteUId:    "0000",
		TicketId:    "123",
		ChannelType: build.TicketChannelType_MAIL,
		Requirement: "urgent",
		Because:     "important",
		State:       build.TicketState_CREATED,
	})

	/*
		client := build.NewGreeterClient(conn)
		response, err := client.SayHello(ctx, &build.HelloRequest{Name: "Patrick"})
	*/

	if err != nil {
		log.Fatalf("could not call method: %v", err)
	}
	log.Printf("Response: %v", response)
}
