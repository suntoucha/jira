package jira

import (
	"io/ioutil"
	"net/http"
	"strconv"
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
	if err := IsValidKey(key); err != nil {
		return Project{}, err
	}

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
	if err := IsValidKey(key); err != nil {
		return Issue{}, err
	}

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

func (cli *Client) IssueByProject(key string, startAt int, maxResults int) (IssueResult, error) {
	if err := IsValidKey(key); err != nil {
		return IssueResult{}, err
	}

	raw, err := cli.get("/rest/api/2/search?jql=project=" + key + "+order+by+key&startAt=" + strconv.Itoa(startAt) + "&maxResults=" + strconv.Itoa(maxResults))
	if err != nil {
		return IssueResult{}, err
	}

	res, err := IssueResultFromJson(raw)
	if err != nil {
		return IssueResult{}, err
	}

	return res, nil
}

func (cli *Client) ExportIssueByProject(key string, maxResults int, e Exporter) error {
	startAt := 0
	res, err := cli.IssueByProject(key, startAt, maxResults)

	for err == nil && len(res.Issues) > 0 {
		for i, raw := range res.Issues {
			e.Export(raw, startAt+i, res.Total)
		}

		startAt += maxResults
		res, err = cli.IssueByProject(key, startAt, maxResults)
	}

	return err
}
