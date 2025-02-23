package requestclient

import (
	"errors"

	"amrita_pyq/cmd/internal/webclient"
	"github.com/anaskhan96/soup"
)

type (
	Resource struct {
		Name string
		Path string
	}

	RequestClient struct {
		WebClient webclient.WebClient
	}
)

var errHTMLFetch = errors.New("failed to fetch the HTML content")

func (rc *RequestClient) GetCoursesReq(url string) ([]Resource, error) {
	res, err := rc.WebClient.FetchHTML(url)
	if err != nil {
		return nil, errHTMLFetch
	}

	doc := soup.HTMLParse(res)
	div := doc.Find("div", "id", "aspect_artifactbrowser_CommunityViewer_div_community-view")
	subs := div.FindAll("div", "class", "artifact-title")

	var subjects []Resource
	for _, item := range subs {
		sub := item.Find("span")
		a := item.Find("a")
		path := a.Attrs()["href"]
		subject := Resource{Name: sub.Text(), Path: path}
		subjects = append(subjects, subject)
	}

	return subjects, nil
}

func (rc *RequestClient) SemChooseReq(url string) ([]Resource, error) {
	res, err := rc.WebClient.FetchHTML(url)
	if err != nil {
		return nil, errHTMLFetch
	}

	doc := soup.HTMLParse(res)
	div := doc.Find("div", "id", "aspect_artifactbrowser_CommunityViewer_div_community-view")

	if div.Error != nil {
		return nil, errors.New("no assessments found on the page")
	}

	ul := div.FindAll("ul")
	var li []soup.Root
	if len(ul) > 1 {
		li = ul[1].FindAll("li")
	} else {
		li = ul[0].FindAll("li")
	}

	var assessments []Resource
	for _, link := range li {
		a := link.Find("a")
		span := a.Find("span")
		path := a.Attrs()["href"]
		assessment := Resource{Name: span.Text(), Path: path}
		assessments = append(assessments, assessment)
	}

	return assessments, nil
}

func (rc *RequestClient) SemTableReq(url string) ([]Resource, error) {
	res, err := rc.WebClient.FetchHTML(url)
	if err != nil {
		return nil, errHTMLFetch
	}

	doc := soup.HTMLParse(res)
	div := doc.Find("div", "id", "aspect_artifactbrowser_CommunityViewer_div_community-view")

	if div.Error != nil {
		return nil, errors.New("no semesters found on the page")
	}

	ul := div.Find("ul")
	li := ul.FindAll("li")
	if len(li) == 0 {
		return nil, errors.New("no semesters found on the page")
	}

	var semesters []Resource
	for _, link := range li {
		a := link.Find("a")
		span := a.Find("span")
		path := a.Attrs()["href"]
		semester := Resource{Name: span.Text(), Path: path}
		semesters = append(semesters, semester)
	}

	return semesters, nil
}

func (rc *RequestClient) SubComReq(url string) ([]Resource, error) {
	res, err := rc.WebClient.FetchHTML(url)
	if err != nil {
		return nil, errHTMLFetch
	}

	doc := soup.HTMLParse(res)
	div := doc.Find("div", "xmlns", "http://di.tamu.edu/DRI/1.0/")
	subComSections := div.FindAll("ul")
	communities := map[soup.Root]bool{}

	for _, subComSection := range subComSections {
		subComs := subComSection.FindAll("li")
		for _, subCom := range subComs {
			_, ok := communities[subCom]
			if !ok {
				communities[subCom] = true
			}
		}
	}

	var subCommunities []Resource

	for com := range communities {
		a := com.Find("a")
		path := a.Attrs()["href"]
		subCommunity := Resource{a.Text(), path}
		subCommunities = append(subCommunities, subCommunity)
	}
	return subCommunities, nil

}

func (rc *RequestClient) YearReq(url string) ([]Resource, error) {
	page, err := rc.WebClient.FetchHTML(url)
	if err != nil {
		return nil, errHTMLFetch
	}

	doc := soup.HTMLParse(page)
	div := doc.Find("div", "class", "file-list")
	subdiv := div.FindAll("div", "class", "file-wrapper")

	var files []Resource
	for _, item := range subdiv {
		title := item.FindAll("div")
		indiv := title[1].Find("div")
		span := indiv.FindAll("span")
		fileName := span[1].Attrs()["title"]
		path := title[0].Find("a").Attrs()["href"]
		file := Resource{Name: fileName, Path: path}
		files = append(files, file)
	}

	return files, nil
}
