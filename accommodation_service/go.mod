module accommodation_service

replace github.com/Booking-Platform/accommodation-booking-webapp/common => ../common

go 1.20

require (
	github.com/Booking-Platform/accommodation-booking-webapp/common v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.11.6
	google.golang.org/grpc v1.55.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
