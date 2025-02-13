/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateCustomerProfileChannelEndpointAssignment'
type CreateCustomerProfileChannelEndpointAssignmentParams struct {
	// The SID of an channel endpoint
	ChannelEndpointSid *string `json:"ChannelEndpointSid,omitempty"`
	// The type of channel endpoint. eg: phone-number
	ChannelEndpointType *string `json:"ChannelEndpointType,omitempty"`
}

func (params *CreateCustomerProfileChannelEndpointAssignmentParams) SetChannelEndpointSid(ChannelEndpointSid string) *CreateCustomerProfileChannelEndpointAssignmentParams {
	params.ChannelEndpointSid = &ChannelEndpointSid
	return params
}
func (params *CreateCustomerProfileChannelEndpointAssignmentParams) SetChannelEndpointType(ChannelEndpointType string) *CreateCustomerProfileChannelEndpointAssignmentParams {
	params.ChannelEndpointType = &ChannelEndpointType
	return params
}

// Create a new Assigned Item.
func (c *ApiService) CreateCustomerProfileChannelEndpointAssignment(CustomerProfileSid string, params *CreateCustomerProfileChannelEndpointAssignmentParams) (*TrusthubV1CustomerProfileChannelEndpointAssignment, error) {
	path := "/v1/CustomerProfiles/{CustomerProfileSid}/ChannelEndpointAssignments"
	path = strings.Replace(path, "{"+"CustomerProfileSid"+"}", CustomerProfileSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChannelEndpointSid != nil {
		data.Set("ChannelEndpointSid", *params.ChannelEndpointSid)
	}
	if params != nil && params.ChannelEndpointType != nil {
		data.Set("ChannelEndpointType", *params.ChannelEndpointType)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1CustomerProfileChannelEndpointAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an Assignment Item Instance.
func (c *ApiService) DeleteCustomerProfileChannelEndpointAssignment(CustomerProfileSid string, Sid string) error {
	path := "/v1/CustomerProfiles/{CustomerProfileSid}/ChannelEndpointAssignments/{Sid}"
	path = strings.Replace(path, "{"+"CustomerProfileSid"+"}", CustomerProfileSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch specific Assigned Item Instance.
func (c *ApiService) FetchCustomerProfileChannelEndpointAssignment(CustomerProfileSid string, Sid string) (*TrusthubV1CustomerProfileChannelEndpointAssignment, error) {
	path := "/v1/CustomerProfiles/{CustomerProfileSid}/ChannelEndpointAssignments/{Sid}"
	path = strings.Replace(path, "{"+"CustomerProfileSid"+"}", CustomerProfileSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1CustomerProfileChannelEndpointAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCustomerProfileChannelEndpointAssignment'
type ListCustomerProfileChannelEndpointAssignmentParams struct {
	// The SID of an channel endpoint
	ChannelEndpointSid *string `json:"ChannelEndpointSid,omitempty"`
	// comma separated list of channel endpoint sids
	ChannelEndpointSids *string `json:"ChannelEndpointSids,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCustomerProfileChannelEndpointAssignmentParams) SetChannelEndpointSid(ChannelEndpointSid string) *ListCustomerProfileChannelEndpointAssignmentParams {
	params.ChannelEndpointSid = &ChannelEndpointSid
	return params
}
func (params *ListCustomerProfileChannelEndpointAssignmentParams) SetChannelEndpointSids(ChannelEndpointSids string) *ListCustomerProfileChannelEndpointAssignmentParams {
	params.ChannelEndpointSids = &ChannelEndpointSids
	return params
}
func (params *ListCustomerProfileChannelEndpointAssignmentParams) SetPageSize(PageSize int) *ListCustomerProfileChannelEndpointAssignmentParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCustomerProfileChannelEndpointAssignmentParams) SetLimit(Limit int) *ListCustomerProfileChannelEndpointAssignmentParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of CustomerProfileChannelEndpointAssignment records from the API. Request is executed immediately.
func (c *ApiService) PageCustomerProfileChannelEndpointAssignment(CustomerProfileSid string, params *ListCustomerProfileChannelEndpointAssignmentParams, pageToken, pageNumber string) (*ListCustomerProfileChannelEndpointAssignmentResponse, error) {
	path := "/v1/CustomerProfiles/{CustomerProfileSid}/ChannelEndpointAssignments"

	path = strings.Replace(path, "{"+"CustomerProfileSid"+"}", CustomerProfileSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChannelEndpointSid != nil {
		data.Set("ChannelEndpointSid", *params.ChannelEndpointSid)
	}
	if params != nil && params.ChannelEndpointSids != nil {
		data.Set("ChannelEndpointSids", *params.ChannelEndpointSids)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCustomerProfileChannelEndpointAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists CustomerProfileChannelEndpointAssignment records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCustomerProfileChannelEndpointAssignment(CustomerProfileSid string, params *ListCustomerProfileChannelEndpointAssignmentParams) ([]TrusthubV1CustomerProfileChannelEndpointAssignment, error) {
	if params == nil {
		params = &ListCustomerProfileChannelEndpointAssignmentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageCustomerProfileChannelEndpointAssignment(CustomerProfileSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []TrusthubV1CustomerProfileChannelEndpointAssignment

	for response != nil {
		records = append(records, response.Results...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListCustomerProfileChannelEndpointAssignmentResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListCustomerProfileChannelEndpointAssignmentResponse)
	}

	return records, err
}

// Streams CustomerProfileChannelEndpointAssignment records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCustomerProfileChannelEndpointAssignment(CustomerProfileSid string, params *ListCustomerProfileChannelEndpointAssignmentParams) (chan TrusthubV1CustomerProfileChannelEndpointAssignment, error) {
	if params == nil {
		params = &ListCustomerProfileChannelEndpointAssignmentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageCustomerProfileChannelEndpointAssignment(CustomerProfileSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan TrusthubV1CustomerProfileChannelEndpointAssignment, 1)

	go func() {
		for response != nil {
			for item := range response.Results {
				channel <- response.Results[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListCustomerProfileChannelEndpointAssignmentResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListCustomerProfileChannelEndpointAssignmentResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListCustomerProfileChannelEndpointAssignmentResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCustomerProfileChannelEndpointAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
