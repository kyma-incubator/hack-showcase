package registration

import (
	"encoding/json"
)

//ServiceDetails - main structure of service register json
type ServiceDetails struct {
	Provider         string            `json:"provider"`
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	ShortDescription string            `json:"shortDescription,omitempty"`
	Identifier       string            `json:"identifier,omitempty"`
	Labels           map[string]string `json:"labels,omitempty"`
	API              *API              `json:"api,omitempty"`
	Events           *Events           `json:"events,omitempty"`
	Documentation    *Documentation    `json:"documentation,omitempty"`
}

type PostServiceResponse struct {
	ID string `json:"id"`
}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type API struct {
	TargetURL                      string                          `json:"targetUrl"`
	Credentials                    *Credentials                    `json:"credentials,omitempty"`
	Spec                           json.RawMessage                 `json:"spec,omitempty"`
	SpecificationURL               string                          `json:"specificationUrl,omitempty"`
	ApiType                        string                          `json:"apiType"`
	SpecificationRequestParameters *SpecificationRequestParameters `json:"specificationRequestParameters,omitempty"`
}

type Credentials struct {
	Oauth          *Oauth          `json:"oauth,omitempty"`
	Basic          *Basic          `json:"basic,omitempty"`
	CertificateGen *CertificateGen `json:"certificateGen,omitempty"`
}

type CSRFInfo struct {
	TokenEndpointURL string `json:"tokenEndpointURL"`
}

type Oauth struct {
	URL          string    `json:"url"`
	ClientID     string    `json:"clientId"`
	ClientSecret string    `json:"clientSecret"`
	CSRFInfo     *CSRFInfo `json:"csrfInfo,omitempty"`
}

type Basic struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	CSRFInfo *CSRFInfo `json:"csrfInfo,omitempty"`
}

type CertificateGen struct {
	CommonName  string    `json:"commonName"`
	Certificate string    `json:"certificate"`
	CSRFInfo    *CSRFInfo `json:"csrfInfo,omitempty"`
}

type Events struct {
	Spec json.RawMessage `json:"spec,omitempty"`
}

type Documentation struct {
	DisplayName string       `json:"displayName"`
	Description string       `json:"description"`
	Type        string       `json:"type"`
	Tags        []string     `json:"tags,omitempty"`
	Docs        []DocsObject `json:"docs,omitempty"`
}

type DocsObject struct {
	Title  string `json:"title"`
	Type   string `json:"type"`
	Source string `json:"source"`
}

type SpecificationRequestParameters struct {
	Headers         *Headers         `json:"headers,omitempty"`
	QueryParameters *QueryParameters `json:"queryParameters,omitempty"`
}

type Headers struct {
	CustomHeader []string `json:"custom-header,omitempty"`
}

type QueryParameters struct {
	Param []string `json:"param,omitempty"`
}
