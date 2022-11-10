// Code generated by protoc-gen-volcengine-sdk
// source: RecordManageService
// DO NOT EDIT!

package live

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/live"
	"github.com/volcengine/volc-sdk-golang/service/live/models/request"
)

func Test_DescribeRecordTaskFileHistory(t *testing.T) {
	instance := live.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.DescribeRecordTaskFileHistoryRequest{
		Vhost:    "your Vhost",
		App:      "your App",
		Stream:   "your Stream",
		DateFrom: "your DateFrom",
		DateTo:   "your DateTo",
		PageNum:  0,
		PageSize: 0,
		Type:     "your Type",
	}

	resp, status, err := instance.DescribeRecordTaskFileHistory(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}
