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

package hook

import (
	Common "github.com/nurture-farm/Contracts/Common/Gen/GoCommon"
	CommunicationEngine "github.com/nurture-farm/Contracts/CommunicationEngine/Gen/GoCommunicationEngine"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/metrics"
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/confluentinc/confluent-kafka-go.v100/kafka"
)

const (
	METRICS_NAME = "NF_CMPS_SEND_MESSAGE"
	HELP_MESSAGE = "Push to kafka metrics metrics"
)

var producer *kafka.Producer

var (
	sendMessageMetrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       METRICS_NAME,
		Help:       HELP_MESSAGE,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	kafkaDeliveryMetrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CMPS_KAFKA_DELIVERY",
		Help:       "Metrics for pushing into kafka",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var lowPriorityTopicNameByChannel map[Common.CommunicationChannel]string

// InitProducer initializes the Kafka producer.
func init() {
	prometheus.MustRegister(sendMessageMetrics)
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":                     viper.GetString("kafka.bootstrap.servers"),
		"acks":                                  cast.ToInt(viper.GetString("kafka.acks")),
		"compression.type":                      viper.GetString("kafka.compression.type"),
		"max.in.flight.requests.per.connection": cast.ToInt(viper.GetString("kafka.max.in.flight.requests.per.connection")),
		//"batch.size": 		   cast.ToInt(viper.GetString("kafka.batch.size")),
		"linger.ms": cast.ToInt(viper.GetString("kafka.linger.ms")),
		//"key.serializer":       viper.GetString("kafka.key.serializer"),
		//"value.serializer": viper.GetString("kafka.value.serializer"),
	})

	if err != nil {
		logger.Panic("Failed to create kafka producer", zap.Error(err))
	}

	producer = p
	initLowPriorityTopicNameByChannelTypeMap()
	go sendMessageCallback(context.Background())
}

func SendMessage(ctx context.Context, events *CommunicationEngine.BulkCommunicationEvent) error {

	for _, event := range events.CommunicationEvents {
		err := pushToKafka(ctx, event)
		if err != nil {
			return err
		}
	}
	return nil
}

func pushToKafka(ctx context.Context, event *CommunicationEngine.CommunicationEvent) error {

	var refID string
	if event.ReferenceId == "" {
		refID = uuid.New().String()
	} else {
		refID = event.ReferenceId
	}
	event.ReferenceId = refID
	if err := sendMessageToKafka(ctx, event); err != nil {
		return err
	}

	return nil
}

func sendMessageToKafka(ctx context.Context, event *CommunicationEngine.CommunicationEvent) error {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(sendMessageMetrics, "SendMessageToKafka", &err, ctx)
	//logger.Info("Sending communication", zap.Any("CommunicationEvent", event))
	topic := lowPriorityTopicNameByChannel[event.GetChannel()[0]]
	if len(topic) == 0 {
		topic = viper.GetString("communication.event.topics.default")
	}
	valueBytes, err := proto.Marshal(event)
	if err != nil {
		logger.Error("Failed to marshal event value", zap.Error(err))
		return err
	}

	var keyBytes []byte
	if event.ReceiverActor != nil {
		keyBytes, err = proto.Marshal(event.ReceiverActor)
		if err != nil {
			logger.Error("Failed to marshal event key", zap.Error(err))
			return err
		}
		//keyBytes = []byte(cast.ToString(event.ReceiverActor.ActorId))
	} else {
		keyBytes, err = proto.Marshal(event.ReceiverActorDetails)
		if err != nil {
			logger.Error("Failed to marshal event key", zap.Error(err))
			return err
		}
		//if event.ReceiverActorDetails.MobileNumber != "" {
		//	keyBytes = []byte(cast.ToString(event.ReceiverActorDetails.MobileNumber))
		//} else if event.ReceiverActorDetails.FcmToken != "" {
		//	keyBytes = []byte(cast.ToString(event.ReceiverActorDetails.FcmToken))
		//} else if event.ReceiverActorDetails.EmailId != "" {
		//	keyBytes = []byte(cast.ToString(event.ReceiverActorDetails.EmailId))
		//}
	}

	message := kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          valueBytes,
		Key:            keyBytes,
	}
	logger.Info("Event", zap.Any("event", event))
	producer.ProduceChannel() <- &message

	return nil
}

func initLowPriorityTopicNameByChannelTypeMap() {
	lowPriorityTopicNameByChannel = make(map[Common.CommunicationChannel]string)
	lowPriorityTopicNameByChannel[Common.CommunicationChannel_NO_CHANNEL] = viper.GetString("communication.event.topics.default")
	lowPriorityTopicNameByChannel[Common.CommunicationChannel_SMS] = viper.GetString("communication.event.topics.sms.low.priority")
	lowPriorityTopicNameByChannel[Common.CommunicationChannel_EMAIL] = viper.GetString("communication.event.topics.email.low.priority")
	lowPriorityTopicNameByChannel[Common.CommunicationChannel_APP_NOTIFICATION] = viper.GetString("communication.event.topics.pn.low.priority")
	lowPriorityTopicNameByChannel[Common.CommunicationChannel_WHATSAPP] = viper.GetString("communication.event.topics.whatsapp.low.priority")
}

func sendMessageCallback(ctx context.Context) {
	prometheus.MustRegister(kafkaDeliveryMetrics)
	for {
		for e := range producer.Events() {
			func() {

				var err error
				defer metrics.Metrics.PushToSummarytMetrics()(kafkaDeliveryMetrics, "SendMessageCallbackKafka", &err, ctx)

				switch ev := e.(type) {
				case *kafka.Message:
					m := ev
					if m.TopicPartition.Error != nil {
						logger.Error("Kafka producer failed to deliver message : ", zap.String("topic", *m.TopicPartition.Topic),
							zap.String("message_key", string(m.Key)), zap.Error(m.TopicPartition.Error))
					} else {
						logger.Info("Kafka producer successfully delivered message")
						//zap.String("topic", *m.TopicPartition.Topic),
						//zap.Int32("partition", m.TopicPartition.Partition), zap.Any("offset", m.TopicPartition.Offset))
					}

				default:
					logger.Error("Ignored event: ", zap.Any("event", ev))
				}
			}()
		}
	}
}
