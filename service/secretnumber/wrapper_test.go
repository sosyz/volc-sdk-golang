package secretnumber

import (
	"testing"
)

const (
	testAk = "your ak"
	testSk = "your sk"
)

func init() {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	DefaultDataCenterInstance.Client.SetAccessKey(testAk)
	DefaultDataCenterInstance.Client.SetSecretKey(testSk)
	DefaultNumberPoolInstance.Client.SetAccessKey(testAk)
	DefaultNumberPoolInstance.Client.SetSecretKey(testSk)
	DefaultMercServiceInstance.Client.SetAccessKey(testAk)
	DefaultMercServiceInstance.Client.SetSecretKey(testSk)
	DefaultConfigServiceInstance.Client.SetAccessKey(testAk)
	DefaultConfigServiceInstance.Client.SetSecretKey(testSk)
}

func TestSecretNumber_BindAXB(t *testing.T) {
	req := &BindAXBRequest{
		NumberPoolNo: "NP161156328504091435",
		PhoneNoA:     "188xxxx1647",
		PhoneNoB:     "137xxxx8258",
		PhoneNoX:     "170xxxx3159",
		ExpireTime:   1632313906,
	}
	result, statusCode, err := DefaultInstance.BindAXB(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSecretNumber_SelectNumberAndBindAXBindAXB(t *testing.T) {
	req := &SelectNumberAndBindAXBRequest{
		NumberPoolNo:    "NP161156328504091435",
		PhoneNoA:        "188xxxx1647",
		PhoneNoB:        "137xxxx8258",
		ExpireTime:      1632313906,
		CityCode:        "010",
		DegradeCityList: "010,020",
	}
	result, statusCode, err := DefaultInstance.SelectNumberAndBindAXB(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSecretNumber_UnbindAXB(t *testing.T) {
	req := &SpecificSubIdRequest{
		NumberPoolNo: "NP161156328504091435",
		SubId:        "S1632298175315938db419",
	}
	result, statusCode, err := DefaultInstance.UnbindAXB(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSecretNumber_QuerySubscription(t *testing.T) {
	req := &SpecificSubIdRequest{
		NumberPoolNo: "NP161156328504091435",
		SubId:        "S163229841631591737db3",
	}
	result, statusCode, err := DefaultInstance.QuerySubscription(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSecretNumber_QuerySubscriptionForList(t *testing.T) {
	req := &QuerySubscriptionForListRequest{
		NumberPoolNo: "NP161156328504091435",
		PhoneNoX:     "170xxxx3159",
		Status:       1,
		Offset:       0,
		Limit:        20,
	}
	result, statusCode, err := DefaultInstance.QuerySubscriptionForList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSecretNumber_UpgradeAXToAXB(t *testing.T) {
	req := &UpgradeAXToAXBRequest{
		NumberPoolNo: "NP161156328504091435",
		SubId:        "S16323073363159890f81b",
		PhoneNoB:     "131xxxx7582",
	}
	result, statusCode, err := DefaultInstance.UpgradeAXToAXB(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSecretNumber_UpdateAXB(t *testing.T) {
	req := &UpdateAXBRequest{
		NumberPoolNo: "NP161156328504091435",
		SubId:        "S16323075803159b97ba7a",
		UpdateType:   "updateExpireTime",
		ExpireTime:   1632313906,
	}
	result, statusCode, err := DefaultInstance.UpdateAXB(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSecretNumber_BindAXN(t *testing.T) {
	req := &BindAXNRequest{
		NumberPoolNo: "NP162981168404095092",
		PhoneNoA:     "188xxxx1647",
		PhoneNoX:     "170xxxx8991",
		PhoneNoB:     "137xxxx8258",
		ExpireTime:   1632313906,
	}
	result, statusCode, err := DefaultInstance.BindAXN(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSecretNumber_SelectNumberAndBindAXN(t *testing.T) {
	req := &SelectNumberAndBindAXNRequest{
		NumberPoolNo: "NP162981168404095092",
		PhoneNoA:     "188xxxx5753",
		PhoneNoB:     "137xxxx8257",
		ExpireTime:   1642672859,
	}
	result, statusCode, err := DefaultInstance.SelectNumberAndBindAXN(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSecretNumber_UpdateAXN(t *testing.T) {
	req := &UpdateAXNRequest{
		NumberPoolNo: "NP162981168404095092",
		SubId:        "S16323109008991825a8b7",
		UpdateType:   "updatePhoneNoA",
		PhoneNoA:     "188xxxx5753",
	}
	result, statusCode, err := DefaultInstance.UpdateAXN(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSecretNumber_UnbindAXN(t *testing.T) {
	req := &SpecificSubIdRequest{
		NumberPoolNo: "NP162981168404095092",
		SubId:        "S16323109008991825a8b7",
	}
	result, statusCode, err := DefaultInstance.UnbindAXN(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSecretNumber_Click2Call(t *testing.T) {
	req := &Click2CallRequest{
		Caller:             "137XXXX8257",
		Callee:             "158XXXX9130",
		CallerNumberPoolNo: "NP163517154204092175",
		CalleeNumberPoolNo: "NP163517154204092175",
	}
	result, statusCode, err := DefaultInstance.Click2Call(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSecretNumber_CancelClick2Call(t *testing.T) {
	req := &CancelClick2CallRequest{
		CallId: "Dccfebdedfe",
	}
	result, statusCode, err := DefaultInstance.CancelClick2Call(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSecretNumber_Click2CallLite(t *testing.T) {
	req := &Click2CallLiteRequest{
		Caller:       "137XXXX8257",
		Callee:       "158XXXX9130",
		NumberPoolNo: "NPXXXXX810901043",
	}
	result, statusCode, err := DefaultInstance.Click2CallLite(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestDataCenter_QueryAudioRecordFileUrl(t *testing.T) {
	req := &QueryAudioRecordFileUrlRequest{
		CallId: "*",
	}
	result, statusCode, err := DefaultDataCenterInstance.QueryAudioRecordFileUrl(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestDataCenter_QueryAudioRecordToTextFileUrl(t *testing.T) {
	req := &QueryAudioRecordToTextFileRequest{
		CallIdList: "Vcc01b1fe30f4868c4b7c9742bdd036f12559,Vcc017784d0be628542aa9a676af3d8fa2e06",
	}
	result, statusCode, err := DefaultDataCenterInstance.QueryAudioRecordToTextFileUrl(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestDataCenter_QueryCallRecordMsg(t *testing.T) {
	req := &QueryCallRecordMsgRequest{
		CallIdList: "S164275060051193662574_1dc1f1b3a8891a8b,S1015_c13f7b27b41e7773",
	}
	result, statusCode, err := DefaultDataCenterInstance.QueryCallRecordMsg(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestNumberPool_CreateNumberPool(t *testing.T) {
	req := &CreateNumberPoolRequest{
		Name:           "测试创建号码池",
		ServiceType:    "100",
		SubServiceType: "102",
	}
	result, statusCode, err := DefaultNumberPoolInstance.CreateNumberPool(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestNumberPool_UpdateNumberPool(t *testing.T) {
	req := &UpdateNumberPoolRequest{
		NumberPoolNo:   "NP164612245301914680",
		Name:           "测试创建号码池2",
		ServiceType:    "200",
		SubServiceType: "201",
	}
	result, statusCode, err := DefaultNumberPoolInstance.UpdateNumberPool(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestNumberPool_NumberPoolList(t *testing.T) {
	req := &NumberPoolListRequest{
		SubServiceType: "101",
		Limit:          "10",
		Offset:         "0",
	}
	result, statusCode, err := DefaultNumberPoolInstance.NumberPoolList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestNumberPool_NumberList(t *testing.T) {
	req := &NumberListRequest{
		Number: "17192359311",
		Limit:  "3",
		Offset: "0",
	}
	result, statusCode, err := DefaultNumberPoolInstance.NumberList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestNumberPool_EnableOrDisableNumber(t *testing.T) {
	req := &EnableOrDisableNumberRequest{
		NumberList: "18792770474,18792770475",
		EnableCode: "1",
	}

	result, statusCode, err := DefaultNumberPoolInstance.EnableOrDisableNumber(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestNumberPool_QueryNumberApplyRecordList(t *testing.T) {
	req := &QueryNumberApplyRecordListRequest{
		ApplyBillId: "NA5345832249",
		Limit:       "10",
		Offset:      "0",
	}

	result, statusCode, err := DefaultNumberPoolInstance.QueryNumberApplyRecordList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestMercService_CreateNumberApplication(t *testing.T) {

	detailList := []NumberApplicationCityItem{
		{
			CityCode:       "010",
			CountryIsoCode: "CN",
			NumberCount:    "1",
		},
	}

	req := &CreateNumberApplicationRequest{
		QualificationNo:               "QUA164679537228220075",
		NumberPoolNo:                  "NP164015715228226293",
		NumberPurpose:                 "1",
		NumberType:                    "1",
		SubServiceType:                "101",
		Remark:                        "remark3",
		NumberApplicationCityItemList: detailList,
	}
	result, statusCode, err := DefaultMercServiceInstance.CreateNumberApplication(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestConfigService_AddQualification(t *testing.T) {
	qualificationMainInfoFormDO := QualificationMainInfoFormDO{
		QualificationEntity:                            "3",
		CertificateThreeInOne:                          1,
		EnterpriseAddress:                              "222",
		LegalRepresentativeName:                        "2",
		LegalRepresentativeId:                          "2",
		UnitSocialCreditCode:                           "2",
		LegalRepresentativeFrontIdPhotoFileCode:        "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		DocOfNumberApplyPhotoFileCode:                  "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		CommitmentLetterOfNetAccessPhotoFileCode:       "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		ThreeInOneBusinessLicensePhotoFileCode:         "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		CodeOfOrganizationCertificate:                  "2",
		BusinessLicensePhotoFileCode:                   "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		CertificateOfOrganizationCodesPhotoFileCode:    "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		CertificateOfTaxationRegistrationPhotoFileCode: "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
	}

	qualificationAdminInfoFormDO := QualificationAdminInfoFormDO{
		Name:                          "2",
		ContactNumber:                 "2",
		IdCardNumber:                  "2",
		IdCardFrontPhotoFileCode:      "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		IdCardPhotoWithPeopleFileCode: "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
	}

	qualificationScenarioInfoFormDOList := []QualificationScenarioInfoFormDO{
		{
			SceneType:         1001,
			Description:       "2",
			ScenarioOfCalling: "2",
		},
	}

	req := &AddQualificationRequest{
		QualificationMainInfoFormDO:         qualificationMainInfoFormDO,
		QualificationAdminInfoFormDO:        qualificationAdminInfoFormDO,
		QualificationScenarioInfoFormDOList: qualificationScenarioInfoFormDOList,
	}

	result, statusCode, err := DefaultConfigServiceInstance.AddQualification(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestConfigService_UpdateQualification(t *testing.T) {
	qualificationMainInfoFormDO := QualificationMainInfoFormDO{
		QualificationEntity:                            "3",
		QualificationNo:                                "QUA164664762128220429",
		CertificateThreeInOne:                          1,
		EnterpriseAddress:                              "22",
		LegalRepresentativeName:                        "2",
		LegalRepresentativeId:                          "2",
		UnitSocialCreditCode:                           "2",
		LegalRepresentativeFrontIdPhotoFileCode:        "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		DocOfNumberApplyPhotoFileCode:                  "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		CommitmentLetterOfNetAccessPhotoFileCode:       "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		ThreeInOneBusinessLicensePhotoFileCode:         "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		CodeOfOrganizationCertificate:                  "2",
		BusinessLicensePhotoFileCode:                   "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		CertificateOfOrganizationCodesPhotoFileCode:    "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		CertificateOfTaxationRegistrationPhotoFileCode: "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
	}

	qualificationAdminInfoFormDO := QualificationAdminInfoFormDO{
		Name:                          "2",
		ContactNumber:                 "2",
		IdCardNumber:                  "2",
		IdCardFrontPhotoFileCode:      "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		IdCardPhotoWithPeopleFileCode: "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
	}

	qualificationScenarioInfoFormDOList := []QualificationScenarioInfoFormDO{
		{
			SceneType:         1001,
			Description:       "2",
			ScenarioOfCalling: "2",
		},
	}

	req := &UpdateQualificationRequest{
		QualificationMainInfoFormDO:         qualificationMainInfoFormDO,
		QualificationAdminInfoFormDO:        qualificationAdminInfoFormDO,
		QualificationScenarioInfoFormDOList: qualificationScenarioInfoFormDOList,
	}

	result, statusCode, err := DefaultConfigServiceInstance.UpdateQualification(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestConfigService_AddQualificationScene(t *testing.T) {

	qualificationScenarioInfoFormDOList := []QualificationScenarioInfoFormDO{
		{
			QualificationNo:   "QUA164664762128220429",
			SceneType:         1001,
			Description:       "22212",
			ScenarioOfCalling: "2222",
		},
	}

	req := &AddQualificationSceneRequest{
		QualificationScenarioInfoFormDOList: qualificationScenarioInfoFormDOList,
	}

	result, statusCode, err := DefaultConfigServiceInstance.AddQualificationScene(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestConfigService_UpdateQualificationScene(t *testing.T) {

	qualificationScenarioInfoFormDOList := []QualificationScenarioInfoFormDO{
		{
			QualificationNo:   "QUA164664762128220429",
			SceneType:         1001,
			Description:       "2221",
			ScenarioOfCalling: "2222",
		},
	}

	req := &UpdateQualificationSceneRequest{
		QualificationScenarioInfoFormDOList: qualificationScenarioInfoFormDOList,
	}

	result, statusCode, err := DefaultConfigServiceInstance.UpdateQualificationScene(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestConfigService_QueryQualification(t *testing.T) {

	qualificationNoList := []string{"QUA164679537228220075"}

	req := &QueryQualificationRequest{
		QualificationNoList: qualificationNoList,
		ApprovalStatus:      "2",
		Limit:               20,
	}

	result, statusCode, err := DefaultConfigServiceInstance.QueryQualification(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
