package application

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const (
	UrlBase    = "https://api.github.com/orgs/%s"
	SemMatcher = "^[0-9]+\\.[0-9]+\\.[0-9]+(-[a-z,A-Z,0-9]*)?$"
	ShaMatcher = "^sha-[0-9a-f]{40}$"
)

type Package struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	Id   int      `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

type Image struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Metadata struct {
		PackageType string `json:"package_type"`
		Container   struct {
			Tags []string `json:"tags"`
		} `json:"container"`
	} `json:"metadata"`
}

type Registry struct {
	org       string
	token     string
	client    *http.Client
	urlPrefix string
}

func NewRegistry(org, token string) *Registry {
	return &Registry{
		org:   org,
		token: token,
		client: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout: 5 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout:   5 * time.Second,
				ResponseHeaderTimeout: 5 * time.Second,
				IdleConnTimeout:       5 * time.Second,
				MaxIdleConns:          100,
				MaxConnsPerHost:       100,
				MaxIdleConnsPerHost:   100,
				TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
			},
		},
		urlPrefix: fmt.Sprintf(UrlBase, org),
	}
}

func (r *Registry) GetPackages(matcher string) ([]Package, error) {
	packages, err := r.getPackages(matcher)
	if err != nil {
		return nil, err
	}
	return packages, nil
}

func (r *Registry) GetTags(matcher, pattern string, depth int, semRange string) ([]Tag, error) {
	packages, err := r.getPackages(matcher)
	if err != nil {
		return nil, err
	}
	tags := make([]Tag, 0, 100)
	for _, pkg := range packages {
		arr, err := r.getTags(pkg.Name, pattern, depth, semRange)
		if err != nil {
			return nil, err
		}
		tag := Tag{
			Id:   pkg.Id,
			Name: pkg.Name,
			Tags: arr,
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (r *Registry) getPackages(matcher string) ([]Package, error) {
	packages := make([]Package, 0, 100)
	for i := 1; ; i++ {
		req, err := http.NewRequest(
			"GET",
			r.urlPrefix+"/packages",
			nil)
		if err != nil {
			return nil, err
		}
		q := req.URL.Query()
		q.Add("package_type", "container")
		q.Add("per_page", "100")
		q.Add("page", fmt.Sprintf("%d", i))
		req.URL.RawQuery = q.Encode()

		header, body, err := r.execHttpReq(req)
		if err != nil {
			return nil, err
		}
		tmpArr := make([]Package, 0, 100)
		json.Unmarshal(body, &tmpArr)
		for _, p := range tmpArr {
			match, _ := regexp.MatchString(matcher, p.Name)
			if match {
				packages = append(packages, p)
			}
		}
		linkVal := header.Get("Link")
		links := strings.Split(linkVal, ",")
		for _, link := range links {
			if strings.Contains(link, "rel=\"next\"") {
				goto loop
			}
		}
		break
	loop:
	}
	return packages, nil
}

func (r *Registry) getTags(name, pattern string, depth int, semRange string) ([]string, error) {
	var matcher string
	if pattern == "sem" {
		matcher = SemMatcher
	} else {
		matcher = ShaMatcher
	}
	tags := make([]string, 0, 100)
	for i := 1; ; i++ {
		req, err := http.NewRequest(
			"GET",
			r.urlPrefix+"/packages/container/"+url.PathEscape(name)+"/versions",
			nil)
		if err != nil {
			return nil, err
		}
		q := req.URL.Query()
		q.Add("per_page", "100")
		q.Add("page", fmt.Sprintf("%d", i))
		req.URL.RawQuery = q.Encode()

		header, body, err := r.execHttpReq(req)
		if err != nil {
			return nil, err
		}
		tmpArr := make([]Image, 0, 100)
		json.Unmarshal(body, &tmpArr)
		for _, i := range tmpArr {
			for _, tag := range i.Metadata.Container.Tags {
				match, _ := regexp.MatchString(matcher, tag)
				if match && r.checkSemRange(semRange, tag, tags) {
					tags = append(tags, tag)
				}
				if depth != 0 && len(tags) >= depth {
					return tags, nil
				}
			}
		}
		linkVal := header.Get("Link")
		links := strings.Split(linkVal, ",")
		for _, link := range links {
			if strings.Contains(link, "rel=\"next\"") {
				goto loop
			}
		}
		break
	loop:
	}
	return tags, nil
}

func (r Registry) checkSemRange(semRange, tag string, tags []string) bool {
	if semRange == "all" {
		return true
	}
	var rex *regexp.Regexp = nil
	if semRange == "major" {
		rex = regexp.MustCompile("^[0-9]+")
	} else if semRange == "minor" {
		rex = regexp.MustCompile("^[0-9]+\\.[0-9]+")
	} else {
		return false
	}

	target := rex.FindString(tag)
	for _, t := range tags {
		check := rex.FindString(t)
		if target == check {
			return false
		}
	}
	return true
}

func (r *Registry) execHttpReq(req *http.Request) (http.Header, []byte, error) {
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.token))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	header := resp.Header
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("HTTP status code: %d, error: %s", resp.StatusCode, body)
	}
	return header, body, nil
}
