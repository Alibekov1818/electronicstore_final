package main

import (
	"computers_service/internal/data"
	db "computers_service/internal/db"
	server "computers_service/internal/grpc_server"
	"computers_service/internal/validator"
	"computers_service/pb"
	"computers_service/pkg/jsonlog"
	"context"
	_ "database/sql"
	_ "log"
	"os"
	"testing"
)

var testServer server.Server
var id int64
var size int

func init() {
	db, _ := db.OpenDB()
	testServer = server.Server{
		Logger: jsonlog.New(os.Stdout, jsonlog.LevelInfo),
		Models: data.NewModels(db),
	}

}

func TestValidateComputer(t *testing.T) {
	v := validator.New()

	testComputer := &data.Computer{ID: 10, Model: "Test", Cpu: "8GB", Memory: "500GB", Price: 10000}

	if data.ValidateComputer(v, testComputer); !v.Valid() {
		t.Errorf("Should be valid")
	}

	testComputer = &data.Computer{ID: 123, Model: "Test", Cpu: "8GB", Memory: "500GB", Price: 0}

	if data.ValidateComputer(v, testComputer); v.Valid() {
		t.Errorf("Should be invalid")
	}
}

func TestInsertComputer(t *testing.T) {
	testComputer := &data.Computer{ID: 10, Model: "Test", Cpu: "8GB", Memory: "500GB", Price: 10000}
	err := testServer.Models.Computers.Insert(testComputer)
	id = testComputer.ID
	if err != nil {
		t.Errorf("%v", err.Error())
	}
}

func TestGetComputer(t *testing.T) {
	_, err := testServer.GetComputer(context.Background(), &pb.ComputerId{Id: id})
	if err != nil {
		t.Errorf("%v", err.Error())
	}
}

func TestGetComputersFirst(t *testing.T) {
	list, err := testServer.GetComputers(context.Background(), &pb.GetComputersRequest{})
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	size = len(list.List)
}

func TestDeleteComputer(t *testing.T) {
	err := testServer.Models.Computers.Delete(id)
	if err != nil {
		t.Errorf("%v", err.Error())
	}
}

func TestGetComputersSecond(t *testing.T) {
	list, err := testServer.GetComputers(context.Background(), &pb.GetComputersRequest{})
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	if len(list.List)+1 != size {
		t.Errorf("Get computers error")
	}
}
