package processer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_page_version(t *testing.T) {
	process := &processer{}
	url := "https://www.facebook.com"
	version, err := process.getHTMLVersion(url)

	assert.NoError(t, err)
	assert.Equal(t, "HTML5", version)
}

func Test_page_version_for_empty_url(t *testing.T) {
	process := &processer{}
	url := ""
	version, err := process.getHTMLVersion(url)

	assert.Empty(t, version)
	assert.ErrorContains(t, err, "!Doctype node is not found.")
}

func Test_page_title_not_found(t *testing.T) {
	process := &processer{}
	url := "test.com"
	title, err := process.getPageTitle(url)
	fmt.Println(title)
	assert.Empty(t, title)
	assert.ErrorContains(t, err, "Unable to Find the Title")
}

func Test_page_title(t *testing.T) {
	process := &processer{}
	url := "https://www.google.com"
	title, _ := process.getPageTitle(url)
	fmt.Println(title)
	assert.Equal(t, title, "Google")
}
