package vod

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/avast/retry-go"
	"github.com/volcengine/volc-sdk-golang/base"
	model_base "github.com/volcengine/volc-sdk-golang/service/base/models/base"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/business"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/response"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/volcengine/volc-sdk-golang/service/vod/upload/consts"
	"github.com/volcengine/volc-sdk-golang/service/vod/upload/model"
)

func (p *Vod) GetSubtitleAuthToken(req *request.VodGetSubtitleInfoListRequest, tokenExpireTime int) (string, error) {
	if len(req.GetVid()) == 0 {
		return "", errors.New("传入的Vid为空")
	}
	query := url.Values{
		"Vid":    []string{req.GetVid()},
		"Status": []string{"Published"},
	}

	if tokenExpireTime > 0 {
		query.Add("X-Expires", strconv.Itoa(tokenExpireTime))
	}

	if getSubtitleInfoAuthToken, err := p.GetSignUrl("GetSubtitleInfoList", query); err == nil {
		ret := map[string]string{}
		ret["GetSubtitleAuthToken"] = getSubtitleInfoAuthToken
		b, err := json.Marshal(ret)
		if err != nil {
			return "", err
		}
		return base64.StdEncoding.EncodeToString(b), nil
	} else {
		return "", err
	}
}

func (p *Vod) GetPrivateDrmAuthToken(req *request.VodGetPrivateDrmPlayAuthRequest, tokenExpireTime int) (string, error) {
	if len(req.GetVid()) == 0 {
		return "", errors.New("传入的Vid为空")
	}
	query := url.Values{
		"Vid": []string{req.GetVid()},
	}

	if len(req.GetPlayAuthIds()) > 0 {
		query.Add("PlayAuthIds", req.GetPlayAuthIds())
	}
	if len(req.GetDrmType()) > 0 {
		query.Add("DrmType", req.GetDrmType())
		switch req.GetDrmType() {
		case "appdevice", "webdevice":
			if len(req.GetUnionInfo()) == 0 {
				return "", errors.New("invalid unionInfo")
			}
		}
	}
	if len(req.GetUnionInfo()) > 0 {
		query.Add("UnionInfo", req.GetUnionInfo())
	}
	if tokenExpireTime > 0 {
		query.Add("X-Expires", strconv.Itoa(tokenExpireTime))
	}

	if getPrivateDrmAuthToken, err := p.GetSignUrl("GetPrivateDrmPlayAuth", query); err == nil {
		return getPrivateDrmAuthToken, nil
	} else {
		return "", err
	}
}

func (p *Vod) CreateSha1HlsDrmAuthToken(expireSeconds int64) (auth string, err error) {
	return p.createHlsDrmAuthToken(DSAHmacSha1, expireSeconds)
}

func (p *Vod) createHlsDrmAuthToken(authAlgorithm string, expireSeconds int64) (string, error) {
	if expireSeconds == 0 {
		return "", errors.New("invalid expire")
	}

	token, err := createAuth(authAlgorithm, Version2, p.ServiceInfo.Credentials.AccessKeyID,
		p.ServiceInfo.Credentials.SecretAccessKey, p.ServiceInfo.Credentials.Region, expireSeconds)
	if err != nil {
		return "", err
	}

	query := url.Values{}
	query.Set("DrmAuthToken", token)
	query.Set("X-Expires", strconv.FormatInt(expireSeconds, 10))
	if getAuth, err := p.GetSignUrl("GetHlsDecryptionKey", query); err == nil {
		return getAuth, nil
	} else {
		return "", err
	}
}

