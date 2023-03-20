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
	fs "github.com/nurture-farm/Contracts/CampaignService/Gen/GoCampaignService"
	Common "github.com/nurture-farm/Contracts/Common/Gen/GoCommon"
	query "github.com/nurture-farm/campaign-service/core/golang/database"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/database/executor"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/database/mappers"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/database/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/facebook/ent/dialect"
	entsql "github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type ScheduleUserJourneyCampaignInterface interface {
	OnRequest(ctx context.Context, request *fs.ScheduleUserJourneyCampaignRequest) *fs.ScheduleUserJourneyCampaignResponse
	OnData(ctx context.Context, request *fs.ScheduleUserJourneyCampaignRequest, response *fs.ScheduleUserJourneyCampaignResponse) *fs.ScheduleUserJourneyCampaignResponse
	OnError(ctx context.Context, request *fs.ScheduleUserJourneyCampaignRequest, response *fs.ScheduleUserJourneyCampaignResponse, err error) *fs.ScheduleUserJourneyCampaignResponse
	OnResponse(ctx context.Context, request *fs.ScheduleUserJourneyCampaignRequest, response *fs.ScheduleUserJourneyCampaignResponse) *fs.ScheduleUserJourneyCampaignResponse
}

type GenericScheduleUserJourneyCampaignExecutor struct {
	ScheduleUserJourneyCampaignInterface ScheduleUserJourneyCampaignInterface
}

type ScheduleUserJourneyCampaignController struct {
}

var ScheduleUserJourneyCampaignExecutor *GenericScheduleUserJourneyCampaignExecutor

func (ge *GenericScheduleUserJourneyCampaignExecutor) OnRequest(ctx context.Context, request *fs.ScheduleUserJourneyCampaignRequest) *fs.ScheduleUserJourneyCampaignResponse {
	return ge.ScheduleUserJourneyCampaignInterface.OnRequest(ctx, request)
}

func (ge *GenericScheduleUserJourneyCampaignExecutor) OnResponse(ctx context.Context, request *fs.ScheduleUserJourneyCampaignRequest, response *fs.ScheduleUserJourneyCampaignResponse) *fs.ScheduleUserJourneyCampaignResponse {
	return ge.ScheduleUserJourneyCampaignInterface.OnResponse(ctx, request, response)
}

func (ge *GenericScheduleUserJourneyCampaignExecutor) OnData(ctx context.Context, request *fs.ScheduleUserJourneyCampaignRequest, response *fs.ScheduleUserJourneyCampaignResponse) *fs.ScheduleUserJourneyCampaignResponse {
	return ge.ScheduleUserJourneyCampaignInterface.OnData(ctx, request, response)
}

func (ge *GenericScheduleUserJourneyCampaignExecutor) OnError(ctx context.Context, request *fs.ScheduleUserJourneyCampaignRequest, response *fs.ScheduleUserJourneyCampaignResponse, err error) *fs.ScheduleUserJourneyCampaignResponse {
	return ge.ScheduleUserJourneyCampaignInterface.OnError(ctx, request, response, err)
}

