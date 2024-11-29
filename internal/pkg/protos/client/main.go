package main

// ПРОСТО ДЛЯ ТЕСТА

import (
	"context"
	authv1 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/protos/gen/go/pushart.auth.v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {

	//var opts []grpc.DialOption

	serverAddr := "localhost:777"

	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//conn, err := grpc.Dial("localhost:777", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//// Code removed for brevity
	//
	client := authv1.NewAuthClient(conn)

	// Note how we are calling the GetBookList method on the server
	// This is available to us through the auto-generated code
	response, err := client.Login(context.Background(), &authv1.LoginRequest{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJiOTczODhhYi03ZGYwLTQ3ZGItOWVjMC00NDI1NDU5ZTlhOTMiLCJ1c2VybmFtZSI6InRlc3RfcnBjIiwicm9sZSI6IlJlYWRlciIsImlzcyI6ImF1dGgtYXBwIiwiZXhwIjoxNzMyNTMzODA3LCJpYXQiOjE3MzI0NDc0MDd9.UR_dXreYHapOny4C-qcNVJDFmXARmhDVmgWDalmzpL8"})

	log.Printf("Response={%v} err={%v}", response, err)

	log.Printf("response is: %v", response)
}