func (p *Vod) GetPlayAuthToken(req *request.VodGetPlayInfoRequest, tokenExpireTime int) (string, error) {
	if len(req.GetVid()) == 0 {
		return "", errors.New("传入的Vid为空")
	}
	query := url.Values{}
	marshaler := protojson.MarshalOptions{
		Multiline:       false,
		AllowPartial:    false,
		UseProtoNames:   true,
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
	}
	jsonData := marshaler.Format(req)
	reqMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonData), &reqMap)
	if err != nil {
		return "", err
	}
	for k, v := range reqMap {
		var sv string
		switch ov := v.(type) {
		case string:
			sv = ov
		case int:
			sv = strconv.FormatInt(int64(ov), 10)
		case uint:
			sv = strconv.FormatUint(uint64(ov), 10)
		case int8:
			sv = strconv.FormatInt(int64(ov), 10)
		case uint8:
			sv = strconv.FormatUint(uint64(ov), 10)
		case int16:
			sv = strconv.FormatInt(int64(ov), 10)
		case uint16:
			sv = strconv.FormatUint(uint64(ov), 10)
		case int32:
			sv = strconv.FormatInt(int64(ov), 10)
		case uint32:
			sv = strconv.FormatUint(uint64(ov), 10)
		case int64:
			sv = strconv.FormatInt(ov, 10)
		case uint64:
			sv = strconv.FormatUint(ov, 10)
		case bool:
			sv = strconv.FormatBool(ov)
		case float32:
			sv = strconv.FormatFloat(float64(ov), 'f', -1, 32)
		case float64:
			sv = strconv.FormatFloat(ov, 'f', -1, 64)
		case []byte:
			sv = string(ov)
		default:
			v2, e2 := json.Marshal(ov)
			if e2 != nil {
				return "", e2
			}
			sv = string(v2)
		}
		query.Set(k, sv)
	}
	if tokenExpireTime > 0 {
		query.Add("X-Expires", strconv.Itoa(tokenExpireTime))
	}
	if getPlayInfoToken, err := p.GetSignUrl("GetPlayInfo", query); err == nil {
		ret := map[string]string{}
		ret["GetPlayInfoToken"] = getPlayInfoToken
		ret["TokenVersion"] = "V2"
		b, err := json.Marshal(ret)
		if err != nil {
			return "", err
		}
		return base64.StdEncoding.EncodeToString(b), nil
	} else {
		return "", err
	}
}

func (p *Vod) UploadObjectWithCallback(filePath string, spaceName string, callbackArgs string, fileName, fileExtension string, funcs string) (*response.VodCommitUploadInfoResponse, int, error) {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return nil, -1, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return nil, -1, err
	}
	return p.UploadMediaInner(file, stat.Size(), spaceName, "object", callbackArgs, funcs, fileName, fileExtension, 0)
}

func (p *Vod) UploadMediaWithCallback(mediaRequset *request.VodUploadMediaRequest) (*response.VodCommitUploadInfoResponse, int, error) {
	file, err := os.Open(filepath.Clean(mediaRequset.GetFilePath()))
	if err != nil {
		return nil, -1, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return nil, -1, err
	}
	return p.UploadMediaInner(file, stat.Size(), mediaRequset.GetSpaceName(), "", mediaRequset.GetCallbackArgs(), mediaRequset.GetFunctions(), mediaRequset.GetFileName(), mediaRequset.GetFileExtension(), mediaRequset.StorageClass)
}

func (p *Vod) UploadMaterialWithCallback(materialRequest *request.VodUploadMaterialRequest) (*response.VodCommitUploadInfoResponse, int, error) {
	file, err := os.Open(filepath.Clean(materialRequest.GetFilePath()))
	if err != nil {
		return nil, -1, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return nil, -1, err
	}
	return p.UploadMediaInner(file, stat.Size(), materialRequest.GetSpaceName(), materialRequest.GetFileType(), materialRequest.GetCallbackArgs(), materialRequest.GetFunctions(), materialRequest.GetFileName(), materialRequest.GetFileExtension(), 0)
}

func (p *Vod) UploadMediaInner(rd io.Reader, size int64, spaceName string, fileType, callbackArgs string, funcs string, fileName, fileExtension string, storageClass int32) (*response.VodCommitUploadInfoResponse, int, error) {
	logId, sessionKey, err, code := p.Upload(rd, size, spaceName, fileType, fileName, fileExtension, storageClass)
	if err != nil {
		return p.fillCommitUploadInfoResponseWhenError(logId, err.Error()), code, err
	}

	commitRequest := &request.VodCommitUploadInfoRequest{
		SpaceName:    spaceName,
		SessionKey:   sessionKey,
		CallbackArgs: callbackArgs,
		Functions:    funcs,
	}

	commitResp, code, err := p.CommitUploadInfo(commitRequest)
	if err != nil {
		return commitResp, code, err
	}
	return commitResp, code, nil
}

func (p *Vod) GetUploadAuthWithExpiredTime(expiredTime time.Duration) (*base.SecurityToken2, error) {
	inlinePolicy := new(base.Policy)
	actions := []string{"vod:ApplyUploadInfo", "vod:CommitUploadInfo"}
	resources := make([]string, 0)
	statement := base.NewAllowStatement(actions, resources)
	inlinePolicy.Statement = append(inlinePolicy.Statement, statement)
	return p.SignSts2(inlinePolicy, expiredTime)
}

