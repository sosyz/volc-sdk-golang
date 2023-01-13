package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type OcrApiResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             interface{}            `json:"data"`
}
