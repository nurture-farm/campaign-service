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

type FindQueryCampaignInterface interface {
	OnRequest(ctx context.Context, request *fs.FindQueryCampaignRequest) *fs.FindQueryCampaignResponse
	OnData(ctx context.Context, request *fs.FindQueryCampaignRequest, response *fs.FindQueryCampaignResponse) *fs.FindQueryCampaignResponse
	OnError(ctx context.Context, request *fs.FindQueryCampaignRequest, response *fs.FindQueryCampaignResponse, err error) *fs.FindQueryCampaignResponse
	OnResponse(ctx context.Context, request *fs.FindQueryCampaignRequest, response *fs.FindQueryCampaignResponse) *fs.FindQueryCampaignResponse
}

type GenericFindQueryCampaignExecutor struct {
	FindQueryCampaignInterface FindQueryCampaignInterface
}

type FindQueryCampaignController struct {
}

var FindQueryCampaignExecutor *GenericFindQueryCampaignExecutor

func (ge *GenericFindQueryCampaignExecutor) OnRequest(ctx context.Context, request *fs.FindQueryCampaignRequest) *fs.FindQueryCampaignResponse {
	return ge.FindQueryCampaignInterface.OnRequest(ctx, request)
}

func (ge *GenericFindQueryCampaignExecutor) OnResponse(ctx context.Context, request *fs.FindQueryCampaignRequest, response *fs.FindQueryCampaignResponse) *fs.FindQueryCampaignResponse {
	return ge.FindQueryCampaignInterface.OnResponse(ctx, request, response)
}

func (ge *GenericFindQueryCampaignExecutor) OnData(ctx context.Context, request *fs.FindQueryCampaignRequest, response *fs.FindQueryCampaignResponse) *fs.FindQueryCampaignResponse {
	return ge.FindQueryCampaignInterface.OnData(ctx, request, response)
}

func (ge *GenericFindQueryCampaignExecutor) OnError(ctx context.Context, request *fs.FindQueryCampaignRequest, response *fs.FindQueryCampaignResponse, err error) *fs.FindQueryCampaignResponse {
	return ge.FindQueryCampaignInterface.OnError(ctx, request, response, err)
}

func (rc *FindQueryCampaignController) OnRequest(ctx context.Context, request *fs.FindQueryCampaignRequest) *fs.FindQueryCampaignResponse {
	return nil
}

func (rc *FindQueryCampaignController) OnResponse(ctx context.Context, request *fs.FindQueryCampaignRequest, response *fs.FindQueryCampaignResponse) *fs.FindQueryCampaignResponse {
	return nil
}

func (rc *FindQueryCampaignController) OnData(ctx context.Context, request *fs.FindQueryCampaignRequest, response *fs.FindQueryCampaignResponse) *fs.FindQueryCampaignResponse {
	return nil
}

func (rc *FindQueryCampaignController) OnError(ctx context.Context, request *fs.FindQueryCampaignRequest, response *fs.FindQueryCampaignResponse, err error) *fs.FindQueryCampaignResponse {
	return nil
}
