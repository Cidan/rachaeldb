package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/Cidan/rachaeldb/api/v1"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
	"google.golang.org/grpc"
)

func main() {
	setupLogging()
	viper.SetDefault("service.port", 5309)
	viper.SetDefault("service.rest", 8080)
	viper.BindEnv("service.rest", "PORT")

	api := pb.New()
	go startGrpc(api)
	go startWeb()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}

func startGrpc(api *pb.API) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("service.port")))
	if err != nil {
		log.Panic().
			Err(err).
			Msg("error listening on gRPC port")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterKVServer(grpcServer, api)
	log.Info().
		Int("port", viper.GetInt("service.port")).
		Msg("Started gRPC server, you're good to go!")

	grpcServer.Serve(lis)
}

func startWeb() {
	mux := runtime.NewServeMux()
	err := pb.RegisterKVHandlerFromEndpoint(
		context.Background(),
		mux,
		fmt.Sprintf("localhost:%d", viper.GetInt("service.port")),
		[]grpc.DialOption{grpc.WithInsecure()})

	if err != nil {
		log.Panic().
			Err(err).
			Msg("error setting up gRPC gateway")
	}
	log.Info().
		Int("port", viper.GetInt("service.rest")).
		Msg("Started REST server, you're good to go!")
	http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("service.rest")), mux)
}

func setupLogging() {
	// If we're in a terminal, pretty print
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		var level zerolog.Level
		switch viper.GetString("level") {
		case "info":
			level = zerolog.InfoLevel
		case "warn":
			level = zerolog.WarnLevel
		case "error":
			level = zerolog.ErrorLevel
		case "debug":
			level = zerolog.DebugLevel
		default:
			level = zerolog.InfoLevel
		}
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).Level(level)
		log.Info().Msg("Detected terminal, pretty logging enabled.")
	}
}
