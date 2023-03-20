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

type AddQueryCampaignInterface interface {
	OnRequest(ctx context.Context, request *fs.AddQueryCampaignRequest) *fs.AddQueryCampaignResponse
	OnError(ctx context.Context, request *fs.AddQueryCampaignRequest, response *fs.AddQueryCampaignResponse, err error) *fs.AddQueryCampaignResponse
	OnResponse(ctx context.Context, request *fs.AddQueryCampaignRequest, response *fs.AddQueryCampaignResponse) *fs.AddQueryCampaignResponse
}

type AddQueryCampaignBulkInterface interface {
	OnRequest(ctx context.Context, request *fs.BulkAddQueryCampaignRequest) *fs.BulkAddQueryCampaignResponse
	OnError(ctx context.Context, request *fs.BulkAddQueryCampaignRequest, response *fs.BulkAddQueryCampaignResponse, err error) *fs.BulkAddQueryCampaignResponse
	OnResponse(ctx context.Context, request *fs.BulkAddQueryCampaignRequest, response *fs.BulkAddQueryCampaignResponse) *fs.BulkAddQueryCampaignResponse
}

type GenericAddQueryCampaignExecutor struct {
	AddQueryCampaignInterface AddQueryCampaignInterface
}

type GenericAddQueryCampaignExecutorBulk struct {
	AddQueryCampaignBulkInterface AddQueryCampaignBulkInterface
}

type AddQueryCampaignController struct {
}

type BulkAddQueryCampaignController struct {
}

var AddQueryCampaignExecutor *GenericAddQueryCampaignExecutor
var BulkAddQueryCampaignExecutor *GenericAddQueryCampaignExecutorBulk

func (ge *GenericAddQueryCampaignExecutor) OnRequest(ctx context.Context, request *fs.AddQueryCampaignRequest) *fs.AddQueryCampaignResponse {
	return ge.AddQueryCampaignInterface.OnRequest(ctx, request)
}

func (ge *GenericAddQueryCampaignExecutor) OnResponse(ctx context.Context, request *fs.AddQueryCampaignRequest, response *fs.AddQueryCampaignResponse) *fs.AddQueryCampaignResponse {
	return ge.AddQueryCampaignInterface.OnResponse(ctx, request, response)
}

func (ge *GenericAddQueryCampaignExecutor) OnError(ctx context.Context, request *fs.AddQueryCampaignRequest, response *fs.AddQueryCampaignResponse, err error) *fs.AddQueryCampaignResponse {
	return ge.AddQueryCampaignInterface.OnError(ctx, request, response, err)
}

func (ge *GenericAddQueryCampaignExecutorBulk) OnRequest(ctx context.Context, request *fs.BulkAddQueryCampaignRequest) *fs.BulkAddQueryCampaignResponse {
	return ge.AddQueryCampaignBulkInterface.OnRequest(ctx, request)
}

func (ge *GenericAddQueryCampaignExecutorBulk) OnResponse(ctx context.Context, request *fs.BulkAddQueryCampaignRequest, response *fs.BulkAddQueryCampaignResponse) *fs.BulkAddQueryCampaignResponse {
	return ge.AddQueryCampaignBulkInterface.OnResponse(ctx, request, response)
}

func (ge *GenericAddQueryCampaignExecutorBulk) OnError(ctx context.Context, request *fs.BulkAddQueryCampaignRequest, response *fs.BulkAddQueryCampaignResponse, err error) *fs.BulkAddQueryCampaignResponse {
	return ge.AddQueryCampaignBulkInterface.OnError(ctx, request, response, err)
}

func (rc *AddQueryCampaignController) OnRequest(ctx context.Context, request *fs.AddQueryCampaignRequest) *fs.AddQueryCampaignResponse {
	return nil
}

func (rc *AddQueryCampaignController) OnResponse(ctx context.Context, request *fs.AddQueryCampaignRequest, response *fs.AddQueryCampaignResponse) *fs.AddQueryCampaignResponse {
	return nil
}

func (rc *AddQueryCampaignController) OnError(ctx context.Context, request *fs.AddQueryCampaignRequest, response *fs.AddQueryCampaignResponse, err error) *fs.AddQueryCampaignResponse {
	return nil
}

func (rc *BulkAddQueryCampaignController) OnRequest(ctx context.Context, request *fs.BulkAddQueryCampaignRequest) *fs.BulkAddQueryCampaignResponse {
	return nil
}

func (rc *BulkAddQueryCampaignController) OnResponse(ctx context.Context, request *fs.BulkAddQueryCampaignRequest, response *fs.BulkAddQueryCampaignResponse) *fs.BulkAddQueryCampaignResponse {
	return nil
}

func (rc *BulkAddQueryCampaignController) OnError(ctx context.Context, request *fs.BulkAddQueryCampaignRequest, response *fs.BulkAddQueryCampaignResponse, err error) *fs.BulkAddQueryCampaignResponse {
	return nil
}