func (rc *ScheduleUserJourneyCampaignController) OnRequest(ctx context.Context, request *fs.ScheduleUserJourneyCampaignRequest) *fs.ScheduleUserJourneyCampaignResponse {

	var engagementVertexId, campaignId int64
	err := validateScheduleUserJourneyCampaignRequest(request)
	if err != nil {
		return mapToScheduleUserJourneyCampaignResponse(request, Common.RequestStatus_BAD_INPUT, err, engagementVertexId, EMPTY, EMPTY, campaignId)
	}
	cronExpression, err := getUserJourneyCampaignCronSchedule(ctx, request)
	if err != nil {
		return mapToScheduleUserJourneyCampaignResponse(request, Common.RequestStatus_BAD_INPUT, err, engagementVertexId, EMPTY, EMPTY, campaignId)
	}
	findCampaignByIdResponse, err := getCampaign(ctx, request)
	if err != nil {
		return mapToScheduleUserJourneyCampaignResponse(request, Common.RequestStatus_BAD_INPUT, err, engagementVertexId, EMPTY, EMPTY, campaignId)
	}
	var status, dbCronExpression, metadata string
	if findCampaignByIdResponse != nil {
		status = findCampaignByIdResponse.Records.Status
		dbCronExpression = findCampaignByIdResponse.Records.CronExpression
		metadata = findCampaignByIdResponse.Records.Attributes
	}
	if request.CampaignId != 0 && status != CONST_DRAFTED {
		if !request.TriggerCampaign && status == "RUNNING" {
			err = updateUserJourneyCampaignStatus(ctx, request.CampaignId, Common.CampaignStatus_HALTED)
			if err != nil {
				return mapToScheduleUserJourneyCampaignResponse(request, Common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, EMPTY, EMPTY, request.CampaignId)
			}
			return mapToScheduleUserJourneyCampaignResponse(request, Common.RequestStatus_SUCCESS, nil, engagementVertexId, EMPTY, EMPTY, request.CampaignId)
		}
		if request.TriggerCampaign && status == "HALTED" {
			err = updateUserJourneyCampaignStatus(ctx, request.CampaignId, Common.CampaignStatus_RUNNING)
			if err != nil {
				return mapToScheduleUserJourneyCampaignResponse(request, Common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, EMPTY, EMPTY, request.CampaignId)
			}
			return mapToScheduleUserJourneyCampaignResponse(request, Common.RequestStatus_SUCCESS, nil, engagementVertexId, dbCronExpression, uuid.New().String(), request.CampaignId)
		}
		return mapToScheduleUserJourneyCampaignResponse(request, Common.RequestStatus_REQUEST_NOT_FULFILLED, fmt.Errorf("CAMPAIGN_ALREADY_RUNNING"), engagementVertexId, EMPTY, EMPTY, campaignId)
	}
	if cronExpression == EMPTY && dbCronExpression != EMPTY {
		cronExpression = dbCronExpression
	}
	if cronExpression == EMPTY && request.TriggerCampaign &&
		(request.Campaign == nil || request.Campaign.UserJourneys == nil || len(request.Campaign.UserJourneys) == 0) {
		return mapToScheduleUserJourneyCampaignResponse(request, Common.RequestStatus_REQUEST_NOT_FULFILLED, fmt.Errorf("MISSING_USER_JOURNEY"), engagementVertexId, EMPTY, EMPTY, campaignId)
	}
	addCampaignReq, userJourneyMetadata, engagementMetadata, err := makeAddUserJourneyCampaignRequest(ctx, request, cronExpression, metadata)
	if err != nil {
		return mapToScheduleUserJourneyCampaignResponse(request, Common.RequestStatus_BAD_INPUT, err, engagementVertexId, EMPTY, EMPTY, campaignId)
	}
	addNewCampaignReq := &fs.AddNewCampaignRequest{
		AddCampaignRequest: addCampaignReq,
	}
	_, txErr := executor.Driver.TransactionRunner(ctx, "OnRequestScheduleUserJourneyCampaign", func(ctx context.Context, txName string, tx dialect.Tx) (res executor.TransactionResult, err error) {

		err = invalidatePreviousUserJourneyCampaignData(ctx, tx, request)
		if err != nil {
			return nil, err
		}
		engagementVertexId, campaignId, err = scheduleUserJourneyCampaign(ctx, tx, request, addCampaignReq, addNewCampaignReq, userJourneyMetadata, engagementMetadata)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	if txErr != nil {
		logger.Error("ScheduleUserJourneyCampaignController OnRequest hook, transaction error",
			zap.Error(txErr), zap.Any("request", request))
		return mapToScheduleUserJourneyCampaignResponse(request, Common.RequestStatus_INTERNAL_ERROR, txErr, engagementVertexId, EMPTY, EMPTY, campaignId)
	}
	return mapToScheduleUserJourneyCampaignResponse(request, Common.RequestStatus_SUCCESS, nil, engagementVertexId, cronExpression, uuid.New().String(), campaignId)
}

func updateUserJourneyCampaignStatus(ctx context.Context, campaignId int64, status Common.CampaignStatus) error {

	updateCampaignReq := &fs.UpdateCampaignRequest{
		Id: campaignId,
		AddCampaignRequest: &fs.AddCampaignRequest{
			Status: status,
		},
		UpdatedByActor: &Common.ActorID{
			ActorType: Common.ActorType_SYSTEM,
			ActorId:   1,
		},
	}
	findCampaignByIdReq := &fs.FindCampaignByIdRequest{
		Id: campaignId,
	}
	_, err := updateCampaignTx(ctx, updateCampaignReq, findCampaignByIdReq)
	return err
}

func getCampaign(ctx context.Context, request *fs.ScheduleUserJourneyCampaignRequest) (*fs.FindCampaignByIdResponse, error) {

	if request.CampaignId == 0 {
		return nil, nil
	}
	findCampaignByIdRequest := &fs.FindCampaignByIdRequest{
		Id: request.CampaignId,
	}
	findCampaignByIdResponse, err := executor.RequestExecutor.ExecuteFindCampaignById(ctx, findCampaignByIdRequest)
	if err != nil {
		return nil, err
	}
	if findCampaignByIdResponse == nil || findCampaignByIdResponse.Records == nil {
		return nil, nil
	}
	return findCampaignByIdResponse, nil
}

func getUserJourneyCampaignCronSchedule(ctx context.Context, request *fs.ScheduleUserJourneyCampaignRequest) (cronSchedule string, err error) {

	defer func() {
		if r := recover(); r != nil {
			logger.Error("ScheduleUserJourneyCampaignController OnRequest hook, Error in calculating cron schedule",
				zap.Any("request", request))
			err = errors.New("GO_PANIC_CRON_SCHEDULE")
		}
	}()
	if request.Campaign == nil || request.Campaign.UserJourneys == nil || len(request.Campaign.UserJourneys) == 0 {
		return
	}
	waitTime, err := convertToWaitTime(request.Campaign.UserJourneys[0].UserJourneyVertex.Edge.WaitTime.WaitFor)
	if err != nil {
		return
	}
	cronSchedule = "*/" + cast.ToString(waitTime/60) + " * * * *"
	return
}

func (rc *ScheduleUserJourneyCampaignController) OnResponse(ctx context.Context, request *fs.ScheduleUserJourneyCampaignRequest, response *fs.ScheduleUserJourneyCampaignResponse) *fs.ScheduleUserJourneyCampaignResponse {
	return nil
}

func (rc *ScheduleUserJourneyCampaignController) OnData(ctx context.Context, request *fs.ScheduleUserJourneyCampaignRequest, response *fs.ScheduleUserJourneyCampaignResponse) *fs.ScheduleUserJourneyCampaignResponse {
	return nil
}

func (rc *ScheduleUserJourneyCampaignController) OnError(ctx context.Context, request *fs.ScheduleUserJourneyCampaignRequest, response *fs.ScheduleUserJourneyCampaignResponse, err error) *fs.ScheduleUserJourneyCampaignResponse {
	return nil
}

func invalidatePreviousUserJourneyCampaignData(ctx context.Context, tx dialect.Tx, request *fs.ScheduleUserJourneyCampaignRequest) error {

	var err error
	campaignId := request.CampaignId
	if request.CampaignId == 0 {
		return nil
	}
	if request.Campaign != nil && request.Campaign.UserJourneys != nil && len(request.Campaign.UserJourneys) > 0 {
		err = deleteUserJourneyVertices(ctx, tx, campaignId)
		if err != nil {
			return err
		}
		err = deleteEdges(ctx, tx, campaignId, "USER_JOURNEY")
		if err != nil {
			return err
		}
	}
	if request.Campaign != nil && request.Campaign.EngagementStartVertex != nil {
		err = deleteEngagementVertices(ctx, tx, campaignId)
		if err != nil {
			return err
		}
		err = deleteEdges(ctx, tx, campaignId, "ENGAGEMENT")
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteEngagementVertices(ctx context.Context, tx dialect.Tx, campaignId int64) error {

	model := MakeDeleteEngagementVerticesVO(campaignId)
	args := DeleteEngagementVerticesArgs(model)
	var rows sql.Result
	err := tx.Exec(ctx, query.QUERY_DeleteEngagementVerticesByCampaignId, args, &rows)
	if err != nil {
		logger.Error("ScheduleUserJourneyCampaignController OnRequest hook, Error could not delete EngagementVertices",
			zap.Error(err), zap.Any("campaignId", campaignId))
		return err
	}
	return nil
}

func deleteUserJourneyVertices(ctx context.Context, tx dialect.Tx, campaignId int64) error {

	model := MakeDeletUserJourneyVerticesVO(campaignId)
	args := DeleteUserJourneyVerticesArgs(model)
	var rows sql.Result
	err := tx.Exec(ctx, query.QUERY_DeleteUserJourneyVerticesByCampaignId, args, &rows)
	if err != nil {
		logger.Error("ScheduleUserJourneyCampaignController OnRequest hook, Error could not delete UserJourneyVertices",
			zap.Error(err), zap.Any("campaignId", campaignId))
		return err
	}
	return nil
}

func deleteEdges(ctx context.Context, tx dialect.Tx, campaignId int64, vertexType string) error {

	model := MakeDeleteEdgesVO(campaignId, vertexType)
	args := DeleteEdgesVOArgs(model)
	var rows sql.Result
	err := tx.Exec(ctx, query.QUERY_DeleteEdgesByCampaignId, args, &rows)
	if err != nil {
		logger.Error("ScheduleUserJourneyCampaignController OnRequest hook, Error could not delete Edges",
			zap.Error(err), zap.Any("campaignId", campaignId))
		return err
	}
	return nil
}

func scheduleUserJourneyCampaign(ctx context.Context, tx dialect.Tx,
	request *fs.ScheduleUserJourneyCampaignRequest, addCampaignReq *fs.AddCampaignRequest, addNewCampaignReq *fs.AddNewCampaignRequest,
	userJourneyMetadata string, engagementMetadata string) (int64, int64, error) {

	var engagementVertexId, campaignId int64

	campaignId, err := addCampaignEntry(ctx, tx, request, addCampaignReq, addNewCampaignReq, userJourneyMetadata, engagementMetadata)
	if err != nil {
		OutputLog("ScheduleUserJourneyCampaignController OnRequest hook, Error could not execute AddCampaign entry", CONST_LOG_LEVEL_ERRROR, err, addNewCampaignReq)
		return engagementVertexId, campaignId, err
	}
	err = addUserJourneys(ctx, tx, request, campaignId)
	if err != nil {
		OutputLog("ScheduleUserJourneyCampaignController OnRequest hook, Error could not add user journeys", CONST_LOG_LEVEL_ERRROR, err, addNewCampaignReq)
		return engagementVertexId, campaignId, err
	}
	engagementVertexId, err = addEngagementJourney(ctx, tx, request, campaignId)
	if err != nil {
		return engagementVertexId, campaignId, err
	}
	//TO-DO: Add perfix operator processing
	return engagementVertexId, campaignId, nil
}

func GetStartEngagementVertex(ctx context.Context, campaignId int64) (int64, error) {

	var engagementVertexId int64
	var rows = entsql.Rows{}
	args := FindEngagementStartVertexByIdArgs(campaignId)
	query := query.QUERY_FindStartEngagementVertexId

	err := executor.Driver.GetDriver().Query(ctx, query, args, &rows)
	if err != nil {
		logger.Error("ScheduleUserJourneyCampaignController, Error could not FindEngagementStartVertexById", zap.Error(err))
		return engagementVertexId, err
	}
	for rows.Next() {
		model := FindEngagementStartVertexVO{}
		err := rows.Scan(&model.ID)
		if err != nil {
			logger.Error("ScheduleUserJourneyCampaignController, Error while fetching rows for FindEngagementStartVertexById", zap.Error(err))
			return engagementVertexId, err
		}
		engagementVertexId = model.ID.Int64
	}
	return engagementVertexId, nil

}

func addUserJourneys(ctx context.Context, tx dialect.Tx,
	request *fs.ScheduleUserJourneyCampaignRequest, campaignId int64) error {

	if request.Campaign == nil || request.Campaign.UserJourneys == nil || len(request.Campaign.UserJourneys) == 0 {
		return nil
	}

	var prevVertexId int64
	var prevEdge *fs.UserJourneyEdge
	var rows sql.Result
	operators := []Common.LogicalOperator{}
	for _, userJourney := range request.Campaign.UserJourneys {
		prevVertexId = 0
		operators = append(operators, userJourney.Operator)
		var node interface{}
		node = userJourney.UserJourneyVertex
	loop:
		for {
			switch node.(type) {
			case *fs.UserJourneyVertex:
				currVertex := node.(*fs.UserJourneyVertex)
				if currVertex == nil {
					break loop
				}
				userJourneyVertexModel := makeAddUserJourneyVertexVO(currVertex, campaignId)
				args := AddUserJourneyVertexArgs(userJourneyVertexModel)
				var rows sql.Result
				err := tx.Exec(ctx, query.QUERY_AddUserJourneyVertex, args, &rows)
				if err != nil {
					logger.Error("ScheduleUserJourneyCampaignController OnRequest hook, Error could not execute AddUserJourneyVertex request",
						zap.Error(err), zap.Any("request", request), zap.Any("model", userJourneyVertexModel))
					return err
				}
				vertexId, err := rows.LastInsertId()
				if err != nil {
					logger.Error("ScheduleUserJourneyCampaignController OnRequest hook, Error in getting last insertedId",
						zap.Error(err), zap.Any("request", request), zap.Any("model", userJourneyVertexModel))
					return err
				}
				if prevVertexId != 0 {
					userJourneyEdgeModel := makeAddEdgeVO(campaignId, "USER_JOURNEY", prevVertexId, vertexId, prevEdge.WaitTime, "WAIT_FOR", nil)
					edgeArgs := AddEdgeArgs(userJourneyEdgeModel)
					err = tx.Exec(ctx, query.QUERY_AddEdge, edgeArgs, &rows)
					if err != nil {
						logger.Error("ScheduleUserJourneyCampaignController OnRequest hook, Error could not execute AddUserJourneyEdge request",
							zap.Error(err), zap.Any("request", request), zap.Any("model", userJourneyEdgeModel))
						return err
					}
				}
				prevVertexId = vertexId
				node = currVertex.Edge
			case *fs.UserJourneyEdge:
				edge := node.(*fs.UserJourneyEdge)
				if edge == nil || edge.EdgeType == Common.CampaignEdgeType_EXIT {
					userJourneyEdgeModel := makeAddEdgeVO(campaignId, "USER_JOURNEY", prevVertexId, 0, nil, "WAIT_FOR", nil)
					edgeArgs := AddEdgeArgs(userJourneyEdgeModel)
					err := tx.Exec(ctx, query.QUERY_AddEdge, edgeArgs, &rows)
					if err != nil {
						logger.Error("ScheduleUserJourneyCampaignController OnRequest hook, Error could not execute AddUserJourneyEdge request",
							zap.Error(err), zap.Any("request", request), zap.Any("model", userJourneyEdgeModel))
						return err
					}
					break loop
				}
				prevEdge = edge
				node = edge.UserJourneyVertex
			}
		}
	}
	return nil
}

func addCampaignEntry(ctx context.Context, tx dialect.Tx,
	request *fs.ScheduleUserJourneyCampaignRequest, addCampaignReq *fs.AddCampaignRequest, addNewCampaignReq *fs.AddNewCampaignRequest,
	userJourneyMetadata string, engagementMetadata string) (int64, error) {

	var campaignId int64
	userJourneys := []*fs.UserJourney{}
	userJourneys = request.Campaign.UserJourneys
	userMetadataList := []*fs.UserMetadata{}
	for index := 0; index < len(userJourneys); index++ {
		userMetadataList = append(userMetadataList, userJourneys[index].UserMetadata)
	}
	if request.CampaignId != 0 {
		err := updateUserJourneyCampaign(ctx, tx, request, addCampaignReq, userJourneyMetadata, engagementMetadata, userMetadataList)
		if err != nil {
			return campaignId, err
		}
		return request.CampaignId, nil
	}
	model := mappers.MakeAddCampaignRequestVO(addCampaignReq, userJourneyMetadata, engagementMetadata, userMetadataList)
	args := executor.AddCampaignArgs(model)
	var rows sql.Result
	err := tx.Exec(ctx, query.QUERY_AddCampaign, args, &rows)
	if err != nil {
		logger.Error("ScheduleUserJourneyCampaignController OnRequest hook, Error could not execute AddCampaign request",
			zap.Error(err), zap.Any("addNewCampaignReq", addNewCampaignReq))
		return campaignId, err
	}
	campaignId, err = rows.LastInsertId()
	if err != nil {
		logger.Error("ScheduleUserJourneyCampaignController OnRequest hook, Error could not get lastInsertedId for AddCampaign",
			zap.Error(err), zap.Any("aaddNewCampaignReq", addNewCampaignReq))
		return campaignId, err
	}
	return campaignId, nil
}

func updateUserJourneyCampaign(ctx context.Context, tx dialect.Tx, request *fs.ScheduleUserJourneyCampaignRequest,
	addCampaignReq *fs.AddCampaignRequest, userJourneyMetadata string, engagementMetadata string, userMetadataList []*fs.UserMetadata) error {

	var rows = entsql.Rows{}
	args := executor.FindCampaignByIdArgs(&fs.FindCampaignByIdRequest{
		Id: request.CampaignId,
	})
	err := tx.Query(ctx, query.QUERY_FindCampaignById, args, &rows)
	if err != nil {
		logger.Error("Error could not ExecuteFindCampaignByIdRequest", zap.Error(err))
		return err
	}
	model := models.FindCampaignByIdResponseVO{}
	for rows.Next() {
		err := rows.Scan(&model.Id, &model.Namespace, &model.Name, &model.Description, &model.CronExpression, &model.Occurrences, &model.CommunicationChannel, &model.Status, &model.Type, &model.ScheduleType, &model.Query, &model.InactionQuery, &model.InactionDuration, &model.Attributes, &model.CreatedByActorid, &model.CreatedByActortype, &model.UpdatedByActorid, &model.UpdatedByActortype, &model.Version, &model.CreatedAt, &model.UpdatedAt, &model.DeletedAt)
		if err != nil {
			logger.Error("Error while fetching rows for ExecuteFindCampaignByIdRequest", zap.Error(err))
			return err
		}
	}
	cronExpression := model.CronExpression.String
	if addCampaignReq.CronExpression != EMPTY && addCampaignReq.CronExpression != cronExpression {
		cronExpression = addCampaignReq.CronExpression
	}
	status := model.Status.String
	if addCampaignReq.Status != Common.CampaignStatus_NO_CAMPAGIN_STATUS && addCampaignReq.Status.String() != status {
		status = addCampaignReq.Status.String()
	}

	updateModel := makeUpdateCampaignVO(addCampaignReq, cronExpression, status, request.CampaignId, userJourneyMetadata, engagementMetadata, userMetadataList)
	args = UpdateCampaignArgs(updateModel)
	var result sql.Result
	err = tx.Exec(ctx, query.QUERY_UpdateCampaign_CronExpression_Status_MetaData, args, &result)
	if err != nil {
		logger.Error("ScheduleUserJourneyCampaignController OnRequest hook, Error could not execute Update request",
			zap.Error(err), zap.Any("addCampaignReq", addCampaignReq))
		return err
	}
	return nil
}

func addEngagementJourney(ctx context.Context, tx dialect.Tx, request *fs.ScheduleUserJourneyCampaignRequest, campaignId int64) (engagementVertexId int64, err error) {

	if request.Campaign == nil || request.Campaign.EngagementStartVertex == nil {
		return GetStartEngagementVertex(ctx, request.CampaignId)
	}
	return addEngagementJourneyUtil(ctx, tx, campaignId, 0, request.Campaign.EngagementStartVertex, nil)
}

func addEngagementJourneyUtil(ctx context.Context, tx dialect.Tx, campaignId int64, prevVrtexId int64, currVertex *fs.EngagementVertex,
	prevEdge *fs.EngagementEdge) (engagementVertexId int64, err error) {

	if currVertex == nil {
		return
	}
	model := makeAddEngagementVertexVO(currVertex, campaignId)
	args := AddEngagementVertexArgs(model)
	var rows sql.Result
	err = tx.Exec(ctx, query.QUERY_AddEngagementVertex, args, &rows)
	if err != nil {
		logger.Error("ScheduleUserJourneyCampaignController OnRequest hook, Error could not execute AddEngagementVertex request",
			zap.Error(err), zap.Any("model", model))
		return
	}
	vertexId, err := rows.LastInsertId()
	if err != nil {
		logger.Error("ScheduleUserJourneyCampaignController OnRequest hook, Error could not get lastInsertedId for AddEngagementVertex",
			zap.Error(err), zap.Any("model", model))
		return
	}
	engagementVertexId = vertexId
	if prevEdge != nil {
		var waitType string
		if prevEdge.WaitTime.WaitFor != EMPTY {
			waitType = "WAIT_FOR"
		} else if prevEdge.WaitTime.WaitTill != EMPTY {
			waitType = "WAIT_TILL"
		}
		engagementModel := makeAddEdgeVO(campaignId, "ENGAGEMENT", prevVrtexId, vertexId, prevEdge.WaitTime, waitType, prevEdge.States)
		args := AddEdgeArgs(engagementModel)
		var rows sql.Result
		err = tx.Exec(ctx, query.QUERY_AddEdge, args, &rows)
		if err != nil {
			logger.Error("ScheduleUserJourneyCampaignController OnRequest hook, Error could not execute AddEngagementVertex request",
				zap.Error(err), zap.Any("model", model))
			return
		}
	}

	if currVertex.Edges != nil {
		for _, edge := range currVertex.Edges {
			_, err = addEngagementJourneyUtil(ctx, tx, campaignId, vertexId, edge.Vertex, edge)
			if err != nil {
				return
			}
		}
	}
	return
}

func validateScheduleUserJourneyCampaignRequest(request *fs.ScheduleUserJourneyCampaignRequest) (err error) {

	logger.Info("ScheduleUserJourneyCampaignController OnRequest hook started", zap.Any("request", request))

	if (request.Name == EMPTY || request.Campaign == nil || request.Namespace == Common.NameSpace_NO_NAMESPACE || request.CreatedByActor == nil) &&
		request.CampaignId == 0 {
		err = errors.New(INVALID_REQUEST)
		return
	}
	//campaign := request.Campaign
	//if (campaign == nil && (request.CampaignId != 0 && !request.TriggerCampaign)) || (campaign != nil && campaign.UserJourneys == nil &&
	//	campaign.EngagementStartVertex == nil && (request.CampaignId != 0 && !request.TriggerCampaign)) {
	//	err = errors.New(INVALID_REQUEST)
	//	return
	//}
	return
}

func mapToScheduleUserJourneyCampaignResponse(request *fs.ScheduleUserJourneyCampaignRequest, status Common.RequestStatus, err error,
	engagementVertexId int64, cronExpression string, referenceId string, campaignId int64) *fs.ScheduleUserJourneyCampaignResponse {

	if err != nil {
		switch status {
		case Common.RequestStatus_BAD_INPUT:
			logger.Error("ScheduleUserJourneyCampaignController request failed, invalid request", zap.Any("request", request), zap.Error(err))
		case Common.RequestStatus_INTERNAL_ERROR:
			logger.Error("ScheduleUserJourneyCampaignController request failed, internal error", zap.Any("request", request), zap.Error(err))
		}
		return &fs.ScheduleUserJourneyCampaignResponse{
			Status: &Common.RequestStatusResult{
				Status:        status,
				ErrorMessages: []string{err.Error()},
			},
		}
	}
	logger.Info("ScheduleUserJourneyCampaignController OnRequest hook completed successfully", zap.Any("request", request))
	return &fs.ScheduleUserJourneyCampaignResponse{
		Status: &Common.RequestStatusResult{
			Status: Common.RequestStatus_SUCCESS,
		},
		CampaignId:         campaignId,
		EngagementVertexId: engagementVertexId,
		CronSchedule:       cronExpression,
		ReferenceId:        referenceId,
	}
}

func makeAddUserJourneyCampaignRequest(ctx context.Context, request *fs.ScheduleUserJourneyCampaignRequest, cronExpression string, metadata string) (*fs.AddCampaignRequest, string, string, error) {

	var userJourneyMetadata, engagementMetadata string
	if metadata != EMPTY {
		userJourneyMetadata, engagementMetadata, _ = mappers.MapMetaData(metadata)
	}
	if request.UserJourneyMetadata != EMPTY {
		userJourneyMetadata = request.UserJourneyMetadata
	}
	if request.EngagementMetadata != EMPTY {
		engagementMetadata = request.EngagementMetadata
	}
	addCampaignRequest := &fs.AddCampaignRequest{
		Namespace:      request.Namespace,
		Name:           request.Name,
		CronExpression: cronExpression,
		Status:         Common.CampaignStatus_DRAFTED,
		Type:           Common.CampaignQueryType_USER_JOURNEY,
		CreatedByActor: request.CreatedByActor,
	}
	if request.TriggerCampaign {
		addCampaignRequest.Status = Common.CampaignStatus_RUNNING
	}
	return addCampaignRequest, userJourneyMetadata, engagementMetadata, nil
}
