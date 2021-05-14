package gotion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	client  *http.Client
	token   string
	version string
	baseUrl string
}

func New(token, version string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	return &Client{
		client,
		token,
		version,
		"https://api.notion.com/v1",
	}
}

func (c *Client) Do(r *http.Request) (*http.Response, error) {
	r.Header.Add("Authorization", "Bearer "+c.token)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Notion-Version", c.version)

	return c.client.Do(r)
}

func (c *Client) DoAndRead(req *http.Request, out interface{}) error {
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	buff, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	msg := new(Response)
	if err = json.Unmarshal(buff, msg); err != nil {
		return err
	}

	if msg.Object == Error {
		return fmt.Errorf("%d: %s: %s", res.StatusCode, msg.Code, msg.Message)
	}

	return json.Unmarshal(buff, out)
}

func (c *Client) GetDatabase(databaseId UUID) (*DatabaseObject, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"/databases/"+databaseId, nil)
	if err != nil {
		return nil, err
	}

	databaseObject := new(DatabaseObject)
	if err = c.DoAndRead(req, databaseObject); err != nil {
		return nil, err
	}

	return databaseObject, nil
}

func (c *Client) ListDatabases() ([]*DatabaseObject, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"/databases?page_size=100", nil)
	if err != nil {
		return nil, err
	}

	databaseList := new(DatabaseList)
	if err = c.DoAndRead(req, databaseList); err != nil {
		return nil, err
	}

	databaseObjects := databaseList.Results
	hasMore := databaseList.HasMore
	for hasMore {

		req, err := http.NewRequest("GET", c.baseUrl+"/databases?page_size=100", nil)
		if err != nil {
			return nil, err
		}

		databaseList := new(DatabaseList)
		if err = c.DoAndRead(req, databaseList); err != nil {
			return nil, err
		}

		databaseObjects = append(databaseObjects, databaseList.Results...)
		hasMore = databaseList.HasMore
	}

	return databaseObjects, nil
}

func (c *Client) GetPage(pageId UUID) (*PageObject, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"/pages/"+pageId, nil)
	if err != nil {
		return nil, err
	}

	pageObject := new(PageObject)
	if err = c.DoAndRead(req, pageObject); err != nil {
		return nil, err
	}

	return pageObject, nil
}

func (c *Client) PostPage(p *PageObject) (*PageObject, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.baseUrl+"/pages", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	pageObject := new(PageObject)
	if err = c.DoAndRead(req, pageObject); err != nil {
		return nil, err
	}

	return pageObject, nil
}

func (c *Client) UpdatePage(pageId UUID, properties map[string]*PropertyValue) (*PageObject, error) {
	body, err := json.Marshal(&PageObject{
		Properties: properties,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.baseUrl+"/pages/"+pageId, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	pageObject := new(PageObject)
	if err = c.DoAndRead(req, pageObject); err != nil {
		return nil, err
	}

	return pageObject, nil
}
