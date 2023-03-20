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

package aws

import (
	"github.com/nurture-farm/campaign-service/zerotouch/golang/metrics"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
	"time"
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

var (
	atheClient *athena.Athena
)

func init() {
	once.Do(func() {
		region := CONST_AWS_REGION
		awscfg := &aws.Config{
			Region:      &region,
			Credentials: credentials.NewStaticCredentials(viper.GetString("AWS_ACCESS_KEY"), viper.GetString("AWS_SECRET_KEY"), ""),
		}
		sess := session.Must(session.NewSession(awscfg))
		atheClient = athena.New(sess, awscfg)
		svc = s3.New(sess)
	})
}

func ExecuteAthenaQuery(ctx context.Context, query string, campaignId int64) (map[string]string, [][]string, error) {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.AthenaQuery_Metrics, "ExecuteAthenaQuery", &err, ctx)
	logger.Info("Executing Athena Query", zap.Any("query", query), zap.Any("campaignId", campaignId))
	athenaResponse, err := RunAthenaQuery(query, campaignId)
	if err != nil {
		logger.Error("Error in Running AthenaQuery", zap.Error(err), zap.Any("campaignId", campaignId), zap.Any("query", query))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.ATHENAQUERY_Error_Metrics, err, ctx)
		return nil, nil, err
	}
	columnMap, rowData := MapAthenaResponse(athenaResponse)
	return columnMap, rowData, nil
}

func RunAthenaQuery(query string, campaignId int64) (*athena.GetQueryResultsOutput, error) {

	database := CONST_DATABASE_REWARDS_GATEWAY
	logger.Info("Executing Athena query", zap.Any("campaignId", campaignId), zap.Any("database", database))
	var s athena.StartQueryExecutionInput
	s.QueryString = &query

	var q athena.QueryExecutionContext
	q.SetDatabase(database)
	s.SetQueryExecutionContext(&q)

	result, err := atheClient.StartQueryExecution(&s)
	if err != nil {
		logger.Error("Error in StartQueryExecution for Athena query", zap.Error(err), zap.Any("campaignId", campaignId), zap.Any("database", database))
		return nil, err
	}

	var qri athena.GetQueryExecutionInput
	qri.SetQueryExecutionId(*result.QueryExecutionId)

	var qrop *athena.GetQueryExecutionOutput
	duration := time.Duration(CONST_ATHENA_QUERY_SLEEP_DURATION) * time.Second

	for {
		qrop, err = atheClient.GetQueryExecution(&qri)
		if err != nil {
			logger.Error("Error in GetQueryExecution for Athena query", zap.Error(err), zap.Any("campaignId", campaignId), zap.Any("database", database))
			return nil, err
		}
		if *qrop.QueryExecution.Status.State == CONST_ATHENA_QUERY_SUCCESS {
			break
		}
		if *qrop.QueryExecution.Status.State == CONST_ATHENA_QUERY_FAILED {
			err = fmt.Errorf("ATHENA_QUERY_FAILED")
			logger.Error("Athena query execution failed", zap.Any("campaignId", campaignId), zap.Any("database", database))
			return nil, err
		}
		logger.Info("Waiting for Athena query execution to finish", zap.Any("campaignId", campaignId), zap.Any("database", database),
			zap.Any("state", *qrop.QueryExecution.Status.State))
		time.Sleep(duration)
	}

	logger.Info("Athena query execution completed", zap.Any("campaignId", campaignId), zap.Any("database", database))
	var ip athena.GetQueryResultsInput
	ip.SetQueryExecutionId(*result.QueryExecutionId)
	logger.Info("Getting Athena query result", zap.Any("campaignId", campaignId), zap.Any("database", database))
	op, err := atheClient.GetQueryResults(&ip)
	if err != nil {
		logger.Error("Error in GetQueryResults for Athena query", zap.Error(err), zap.Any("campaignId", campaignId), zap.Any("database", database))
		return nil, err
	}
	logger.Info("Athena query result success", zap.Any("campaignId", campaignId), zap.Any("database", database))
	nextToken := op.NextToken
	for nextToken != nil {
		logger.Info("Proccessing Athena query result from next Token", zap.Any("campaignId", campaignId), zap.Any("database", database))
		queryResults, _ := atheClient.GetQueryResults(&athena.GetQueryResultsInput{
			NextToken:        nextToken,
			QueryExecutionId: result.QueryExecutionId,
		})
		nextToken = queryResults.NextToken
		op.ResultSet.Rows = append(op.ResultSet.Rows, queryResults.ResultSet.Rows...)
		logger.Info("Total Number of rows received ", zap.Any("Rows", len(op.ResultSet.Rows)))
	}
	logger.Info("Total number of rows returned in query ", zap.Any("Rows", len(op.ResultSet.Rows)-1))
	logger.Info("Athena query execution successful!", zap.Any("campaignId", campaignId), zap.Any("database", database))
	return op, nil
}

func MapAthenaResponse(response *athena.GetQueryResultsOutput) (map[string]string, [][]string) {

	columnMap := make(map[string]string)
	rowsData := [][]string{}
	dataStartIndex := 0
	for _, columnInfo := range response.ResultSet.ResultSetMetadata.ColumnInfo {
		name := *columnInfo.Name
		if name != "" && strings.HasPrefix(name, "_") {
			dataStartIndex++
			continue
		}
		columnMap[name] = *columnInfo.Type
	}
	for _, row := range response.ResultSet.Rows {
		rowData := []string{}
		for index, data := range row.Data {
			value := data.VarCharValue
			if index >= dataStartIndex {
				if value != nil {
					rowData = append(rowData, *data.VarCharValue)
				} else {
					rowData = append(rowData, CONST_NIL)
				}
			}
		}
		rowsData = append(rowsData, rowData)
	}
	return columnMap, rowsData
}
