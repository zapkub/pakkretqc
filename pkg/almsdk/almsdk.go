package almsdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

func NewALMError(resp *http.Response) error {
	var almerr ALMError
	var err error
	almerr.Body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	almerr.Code = resp.StatusCode
	return &almerr
}

type ALMError struct {
	Body []byte
	Code int
}

func (a *ALMError) Error() string {
	return fmt.Sprintf("%d: %s", a.Code, string(a.Body))
}

type Client struct {
	client *http.Client
	config *ClientOptions
	token  string
}

type ClientOptions struct {
	Endpoint string
}

func New(opt *ClientOptions) *Client {
	cookieJar, _ := cookiejar.New(nil)
	return &Client{
		client: &http.Client{
			Transport:     http.DefaultTransport,
			CheckRedirect: http.DefaultClient.CheckRedirect,
			Jar:           cookieJar,
		},
		config: opt,
	}
}

func join(endpoint string, pathname ...string) *url.URL {
	u, err := url.Parse(endpoint)
	if err != nil {
		panic(fmt.Sprintf("invalid url for alm endpoint %s: %+v", endpoint, err))
	}
	u.Path = path.Join("qcbin/api", path.Join(pathname...))
	return u
}

type sessionCookieContext struct{}

func (c *Client) setTokenToRequest(ctx context.Context, req *http.Request) {
	if token, ok := ctx.Value(sessionCookieContext{}).(string); ok {
		req.Header.Set("Authorization", "Basic "+token)
		c.Authenticate(ctx, token)
	}
}

var InvalidCredential = fmt.Errorf("invalid credential")

func (c *Client) Authenticate(ctx context.Context, authtoken string) error {

	req, err := http.NewRequest("POST", join(c.config.Endpoint, "authentication/sign-in").String(), nil)
	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", authtoken))
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		respb, _ := ioutil.ReadAll(resp.Body)
		if resp.StatusCode == 401 {
			return InvalidCredential
		}
		return fmt.Errorf("login return %d: %s", resp.StatusCode, string(respb))
	}
	c.token = authtoken
	c.client.Jar.SetCookies(join(c.config.Endpoint), resp.Cookies())
	return nil
}

type Projects struct {
	Name string `json:"name"`
}

func (c *Client) Projects(ctx context.Context, domain string) ([]Projects, error) {
	var req, err = http.NewRequest("GET", join(c.config.Endpoint, "domains", domain, "projects").String(), nil)
	req.Header.Set("Accept", "application/json")
	c.setTokenToRequest(ctx, req)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		message, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected error: /domains/%s/projects return %d status\n%s", domain, resp.StatusCode, string(message))
	}

	type respbody struct {
		Results []Projects `json:"results"`
	}
	var body respbody
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}
	return body.Results, nil
}

type Domains struct {
	Name     string        `json:"name"`
	Projects []interface{} `json:"projects"`
}

func (c *Client) Domains(ctx context.Context) ([]*Domains, error) {
	var req, _ = http.NewRequest("GET", join(c.config.Endpoint, "domains").String(), nil)
	req.Header.Set("Accept", "application/json")
	c.setTokenToRequest(ctx, req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		return nil, fmt.Errorf("unexpected error: /domains return %d status", resp.StatusCode)
	}

	type respbody struct {
		Results []*Domains `json:"results"`
	}
	var body respbody
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}
	return body.Results, nil
}

type Defect struct {
	DevComments  string   `json:"dev-comments"`
	Description  string   `json:"description"`
	LastModified *ALMTime `json:"last-modified"`
	CreationTime *ALMTime `json:"creation-time"`
	Status       string   `json:"user-46"`
	Owner        string   `json:"owner"`
	Severity     string   `json:"severity"`
	Name         string   `json:"name"`
	ID           int      `json:"id"`
}

const almTimeLayout = "2006-01-02 15:04:05"
const almDateLayout = "2006-01-02"

type ALMTime struct {
	t time.Time
}

func (a *ALMTime) UnmarshalJSON(b []byte) error {
	var err error
	var layout = almTimeLayout
	var data = strings.Trim(string(b), "\"")
	if len(strings.Split(data, " ")) == 1 {
		layout = almDateLayout
	}

	if len(strings.Split(data, ":")) == 2 {
		data = data + ":00"
	}

	if a.t, err = time.Parse(layout, data); err != nil {
		return err
	}
	return nil
}

func (a *ALMTime) Time() time.Time {
	return a.t
}

func (a *ALMTime) MarshalJSON() ([]byte, error) {
	return a.t.MarshalJSON()
}

func (c *Client) Defect(ctx context.Context, domain, project, id string) (*Defect, error) {

	var req, _ = http.NewRequest("GET", join(c.config.Endpoint, "domains", domain, "projects", project, "defects", id).String(), nil)
	req.Header.Set("Accept", "application/json")
	c.setTokenToRequest(ctx, req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 300 {
		message, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected error: /domains/%s/projects/%s/defects return %d status\n%s", domain, project, resp.StatusCode, string(message))
	}
	var deflect Defect
	err = json.NewDecoder(resp.Body).Decode(&deflect)
	return &deflect, nil
}
func (c *Client) Defects(ctx context.Context, domain, project string, limit, offset int, orderFlag string) ([]*Defect, int, error) {

	var req, _ = http.NewRequest("GET", join(c.config.Endpoint, "domains", domain, "projects", project, "defects").String(), nil)
	req.Header.Set("Accept", "application/json")
	q := req.URL.Query()
	q.Add("order-by", orderFlag)
	q.Add("limit", strconv.Itoa(limit))
	req.URL.RawQuery = q.Encode()
	c.setTokenToRequest(ctx, req)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		message, _ := ioutil.ReadAll(resp.Body)
		return nil, 0, fmt.Errorf("unexpected error: /domains/%s/projects/%s/defects return %d status\n%s", domain, project, resp.StatusCode, string(message))
	}

	type respbody struct {
		Data []*Defect `json:"data"`
	}
	var body respbody
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, 0, err
	}
	return body.Data, 0, nil

}

type Attachment struct {
	Type         string      `json:"type"`
	LastModified string      `json:"last-modified"`
	VcCurVer     interface{} `json:"vc-cur-ver"`
	VcUserName   interface{} `json:"vc-user-name"`
	Name         string      `json:"name"`
	FileSize     int         `json:"file-size"`
	RefSubtype   int         `json:"ref-subtype"`
	Description  interface{} `json:"description"`
	ID           int         `json:"id"`
	RefType      string      `json:"ref-type"`
	Entity       struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
	} `json:"entity"`
}

func (c *Client) Attachments(ctx context.Context, domain, project string, query string, limit, offset int) ([]*Attachment, error) {
	var req, _ = http.NewRequest("GET", join(c.config.Endpoint, "domains", domain, "projects", project, "attachments").String(), nil)
	q := req.URL.Query()
	q.Add("query", fmt.Sprintf("\"%s\"", query))
	if limit <= 0 {
		limit = 20
	}
	if offset <= 0 {
		offset = 0
	}
	q.Add("limit", strconv.Itoa(limit))
	q.Add("offset", strconv.Itoa(offset))
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, NewALMError(resp)
	}
	type respbody struct {
		Data []*Attachment `json:"data"`
	}
	var body respbody
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}
	return body.Data, nil
}

func (c *Client) Attachment(ctx context.Context, domain, project string, id string, w io.Writer) error {
	var req, _ = http.NewRequest("GET", join(c.config.Endpoint, "domains", domain, "projects", project, "attachments", id).String(), nil)
	req.Header.Set("Accept", "application/json")
	c.setTokenToRequest(ctx, req)
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
