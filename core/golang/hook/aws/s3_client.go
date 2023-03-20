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
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"go.uber.org/zap"
	"log"
	"sync"
)

const (
	CONST_AWS_REGION                  = "ap-south-1"
	CONST_ATHENA_QUERY_SLEEP_DURATION = 2 // in seconds
	CONST_ATHENA_QUERY_SUCCESS        = "SUCCEEDED"
	CONST_ATHENA_QUERY_FAILED         = "FAILED"
	CONST_DATABASE_REWARDS_GATEWAY    = "rewards_gateway"
	CONST_NIL                         = "NIL"
)

var (
	once sync.Once
	svc  *s3.S3
)

func PutObjectInBucket(data []byte, bucket string, key string) error {

	putObjectInput := &s3.PutObjectInput{
		Body:   bytes.NewReader(data),
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	_, err := svc.PutObject(putObjectInput)
	if err != nil {
		logger.Error("Error in uploading data", zap.Error(err), zap.Any("bucket", bucket), zap.Any("key", key))
		return err
	}
	log.Printf("Successfully uploaded data to %q/%q\n\n", bucket, key)
	return nil
}

func main() {

	data := []byte("Hello, S3!")

	// Specify the bucket and object key
	bucket := "afs-prod-client/push_notification_campaign"
	key := "hello.txt"

	err := PutObjectInBucket(data, bucket, key)
	if err != nil {
		log.Fatalf("Failed to upload data, %v", err)
	}
	log.Printf("Successfully uploaded data to %q/%q\n\n", bucket, key)
}
