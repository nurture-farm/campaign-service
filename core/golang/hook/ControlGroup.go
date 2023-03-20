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
	"code.nurture.farm/platform/CampaignService/zerotouch/golang/database/mappers"
	"encoding/binary"
	"encoding/json"
	"errors"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func getBloomFilter(request *fs.FindControlGroupByCampaignIdResponse) (*bloom.BloomFilter, error) {
	var bloomFilterSet []uint64
	attributes := mappers.MapControlGroupAttributes(request.Records.Attributes)
	if request.Records.BloomFilter != nil {
		err := json.Unmarshal(request.Records.BloomFilter, &bloomFilterSet)
		if err != nil {
			logger.Error("Error while Unmarshalling bloomFilterSet, ControlGroupAttributes ", zap.Error(err), zap.Any("AttributeList", request.Records.BloomFilter))
			return nil, err
		}
	}

	if attributes != nil {
		var k uint
		var capacity uint
		if attributes[CONST_K] != "" {
			k = getK(attributes[CONST_K])
		}
		if attributes[CONST_CAPACITY] != "" {
			capacity = getCapacity(attributes[CONST_CAPACITY])
		}
		if k != 0 && capacity != 0 {
			return bloom.FromWithM(bloomFilterSet, capacity, k), nil
		}
	}

	return nil, errors.New("error in getting Bloom Filter")
}

func getK(k string) uint {
	return cast.ToUint(k)
}

func getCapacity(capacity string) uint {
	return cast.ToUint(capacity)
}

func GetIndexOfId(rowData []string) int {
	for k, columnData := range rowData {
		if columnData == "id" {
			return k
		}
	}
	return -1
}
func convertToInt64(userId string) (int64, error) {
	n, err := cast.ToInt64E(userId)
	if err == nil {
		return n, nil
	} else {
		return 0, err
	}
}
func GenerateRandomUserIds(userIds []int64, controlGroupPercentage int32) []int64 {
	totalNoOfUserIds := int32(len(userIds))
	var noOfControlGroupUserIds int32
	noOfControlGroupUserIds = (controlGroupPercentage * totalNoOfUserIds) / 100
	rand.Seed(time.Now().Unix())
	randomList := rand.Perm(int(totalNoOfUserIds))
	controlGroupUserIds := make([]int64, noOfControlGroupUserIds)
	for i := 0; i < int(noOfControlGroupUserIds); i++ {
		controlGroupUserIds[i] = userIds[randomList[i]]
	}
	return controlGroupUserIds
}
func BloomFilter(controlGroupIds []int64) *bloom.BloomFilter {
	filter := bloom.NewWithEstimates(uint(len(controlGroupIds)), 0.01)
	for _, actorId := range controlGroupIds {
		filter.Add(encodeToByteSequence(actorId))
	}
	return filter
}
func IsPresent(bloomFilter *bloom.BloomFilter, userId int64) bool {
	return bloomFilter.Test(encodeToByteSequence(userId))
}
func encodeToByteSequence(n int64) []byte {
	i := uint32(n)
	byteSequence := make([]byte, 4)
	binary.BigEndian.PutUint32(byteSequence, i)
	return byteSequence
}
