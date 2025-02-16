package service

import (
	"context"
	"net"
	"pcbook/pb"
	"pcbook/sample"
	"pcbook/serializer"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopService, serverAddress := startTestLaptopService(t)

	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()

	expectedID := laptop.Id

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.Id)

	other, err := laptopService.Store.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	requireSameLaptop(t, laptop, other)
}

func startTestLaptopService(t *testing.T) (*LaptopService, string) {
	laptopService := NewLaptopService(NewInMemoryLaptopStore())

	grpcServer := grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer, laptopService)

	listener, err := net.Listen("tcp", ":0")

	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return laptopService, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)

	return pb.NewLaptopServiceClient(conn)
}

func requireSameLaptop(t *testing.T, laptop1 *pb.Laptop, laptop2 *pb.Laptop) {
	json1, err := serializer.ProtobufToJSON(laptop1)

	require.NoError(t, err)

	json2, err := serializer.ProtobufToJSON(laptop2)

	require.NoError(t, err)

	require.Equal(t, json1, json2)
}
