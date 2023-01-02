package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/xdot2012/simple-bank/api"
	db "github.com/xdot2012/simple-bank/db/sqlc"
	"github.com/xdot2012/simple-bank/gapi"
	"github.com/xdot2012/simple-bank/pb"
	"github.com/xdot2012/simple-bank/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot Load Configuration:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot Connect to Database", err)
	}

	store := db.NewStore(conn)
	runGrpcServer(config, store)

}

func runGrpcServer(config util.Config, store db.Store) {

	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Could not Start Server")
	}
}
