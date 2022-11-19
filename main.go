package main

import (
	broker2 "InnoTaxi-Driver/internal/broker"
	"InnoTaxi-Driver/internal/config"
	handlers "InnoTaxi-Driver/internal/handler"
	"InnoTaxi-Driver/internal/middleware"
	"InnoTaxi-Driver/internal/repository"
	"InnoTaxi-Driver/internal/service"
	"InnoTaxi-Driver/pkg/api"
	pb "InnoTaxi-Driver/pkg/grpc"
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/segmentio/kafka-go"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

//@title Driver service
//@version 1.0
//description  driver service for taxi app

//@host localhost:8081
//@BasePath /

//@securityDefinitions.apikey SignIn
//@in header
//@name Authorization
func main() {

	cfg, err := config.GetCfg()
	if err != nil {
		log.Fatalf("Error while loading config: %v", err)
	}
	repos, err := repository.NewRepository(cfg)
	if err != nil {
		log.Fatalf("Error while loading repository: %v", err)
	}
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.BROKER_HOST+cfg.BROKER_PORT, "topic", 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	broker := broker2.KafkaClient{Connection: conn}
	service := service.NewService(repos, broker)
	handler := handlers.NewHandler(*service)
	startGrpcServer(*service, cfg)
	startServer(handler, cfg)
	<-wait
}

func startServer(handler handlers.Handler, cfg config.Config) {
	r := gin.Default()
	r.Use(middleware.SetUserStatus())
	r.POST("/register", middleware.EnsureNotLoggedIn(), handler.Register)
	r.POST("/login", middleware.EnsureNotLoggedIn(), handler.Login)
	r.GET("/logout", middleware.EnsureLoggedIn(), handler.Logout)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/setStatus", middleware.EnsureLoggedIn(), handler.SetDriverStatus)
	go func() {
		err := r.Run(cfg.REST_PORT)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}()
}

func startGrpcServer(service service.Service, cfg config.Config) {
	lis, err := net.Listen("tcp", cfg.RPC_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	pb.RegisterDriverServiceServer(server, &api.DriverServer{Service: service})
	go func() {
		err := server.Serve(lis)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}()
}
