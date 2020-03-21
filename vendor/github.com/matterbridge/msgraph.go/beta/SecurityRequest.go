// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// SecurityRequestBuilder is request builder for Security
type SecurityRequestBuilder struct{ BaseRequestBuilder }

// Request returns SecurityRequest
func (b *SecurityRequestBuilder) Request() *SecurityRequest {
	return &SecurityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SecurityRequest is request for Security
type SecurityRequest struct{ BaseRequest }

// Get performs GET request for Security
func (r *SecurityRequest) Get(ctx context.Context) (resObj *Security, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Security
func (r *SecurityRequest) Update(ctx context.Context, reqObj *Security) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Security
func (r *SecurityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Alerts returns request builder for Alert collection
func (b *SecurityRequestBuilder) Alerts() *SecurityAlertsCollectionRequestBuilder {
	bb := &SecurityAlertsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/alerts"
	return bb
}

// SecurityAlertsCollectionRequestBuilder is request builder for Alert collection
type SecurityAlertsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Alert collection
func (b *SecurityAlertsCollectionRequestBuilder) Request() *SecurityAlertsCollectionRequest {
	return &SecurityAlertsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Alert item
func (b *SecurityAlertsCollectionRequestBuilder) ID(id string) *AlertRequestBuilder {
	bb := &AlertRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SecurityAlertsCollectionRequest is request for Alert collection
type SecurityAlertsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Alert collection
func (r *SecurityAlertsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]Alert, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Alert
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []Alert
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for Alert collection
func (r *SecurityAlertsCollectionRequest) Get(ctx context.Context) ([]Alert, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for Alert collection
func (r *SecurityAlertsCollectionRequest) Add(ctx context.Context, reqObj *Alert) (resObj *Alert, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// CloudAppSecurityProfiles returns request builder for CloudAppSecurityProfile collection
func (b *SecurityRequestBuilder) CloudAppSecurityProfiles() *SecurityCloudAppSecurityProfilesCollectionRequestBuilder {
	bb := &SecurityCloudAppSecurityProfilesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/cloudAppSecurityProfiles"
	return bb
}

// SecurityCloudAppSecurityProfilesCollectionRequestBuilder is request builder for CloudAppSecurityProfile collection
type SecurityCloudAppSecurityProfilesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CloudAppSecurityProfile collection
func (b *SecurityCloudAppSecurityProfilesCollectionRequestBuilder) Request() *SecurityCloudAppSecurityProfilesCollectionRequest {
	return &SecurityCloudAppSecurityProfilesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CloudAppSecurityProfile item
func (b *SecurityCloudAppSecurityProfilesCollectionRequestBuilder) ID(id string) *CloudAppSecurityProfileRequestBuilder {
	bb := &CloudAppSecurityProfileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SecurityCloudAppSecurityProfilesCollectionRequest is request for CloudAppSecurityProfile collection
type SecurityCloudAppSecurityProfilesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for CloudAppSecurityProfile collection
func (r *SecurityCloudAppSecurityProfilesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]CloudAppSecurityProfile, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []CloudAppSecurityProfile
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []CloudAppSecurityProfile
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for CloudAppSecurityProfile collection
func (r *SecurityCloudAppSecurityProfilesCollectionRequest) Get(ctx context.Context) ([]CloudAppSecurityProfile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for CloudAppSecurityProfile collection
func (r *SecurityCloudAppSecurityProfilesCollectionRequest) Add(ctx context.Context, reqObj *CloudAppSecurityProfile) (resObj *CloudAppSecurityProfile, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DomainSecurityProfiles returns request builder for DomainSecurityProfile collection
func (b *SecurityRequestBuilder) DomainSecurityProfiles() *SecurityDomainSecurityProfilesCollectionRequestBuilder {
	bb := &SecurityDomainSecurityProfilesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/domainSecurityProfiles"
	return bb
}

// SecurityDomainSecurityProfilesCollectionRequestBuilder is request builder for DomainSecurityProfile collection
type SecurityDomainSecurityProfilesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DomainSecurityProfile collection
func (b *SecurityDomainSecurityProfilesCollectionRequestBuilder) Request() *SecurityDomainSecurityProfilesCollectionRequest {
	return &SecurityDomainSecurityProfilesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DomainSecurityProfile item
func (b *SecurityDomainSecurityProfilesCollectionRequestBuilder) ID(id string) *DomainSecurityProfileRequestBuilder {
	bb := &DomainSecurityProfileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SecurityDomainSecurityProfilesCollectionRequest is request for DomainSecurityProfile collection
type SecurityDomainSecurityProfilesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DomainSecurityProfile collection
func (r *SecurityDomainSecurityProfilesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DomainSecurityProfile, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DomainSecurityProfile
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []DomainSecurityProfile
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for DomainSecurityProfile collection
func (r *SecurityDomainSecurityProfilesCollectionRequest) Get(ctx context.Context) ([]DomainSecurityProfile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DomainSecurityProfile collection
func (r *SecurityDomainSecurityProfilesCollectionRequest) Add(ctx context.Context, reqObj *DomainSecurityProfile) (resObj *DomainSecurityProfile, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// FileSecurityProfiles returns request builder for FileSecurityProfile collection
func (b *SecurityRequestBuilder) FileSecurityProfiles() *SecurityFileSecurityProfilesCollectionRequestBuilder {
	bb := &SecurityFileSecurityProfilesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/fileSecurityProfiles"
	return bb
}

// SecurityFileSecurityProfilesCollectionRequestBuilder is request builder for FileSecurityProfile collection
type SecurityFileSecurityProfilesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for FileSecurityProfile collection
func (b *SecurityFileSecurityProfilesCollectionRequestBuilder) Request() *SecurityFileSecurityProfilesCollectionRequest {
	return &SecurityFileSecurityProfilesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for FileSecurityProfile item
func (b *SecurityFileSecurityProfilesCollectionRequestBuilder) ID(id string) *FileSecurityProfileRequestBuilder {
	bb := &FileSecurityProfileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SecurityFileSecurityProfilesCollectionRequest is request for FileSecurityProfile collection
type SecurityFileSecurityProfilesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for FileSecurityProfile collection
func (r *SecurityFileSecurityProfilesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]FileSecurityProfile, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []FileSecurityProfile
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []FileSecurityProfile
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for FileSecurityProfile collection
func (r *SecurityFileSecurityProfilesCollectionRequest) Get(ctx context.Context) ([]FileSecurityProfile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for FileSecurityProfile collection
func (r *SecurityFileSecurityProfilesCollectionRequest) Add(ctx context.Context, reqObj *FileSecurityProfile) (resObj *FileSecurityProfile, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// HostSecurityProfiles returns request builder for HostSecurityProfile collection
func (b *SecurityRequestBuilder) HostSecurityProfiles() *SecurityHostSecurityProfilesCollectionRequestBuilder {
	bb := &SecurityHostSecurityProfilesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/hostSecurityProfiles"
	return bb
}

// SecurityHostSecurityProfilesCollectionRequestBuilder is request builder for HostSecurityProfile collection
type SecurityHostSecurityProfilesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for HostSecurityProfile collection
func (b *SecurityHostSecurityProfilesCollectionRequestBuilder) Request() *SecurityHostSecurityProfilesCollectionRequest {
	return &SecurityHostSecurityProfilesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for HostSecurityProfile item
func (b *SecurityHostSecurityProfilesCollectionRequestBuilder) ID(id string) *HostSecurityProfileRequestBuilder {
	bb := &HostSecurityProfileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SecurityHostSecurityProfilesCollectionRequest is request for HostSecurityProfile collection
type SecurityHostSecurityProfilesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for HostSecurityProfile collection
func (r *SecurityHostSecurityProfilesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]HostSecurityProfile, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []HostSecurityProfile
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []HostSecurityProfile
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for HostSecurityProfile collection
func (r *SecurityHostSecurityProfilesCollectionRequest) Get(ctx context.Context) ([]HostSecurityProfile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for HostSecurityProfile collection
func (r *SecurityHostSecurityProfilesCollectionRequest) Add(ctx context.Context, reqObj *HostSecurityProfile) (resObj *HostSecurityProfile, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// IPSecurityProfiles returns request builder for IPSecurityProfile collection
func (b *SecurityRequestBuilder) IPSecurityProfiles() *SecurityIPSecurityProfilesCollectionRequestBuilder {
	bb := &SecurityIPSecurityProfilesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/ipSecurityProfiles"
	return bb
}

// SecurityIPSecurityProfilesCollectionRequestBuilder is request builder for IPSecurityProfile collection
type SecurityIPSecurityProfilesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for IPSecurityProfile collection
func (b *SecurityIPSecurityProfilesCollectionRequestBuilder) Request() *SecurityIPSecurityProfilesCollectionRequest {
	return &SecurityIPSecurityProfilesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for IPSecurityProfile item
func (b *SecurityIPSecurityProfilesCollectionRequestBuilder) ID(id string) *IPSecurityProfileRequestBuilder {
	bb := &IPSecurityProfileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SecurityIPSecurityProfilesCollectionRequest is request for IPSecurityProfile collection
type SecurityIPSecurityProfilesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for IPSecurityProfile collection
func (r *SecurityIPSecurityProfilesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]IPSecurityProfile, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []IPSecurityProfile
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []IPSecurityProfile
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for IPSecurityProfile collection
func (r *SecurityIPSecurityProfilesCollectionRequest) Get(ctx context.Context) ([]IPSecurityProfile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for IPSecurityProfile collection
func (r *SecurityIPSecurityProfilesCollectionRequest) Add(ctx context.Context, reqObj *IPSecurityProfile) (resObj *IPSecurityProfile, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ProviderTenantSettings returns request builder for ProviderTenantSetting collection
func (b *SecurityRequestBuilder) ProviderTenantSettings() *SecurityProviderTenantSettingsCollectionRequestBuilder {
	bb := &SecurityProviderTenantSettingsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/providerTenantSettings"
	return bb
}

// SecurityProviderTenantSettingsCollectionRequestBuilder is request builder for ProviderTenantSetting collection
type SecurityProviderTenantSettingsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ProviderTenantSetting collection
func (b *SecurityProviderTenantSettingsCollectionRequestBuilder) Request() *SecurityProviderTenantSettingsCollectionRequest {
	return &SecurityProviderTenantSettingsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ProviderTenantSetting item
func (b *SecurityProviderTenantSettingsCollectionRequestBuilder) ID(id string) *ProviderTenantSettingRequestBuilder {
	bb := &ProviderTenantSettingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SecurityProviderTenantSettingsCollectionRequest is request for ProviderTenantSetting collection
type SecurityProviderTenantSettingsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ProviderTenantSetting collection
func (r *SecurityProviderTenantSettingsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ProviderTenantSetting, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ProviderTenantSetting
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ProviderTenantSetting
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for ProviderTenantSetting collection
func (r *SecurityProviderTenantSettingsCollectionRequest) Get(ctx context.Context) ([]ProviderTenantSetting, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ProviderTenantSetting collection
func (r *SecurityProviderTenantSettingsCollectionRequest) Add(ctx context.Context, reqObj *ProviderTenantSetting) (resObj *ProviderTenantSetting, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// SecureScoreControlProfiles returns request builder for SecureScoreControlProfile collection
func (b *SecurityRequestBuilder) SecureScoreControlProfiles() *SecuritySecureScoreControlProfilesCollectionRequestBuilder {
	bb := &SecuritySecureScoreControlProfilesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/secureScoreControlProfiles"
	return bb
}

// SecuritySecureScoreControlProfilesCollectionRequestBuilder is request builder for SecureScoreControlProfile collection
type SecuritySecureScoreControlProfilesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SecureScoreControlProfile collection
func (b *SecuritySecureScoreControlProfilesCollectionRequestBuilder) Request() *SecuritySecureScoreControlProfilesCollectionRequest {
	return &SecuritySecureScoreControlProfilesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SecureScoreControlProfile item
func (b *SecuritySecureScoreControlProfilesCollectionRequestBuilder) ID(id string) *SecureScoreControlProfileRequestBuilder {
	bb := &SecureScoreControlProfileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SecuritySecureScoreControlProfilesCollectionRequest is request for SecureScoreControlProfile collection
type SecuritySecureScoreControlProfilesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SecureScoreControlProfile collection
func (r *SecuritySecureScoreControlProfilesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]SecureScoreControlProfile, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SecureScoreControlProfile
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []SecureScoreControlProfile
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for SecureScoreControlProfile collection
func (r *SecuritySecureScoreControlProfilesCollectionRequest) Get(ctx context.Context) ([]SecureScoreControlProfile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for SecureScoreControlProfile collection
func (r *SecuritySecureScoreControlProfilesCollectionRequest) Add(ctx context.Context, reqObj *SecureScoreControlProfile) (resObj *SecureScoreControlProfile, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// SecureScores returns request builder for SecureScore collection
func (b *SecurityRequestBuilder) SecureScores() *SecuritySecureScoresCollectionRequestBuilder {
	bb := &SecuritySecureScoresCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/secureScores"
	return bb
}

// SecuritySecureScoresCollectionRequestBuilder is request builder for SecureScore collection
type SecuritySecureScoresCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SecureScore collection
func (b *SecuritySecureScoresCollectionRequestBuilder) Request() *SecuritySecureScoresCollectionRequest {
	return &SecuritySecureScoresCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SecureScore item
func (b *SecuritySecureScoresCollectionRequestBuilder) ID(id string) *SecureScoreRequestBuilder {
	bb := &SecureScoreRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SecuritySecureScoresCollectionRequest is request for SecureScore collection
type SecuritySecureScoresCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SecureScore collection
func (r *SecuritySecureScoresCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]SecureScore, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SecureScore
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []SecureScore
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for SecureScore collection
func (r *SecuritySecureScoresCollectionRequest) Get(ctx context.Context) ([]SecureScore, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for SecureScore collection
func (r *SecuritySecureScoresCollectionRequest) Add(ctx context.Context, reqObj *SecureScore) (resObj *SecureScore, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// SecurityActions returns request builder for SecurityAction collection
func (b *SecurityRequestBuilder) SecurityActions() *SecuritySecurityActionsCollectionRequestBuilder {
	bb := &SecuritySecurityActionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/securityActions"
	return bb
}

// SecuritySecurityActionsCollectionRequestBuilder is request builder for SecurityAction collection
type SecuritySecurityActionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SecurityAction collection
func (b *SecuritySecurityActionsCollectionRequestBuilder) Request() *SecuritySecurityActionsCollectionRequest {
	return &SecuritySecurityActionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SecurityAction item
func (b *SecuritySecurityActionsCollectionRequestBuilder) ID(id string) *SecurityActionRequestBuilder {
	bb := &SecurityActionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SecuritySecurityActionsCollectionRequest is request for SecurityAction collection
type SecuritySecurityActionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SecurityAction collection
func (r *SecuritySecurityActionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]SecurityAction, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SecurityAction
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []SecurityAction
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for SecurityAction collection
func (r *SecuritySecurityActionsCollectionRequest) Get(ctx context.Context) ([]SecurityAction, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for SecurityAction collection
func (r *SecuritySecurityActionsCollectionRequest) Add(ctx context.Context, reqObj *SecurityAction) (resObj *SecurityAction, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// TiIndicators returns request builder for TiIndicator collection
func (b *SecurityRequestBuilder) TiIndicators() *SecurityTiIndicatorsCollectionRequestBuilder {
	bb := &SecurityTiIndicatorsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tiIndicators"
	return bb
}

// SecurityTiIndicatorsCollectionRequestBuilder is request builder for TiIndicator collection
type SecurityTiIndicatorsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TiIndicator collection
func (b *SecurityTiIndicatorsCollectionRequestBuilder) Request() *SecurityTiIndicatorsCollectionRequest {
	return &SecurityTiIndicatorsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TiIndicator item
func (b *SecurityTiIndicatorsCollectionRequestBuilder) ID(id string) *TiIndicatorRequestBuilder {
	bb := &TiIndicatorRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SecurityTiIndicatorsCollectionRequest is request for TiIndicator collection
type SecurityTiIndicatorsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TiIndicator collection
func (r *SecurityTiIndicatorsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]TiIndicator, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TiIndicator
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []TiIndicator
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for TiIndicator collection
func (r *SecurityTiIndicatorsCollectionRequest) Get(ctx context.Context) ([]TiIndicator, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for TiIndicator collection
func (r *SecurityTiIndicatorsCollectionRequest) Add(ctx context.Context, reqObj *TiIndicator) (resObj *TiIndicator, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// UserSecurityProfiles returns request builder for UserSecurityProfile collection
func (b *SecurityRequestBuilder) UserSecurityProfiles() *SecurityUserSecurityProfilesCollectionRequestBuilder {
	bb := &SecurityUserSecurityProfilesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userSecurityProfiles"
	return bb
}

// SecurityUserSecurityProfilesCollectionRequestBuilder is request builder for UserSecurityProfile collection
type SecurityUserSecurityProfilesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UserSecurityProfile collection
func (b *SecurityUserSecurityProfilesCollectionRequestBuilder) Request() *SecurityUserSecurityProfilesCollectionRequest {
	return &SecurityUserSecurityProfilesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UserSecurityProfile item
func (b *SecurityUserSecurityProfilesCollectionRequestBuilder) ID(id string) *UserSecurityProfileRequestBuilder {
	bb := &UserSecurityProfileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SecurityUserSecurityProfilesCollectionRequest is request for UserSecurityProfile collection
type SecurityUserSecurityProfilesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for UserSecurityProfile collection
func (r *SecurityUserSecurityProfilesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]UserSecurityProfile, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []UserSecurityProfile
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []UserSecurityProfile
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for UserSecurityProfile collection
func (r *SecurityUserSecurityProfilesCollectionRequest) Get(ctx context.Context) ([]UserSecurityProfile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for UserSecurityProfile collection
func (r *SecurityUserSecurityProfilesCollectionRequest) Add(ctx context.Context, reqObj *UserSecurityProfile) (resObj *UserSecurityProfile, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}