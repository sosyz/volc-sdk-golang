package ipaas

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

var AK = "*****"
var SK = "*******=="
var HOST = "**********"
var client = NewInstance()
var productid = "*********"
var objectlist = []string{"i-********"}
var objecttype = "instance"
var hostlist = []string{"h-*********"}

func TestSample(t *testing.T) {
	ak := "**********"
	sk := "***********=="
	host := "*********"
	DefaultInstance := NewInstance()
	DefaultInstance.Client.SetAccessKey(ak)
	DefaultInstance.Client.SetSecretKey(sk)
	DefaultInstance.SetHost(host)

	ctx := context.Background()

	respProduct, err := DefaultInstance.ListProduct(ctx, &ListProductQuery{})
	if err != nil {
		t.Logf("error occur %v", err)
	}
	fmt.Println(respProduct)

	second := int32(100)
	token1, err := DefaultInstance.AcquireIdempotentToken(ctx, &AcquireIdempotentTokenBody{
		TimeoutSeconds: &second,
	})
	if err != nil {
		t.Logf("error occur %v", err)
	}
	tokenInfo := token1.Result.Token

	// 订购、退订
	deviceName := "你说啥"
	dcDetail := "boe-tjjnsx"
	region := "cn-north"
	dc := []*string{&dcDetail}
	respDevice, err := DefaultInstance.CreateDevices(ctx, &CreateDevicesReq{
		CreateDevicesQuery: &CreateDevicesQuery{
			XIPaaSIdempotentToken: *tokenInfo,
		},
		CreateDevicesBody: &CreateDevicesBody{
			ProductID:       "***********",
			DeviceCount:     3,
			DeviceName:      &deviceName,
			DeviceType:      "bm",
			DevicePackageID: "CloudHost_ARMNode_8c12g_permonth_inner",
			Region:          &region,
			DC:              dc,
			ImageConfig:     CreateDevicesBodyImageConfig{},
			NetConfig: CreateDevicesBodyNetConfig{
				ISP: 7,
			},
			DeviceCharge: CreateDevicesBodyDeviceCharge{
				DeviceChargeMode: CreateDevicesBodyDeviceChargeMode{
					DeviceChargeType: 3,
				},
				NetworkChargeMode: CreateDevicesBodyDeviceChargeNetworkChargeMode{
					BandWidthChargeType: "traffic",
				},
			},
		},
	})
	if err != nil {
		t.Logf("error occur %v", err)
	}
	fmt.Println(respDevice)

	detail11 := true
	respList11, err := DefaultInstance.ListHost(ctx, &ListHostQuery{
		ProductID: "**********",
		Detail:    &detail11,
	})
	if err != nil {
		t.Logf("error occur %v", err)
	}
	fmt.Println(respList11)

	second2 := int32(100)
	token2, err := DefaultInstance.AcquireIdempotentToken(ctx, &AcquireIdempotentTokenBody{
		TimeoutSeconds: &second2,
	})
	if err != nil {
		t.Logf("error occur %v", err)
	}
	tokenInfo2 := token2.Result.Token
	respDeviceDel, err := DefaultInstance.DeleteDevices(ctx, &DeleteDevicesReq{
		DeleteDevicesBody: &DeleteDevicesBody{
			ProductID:    "**********",
			DeviceIDList: []string{"h-***********"},
		},
		DeleteDevicesQuery: &DeleteDevicesQuery{
			XIPaaSIdempotentToken: *tokenInfo2,
		},
	})
	if err != nil {
		t.Logf("error occur %v", err)
	}
	fmt.Println(respDeviceDel)

	detail := true
	respList, err := DefaultInstance.ListHost(ctx, &ListHostQuery{
		ProductID: "**********",
		Detail:    &detail,
	})
	if err != nil {
		t.Logf("error occur %v", err)
	}
	fmt.Println(respList)
}

