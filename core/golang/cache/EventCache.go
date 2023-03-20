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

package cache

import (
	common "github.com/nurture-farm/Contracts/Common/Gen/GoCommon"
	ep "github.com/nurture-farm/Contracts/EventPortal/Gen/GoEventPortal"
	"context"
	"errors"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"sync"
	"time"
)

const (
	EVENT_PORTAL_GRPC_END_POINT = "event_portal_grpc_end_point"
	UNDERSCORE                  = "_"
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

type eventCache struct {
	eventCache        sync.Map
	eventPortalClient ep.EventPortalClient
}

var EventCahce *eventCache

func InitCache() {
	EventCahce = &eventCache{}
	EventCahce.eventPortalClient = createEventPortalGrpcClient()
	populateCache()
	triggerPopulateCache()
}

func triggerPopulateCache() {
	ticker := time.NewTicker(3 * time.Hour)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				populateCache()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func populateCache() {
	EventCahce.populateEventCache()
	logger.Info("Refreshed EventCache!")
}

func createEventPortalGrpcClient() ep.EventPortalClient {

	eventPortlURL := viper.GetString(EVENT_PORTAL_GRPC_END_POINT)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(eventPortlURL, grpc.WithInsecure())
	if err != nil {
		logger.Fatal("Could not connect to Event Portal", zap.Error(err))
	}

	return ep.NewEventPortalClient(conn)
}

func (ec *eventCache) populateEventCache() {

	request := &ep.FilterEventsRequest{
		Limit:      10000,
		PageNumber: 1,
	}

	for enumValue, namespace := range common.NameSpace_name {
		if enumValue == 0 {
			continue
		}
		request.Namespace = common.NameSpace(common.NameSpace_value[namespace])

		response, err := ec.eventPortalClient.ExecuteFilterEvents(context.Background(), request)
		if err != nil {
			logger.Fatal("Error when calling ExecuteFilterEvents", zap.Error(err))
		}
		if response == nil || response.Status.Status != common.RequestStatus_SUCCESS {
			logger.Fatal("Received invalid response from EventPortal", zap.Any("response", response))
		}
		for _, record := range response.Record {
			ec.eventCache.Store(namespace+UNDERSCORE+record.Name, record.Index)
		}
	}
}

func (ec *eventCache) GetEventIndex(namspace string, name string) (int32, error) {

	var index int32
	Iindex, ok := ec.eventCache.Load(namspace + UNDERSCORE + name)
	if !ok {
		return index, errors.New("EVENT_NOT_FOUND")
	}
	index = cast.ToInt32(Iindex)
	return index, nil
}
