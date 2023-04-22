package main

import (
	"api_gateway/config"
	"api_gateway/proto/accommodation_reserve_service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	client := accommodation_reserve_service.NewGreeterServiceClient(conn)
	err = accommodation_reserve_service.RegisterGreeterServiceHandlerClient(
}
