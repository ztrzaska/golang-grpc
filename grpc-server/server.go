package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-server/api"
	pb "grpc-server/api/product/v1"
	"net"
	"os"
	"os/signal"
)

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	errCh := make(chan error)
	port := 9001

	go func() {
		errCh <- StartServer(port)
	}()

	select {
	case <-sigCh:
		log.Info().Msg("interrupt signal - goodbye :)")
		return
	case err := <-errCh:
		if err != nil {
			log.Error().Err(err).Msgf("error starting server on %d", port)
			log.Fatal().Msg("cannot start application. Exiting.")
		}
	}
}

func StartServer(port int) error {
	log.Info().Msg("starting new grpc server")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal().Err(err).Msgf("error starting server on %d", port)
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterServiceServer(grpcServer, &api.Server{})
	reflection.Register(grpcServer)
	log.Printf("server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal().Err(err).Msg("can not start grpc server")
		return err
	}
	return nil
}
