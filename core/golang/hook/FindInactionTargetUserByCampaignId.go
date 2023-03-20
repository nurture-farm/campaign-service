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
	"context"
)

type FindInactionTargetUserByCampaignIdInterface interface {
	OnRequest(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest) *fs.FindInactionTargetUserByCampaignIdResponse
	OnData(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest, response *fs.FindInactionTargetUserByCampaignIdResponse) *fs.FindInactionTargetUserByCampaignIdResponse
	OnError(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest, response *fs.FindInactionTargetUserByCampaignIdResponse, err error) *fs.FindInactionTargetUserByCampaignIdResponse
	OnResponse(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest, response *fs.FindInactionTargetUserByCampaignIdResponse) *fs.FindInactionTargetUserByCampaignIdResponse
}

type GenericFindInactionTargetUserByCampaignIdExecutor struct {
	FindInactionTargetUserByCampaignIdInterface FindInactionTargetUserByCampaignIdInterface
}

type FindInactionTargetUserByCampaignIdController struct {
}

var FindInactionTargetUserByCampaignIdExecutor *GenericFindInactionTargetUserByCampaignIdExecutor

func (ge *GenericFindInactionTargetUserByCampaignIdExecutor) OnRequest(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest) *fs.FindInactionTargetUserByCampaignIdResponse {
	return ge.FindInactionTargetUserByCampaignIdInterface.OnRequest(ctx, request)
}

func (ge *GenericFindInactionTargetUserByCampaignIdExecutor) OnResponse(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest, response *fs.FindInactionTargetUserByCampaignIdResponse) *fs.FindInactionTargetUserByCampaignIdResponse {
	return ge.FindInactionTargetUserByCampaignIdInterface.OnResponse(ctx, request, response)
}

func (ge *GenericFindInactionTargetUserByCampaignIdExecutor) OnData(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest, response *fs.FindInactionTargetUserByCampaignIdResponse) *fs.FindInactionTargetUserByCampaignIdResponse {
	return ge.FindInactionTargetUserByCampaignIdInterface.OnData(ctx, request, response)
}

func (ge *GenericFindInactionTargetUserByCampaignIdExecutor) OnError(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest, response *fs.FindInactionTargetUserByCampaignIdResponse, err error) *fs.FindInactionTargetUserByCampaignIdResponse {
	return ge.FindInactionTargetUserByCampaignIdInterface.OnError(ctx, request, response, err)
}

func (rc *FindInactionTargetUserByCampaignIdController) OnRequest(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest) *fs.FindInactionTargetUserByCampaignIdResponse {
	return nil
}

func (rc *FindInactionTargetUserByCampaignIdController) OnResponse(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest, response *fs.FindInactionTargetUserByCampaignIdResponse) *fs.FindInactionTargetUserByCampaignIdResponse {
	return nil
}

func (rc *FindInactionTargetUserByCampaignIdController) OnData(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest, response *fs.FindInactionTargetUserByCampaignIdResponse) *fs.FindInactionTargetUserByCampaignIdResponse {
	return nil
}

func (rc *FindInactionTargetUserByCampaignIdController) OnError(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest, response *fs.FindInactionTargetUserByCampaignIdResponse, err error) *fs.FindInactionTargetUserByCampaignIdResponse {
	return nil
}
