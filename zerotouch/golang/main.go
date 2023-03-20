package main

import (
	proto "github.com/nurture-farm/Contracts/CampaignService/Gen/GoCampaignService"
	"github.com/nurture-farm/campaign-service/core/golang/hook"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/database"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/setup"
	"flag"
	"fmt"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.elastic.co/apm/module/apmgrpc"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func runMonitoring(grpcServer *grpc.Server) {
	// register prometheus
	grpc_prometheus.Register(grpcServer)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", "0.0.0.0", 7805), nil)
	if err != nil {
		logger.Panic("Unable to start prometheus handler", zap.Error(err))
	}
}

func registerAsWorker() client.Client {

	w := worker.New(setup.WorkflowClient, database.TemporalTaskQueue, worker.Options{})
	w.RegisterActivity(setup.CampaignServiceActivities)

	logger.Info("Starting CMPSWorker", zap.Any("worker", w))
	workerErr := w.Run(worker.InterruptCh())
	if workerErr != nil {
		logger.Panic("Unable to start activity worker", zap.Error(workerErr))
	}
	return setup.WorkflowClient
}

func main() {

	port := flag.Int("port", 7800, "Port for GRPC server to listen")
	flag.Parse()
	logger.Info("Starting Farm Service Service!")
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		logger.Fatal("Unable to listen on port", zap.Int("port", *port), zap.Error(err))
	}

	grpcServer := grpc.NewServer(
		grpc.ChainStreamInterceptor(grpc_prometheus.StreamServerInterceptor, apmgrpc.NewStreamServerInterceptor()),
		grpc.ChainUnaryInterceptor(grpc_prometheus.UnaryServerInterceptor, apmgrpc.NewUnaryServerInterceptor()))
	proto.RegisterCampaignServiceServer(grpcServer, setup.CampaignService)
	logger.Info("Registered server",
		zap.Any("grpcServer", grpcServer), zap.Any("listener", lis), zap.Int("port", *port))

	// on GRPC services
	go runMonitoring(grpcServer)

	// register worker
	go func() {
		c := registerAsWorker()
		defer c.Close()
	}()
	//initCostControl()
	hook.PreStartUpHook()

	// Start server
	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Fatal("Unable to listen on service", zap.Int("port", *port), zap.Error(err))
	}

	hook.PostStartUpHook()
}

//func initCostControl() {
//	cc.Configure(cc_models.Config{
//		Timeout:                    2000 * time.Millisecond,
//		ApiKey:                     viper.GetString("grafana_api_key"),
//		GrafanaHost:                viper.GetString("grafana_host"),
//		GrafanaRulesDirectory:      viper.GetString("grafana_rules_directory"),
//		AlertStatusRefreshStrategy: 1,
//		AlertStatusTTL:             240 * time.Second,
//	})
//}
