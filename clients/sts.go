// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clients

import (
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
)

func NewSTSClient(sdkConfig *sdk.Config, region, key, secret string) (*STSClient, error) {
	creds, err := chainedCreds(key, secret)
	if err != nil {
		return nil, err
	}
	// We hard-code a region here because there's only one RAM endpoint regardless of the
	// region you're in.
	// client, err := sts.NewClientWithOptions("us-east-1", sdkConfig, creds)
	client, err := sts.NewClientWithOptions(region, sdkConfig, creds)
	if err != nil {
		return nil, err
	}
	return &STSClient{client: client}, nil
}

type STSClient struct {
	client *sts.Client
}

func (c *STSClient) AssumeRole(roleSessionName, roleARN string, expireDuration time.Duration) (*sts.AssumeRoleResponse, error) {
	assumeRoleReq := sts.CreateAssumeRoleRequest()
	assumeRoleReq.RoleArn = roleARN
	assumeRoleReq.RoleSessionName = roleSessionName
	assumeRoleReq.DurationSeconds = requests.NewInteger64(int64(expireDuration.Seconds()))
	return c.client.AssumeRole(assumeRoleReq)
}
