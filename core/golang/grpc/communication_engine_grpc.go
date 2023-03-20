/*
 *  Copyright 2023 NURTURE AGTECH PVT LTD
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

/*
 *  Copyright 2023 NURTURE AGTECH PVT LTD
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package grpc

import (
	ce "github.com/nurture-farm/Contracts/CommunicationEngine/Gen/GoCommunicationEngine"
	"context"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

const (
	COMMUNICATION_ENGINE_GRPC_END_POINT = "communication_engine_grpc_end_point"
)

var logger = getLogger()

func getLogger() *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	return logger
}

type communicationEnginePlatformGrpcClient struct {
	communicationEnginePlatformClient ce.CommunicationEnginePlatformClient
}

var CommunicationEnginePlatformGrpcClient *communicationEnginePlatformGrpcClient

func init() {
	CommunicationEnginePlatformGrpcClient = &communicationEnginePlatformGrpcClient{
		communicationEnginePlatformClient: createCommunicationEnginePlatformGrpcClient(),
	}
}

func createCommunicationEnginePlatformGrpcClient() ce.CommunicationEnginePlatformClient {

	ceURL := viper.GetString(COMMUNICATION_ENGINE_GRPC_END_POINT)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ceURL, grpc.WithInsecure())
	if err != nil {
		logger.Fatal("Could not connect to Event Portal", zap.Error(err))
	}

	return ce.NewCommunicationEnginePlatformClient(conn)
}

func (cep *communicationEnginePlatformGrpcClient) SearchMessageAcknowledgement(ctx context.Context, referenceId string) (*ce.MessageAcknowledgementResponse, error) {

	request := &ce.MessageAcknowledgementRequest{
		ReferenceId: referenceId,
	}
	response, err := cep.communicationEnginePlatformClient.SearchMessageAcknowledgements(ctx, request)
	if err != nil {
		logger.Error("Error in making call to Communication Engine", zap.Error(err), zap.Any("referenceID", referenceId),
			zap.Any("response", response))
		return nil, err
	}
	return response, nil
}
