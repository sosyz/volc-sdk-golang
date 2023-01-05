package businessSecurity

import (
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"testing"
)

const (
	Ak = "ak" // write your access key
	Sk = "sk" // write your secret key
)

func init() {
	DefaultInstance.Client.SetAccessKey(Ak)
	DefaultInstance.Client.SetSecretKey(Sk)
}

func RiskDetection(appId int64, service string, parameters string) {
	res, err := DefaultInstance.RiskDetection(&RiskDetectionRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func TextRiskDetection(appId int64, service string, parameters string) {
	res, err := DefaultInstance.TextRisk(&RiskDetectionRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func AsyncImageRiskDetection(appId int64, service string, parameters string) {
	res, err := DefaultInstance.AsyncImageRisk(&AsyncRiskDetectionRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func AsyncRiskDetection(appId int64, service string, parameters string) {
	res, err := DefaultInstance.AsyncRiskDetection(&AsyncRiskDetectionRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func GetImageResult(dataId, service string, appId int64) (*ImageResultResponse, error) {
	return DefaultInstance.GetImageResult(&VideoResultRequest{
		DataId:  dataId,
		AppId:   appId,
		Service: service,
	})
}

func ImageRisk(appId int64, service string, parameters string) (*ImageResultResponse, error) {
	return DefaultInstance.ImageContentRisk(&RiskDetectionRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
}

func AsyncVideoRiskDetection(appId int64, service string, parameters string) (*AsyncVideoRiskResponse, error) {
	return DefaultInstance.AsyncVideoRisk(&AsyncRiskDetectionRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
}

func GetVideoResult(dataId, service string, appId int64) (*VideoResultResponse, error) {
	return DefaultInstance.VideoResult(&VideoResultRequest{
		DataId:  dataId,
		AppId:   appId,
		Service: service,
	})
}

func RiskResult(appId int64, service string, startTime, endTime, pageSize, pageNum int64) {
	res, err := DefaultInstance.RiskResult(&RiskResultRequest{
		AppId:     appId,   // write your app id
		Service:   service, // write business security service
		StartTime: startTime,
		EndTime:   endTime,
		Page: Page{
			PageSize: pageSize,
			PageNum:  pageNum,
		},
	})
	fmt.Println(err)
	if res != nil {
		fmt.Printf("result %+v\n", *res)
	}
}

func DataReport(appId int64, service string, parameters string) {
	res, err := DefaultInstance.DataReport(&DataReportRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func AsyncAudioRiskDetection(appId int64, service string, parameters string) (*AsyncVideoRiskResponse, error) {
	return DefaultInstance.AsyncAudioRisk(&AsyncRiskDetectionRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
}

func GetAudioResult(dataId, service string, appId int64) (*AudioResultResponse, error) {
	return DefaultInstance.GetAudioResult(&VideoResultRequest{
		DataId:  dataId,
		AppId:   appId,
		Service: service,
	})
}

func AsyncVideoLiveRiskDetection(appId int64, service string, parameters string) (*AsyncVideoRiskResponse, error) {
	return DefaultInstance.AsyncLiveVideoRisk(&AsyncRiskDetectionRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
}

func GetVideoLiveResult(dataId, service string, appId int64) (*VideoResultResponse, error) {
	return DefaultInstance.GetVideoLiveResult(&VideoResultRequest{
		DataId:  dataId,
		AppId:   appId,
		Service: service,
	})
}

func AsyncAudioLiveRiskDetection(appId int64, service string, parameters string) (*AsyncVideoRiskResponse, error) {
	return DefaultInstance.AsyncLiveAudioRisk(&AsyncRiskDetectionRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
}

func GetAudioLiveResult(dataId, service string, appId int64) (*AudioResultResponse, error) {
	return DefaultInstance.GetAudioLiveResult(&VideoResultRequest{
		DataId:  dataId,
		AppId:   appId,
		Service: service,
	})
}

func TextSliceRisk(appId int64, service, params string) (*TextSliceResultResponse, error) {
	return DefaultInstance.TextSliceRisk(&RiskDetectionRequest{
		AppId:      appId,
		Service:    service,
		Parameters: params,
	})
}

func NewCustomContents(appId int64, service, name, description, decision string, matchType int) (*AsyncRiskDetectionResponse, error) {
	return DefaultInstance.CreateCustomContents(&NewCustomContentsReq{
		AppID:       appId,
		Service:     service,
		Name:        name,
		Description: description,
		Decision:    decision,
		MatchType:   matchType,
	})
}

func EnableCustomContents(appId int64, name string) (*AsyncRiskDetectionResponse, error) {
	return DefaultInstance.EnableCustomContents(&UpdateContentReq{
		AppID: appId,
		Name:  name,
	})
}

func DisableCustomContents(appId int64, name string) (*AsyncRiskDetectionResponse, error) {
	return DefaultInstance.DisableCustomContents(&UpdateContentReq{
		AppID: appId,
		Name:  name,
	})
}

func DeleteCustomContents(appId int64, name string) (*AsyncRiskDetectionResponse, error) {
	return DefaultInstance.DeleteCustomContents(&UpdateContentReq{
		AppID: appId,
		Name:  name,
	})
}

func UploadCustomContents(appId int64, name string, contents []string, modifyType int) (*AsyncRiskDetectionResponse, error) {
	return DefaultInstance.UploadCustomContents(&UpdateContentReq{
		AppID:      appId,
		Name:       name,
		Contents:   contents,
		ModifyType: modifyType,
	})
}

func AsyncImageRiskV2(appId int64, service, params string) (*AsyncVideoRiskResponse, error) {
	return DefaultInstance.AsyncImageRiskV2(&AsyncRiskDetectionRequest{
		AppId:      appId,
		Service:    service,
		Parameters: params,
	})
}

func TestAsyncImageRiskV2(t *testing.T) {
	resp, err := AsyncImageRiskV2(5461, "image_content_risk", "")
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestBusinessSecurity_RiskResult(t *testing.T) {
	RiskResult(3332, "login", 1615535000, 1615540603, 10, 1)
}

func TestBusinessSecurity_RiskDetection(t *testing.T) {
	RiskDetection(3332, "login", "{\"operate_time\": 1615540603, \"uid\":123444}")
}

func ElementVerify(appId int64, service string, parameters string) {
	res, err := DefaultInstance.ElementVerify(&ElementVerifyRequest{
		AppId:      appId,      // write your app id
		Service:    service,    // write business security service
		Parameters: parameters, // write your parameters
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}
func TestBusinessSecurity_ElementVerify(t *testing.T) {
	ElementVerify(3332, "idcard_two_element_verify", "{\"operate_time\": 1615540603, \"idcard_no\": \"\", \"idcard_name\":\"\"}")
}

func TestBusinessSecurity_ElementVerifyV2(t *testing.T) {
	appId := int64(3332)
	service := "mobile_two_element_verify"
	parameters := "{\"operate_time\": 1635321212,\"mobile\":\"\",\"idcard_name\":\"\"}"

	res, err := DefaultInstance.ElementVerifyV2(&ElementVerifyRequest{
		AppId:      appId,
		Service:    service,
		Parameters: parameters,
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", *res)
}

func TestBusinessSecurity_MobileStatus(t *testing.T) {
	appId := int64(3332)
	service := "mobile_empty_status"
	parameters := "{\"operate_time\":1617960951,\"mobile\":\"12312341234\"}"

	res, err := DefaultInstance.MobileStatus(&MobileStatusRequest{
		AppId:      appId,
		Service:    service,
		Parameters: parameters,
	})
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("%v", *res)
}

func TestBusinessSecurity_MobileStatusV2(t *testing.T) {
	appId := int64(3332)
	service := "mobile_address"
	parameters := "{\"operate_time\":1617960951,\"mobile\":\"12312341234\"}"

	res, err := DefaultInstance.MobileStatusV2(&MobileStatusRequest{
		AppId:      appId,
		Service:    service,
		Parameters: parameters,
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", *res)
}

func TestNewCustomContents(t *testing.T) {
	resp, err := NewCustomContents(5461, "text_risk", "text", "text", "BLOCK1", 1)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", *resp)
}

func TestEnableCustomContents(t *testing.T) {
	resp, err := NewCustomContents(5461, "text_risk", "", "是的是的", "BLOCK", 1)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", *resp)
}

func TestVideoResultGet(t *testing.T) {
	resp, err := GetVideoResult("wangyifan46", "video_risk", 265193)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", *resp)
}

func TestAudioResultGet(t *testing.T) {
	resp, err := GetAudioResult("test009284", "audio_risk", 3332)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", *resp)
}

func TestTextSliceRisk(t *testing.T) {
	resp, err := TextSliceRisk(334361, "text_risk", "{\"operate_time\":\"1663895098\",\"biztype\":\"kmvkd\",\"text\":\"测试文本\",\"account_id\":\"1663895098507\"}")
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("%v", resp.String())
}

func TestBusinessSecurity_ElementVerifyEncrypted(t *testing.T) {
	type fields struct {
		Client *base.Client
		retry  bool
	}
	type args struct {
		encryptedType string
		secretKey     string
		req           *ElementVerifyRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ElementVerifyResponseV2
		wantErr bool
	}{
		{
			name: "test_element_verity_encrypted_01",
			fields: fields{
				Client: DefaultInstance.Client,
				retry:  true,
			},
			args: args{
				encryptedType: "AES",
				secretKey:     "your private key",
				req: &ElementVerifyRequest{
					AppId:   10067295,
					Service: "mobile_three_element_verify",
					// json string
					Parameters: "{\"operate_time\": 1635321212,\"mobile\":\"18400457239\",\"idcard_name\":\"xxxxx\", \"idcard_no\":\"630105199705251716\",\"bank_acc\":\"6214832935433257\"}",
				},
			},
			want: &ElementVerifyResponseV2{
				Code: 1012,
				Data: ElementVerifyDataV2{
					Status: 0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &BusinessSecurity{
				Client: tt.fields.Client,
				retry:  tt.fields.retry,
			}
			got, err := p.ElementVerifyEncrypted(tt.args.encryptedType, tt.args.secretKey, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ElementVerifyEncrypted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Code != tt.want.Code {
				t.Errorf("ElementVerifyEncrypted()  want code = %v, but get %v", tt.want.Code, got.Code)
				return
			}
			if got.Data.Status != tt.want.Data.Status {
				t.Errorf("ElementVerifyEncrypted()  want status = %v, but get %v", tt.want.Data.Status, got.Data.Status)
				return
			}
			t.Log(got)
		})
	}
}
