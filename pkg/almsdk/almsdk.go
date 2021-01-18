package almsdk

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	client *http.Client
	config *ClientOptions
}

type ClientOptions struct {
	Endpoint string
}

func New(opt *ClientOptions) *Client {
	return &Client{
		client: &http.Client{
			Transport: http.DefaultTransport,
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

func AppendSessionCookieToContext(ctx context.Context, r *http.Request) context.Context {
	return context.WithValue(ctx, sessionCookieContext{}, r.Cookies())
}
func setCookieToRequest(ctx context.Context, req *http.Request) {
	if cookies, ok := ctx.Value(sessionCookieContext{}).([]*http.Cookie); ok {
		for _, cook := range cookies {
			req.AddCookie(cook)
		}
	}
}

func (c *Client) Authenticate(ctx context.Context, user, password string) ([]*http.Cookie, error) {

	req, err := http.NewRequest("POST", join(c.config.Endpoint, "authentication/sign-in").String(), nil)
	req.Header.Add("Content-Type", "application/xml")
	token := base64.URLEncoding.EncodeToString([]byte(user + ":" + password))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		respb, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("login return %d: %s", resp.StatusCode, string(respb))
	}
	return resp.Cookies(), nil
}

type Projects struct {
	Name string `json:"name"`
}

func (c *Client) Projects(ctx context.Context, domain string) ([]Projects, error) {

	var req, err = http.NewRequest("GET", join(c.config.Endpoint, "domains", domain, "projects").String(), nil)
	req.Header.Set("Accept", "application/json")
	setCookieToRequest(ctx, req)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

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
	setCookieToRequest(ctx, req)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 300 {
		return nil, fmt.Errorf("unexpected error: /domains return %d status", resp.StatusCode)
	}

	type respbody struct {
		Results []*Domains `json:"results"`
	}
	var body respbody
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}
	return body.Results, nil
}

type Deflect struct {
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

func (c *Client) Deflects(ctx context.Context, domain, project string, limit, offset int, orderFlag string) ([]*Deflect, int, error) {

	var req, _ = http.NewRequest("GET", join(c.config.Endpoint, "domains", domain, "projects", project, "defects").String(), nil)
	req.Header.Set("Accept", "application/json")
	q := req.URL.Query()
	q.Add("order-by", orderFlag)
	q.Add("limit", strconv.Itoa(limit))
	req.URL.RawQuery = q.Encode()
	setCookieToRequest(ctx, req)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	if resp.StatusCode > 300 {
		message, _ := ioutil.ReadAll(resp.Body)
		return nil, 0, fmt.Errorf("unexpected error: /domains/%s/projects/%s/defects return %d status\n%s", domain, project, resp.StatusCode, string(message))
	}

	type respbody struct {
		Data []*Deflect `json:"data"`
	}
	var body respbody
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, 0, err
	}
	return body.Data, 0, nil

}