func TestIPaaS_AdbCommand(t *testing.T) {
	client.Client.SetAccessKey(AK)
	client.Client.SetSecretKey(SK)
	client.SetHost(HOST)

	ctx := context.Background()
	command := "ls -ahl"

	resp, err := client.AdbCommand(ctx, &AdbCommandBody{
		ProductID:    productid,
		Command:      command,
		ObjectType:   objecttype,
		ObjectIDList: objectlist,
	})
	if err != nil {
		t.Logf("error occur %v", err)
	}
	fmt.Println(resp)

	t.Logf("%+v", resp)
	t.Logf("%+v", resp.ResponseMetadata)
	t.Logf("%+v", resp.Result)
	t.Logf("%#v", resp.Result.PassedIDList)
	t.Logf("%#v", resp.Result.FailedIDList)
	fmt.Println(resp.Result.PassedIDList)
}

func TestIPaaS_ListHost(t *testing.T) {
	client.Client.SetAccessKey(AK)
	client.Client.SetSecretKey(SK)
	client.SetHost(HOST)

	ctx := context.Background()

	resp, err := client.ListHost(ctx, &ListHostQuery{
		ProductID: productid,
	})
	if err != nil {
		t.Logf("error occur %v", err)
	}
	fmt.Println(resp)

	t.Logf("%+v", resp)
	t.Logf("%#v", resp)
	t.Logf("%#v", *resp.Result.Row[0].Bandwidth)
	t.Logf("type of %T", resp.Result.Row)
	//for _, item := range resp.Result.Row {
	//	fmt.Println("Bandwidth:", *item.Bandwidth)
	//}
	//for _, item := range resp.Result.Row {
	//	v := reflect.ValueOf(item).Elem()
	//	t := v.Type()
	//
	//	for i := 0; i < v.NumField(); i++ {
	//		field := v.Field(i)
	//		fieldName := t.Field(i).Name
	//		fieldValue := fmt.Sprintf("%v", field.Interface())
	//
	//		fmt.Println(fieldName+":", fieldValue)
	//	}
	//
	//	fmt.Println()
	//}

	for _, item := range resp.Result.Row {
		{
			v := reflect.ValueOf(item).Elem()
			t := v.Type()

			for i := 0; i < v.NumField(); i++ {
				field := v.Field(i)
				fieldName := t.Field(i).Name
				fieldValue := fmt.Sprintf("%v", field.Elem())

				fmt.Println(fieldName+":", fieldValue)
			}

			fmt.Println()
		}
	}

}

func TestIPaaS_RebootHost(t *testing.T) {
	client.Client.SetAccessKey(AK)
	client.Client.SetSecretKey(SK)
	client.SetHost(HOST)

	ctx := context.Background()
	value := true

	resp, err := client.RebootHost(ctx, &RebootHostBody{
		HostIDList: []string{"h-**********"},
		ProductID:  productid,
		Force:      &value,
	})
	if err != nil {
		t.Logf("error occur %v", err)
	}

	t.Logf("%+v", resp)
	t.Logf("%#v", resp)
	v := reflect.ValueOf(resp.Result).Elem()
	p := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := p.Field(i).Name

		if field.Kind() == reflect.Slice {
			for j := 0; j < field.Len(); j++ {
				elem := field.Index(j)
				elemValue := fmt.Sprintf("%v", elem.Elem())

				fmt.Printf("%s[%d]: %s\n", fieldName, j, elemValue)
			}
		}
	}

}

