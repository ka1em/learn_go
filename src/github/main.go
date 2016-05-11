package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"net/http"
	"strings"
)

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
}