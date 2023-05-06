module github.com/Booking-Platform/accommodation-booking-webapp/common

go 1.19

replace github.com/Booking-Platform/accommodation-booking-webapp/common => ../common

require (
	github.com/golang/protobuf v1.5.3
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.9.0
	github.com/nats-io/nats.go v1.13.1-0.20220308171302-2f2f6968e98d
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.30.0
)

require (
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
)
