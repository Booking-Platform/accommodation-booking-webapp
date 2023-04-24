module github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service

go 1.19

replace github.com/Booking-Platform/accommodation-booking-webapp/common => ../common

require (
	github.com/Booking-Platform/accommodation-booking-webapp/common v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.11.4
	google.golang.org/grpc v1.54.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230223222841-637eb2293923 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
