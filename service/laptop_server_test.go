package service_test

import (
	"context"
	"pcbook/pb"
	"pcbook/sample"
	"pcbook/service"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidID := sample.NewLaptop()
	laptopInvalidID.Id = "Invalid-id"

	laptopDuplicateID := sample.NewLaptop()
	storeDulicateID := service.NewInMemoryLaptopStore()
	err := storeDulicateID.Save(laptopDuplicateID)
	require.Nil(t, err)

	testCases := []struct {
		name   string
		laptop *pb.Laptop
		store  service.LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_no_id",
			laptop: laptopNoID,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure_invalid_id",
			laptop: laptopInvalidID,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failure_exist_id",
			laptop: laptopDuplicateID,
			store:  storeDulicateID,
			code:   codes.AlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			laptopService := service.NewLaptopService(tc.store)

			response, err := laptopService.CreateLaptop(context.Background(), req)

			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, response)
				require.NotEmpty(t, response.Id)
				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, response.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, response)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}
