package rss

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"net/http"
)

// Do a GET HTTP Request. Using Context, and URL Fetch.
func GET(r *http.Request) (resp *http.Response, err error){
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

}
