package api

import (
	"InnoTaxi-Driver/internal/service"
	pb "InnoTaxi-Driver/pkg/grpc"
	"context"
)

type DriverServer struct {
	pb.UnimplementedDriverServiceServer
	Service service.Service
}

func (s *DriverServer) GetAllDrivers(ctx context.Context, in *pb.Request) (*pb.AllDrivers, error) {
	drivers, err := s.Service.GetAllDrivers(ctx)
	grpcDrivers := make([]*pb.Driver, 0)

	for _, driver := range drivers {
		grpcDriver := pb.Driver{}
		grpcDriver.Email = driver.Email
		grpcDriver.Phone = driver.Phone
		grpcDriver.Id = driver.Id
		grpcDriver.TaxiType = driver.TaxiType
		grpcDriver.Rating = driver.Rating
		grpcDriver.Status = driver.Status
		grpcDrivers = append(grpcDrivers, &grpcDriver)
	}
	if err != nil {
		return &pb.AllDrivers{}, err
	}
	return &pb.AllDrivers{Drivers: grpcDrivers}, nil
}

func (s *DriverServer) GetAllFreeDrivers(ctx context.Context, in *pb.Request) (*pb.AllFreeDrivers, error) {
	drivers, err := s.Service.GetAllFreeDrivers(ctx)
	grpcDrivers := make([]*pb.Driver, 0)

	for _, driver := range drivers {
		grpcDriver := pb.Driver{}

		grpcDriver.Email = driver.Email
		grpcDriver.Phone = driver.Phone
		grpcDriver.Id = driver.Id
		grpcDriver.TaxiType = driver.TaxiType
		grpcDriver.Rating = driver.Rating
		grpcDriver.Status = driver.Status
		grpcDrivers = append(grpcDrivers, &grpcDriver)
	}
	if err != nil {
		return &pb.AllFreeDrivers{}, err
	}
	return &pb.AllFreeDrivers{Drivers: grpcDrivers}, nil
}
