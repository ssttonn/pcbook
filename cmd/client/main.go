package main

import (
	"context"
	"flag"
	"log"
	"pcbook/pb"
	"pcbook/sample"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()

	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	laptop.Id = ""

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	res, err := laptopClient.CreateLaptop(context, req)

	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("laptop has already existed")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}

		return
	}

	log.Printf("laptop created successfully with id: %s", res.Id)
}
