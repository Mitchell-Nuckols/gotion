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
	nextCursor := databaseList.NextCursor
	for hasMore {

		req, err := http.NewRequest("GET", c.baseUrl+"/databases?page_size=100&next_cursor="+nextCursor, nil)
		if err != nil {
			return nil, err
		}

		databaseList := new(DatabaseList)
		if err = c.DoAndRead(req, databaseList); err != nil {
			return nil, err
		}

		databaseObjects = append(databaseObjects, databaseList.Results...)
		hasMore = databaseList.HasMore
		nextCursor = databaseList.NextCursor
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
	body, err := json.Marshal(struct {
		Properties map[string]*PropertyValue `json:"properties"`
	}{
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

func (c *Client) GetBlockChildren(blockId UUID) ([]*BlockObject, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"/blocks/"+blockId+"/children?page_size=100", nil)
	if err != nil {
		return nil, err
	}

	blockList := new(BlockList)
	if err = c.DoAndRead(req, blockList); err != nil {
		return nil, err
	}

	blockObjects := blockList.Results
	hasMore := blockList.HasMore
	nextCursor := blockList.NextCursor
	for hasMore {

		req, err := http.NewRequest("GET", c.baseUrl+"/blocks/"+blockId+"/children?page_size=100&next_cursor="+nextCursor, nil)
		if err != nil {
			return nil, err
		}

		blockList := new(BlockList)
		if err = c.DoAndRead(req, blockList); err != nil {
			return nil, err
		}

		blockObjects = append(blockObjects, blockList.Results...)
		hasMore = blockList.HasMore
		nextCursor = blockList.NextCursor
	}

	return blockObjects, nil
}

func (c *Client) AppendBlockChildren(blockId UUID, children []*BlockObject) (*BlockObject, error) {
	body, err := json.Marshal(struct {
		Children []*BlockObject `json:"children"`
	}{
		Children: children,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.baseUrl+"/blocks/"+blockId+"/children", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	blockObject := new(BlockObject)
	if err = c.DoAndRead(req, blockObject); err != nil {
		return nil, err
	}

	return blockObject, nil
}

func (c *Client) GetUser(userId UUID) (*UserObject, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"/users/"+userId, nil)
	if err != nil {
		return nil, err
	}

	userObject := new(UserObject)
	if err = c.DoAndRead(req, userObject); err != nil {
		return nil, err
	}

	return userObject, nil
}

func (c *Client) ListUsers() ([]*UserObject, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"/users?page_size=100", nil)
	if err != nil {
		return nil, err
	}

	userList := new(UserList)
	if err = c.DoAndRead(req, userList); err != nil {
		return nil, err
	}

	userObjects := userList.Results
	hasMore := userList.HasMore
	nextCursor := userList.NextCursor
	for hasMore {

		req, err := http.NewRequest("GET", c.baseUrl+"/users?page_size=100&next_cursor="+nextCursor, nil)
		if err != nil {
			return nil, err
		}

		userList := new(UserList)
		if err = c.DoAndRead(req, userList); err != nil {
			return nil, err
		}

		userObjects = append(userObjects, userList.Results...)
		hasMore = userList.HasMore
		nextCursor = userList.NextCursor
	}

	return userObjects, nil
}
