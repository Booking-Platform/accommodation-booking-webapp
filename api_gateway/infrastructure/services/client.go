package services

import (
	reservation "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_reserve_service"
	accommodation "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_service"
	auth "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/auth_service"
	rating "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/rating_service"
	user "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/user_info_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewReservationClient(address string) reservation.AccommodationReserveServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return reservation.NewAccommodationReserveServiceClient(conn)
}

func NewAccommodationClient(address string) accommodation.AccommodationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return accommodation.NewAccommodationServiceClient(conn)
}

func NewUserClient(address string) user.UserInfoServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Shipping service: %v", err)
	}
	return user.NewUserInfoServiceClient(conn)
}

func NewAuthClient(adress string) auth.AuthServiceClient {
	conn, err := getConnection(adress)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Shipping service: %v", err)
	}
	return auth.NewAuthServiceClient(conn)
}

func NewRatingClient(adress string) rating.RatingServiceClient {
	conn, err := getConnection(adress)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Shipping service: %v", err)
	}
	return rating.NewRatingServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
