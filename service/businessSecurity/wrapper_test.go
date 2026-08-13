package businessSecurity

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

const (
	Ak = "" // fill in your access key
	Sk = "" // fill in your secret key
)

func init() {
	DefaultInstance.Client.SetAccessKey(Ak)
	DefaultInstance.Client.SetSecretKey(Sk)

	SecuritySourceInstance.Client.SetAccessKey(Ak)
	SecuritySourceInstance.Client.SetSecretKey(Sk)

}
func TestOpenProduct(t *testing.T) {
	resp, err := DefaultInstance.OpenProduct(&OpenProductReq{
		PRDContentRisk,
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", *resp)
}
func TestCheckProduct(t *testing.T) {
	resp, err := DefaultInstance.CheckProductStatus(&CheckProductStatusReq{
		PRDContentRisk,
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", *resp)
}
func TestOpenService(t *testing.T) {
	resp, err := DefaultInstance.EnableService(&EnableServiceReq{
		Service: ServiceContentTextRisk,
		AppID:   0,
		Product: PRDContentRisk,
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", *resp)
}
func TestCheckServiceStatus(t *testing.T) {
	resp, err := DefaultInstance.CheckServiceStatus(&CheckServiceStatusReq{
		Service: ServiceContentTextRisk,
		AppID:   0,
		Product: PRDContentRisk,
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", *resp)
}
func TestBusinessSecurity_ActivateRiskBasePackage(t *testing.T) {
	req := &ActivateRiskBasePackageReq{
		PackageId:       "Ab123456",
		TotalPackageNum: 1,
		PackageSeq:      1,
		DataType:        "1",
		Data:            []string{"cfcd208495d565ef66e7dff9f98764da"},
	}
	resp, err := DefaultInstance.ActivateRiskBasePackage(req)
	fmt.Println(resp, err)
}

func TestBusinessSecurity_ActivateRiskSampleData(t *testing.T) {
	list := make([]SampleData, 0)
	list = append(list, SampleData{
		Id:           "cfcd208495d565ef66e7dff9f98764da",
		ReachType:    "X",
		LaunchStatus: "2",
	})
	req := &ActivateRiskSampleDataReq{
		PackageId:       "A11",
		TotalPackageNum: 1,
		PackageSeq:      1,
		DataType:        "1",
		BusinessType:    "A1",
		Data:            list,
	}
	resp, err := DefaultInstance.ActivateRiskSampleData(req)
	fmt.Println(resp, err)
}

// ============ 图片/视频内容检测接口测试 ============
// 注意：以下用例与均为真实网络调用，
// 需在上方 Ak/Sk 常量填入有效的火山引擎密钥后才能跑通。

// TestImageContentRiskV2 图片实时同步检测
func TestImageContentRiskV2(t *testing.T) {
	params, _ := json.Marshal(map[string]interface{}{
		"operate_time": time.Now().Unix(),
		"account_id":   "",
		"url":          "",
		"data_id":      "",
		"biztype":      "",
	})
	req := &RiskDetectionRequest{
		AppId:      0,
		Service:    ServiceContentImageContentRisk,
		Parameters: string(params),
	}
	resp, err := DefaultInstance.ImageContentRiskV2(req)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	respJSON, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("%s", respJSON)
}

// TestAsyncImageRiskV2 图片异步提交检测（V2）
func TestAsyncImageRiskV2(t *testing.T) {
	params, _ := json.Marshal(map[string]interface{}{
		"operate_time": time.Now().Unix(),
		"account_id":   "",
		"url":          "",
		"data_id":      "",
		"biztype":      "",
	})
	req := &AsyncRiskDetectionRequest{
		AppId:      0,
		Service:    ServiceContentImageContentRisk,
		Parameters: string(params),
	}
	resp, err := DefaultInstance.AsyncImageRiskV2(req)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	respJSON, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("%s", respJSON)
}

// TestGetImageResultV2 拉取图片异步检测结果（V2）
func TestGetImageResultV2(t *testing.T) {
	req := &VideoResultRequest{
		DataId:  "",
		AppId:   0,
		Service: ServiceContentImageContentRisk,
	}
	resp, err := DefaultInstance.GetImageResultV2(req)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	respJSON, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("%s", respJSON)
}

// TestAsyncVideoRisk 视频异步提交检测
func TestAsyncVideoRisk(t *testing.T) {
	params, _ := json.Marshal(map[string]interface{}{
		"operate_time": time.Now().Unix(),
		"account_id":   "",
		"url":          "",
		"data_id":      "",
		"biztype":      "",
	})
	req := &AsyncRiskDetectionRequest{
		AppId:      0,
		Service:    ServiceContentVideoRisk,
		Parameters: string(params),
	}
	resp, err := DefaultInstance.AsyncVideoRisk(req)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	respJSON, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("%s", respJSON)
}

// TestVideoResult 拉取视频异步检测结果
func TestVideoResult(t *testing.T) {
	req := &VideoResultRequest{
		DataId:  "",
		AppId:   0,
		Service: ServiceContentVideoRisk,
	}
	resp, err := DefaultInstance.VideoResult(req)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	respJSON, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("%s", respJSON)
}