func (p *Vod) GetUploadAuth() (*base.SecurityToken2, error) {
	return p.GetUploadAuthWithExpiredTime(time.Hour)
}

func (p *Vod) fillCommitUploadInfoResponseWhenError(logId, errMsg string) *response.VodCommitUploadInfoResponse {
	commitUploadInfoRespone := &response.VodCommitUploadInfoResponse{
		ResponseMetadata: &model_base.ResponseMetadata{
			RequestId: logId,
			Service:   "vod",
			Error:     &model_base.ResponseError{Message: errMsg},
		},
	}
	return commitUploadInfoRespone
}

func (p *Vod) Upload(rd io.Reader, size int64, spaceName string, fileType string, fileName, fileExtension string, storageClass int32) (string, string, error, int) {
	if size == 0 {
		return "", "", fmt.Errorf("file size is zero"), http.StatusBadRequest
	}

	applyRequest := &request.VodApplyUploadInfoRequest{SpaceName: spaceName, FileType: fileType, FileName: fileName, FileExtension: fileExtension, StorageClass: storageClass}

	resp, code, err := p.ApplyUploadInfo(applyRequest)
	logId := resp.GetResponseMetadata().GetRequestId()
	if err != nil {
		return logId, "", err, code
	}

	if resp.ResponseMetadata.Error != nil && resp.ResponseMetadata.Error.Code != "0" {
		return logId, "", fmt.Errorf("%+v", resp.ResponseMetadata.Error), code
	}

	uploadAddress := resp.GetResult().GetData().GetUploadAddress()
	if uploadAddress != nil {
		if len(uploadAddress.GetUploadHosts()) == 0 {
			return logId, "", fmt.Errorf("no tos host found"), http.StatusBadRequest
		}
		if len(uploadAddress.GetStoreInfos()) == 0 && (uploadAddress.GetStoreInfos()[0] == nil) {
			return logId, "", fmt.Errorf("no store info found"), http.StatusBadRequest
		}

		tosHost := uploadAddress.GetUploadHosts()[0]
		oid := uploadAddress.StoreInfos[0].GetStoreUri()
		sessionKey := uploadAddress.GetSessionKey()
		auth := uploadAddress.GetStoreInfos()[0].GetAuth()
		client := &http.Client{}

		if int(size) < consts.MinChunckSize {
			bts, err := ioutil.ReadAll(rd)
			if err != nil {
				return logId, "", err, http.StatusBadRequest
			}
			if err := p.directUpload(tosHost, oid, auth, bts, client, storageClass); err != nil {
				return logId, "", err, http.StatusBadRequest
			}
		} else {
			uploadPart := model.UploadPartCommon{
				TosHost: tosHost,
				Oid:     oid,
				Auth:    auth,
			}
			if err := p.chunkUpload(rd, uploadPart, client, size, true, storageClass); err != nil {
				return logId, "", err, http.StatusBadRequest
			}
		}
		return oid, sessionKey, nil, http.StatusOK
	}
	return logId, "", errors.New("upload address not exist"), http.StatusBadRequest
}

