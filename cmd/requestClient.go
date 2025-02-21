package cmd

import (
	"errors"
	"github.com/anaskhan96/soup"
)

var errHTMLFetch error = errors.New("failed to fetch the HTML content")

func getCoursesReq(url string) ([]resource, error) {

	res, err := fetchHTML(url)

	if err != nil {
		return nil, errHTMLFetch
	}

	doc := soup.HTMLParse(res)
	div := doc.Find("div", "id", "aspect_artifactbrowser_CommunityViewer_div_community-view")

	subs := div.FindAll("div", "class", "artifact-title")

	var subjects []resource

	for _, item := range subs {
		sub := item.Find("span")
		a := item.Find("a")
		path := a.Attrs()["href"]
		subject := resource{sub.Text(), path}
		subjects = append(subjects, subject)
	}

	return subjects, nil
}

func semChooseReq(url string) ([]resource, error) {

	res, err := fetchHTML(url)

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

	var assesments []resource

	for _, link := range li {
		a := link.Find("a")
		span := a.Find("span")
		path := link.Find("a").Attrs()["href"]
		assesment := resource{span.Text(), path}
		assesments = append(assesments, assesment)
	}

	return assesments, nil
}

func semTableReq(url string) ([]resource, error) {

	res, err := fetchHTML(url)

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

	var semesters []resource

	for _, link := range li {
		a := link.Find("a")
		span := a.Find("span")
		path := link.Find("a").Attrs()["href"]
		semester := resource{span.Text(), path}
		semesters = append(semesters, semester)
	}

	return semesters, nil

}

func yearReq(url string) ([]resource, error) {
	page, err := fetchHTML(url)

	if err != nil {
		return nil, errHTMLFetch
	}

	doc := soup.HTMLParse(page)
	div := doc.Find("div", "class", "file-list")

	subdiv := div.FindAll("div", "class", "file-wrapper")

	var files []resource

	for _, item := range subdiv {
		title := item.FindAll("div")
		indiv := title[1].Find("div")
		span := indiv.FindAll("span")
		fileName := span[1].Attrs()["title"]
		path := title[0].Find("a").Attrs()["href"]
		file := resource{fileName, path}
		files = append(files, file)
	}

	return files, nil

}

func subComReq(url string) ([]resource, error) {
	res, err := fetchHTML(url)

	if err != nil {
		return nil, errHTMLFetch
	}

	doc := soup.HTMLParse(res)
	div := doc.Find("div", "xmlns", "http://di.tamu.edu/DRI/1.0/")

	uls := div.FindAll("ul")
	communities := map[soup.Root]bool{} // set

	for _, ul := range uls {
		lis := ul.FindAll("li")
		for _, li := range lis {
			_, ok := communities[li]
			if !ok {
				communities[li] = true
			}
		}
	}

	var subCommunities []resource

	for com := range communities {
		a := com.Find("a")
		path := a.Attrs()["href"]
		subCommunity := resource{a.Text(), path}
		subCommunities = append(subCommunities, subCommunity)
	}
	return subCommunities, nil
}
