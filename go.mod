module code.nurture.farm/platform/CampaignService

go 1.14

require (
	contrib.go.opencensus.io/integrations/ocsql v0.1.7
	github.com/aws/aws-sdk-go v1.40.4
	github.com/bits-and-blooms/bloom/v3 v3.3.1
	github.com/facebook/ent v0.5.1
	github.com/go-sql-driver/mysql v1.5.1-0.20200311113236-681ffa848bae
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/nurture-farm/Contracts v0.0.5
	github.com/nurture-farm/go-common v0.0.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/prometheus/client_golang v1.12.1
	github.com/spf13/cast v1.3.0
	github.com/spf13/viper v1.7.0
	go.elastic.co/apm/module/apmgrpc v1.14.0
	go.elastic.co/apm/module/apmsql v1.14.0
	go.temporal.io/api v1.2.0
	go.temporal.io/sdk v1.2.0
	go.uber.org/zap v1.15.0
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	google.golang.org/grpc v1.41.0
	google.golang.org/grpc/examples v0.0.0-20220830002355-1dd025639203 // indirect
	google.golang.org/protobuf v1.27.1
	gopkg.in/confluentinc/confluent-kafka-go.v100 v100.0.0-20190207121235-de05b1069faa
)
