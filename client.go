package jira

import (
	"io/ioutil"
	"net/http"
)

type Client struct {
	Host  string
	User  string
	Token string
}

func (cli *Client) setAuth(req *http.Request) {
	req.SetBasicAuth(cli.User, cli.Token)
}

func (cli *Client) get(resource string) ([]byte, error) {
	client := &http.Client{}

	url := "https://" + cli.Host + resource
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	cli.setAuth(req)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyText, nil
}

func (cli *Client) Project(key string) (Project, error) {
	raw, err := cli.get("/rest/api/2/project/" + key + "/?expand=description,lead,issueTypes,url,projectKeys,permissions,insight")
	if err != nil {
		return Project{}, err
	}

	prj, err := ProjectFromJson(raw)
	if err != nil {
		return Project{}, err
	}

	return prj, nil
}

func (cli *Client) Issue(key string) (Issue, error) {
	raw, err := cli.get("/rest/api/2/issue/" + key)
	if err != nil {
		return Issue{}, err
	}

	iss, err := IssueFromJson(raw)
	if err != nil {
		return Issue{}, err
	}

	return iss, nil
}

func (cli *Client) IssueCursor(cur *IssueCursor) (IssueList, error) {
	raw, err := cli.get(cur.Resourse())
	if err != nil {
		return nil, err
	}

	list, err := IssueListFromJson(raw)
	if err != nil {
		return nil, err
	}

	cur.Next()
	return list, nil
}