func (p *Vod) directUpload(tosHost string, oid string, auth string, fileBytes []byte, client *http.Client, storageClass int32) error {
	checkSum := fmt.Sprintf("%08x", crc32.ChecksumIEEE(fileBytes))
	url := fmt.Sprintf("http://%s/%s", tosHost, oid)
	req, err := http.NewRequest("PUT", url, bytes.NewReader(fileBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-CRC32", checkSum)
	req.Header.Set("Authorization", auth)

	if storageClass == int32(business.StorageClassType_Archive) {
		req.Header.Set("X-Upload-Storage-Class", "archive")
	}

	rsp, err := client.Do(req)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return err
	}
	res := &model.UploadPartCommonResponse{}
	if err := json.Unmarshal(b, res); err != nil {
		return err
	}
	if res.Success != 0 {
		return errors.New(res.Error.Message)
	}
	return nil
}

func (p *Vod) chunkUpload(rd io.Reader, uploadPart model.UploadPartCommon, client *http.Client, size int64, isLargeFile bool, storageClass int32) error {
	uploadID, err := p.initUploadPart(uploadPart.TosHost, uploadPart.Oid, uploadPart.Auth, client, isLargeFile, storageClass)
	if err != nil {
		return err
	}
	cur := make([]byte, consts.MinChunckSize)
	parts := make([]string, 0)

	num := size / consts.MinChunckSize
	cnt := 0
	lastNum := int(num) - 1

	// 读 n-1 片并上传上去
	var part string
	for i := 0; i < lastNum; i++ {
		n, err := io.ReadFull(rd, cur)
		if err != nil {
			return err
		}
		cnt += n
		partNumber := i
		if isLargeFile {
			partNumber++
		}
		err = retry.Do(func() error {
			part, err = p.uploadPart(uploadPart, uploadID, partNumber, cur, client, isLargeFile, storageClass)
			return err
		}, retry.Attempts(3))
		if err != nil {
			return err
		}
		parts = append(parts, part)
	}
	// 读 n 和 n+1片（如有）上传上去
	bts, err := ioutil.ReadAll(rd)
	if err != nil {
		return err
	}
	total := len(bts) + cnt
	if total != int(size) {
		return errors.New(fmt.Sprintf("last part download size mismatch ,download %d , size %d", total, size))
	}
	if isLargeFile {
		lastNum++
	}
	err = retry.Do(func() error {
		part, err = p.uploadPart(uploadPart, uploadID, lastNum, bts, client, isLargeFile, storageClass)
		return err
	}, retry.Attempts(3))
	if err != nil {
		return err
	}
	parts = append(parts, part)
	return p.uploadMergePart(uploadPart, uploadID, parts, client, isLargeFile, storageClass)
}

func (p *Vod) initUploadPart(tosHost string, oid string, auth string, client *http.Client, isLargeFile bool, storageClass int32) (string, error) {
	url := fmt.Sprintf("http://%s/%s?uploads", tosHost, oid)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", auth)
	if isLargeFile {
		req.Header.Set("X-Storage-Mode", "gateway")
	}
	if storageClass == int32(business.StorageClassType_Archive) {
		req.Header.Set("X-Upload-Storage-Class", "archive")
	}
	rsp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return "", err
	}
	res := &model.UploadPartResponse{}
	if err := json.Unmarshal(b, res); err != nil {
		return "", err
	}
	if res.Success != 0 {
		return "", errors.New(res.Error.Message)
	}
	return res.PayLoad.UploadID, nil
}

func (p *Vod) uploadPart(uploadPart model.UploadPartCommon, uploadID string, partNumber int, data []byte, client *http.Client, isLargeFile bool, storageClass int32) (string, error) {
	url := fmt.Sprintf("http://%s/%s?partNumber=%d&uploadID=%s", uploadPart.TosHost, uploadPart.Oid, partNumber, uploadID)
	checkSum := fmt.Sprintf("%08x", crc32.ChecksumIEEE(data))
	req, err := http.NewRequest("PUT", url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-CRC32", checkSum)
	req.Header.Set("Authorization", uploadPart.Auth)
	if isLargeFile {
		req.Header.Set("X-Storage-Mode", "gateway")
	}
	if storageClass == int32(business.StorageClassType_Archive) {
		req.Header.Set("X-Upload-Storage-Class", "archive")
	}

	rsp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return "", err
	}
	res := &model.UploadPartResponse{}
	if err := json.Unmarshal(b, res); err != nil {
		return "", err
	}
	if res.Success != 0 {
		return "", errors.New(res.Error.Message)
	}
	return checkSum, nil
}

func (p *Vod) uploadMergePart(uploadPart model.UploadPartCommon, uploadID string, checkSum []string, client *http.Client, isLargeFile bool, storageClass int32) error {
	url := fmt.Sprintf("http://%s/%s?uploadID=%s", uploadPart.TosHost, uploadPart.Oid, uploadID)
	body, err := p.genMergeBody(checkSum)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewReader([]byte(body)))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", uploadPart.Auth)
	if isLargeFile {
		req.Header.Set("X-Storage-Mode", "gateway")
	}
	if storageClass == int32(business.StorageClassType_Archive) {
		req.Header.Set("X-Upload-Storage-Class", "archive")
	}

	rsp, err := client.Do(req)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return err
	}
	res := &model.UploadMergeResponse{}
	if err := json.Unmarshal(b, res); err != nil {
		return err
	}
	if res.Success != 0 {
		return errors.New(res.Error.Message)
	}
	return nil
}

func (p *Vod) genMergeBody(checkSum []string) (string, error) {
	if len(checkSum) == 0 {
		return "", errors.New("body crc32 empty")
	}
	s := make([]string, len(checkSum))
	for partNumber, crc := range checkSum {
		s[partNumber] = fmt.Sprintf("%d:%s", partNumber, crc)
	}
	return strings.Join(s, ","), nil
}
