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

func (cli *Client) Project(key string) (JiraProject, error) {
	if err := IsValidKey(key); err != nil {
		return JiraProject{}, err
	}

	raw, err := cli.get("/rest/api/2/project/" + key + "/?expand=description,lead,issueTypes,url,projectKeys,permissions,insight")
	if err != nil {
		return JiraProject{}, err
	}

	prj, err := JiraProjectFromJson(raw)
	if err != nil {
		return JiraProject{}, err
	}

	return prj, nil
}

func (cli *Client) Issue(key string) (JiraIssue, error) {
	if err := IsValidKey(key); err != nil {
		return JiraIssue{}, err
	}

	raw, err := cli.get("/rest/api/2/issue/" + key)
	if err != nil {
		return JiraIssue{}, err
	}

	iss, err := JiraIssueFromJson(raw)
	if err != nil {
		return JiraIssue{}, err
	}

	return iss, nil
}

func (cli *Client) IssueByProject(key string, startAt int, maxResults int) (JiraIssueResult, error) {
	if err := IsValidKey(key); err != nil {
		return JiraIssueResult{}, err
	}

	raw, err := cli.get("/rest/api/2/search?jql=project=" + key + "+order+by+key&startAt=" + strconv.Itoa(startAt) + "&maxResults=" + strconv.Itoa(maxResults))
	if err != nil {
		return JiraIssueResult{}, err
	}

	res, err := JiraIssueResultFromJson(raw)
	if err != nil {
		return JiraIssueResult{}, err
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
