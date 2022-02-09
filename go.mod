module booking-center-service

go 1.16

replace booking-libs => gitlab.com/booking8/booking-libs v0.0.0-20211129092319-02399c72fee1

replace booking-identity-management => gitlab.com/booking8/booking-identity-management v0.0.0-20220103154609-7c85c51ca421

require (
	booking-identity-management v0.0.0-20220103154609-7c85c51ca421
	booking-libs v0.0.0-20211129092319-02399c72fee1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jinzhu/copier v0.3.2
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/rakyll/statik v0.1.7
	github.com/rs/cors v1.8.0
	github.com/sarulabs/di v2.0.0+incompatible
	github.com/spf13/viper v1.9.0
	go.uber.org/zap v1.17.0
	google.golang.org/genproto v0.0.0-20211118181313-81c1377c94b1
	google.golang.org/grpc v1.42.0
)

require (
	go.mongodb.org/mongo-driver v1.8.0
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/text v0.3.7 // indirect
)
