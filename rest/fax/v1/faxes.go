/*
 * Twilio - Fax
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
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateFax'
type CreateFaxParams struct {
	// The number the fax was sent from. Can be the phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format or the SIP `from` value. The caller ID displayed to the recipient uses this value. If this is a phone number, it must be a Twilio number or a verified outgoing caller id from your account. If `to` is a SIP address, this can be any alphanumeric string (and also the characters `+`, `_`, `.`, and `-`), which will be used in the `from` header of the SIP request.
	From *string `json:"From,omitempty"`
	// The URL of the PDF that contains the fax. See our [security](https://www.twilio.com/docs/usage/security) page for information on how to ensure the request for your media comes from Twilio.
	MediaUrl *string `json:"MediaUrl,omitempty"`
	// The [Fax Quality value](https://www.twilio.com/docs/fax/api/fax-resource#fax-quality-values) that describes the fax quality. Can be: `standard`, `fine`, or `superfine` and defaults to `fine`.
	Quality *string `json:"Quality,omitempty"`
	// The password to use with `sip_auth_username` to authenticate faxes sent to a SIP address.
	SipAuthPassword *string `json:"SipAuthPassword,omitempty"`
	// The username to use with the `sip_auth_password` to authenticate faxes sent to a SIP address.
	SipAuthUsername *string `json:"SipAuthUsername,omitempty"`
	// The URL we should call using the `POST` method to send [status information](https://www.twilio.com/docs/fax/api/fax-resource#fax-status-callback) to your application when the status of the fax changes.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// Whether to store a copy of the sent media on our servers for later retrieval. Can be: `true` or `false` and the default is `true`.
	StoreMedia *bool `json:"StoreMedia,omitempty"`
	// The phone number to receive the fax in [E.164](https://www.twilio.com/docs/glossary/what-e164) format or the recipient's SIP URI.
	To *string `json:"To,omitempty"`
	// How long in minutes from when the fax is initiated that we should try to send the fax.
	Ttl *int `json:"Ttl,omitempty"`
}

func (params *CreateFaxParams) SetFrom(From string) *CreateFaxParams {
	params.From = &From
	return params
}
func (params *CreateFaxParams) SetMediaUrl(MediaUrl string) *CreateFaxParams {
	params.MediaUrl = &MediaUrl
	return params
}
func (params *CreateFaxParams) SetQuality(Quality string) *CreateFaxParams {
	params.Quality = &Quality
	return params
}
func (params *CreateFaxParams) SetSipAuthPassword(SipAuthPassword string) *CreateFaxParams {
	params.SipAuthPassword = &SipAuthPassword
	return params
}
func (params *CreateFaxParams) SetSipAuthUsername(SipAuthUsername string) *CreateFaxParams {
	params.SipAuthUsername = &SipAuthUsername
	return params
}
func (params *CreateFaxParams) SetStatusCallback(StatusCallback string) *CreateFaxParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreateFaxParams) SetStoreMedia(StoreMedia bool) *CreateFaxParams {
	params.StoreMedia = &StoreMedia
	return params
}
func (params *CreateFaxParams) SetTo(To string) *CreateFaxParams {
	params.To = &To
	return params
}
func (params *CreateFaxParams) SetTtl(Ttl int) *CreateFaxParams {
	params.Ttl = &Ttl
	return params
}

// Create a new fax to send to a phone number or SIP endpoint.
func (c *ApiService) CreateFax(params *CreateFaxParams) (*FaxV1Fax, error) {
	path := "/v1/Faxes"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.MediaUrl != nil {
		data.Set("MediaUrl", *params.MediaUrl)
	}
	if params != nil && params.Quality != nil {
		data.Set("Quality", *params.Quality)
	}
	if params != nil && params.SipAuthPassword != nil {
		data.Set("SipAuthPassword", *params.SipAuthPassword)
	}
	if params != nil && params.SipAuthUsername != nil {
		data.Set("SipAuthUsername", *params.SipAuthUsername)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StoreMedia != nil {
		data.Set("StoreMedia", fmt.Sprint(*params.StoreMedia))
	}
	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}
	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FaxV1Fax{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific fax and its associated media.
func (c *ApiService) DeleteFax(Sid string) error {
	path := "/v1/Faxes/{Sid}"
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

// Fetch a specific fax.
func (c *ApiService) FetchFax(Sid string) (*FaxV1Fax, error) {
	path := "/v1/Faxes/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FaxV1Fax{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFax'
type ListFaxParams struct {
	// Retrieve only those faxes sent from this phone number, specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.
	From *string `json:"From,omitempty"`
	// Retrieve only those faxes sent to this phone number, specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.
	To *string `json:"To,omitempty"`
	// Retrieve only those faxes with a `date_created` that is before or equal to this value, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreatedOnOrBefore *time.Time `json:"DateCreatedOnOrBefore,omitempty"`
	// Retrieve only those faxes with a `date_created` that is later than this value, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreatedAfter *time.Time `json:"DateCreatedAfter,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListFaxParams) SetFrom(From string) *ListFaxParams {
	params.From = &From
	return params
}
func (params *ListFaxParams) SetTo(To string) *ListFaxParams {
	params.To = &To
	return params
}
func (params *ListFaxParams) SetDateCreatedOnOrBefore(DateCreatedOnOrBefore time.Time) *ListFaxParams {
	params.DateCreatedOnOrBefore = &DateCreatedOnOrBefore
	return params
}
func (params *ListFaxParams) SetDateCreatedAfter(DateCreatedAfter time.Time) *ListFaxParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListFaxParams) SetPageSize(PageSize int) *ListFaxParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListFaxParams) SetLimit(Limit int) *ListFaxParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Fax records from the API. Request is executed immediately.
func (c *ApiService) PageFax(params *ListFaxParams, pageToken, pageNumber string) (*ListFaxResponse, error) {
	path := "/v1/Faxes"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}
	if params != nil && params.DateCreatedOnOrBefore != nil {
		data.Set("DateCreatedOnOrBefore", fmt.Sprint((*params.DateCreatedOnOrBefore).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreatedAfter", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
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

	ps := &ListFaxResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Fax records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListFax(params *ListFaxParams) ([]FaxV1Fax, error) {
	if params == nil {
		params = &ListFaxParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageFax(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []FaxV1Fax

	for response != nil {
		records = append(records, response.Faxes...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListFaxResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListFaxResponse)
	}

	return records, err
}

// Streams Fax records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamFax(params *ListFaxParams) (chan FaxV1Fax, error) {
	if params == nil {
		params = &ListFaxParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageFax(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan FaxV1Fax, 1)

	go func() {
		for response != nil {
			for item := range response.Faxes {
				channel <- response.Faxes[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListFaxResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListFaxResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListFaxResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFaxResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateFax'
type UpdateFaxParams struct {
	// The new [status](https://www.twilio.com/docs/fax/api/fax-resource#fax-status-values) of the resource. Can be only `canceled`. This may fail if transmission has already started.
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateFaxParams) SetStatus(Status string) *UpdateFaxParams {
	params.Status = &Status
	return params
}

// Update a specific fax.
func (c *ApiService) UpdateFax(Sid string, params *UpdateFaxParams) (*FaxV1Fax, error) {
	path := "/v1/Faxes/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FaxV1Fax{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
