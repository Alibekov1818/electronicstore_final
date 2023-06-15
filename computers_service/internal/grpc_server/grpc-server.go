package server

import (
	"computers_service/internal/data"
	"computers_service/internal/db"
	"computers_service/internal/validator"
	pb "computers_service/pb"
	"computers_service/pkg/config"
	"computers_service/pkg/jsonlog"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"sync"
)

type Server struct {
	pb.UnimplementedComputersServiceServer
	Models data.Models
	Wg     sync.WaitGroup
	Logger *jsonlog.Logger
}

func (s *Server) GetComputer(ctx context.Context, r *pb.ComputerId) (*pb.Computer, error) {
	computer, err := s.Models.Computers.Get(r.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Phone not found")
	}
	return &pb.Computer{Id: computer.ID, Model: computer.Model, Cpu: computer.Cpu, Memory: computer.Memory, Price: computer.Price}, nil
}

func (s *Server) GetComputers(context.Context, *pb.GetComputersRequest) (*pb.ComputerList, error) {
	list, err := s.Models.Computers.SelectAll()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can not get list")
	}
	var computers []*pb.Computer

	for i := 0; i < len(list); i++ {
		computers = append(computers, &pb.Computer{
			Id:     list[i].ID,
			Model:  list[i].Model,
			Cpu:    list[i].Cpu,
			Memory: list[i].Memory,
			Price:  list[i].Price,
		})
	}
	return &pb.ComputerList{List: computers}, nil
}

func (s *Server) CreateComputer(ctx context.Context, r *pb.Computer) (*pb.Computer, error) {
	newComputer := &data.Computer{Model: r.Model, Cpu: r.Cpu, Memory: r.Memory, Price: r.Price}
	v := validator.New()
	if data.ValidateComputer(v, newComputer); !v.Valid() {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid data")
	}
	err := s.Models.Computers.Insert(newComputer)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can't insert data")
	}
	return r, nil
}

func (s *Server) DeleteComputer(ctx context.Context, r *pb.ComputerId) (*pb.Computer, error) {
	err := s.Models.Computers.Delete(r.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can't delete computer or not found")
	}
	return &pb.Computer{Id: r.Id}, nil
}

func Start() {

	conf := config.GetConfig()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := db.OpenDB()

	if err != nil {
		logger.PrintFatal(err, nil)
	}

	logger.PrintInfo("database connection pool established", nil)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))

	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	server := grpc.NewServer()

	reflection.Register(server)

	pb.RegisterComputersServiceServer(server, &Server{
		Logger: logger,
		Models: data.NewModels(db),
	})

	log.Printf("Server listening at %d", conf.Port)

	//server.Serve(lis)
	go serveGoroutine(server, lis)
}

func serveGoroutine(s *grpc.Server, lis net.Listener) {
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
