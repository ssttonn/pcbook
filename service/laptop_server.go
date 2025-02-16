package service

import (
	"context"
	"errors"
	"log"
	"pcbook/pb"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopService struct {
	Store LaptopStore
	pb.UnimplementedLaptopServiceServer
}

func NewLaptopService(store LaptopStore) *LaptopService {
	return &LaptopService{
		Store: store,
	}
}

func (service *LaptopService) CreateLaptop(
	ctx context.Context,
	request *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {
	laptop := request.GetLaptop()
	log.Printf("receive a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)

		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Laptop ID is not a valid UUID, %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Can't generate a new UUID for laptop: %v", err)
		}

		laptop.Id = id.String()
	}

	if ctx.Err() == context.Canceled {
		log.Print("request is cancelled")
		return nil, status.Error(codes.Canceled, "context is cancelled")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Print("deadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}

	err := service.Store.Save(laptop)
	if err != nil {
		code := codes.Internal

		if errors.Is(err, ErrAlreadyExist) {
			code = codes.AlreadyExists
		}

		return nil, status.Errorf(code, "Can't save laptop to the store: %v", err)
	}

	log.Printf("Saved laptop with id: %s", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}

	return res, nil
}
