package businessSecurity

import (
	"encoding/json"
	"fmt"
)

// Synchronous risk detection
// 风险识别实时接口
func (p *BusinessSecurity) RiskDetection(req *RiskDetectionRequest) (*RiskDetectionResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("RiskDetectionRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("RiskDetection", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("RiskDetection", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("RiskDetection: fail to do request, %v", err)
			}
			result := new(RiskDetectionResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("RiskDetection: fail to do request, %v", err)
	}
	result := new(RiskDetectionResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Asynchronous risk detection
// 风险识别异步接口
func (p *BusinessSecurity) AsyncRiskDetection(req *AsyncRiskDetectionRequest) (*AsyncRiskDetectionResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AsyncRiskDetectionRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("AsyncRiskDetection", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("AsyncRiskDetection", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AsyncRiskDetection: fail to do request, %v", err)
			}
			result := new(AsyncRiskDetectionResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AsyncRiskDetection: fail to do request, %v", err)
	}
	result := new(AsyncRiskDetectionResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Risk result
// 风险识别结果获取接口
func (p *BusinessSecurity) RiskResult(req *RiskResultRequest) (*RiskResultResponse, error) {
	respBody, _, err := p.Client.Query("RiskResult", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("RiskResult", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("RiskResult: fail to do request, %v", err)
			}
			result := new(RiskResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AsyncRiskDetection: fail to do request, %v", err)
	}
	result := new(RiskResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) DataReport(req *DataReportRequest) (*DataReportResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AsyncRiskDetectionRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("DataReport", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("DataReport", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AsyncRiskDetection: fail to do request, %v", err)
			}
			result := new(DataReportResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("DataReport: fail to do request, %v", err)
	}
	result := new(DataReportResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Asynchronous video risk detection
// 内容安全视频异步接口
func (p *BusinessSecurity) AsyncVideoRisk(req *AsyncRiskDetectionRequest) (*AsyncVideoRiskResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AsyncVideoRisk: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("AsyncVideoRisk", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("AsyncVideoRisk", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AsyncVideoRisk: fail to do request, %v", err)
			}
			result := new(AsyncVideoRiskResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AsyncRiskDetection: fail to do request, %v", err)
	}
	result := new(AsyncVideoRiskResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Deprecated: use ElementVerifyV2 instead
// 已废弃，请使用ElementVerifyV2
func (p *BusinessSecurity) ElementVerify(req *ElementVerifyRequest) (*ElementVerifyResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ElementVerifyRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("ElementVerify", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("ElementVerify", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ElementVerify: fail to do request, %v", err)
			}
			result := new(ElementVerifyResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ElementVerify: fail to do request, %v", err)
	}
	result := new(ElementVerifyResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Risk result
// 风险识别结果获取接口
func (p *BusinessSecurity) VideoResult(req *VideoResultRequest) (*VideoResultResponse, error) {
	respBody, _, err := p.Client.Query("VideoResult", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("VideoResult", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("VideoResult: fail to do request, %v", err)
			}
			result := new(VideoResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("VideoResult: fail to do request, %v", err)
	}
	result := new(VideoResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Asynchronous video risk detection
// 内容安全图片异步接口
func (p *BusinessSecurity) AsyncImageRisk(req *AsyncRiskDetectionRequest) (*AsyncVideoRiskResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AsyncImageRisk: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("AsyncImageRisk", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("AsyncImageRisk", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AsyncImageRisk: fail to do request, %v", err)
			}
			result := new(AsyncVideoRiskResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AsyncImageRisk: fail to do request, %v", err)
	}
	result := new(AsyncVideoRiskResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) AsyncImageRiskV2(req *AsyncRiskDetectionRequest) (*AsyncVideoRiskResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AsyncImageRiskV2: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("AsyncImageRiskV2", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("AsyncImageRiskV2", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AsyncImageRiskV2: fail to do request, %v", err)
			}
			result := new(AsyncVideoRiskResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AsyncImageRiskV2: fail to do request, %v", err)
	}
	result := new(AsyncVideoRiskResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Risk result
// 风险识别结果获取接口
func (p *BusinessSecurity) GetImageResult(req *VideoResultRequest) (*ImageResultResponse, error) {
	respBody, _, err := p.Client.Query("GetImageResult", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("GetImageResult", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("GetImageResult: fail to do request, %v", err)
			}
			result := new(ImageResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetImageResult: fail to do request, %v", err)
	}
	result := new(ImageResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) GetImageResultV2(req *VideoResultRequest) (*ImageResultResponse, error) {
	respBody, _, err := p.Client.Query("GetImageResultV2", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("GetImageResultV2", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("GetImageResultV2: fail to do request, %v", err)
			}
			result := new(ImageResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetImageResultV2: fail to do request, %v", err)
	}
	result := new(ImageResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// image risk deteciton
// 内容安全图片实时接口
func (p *BusinessSecurity) ImageContentRisk(req *RiskDetectionRequest) (*ImageResultResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ImageContentRisk: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("ImageContentRisk", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("ImageContentRisk", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ImageContentRisk: fail to do request, %v", err)
			}
			result := new(ImageResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ImageContentRisk: fail to do request, %v", err)
	}
	result := new(ImageResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// 要素验证
func (p *BusinessSecurity) ElementVerifyV2(req *ElementVerifyRequest) (*ElementVerifyResponseV2, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ElementVerifyRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("ElementVerifyV2", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("ElementVerifyV2", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ElementVerify: fail to do request, %v", err)
			}
			result := new(ElementVerifyResponseV2)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ImageContentRisk: fail to do request, %v", err)
	}
	result := new(ElementVerifyResponseV2)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ElementVerifyEncrypted 加密要素验证
// encryptedType - 加密类型，例如：AES
// secretKey - 秘钥(需要申请)
func (p *BusinessSecurity) ElementVerifyEncrypted(encryptedType string, secretKey string, req *ElementVerifyRequest) (*ElementVerifyResponseV2, error) {
	if encryptedType == "" || secretKey == "" {
		return nil, fmt.Errorf("ElementVerifyEncrypted: encryptedType or secretKey empty")
	}
	if req == nil {
		return nil, fmt.Errorf("ElementVerifyEncrypted: requset is nil")
	}
	req.EncryptedType = encryptedType
	parameters, err := AesCBCEncryptWithBase64(secretKey, req.Parameters)
	if err != nil {
		return nil, fmt.Errorf("ElementVerifyEncrypted: fail encrypt parameters")
	}
	req.Parameters = parameters

	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ElementVerifyEncryptedRequest: fail to marshal request, %v", err)
	}
	respBody, _, err := p.Client.Json("ElementVerifyEncrypted", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("ElementVerifyEncrypted", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ElementVerifyEncrypted: fail to do request, %v", err)
			}
			result := new(ElementVerifyResponseV2)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ElementVerifyEncrypted: fail to do request, %v", err)
	}
	result := new(ElementVerifyResponseV2)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// text risk detection
// 内容安全文本实时接口
func (p *BusinessSecurity) TextRisk(req *RiskDetectionRequest) (*TextResultResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("TextRisk: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("TextRisk", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("TextRisk", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("TextRisk: fail to do request, %v", err)
			}
			result := new(TextResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("TextRisk: fail to do request, %v", err)
	}
	result := new(TextResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// text slice risk detection
// 内容安全文本切片检测实时接口
func (p *BusinessSecurity) TextSliceRisk(req *RiskDetectionRequest) (*TextSliceResultResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("TextSliceRisk: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("TextSliceRisk", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("TextSliceRisk", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("TextSliceRisk: fail to do request, %v", err)
			}
			result := new(TextSliceResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("TextSliceRisk: fail to do request, %v", err)
	}
	result := new(TextSliceResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Deprecated: use MobileStatusV2 instead
// 已废弃，请使用MobileStatusV2
func (p *BusinessSecurity) MobileStatus(req *MobileStatusRequest) (*MobileStatusResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("MobileSecondSaleRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("MobileStatus", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("MobileStatus", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("MobileStatus: fail to do request, %v", err)
			}
			result := new(MobileStatusResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
	}
	result := new(MobileStatusResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Asynchronous audio risk detection
// 内容安全音频异步接口
func (p *BusinessSecurity) AsyncAudioRisk(req *AsyncRiskDetectionRequest) (*AsyncVideoRiskResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AsyncAudioRisk: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("AsyncAudioRisk", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("AsyncAudioRisk", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AsyncImageRisk: fail to do request, %v", err)
			}
			result := new(AsyncVideoRiskResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AsyncAudioRisk: fail to do request, %v", err)
	}
	result := new(AsyncVideoRiskResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// 号码状态
func (p *BusinessSecurity) MobileStatusV2(req *MobileStatusRequest) (*MobileStatusResponseV2, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("MobileSecondSaleRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("MobileStatusV2", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("MobileStatusV2", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("MobileStatusV2: fail to do request, %v", err)
			}
			result := new(MobileStatusResponseV2)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("MobileStatus: fail to do request, %v", err)
	}
	result := new(MobileStatusResponseV2)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// audio Risk result
// 风险识别结果获取接口
func (p *BusinessSecurity) GetAudioResult(req *VideoResultRequest) (*AudioResultResponse, error) {
	respBody, _, err := p.Client.Query("GetAudioResult", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("GetAudioResult", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("GetAudioResult: fail to do request, %v", err)
			}
			result := new(AudioResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetAudioResult: fail to do request, %v", err)
	}
	result := new(AudioResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Asynchronous video risk detection
// 内容安全图片异步接口
func (p *BusinessSecurity) AsyncLiveVideoRisk(req *AsyncRiskDetectionRequest) (*AsyncVideoRiskResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AsyncLiveVideoRisk: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("AsyncLiveVideoRisk", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("AsyncLiveVideoRisk", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AsyncLiveVideoRisk: fail to do request, %v", err)
			}
			result := new(AsyncVideoRiskResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AsyncLiveVideoRisk: fail to do request, %v", err)
	}
	result := new(AsyncVideoRiskResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Risk result
// 风险识别结果获取接口
func (p *BusinessSecurity) GetVideoLiveResult(req *VideoResultRequest) (*VideoResultResponse, error) {
	respBody, _, err := p.Client.Query("GetVideoLiveResul", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("GetVideoLiveResult", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("GetVideoLiveResult: fail to do request, %v", err)
			}
			result := new(VideoResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetImageResult: fail to do request, %v", err)
	}
	result := new(VideoResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Asynchronous audio risk detection
// 内容安全音频异步接口
func (p *BusinessSecurity) AsyncLiveAudioRisk(req *AsyncRiskDetectionRequest) (*AsyncVideoRiskResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AsyncLiveAudioRisk: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("AsyncLiveAudioRisk", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("AsyncLiveAudioRisk", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AsyncLiveAudioRisk: fail to do request, %v", err)
			}
			result := new(AsyncVideoRiskResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AsyncLiveAudioRisk: fail to do request, %v", err)
	}
	result := new(AsyncVideoRiskResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// audio Risk result
// 风险识别结果获取接口
func (p *BusinessSecurity) GetAudioLiveResult(req *VideoResultRequest) (*AudioResultResponse, error) {
	respBody, _, err := p.Client.Query("GetAudioLiveResult", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("GetAudioLiveResult", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("GetAudioLiveResult: fail to do request, %v", err)
			}
			result := new(AudioResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetAudioLiveResult: fail to do request, %v", err)
	}
	result := new(AudioResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// create custom contents
// 创建文本自定义库
func (p *BusinessSecurity) CreateCustomContents(req *NewCustomContentsReq) (*AsyncRiskDetectionResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("CreateCustomContents: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("CreateCustomContents", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("CreateCustomContents", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("CreateCustomContents: fail to do request, %v", err)
			}
			result := new(AsyncRiskDetectionResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("CreateCustomContents: fail to do request, %v", err)
	}
	result := new(AsyncRiskDetectionResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// enable custom contents
// 启用文本自定义库
func (p *BusinessSecurity) EnableCustomContents(req *UpdateContentReq) (*AsyncRiskDetectionResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("EnableCustomContents: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("EnableCustomContents", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("EnableCustomContents", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("EnableCustomContents: fail to do request, %v", err)
			}
			result := new(AsyncRiskDetectionResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("EnableCustomContents: fail to do request, %v", err)
	}
	result := new(AsyncRiskDetectionResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// disable custom contents
// 禁用文本自定义库
func (p *BusinessSecurity) DisableCustomContents(req *UpdateContentReq) (*AsyncRiskDetectionResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("DisableCustomContents: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("DisableCustomContents", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("DisableCustomContents", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("DisableCustomContents: fail to do request, %v", err)
			}
			result := new(AsyncRiskDetectionResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("DisableCustomContents: fail to do request, %v", err)
	}
	result := new(AsyncRiskDetectionResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// delete custom contents
// 禁用文本自定义库
func (p *BusinessSecurity) DeleteCustomContents(req *UpdateContentReq) (*AsyncRiskDetectionResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("DeleteCustomContents: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("DeleteCustomContents", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("DeleteCustomContents", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("DeleteCustomContents: fail to do request, %v", err)
			}
			result := new(AsyncRiskDetectionResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("DisableCustomContents: fail to do request, %v", err)
	}
	result := new(AsyncRiskDetectionResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// upload custom contents
// 上传文本自定义库文本
func (p *BusinessSecurity) UploadCustomContents(req *UpdateContentReq) (*AsyncRiskDetectionResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("UploadCustomContents: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("UploadCustomContents", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("UploadCustomContents", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("UploadCustomContents: fail to do request, %v", err)
			}
			result := new(AsyncRiskDetectionResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("UploadCustomContents: fail to do request, %v", err)
	}
	result := new(AsyncRiskDetectionResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}
