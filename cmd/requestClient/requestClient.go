package requestClient

import (
	"errors"

	"amrita_pyq/cmd/configs"
	"amrita_pyq/cmd/helpers"
	"amrita_pyq/cmd/model"

	"github.com/anaskhan96/soup"
)

var errHTMLFetch error = errors.New("failed to fetch the HTML content")

func GetCoursesReq(url string) ([]model.Resource, error) {

	res, err := helpers.FetchHTML(url)

	if err != nil {
		return nil, errHTMLFetch
	}

	doc := soup.HTMLParse(res)
	div := doc.Find("div", "id", "aspect_artifactbrowser_CommunityViewer_div_community-view")

	subs := div.FindAll("div", "class", "artifact-title")

	var subjects []model.Resource

	for _, item := range subs {
		sub := item.Find("span")
		a := item.Find("a")
		path := a.Attrs()["href"]
		subject := model.Resource{Name: sub.Text(), Path: path}
		subjects = append(subjects, subject)
	}

	return subjects, nil
}

func SemChooseReq(url string) ([]model.Resource, error) {

	res, err := helpers.FetchHTML(url)

	if err != nil {
		return nil, errHTMLFetch
	}

	doc := soup.HTMLParse(res)
	div := doc.Find("div", "id", "aspect_artifactbrowser_CommunityViewer_div_community-view")

	if div.Error != nil {
		return nil, errors.New("no assesments found on the page")
	}

	ul := div.FindAll("ul")
	var li []soup.Root

	if len(ul) > 1 {
		li = ul[1].FindAll("li")
	} else {
		li = ul[0].FindAll("li")
	}

	var assesments []model.Resource

	for _, link := range li {
		a := link.Find("a")
		span := a.Find("span")
		path := link.Find("a").Attrs()["href"]
		assesment := model.Resource{Name: span.Text(), Path: path}
		assesments = append(assesments, assesment)
	}

	return assesments, nil
}

func SemTableReq(url string) ([]model.Resource, error) {

	res, err := helpers.FetchHTML(url)

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

	var semesters []model.Resource

	for _, link := range li {
		a := link.Find("a")
		span := a.Find("span")
		path := link.Find("a").Attrs()["href"]
		semester := model.Resource{Name: span.Text(), Path: path}
		semesters = append(semesters, semester)
	}

	return semesters, nil

}

func YearReq(url string) ([]model.Resource, error) {

	res, err := helpers.FetchHTML(url)

	if err != nil {
		return nil, errHTMLFetch
	}

	doc := soup.HTMLParse(res)
	div := doc.Find("div", "xmlns", "http://di.tamu.edu/DRI/1.0/")

	ul := div.Find("ul")
	li := ul.Find("li")
	hyper := li.Find("a").Attrs()["href"]

	url = configs.BASE_URL + hyper
	page, err := helpers.FetchHTML(url)

	if err != nil {
		return nil, errHTMLFetch
	}

	doc = soup.HTMLParse(page)
	div = doc.Find("div", "class", "file-list")

	subdiv := div.FindAll("div", "class", "file-wrapper")

	var files []model.Resource

	for _, item := range subdiv {
		title := item.FindAll("div")
		indiv := title[1].Find("div")
		span := indiv.FindAll("span")
		fileName := span[1].Attrs()["title"]
		path := title[0].Find("a").Attrs()["href"]
		file := model.Resource{Name: fileName, Path: path}
		files = append(files, file)
	}

	return files, nil

}