func TestForCloudGame(t *testing.T) {
	ak := "************"
	sk := "***********=="
	host := "*************"
	DefaultInstance := NewInstance()
	DefaultInstance.Client.SetAccessKey(ak)
	DefaultInstance.Client.SetSecretKey(sk)
	DefaultInstance.SetHost(host)

	ctx := context.Background()
	second := int32(100)
	token1, err := DefaultInstance.AcquireIdempotentToken(ctx, &AcquireIdempotentTokenBody{
		TimeoutSeconds: &second,
	})
	if err != nil {
		t.Logf("error occur %v", err)
	}
	tokenInfo := token1.Result.Token

	// 订购、退订
	dcDetail := "****-*****"
	dc := []*string{&dcDetail}
	imageID := "img-******"
	isPublic := true
	bandwidth := int32(5)
	respDevice, err := DefaultInstance.CreateDevices(ctx, &CreateDevicesReq{
		CreateDevicesQuery: &CreateDevicesQuery{
			XIPaaSIdempotentToken: *tokenInfo,
		},
		CreateDevicesBody: &CreateDevicesBody{
			ProductID:   "*********",
			DeviceCount: 1,
			DeviceType:  "container",

			// CloudHost_ARMNode_2c3g_daily_inner
			//CloudHost_ARMNode_3c4g_daily_inner
			//CloudHost_ARMNode_4c6g_daily_inner
			//CloudHost_ARMNode_8c12g_daily_inner
			DevicePackageID: "CloudHost_ARMNode_2c3g_daily_inner",
			DC:              dc,
			ImageConfig: CreateDevicesBodyImageConfig{
				ImageID:       &imageID,
				IsPublicImage: &isPublic,
			},
			NetConfig: CreateDevicesBodyNetConfig{
				ISP:       7,
				Bandwidth: &bandwidth,
			},
			DeviceCharge: CreateDevicesBodyDeviceCharge{
				DeviceChargeMode: CreateDevicesBodyDeviceChargeMode{
					DeviceChargeType: 2,
				},
				// 'traffic','daily_peak','bandwidth','95th_percentile'
				NetworkChargeMode: CreateDevicesBodyDeviceChargeNetworkChargeMode{
					BandWidthChargeType: "daily_peak",
				},
			},
		},
	})
	if err != nil {
		t.Logf("error occur %v", err)
	}
	fmt.Println(respDevice)

	count := int32(1)
	resp2, err := DefaultInstance.GetInfoAfterOrder(ctx, &GetInfoAfterOrderBody{
		ProductID: "**********",
		OrderNo:   *respDevice.Result.OrderNo,
		Count:     &count,
	})
	if err != nil {
		t.Logf("error occur %v", err)
	}
	fmt.Println(resp2)
}

func TestReconfigureDevicesPackage(t *testing.T) {
	ak := "**"
	sk := "**"
	host := ""
	DefaultInstance := NewInstance()
	DefaultInstance.Client.SetAccessKey(ak)
	DefaultInstance.Client.SetSecretKey(sk)
	DefaultInstance.SetHost(host)

	ctx := context.Background()
	resp, err := DefaultInstance.ReconfigureDevicesPackage(ctx, &ReconfigureDevicesPackageBody{
		ProductID: "*",
		ReconfigureParam: []ReconfigureDevicesPackageBodyReconfigureParamItem{
			{
				DeviceID: "h-*",
				InstancesSpec: []ReconfigureDevicesPackageBodyReconfigureParamPropertiesItemsItem{
					{
						BandwidthMbps:      10,
						DisplayResolutionX: 720,
						DisplayResolutionY: 1080,
						ImageID:            "img-*",
						InstanceID:         "i-*",
						Properties:         nil,
					},
				},
				PackageID: "pack-*",
				TimeoutSeconds: func() *int32 {
					var a int32 = 500
					return &a
				}(),
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%s", func() string { s, _ := json.Marshal(resp); return string(s) }())
}

func TestListTaskDetail(t *testing.T) {
	ak := "**"
	sk := "**"
	host := ""
	DefaultInstance := NewInstance()
	DefaultInstance.Client.SetAccessKey(ak)
	DefaultInstance.Client.SetSecretKey(sk)
	DefaultInstance.SetHost(host)

	ctx := context.Background()

	resp, err := DefaultInstance.ListTaskInfo(ctx, &ListTaskInfoQuery{
		ProductID:    "fake",
		Count:        nil,
		CreateAfter:  nil,
		CreateBefore: nil,
		DeviceID:     nil,
		InstanceID:   nil,
		JobID:        nil,
		Offset:       nil,
		StatusIn:     nil,
		TaskID:       nil,
		TypeIn:       nil,
		UpdateAfter:  nil,
		UpdateBefore: nil,
	})

	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%s", func() string { s, _ := json.Marshal(resp); return string(s) }())

}
