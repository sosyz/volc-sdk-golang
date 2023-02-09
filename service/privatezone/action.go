// Code Generated by gadget/xsdk, DO NOT EDIT

package privatezone

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// Group: Other
// 其他API

func (c *Client) ListRegions(ctx context.Context, data *ListRegionsRequest) (*ListRegionsResponse, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/?Action=ListRegions", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payload ListRegionsResponse
	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

// Group: PrivateZones
// 私有域名管理

func (c *Client) BindVPC(ctx context.Context, data *BindVPCRequest) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/?Action=BindVPC", bytes.NewReader(body))
	if err != nil {
		return err
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) CreatePrivateZone(ctx context.Context, data *CreatePrivateZoneRequest) (*CreatePrivateZoneResponse, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/?Action=CreatePrivateZone", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payload CreatePrivateZoneResponse
	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) DeletePrivateZone(ctx context.Context, data *DeletePrivateZoneRequest) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/?Action=DeletePrivateZone", bytes.NewReader(body))
	if err != nil {
		return err
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) IncBindVPC(ctx context.Context, data *IncBindVPCRequest) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/?Action=IncBindVPC", bytes.NewReader(body))
	if err != nil {
		return err
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) ListBindVPC(ctx context.Context, data *ListBindVPCRequest) (*ListBindVPCResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/?Action=ListBindVPC", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if v := data.ZID; v != nil {
		q.Add("ZID", *v)
	}
	if v := data.Region; v != nil {
		q.Add("Region", *v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payload ListBindVPCResponse
	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) ListPrivateZones(ctx context.Context, data *ListPrivateZonesRequest) (*ListPrivateZonesResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/?Action=ListPrivateZones", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if v := data.Region; v != nil {
		q.Add("Region", *v)
	}
	if v := data.RecursionMode; v != nil {
		q.Add("RecursionMode", *v)
	}
	if v := data.LineMode; v != nil {
		q.Add("LineMode", *v)
	}
	if v := data.PageNumber; v != nil {
		q.Add("PageNumber", *v)
	}
	if v := data.PageSize; v != nil {
		q.Add("PageSize", *v)
	}
	if v := data.SearchMode; v != nil {
		q.Add("SearchMode", *v)
	}
	if v := data.KeyWord; v != nil {
		q.Add("KeyWord", *v)
	}
	if v := data.VpcID; v != nil {
		q.Add("VpcID", *v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payload ListPrivateZonesResponse
	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) QueryPrivateZone(ctx context.Context, data *QueryPrivateZoneRequest) (*QueryPrivateZoneResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/?Action=QueryPrivateZone", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if v := data.ZID; v != nil {
		q.Add("ZID", *v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payload QueryPrivateZoneResponse
	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) UpdatePrivateZone(ctx context.Context, data *UpdatePrivateZoneRequest) (*UpdatePrivateZoneResponse, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/?Action=UpdatePrivateZone", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payload UpdatePrivateZoneResponse
	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

// Group: Records
// PrivateZone解析记录管理

func (c *Client) BatchCreateRecord(ctx context.Context, data *BatchCreateRecordRequest) (*BatchCreateRecordResponse, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/?Action=BatchCreateRecord", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payload BatchCreateRecordResponse
	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) BatchDeleteRecord(ctx context.Context, data *BatchDeleteRecordRequest) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/?Action=BatchDeleteRecord", bytes.NewReader(body))
	if err != nil {
		return err
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) BatchUpdateRecord(ctx context.Context, data *BatchUpdateRecordRequest) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/?Action=BatchUpdateRecord", bytes.NewReader(body))
	if err != nil {
		return err
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) CreateRecord(ctx context.Context, data *CreateRecordRequest) (*CreateRecordResponse, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/?Action=CreateRecord", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payload CreateRecordResponse
	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) DeleteRecord(ctx context.Context, data *DeleteRecordRequest) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/?Action=DeleteRecord", bytes.NewReader(body))
	if err != nil {
		return err
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) ListRecordAttributes(ctx context.Context, data *ListRecordAttributesRequest) (*ListRecordAttributesResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/?Action=ListRecordAttributes", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if v := data.ZID; v != nil {
		q.Add("ZID", *v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payload ListRecordAttributesResponse
	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) ListRecordSets(ctx context.Context, data *ListRecordSetsRequest) (*ListRecordSetsResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/?Action=ListRecordSets", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if v := data.PageSize; v != nil {
		q.Add("PageSize", *v)
	}
	if v := data.PageNumber; v != nil {
		q.Add("PageNumber", *v)
	}
	if v := data.ZID; v != nil {
		q.Add("ZID", *v)
	}
	if v := data.Host; v != nil {
		q.Add("Host", *v)
	}
	if v := data.RecordSetID; v != nil {
		q.Add("RecordSetID", *v)
	}
	if v := data.SearchMode; v != nil {
		q.Add("SearchMode", *v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payload ListRecordSetsResponse
	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) ListRecords(ctx context.Context, data *ListRecordsRequest) (*ListRecordsResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/?Action=ListRecords", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if v := data.ZID; v != nil {
		q.Add("ZID", *v)
	}
	if v := data.PageSize; v != nil {
		q.Add("PageSize", *v)
	}
	if v := data.SearchMode; v != nil {
		q.Add("SearchMode", *v)
	}
	if v := data.Host; v != nil {
		q.Add("Host", *v)
	}
	if v := data.Name; v != nil {
		q.Add("Name", *v)
	}
	if v := data.Value; v != nil {
		q.Add("Value", *v)
	}
	if v := data.Line; v != nil {
		q.Add("Line", *v)
	}
	if v := data.Type; v != nil {
		q.Add("Type", *v)
	}
	if v := data.LastOperator; v != nil {
		q.Add("LastOperator", *v)
	}
	if v := data.PageNumber; v != nil {
		q.Add("PageNumber", *v)
	}
	if v := data.SearchOrder; v != nil {
		q.Add("SearchOrder", *v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payload ListRecordsResponse
	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) UpdateRecord(ctx context.Context, data *UpdateRecordRequest) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/?Action=UpdateRecord", bytes.NewReader(body))
	if err != nil {
		return err
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) UpdateRecordSet(ctx context.Context, data *UpdateRecordSetRequest) (*UpdateRecordSetResponse, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/?Action=UpdateRecordSet", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payload UpdateRecordSetResponse
	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&payload); err != nil {
		return nil, err
	}
	return &payload, nil
}