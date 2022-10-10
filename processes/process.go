package processer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	dto "github.com/KanishkaRR/web-scraper/dtos"
	"github.com/gocolly/colly"
	"golang.org/x/net/html"
)

type Process interface {
	ProcessWebPage(url string) (*dto.PageInfo, error)
}

type processer struct {
}

func NewProcessor() Process {
	return &processer{}
}

func (p *processer) ProcessWebPage(url string) (*dto.PageInfo, error) {
	pageInfo := &dto.PageInfo{}
	version, err := p.getHTMLVersion(url)

	if err != nil {
		return nil, fmt.Errorf("page cannot be found")
	}

	pageInfo.Url = url
	pageInfo.PageVersion = version

	title, err := p.getPageTitle(url)
	pageInfo.Title = title

	heading, err := p.getHeadings(url)
	pageInfo.Headings = heading

	linkdetails := p.getLinkDetails(url)
	pageInfo.LinkDetails = linkdetails

	pageInfo.HasLoginForm = p.isLoginForm(url)

	return pageInfo, nil
}

func (p *processer) getHTMLVersion(url string) (string, error) {
	htmlText, err := getFisrtLineFromWebPage(url)
	if err != nil {
		fmt.Errorf("Cannot retrive information from webpage. ")
	}
	reader := strings.NewReader(htmlText)
	tokenizer := html.NewTokenizer(reader)
	docType := ""

loop:
	for {
		token := tokenizer.Next()
		switch token {
		case html.ErrorToken:
			err := tokenizer.Err()
			if errors.Is(err, io.EOF) {
				return "", fmt.Errorf("!Doctype node is not found.")
			}

			return "", fmt.Errorf("unable to process the document, %v", err)

		case html.DoctypeToken:
			docType = string(tokenizer.Text())
			break loop
		}
	}

	if docType == "html" {
		return "HTML5", nil
	}

	return "", fmt.Errorf("Unable to parse the Doctype node %q", docType)
}

func getFisrtLineFromWebPage(url string) (string, error) {
	fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	text := ""
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		text = scanner.Text()
		break
	}
	return text, nil
}

func (p *processer) getPageTitle(url string) (string, error) {
	collect := colly.NewCollector()
	title := ""

	collect.OnHTML("title", func(e *colly.HTMLElement) {
		title = e.Text
	})

	collect.Visit(url)

	if title == "" {
		return "", fmt.Errorf("Unable to Find the Title")
	}
	return title, nil
}

func (p *processer) getHeadings(url string) (map[string]int, error) {

	collect := colly.NewCollector()
	headings := map[string]int{
		"h1": 0,
		"h2": 0,
		"h3": 0,
		"h4": 0,
		"h5": 0,
		"h6": 0,
	}
	collect.OnHTML("h1", func(e *colly.HTMLElement) {
		headings["h1"]++
	})

	collect.OnHTML("h2", func(e *colly.HTMLElement) {
		headings["h2"]++
	})

	collect.OnHTML("h3", func(e *colly.HTMLElement) {
		headings["h3"]++
	})

	collect.OnHTML("h4", func(e *colly.HTMLElement) {
		headings["h4"]++
	})

	collect.OnHTML("h5", func(e *colly.HTMLElement) {
		headings["h5"]++
	})

	collect.Visit(url)

	for key, element := range headings {
		fmt.Println("Header:", key, "=>", "No of Elements:", element)
	}

	return headings, nil
}

func (p *processer) getLinkDetails(url string) dto.LinkDetails {
	linkDetails := &dto.LinkDetails{}
	noOfLinks := 0
	active := 0
	internal := 0
	//collect := colly.NewCollector()
	// collect.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	// Extract the link from the anchor HTML element
	// 	link := e.Attr("href")
	// 	noOfLinks++
	// 	println(noOfLinks)

	// 	if string(link)[0:1] == "#" {
	// 		internal++
	// 	}
	// 	err := collect.Visit(e.Request.AbsoluteURL(link))
	// 	if err != nil {
	// 		active++
	// 	}
	// })
	// collect.Visit(url)

	linkDetails.Active = active
	linkDetails.Internal = internal
	linkDetails.NoofLinks = noOfLinks
	return *linkDetails
}

func (p *processer) isLoginForm(url string) bool {
	passwordExist := false
	textFieldExist := false
	buttonExist := false
	collect := colly.NewCollector()
	collect.OnHTML("input", func(e *colly.HTMLElement) {
		p := e.Attr("type")
		if p == "password" {
			passwordExist = true
		}

		if p == "text" {
			textFieldExist = true
		}

	})

	collect.OnHTML("button", func(e *colly.HTMLElement) {
		name := e.Attr("name")

		if name == "login" || name == "signin" {
			buttonExist = true
		}

	})
	collect.Visit(url)

	if passwordExist && textFieldExist && buttonExist {
		return true
	}
	return false
}
