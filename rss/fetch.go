package rss

import (
	"curiosity/cache"
	_http "curiosity/http"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// THis keys are used for identified cache.
const (
	keyLastMod= "lastmod"
	keyEtag = "etag"

)

func getCacheHeaders(r *http.Request, key string) (headers map[string]string, err error) {
	item, err := cache.GetItem(r, key)
	if err != nil {
		return nil, err
	}
	h := strings.Split(string(item.Value), "|")
	headers = make(map[string]string)
	headers[keyEtag] = h[0]
	headers[keyLastMod] = h[1]

	return headers, nil
}

func setCacheHeaders(r *http.Request, h http.Header, url string){
	/*
		package main

		import (
			"fmt"
			"time"
		)

		func main() {

			// "Mon, 01/02/06, 03:04PM"
			layout := time.RFC1123
			str := "Tue, 06 Feb 2018 17:34:11 GMT"

			t, err := time.Parse(layout, str)

			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(t)

			fmt.Println(t.String())
			fmt.Println(t.Format(time.RFC1123))
		}

	*/
	cxt := appengine.NewContext(r)

	// parsing time
	t, err := time.Parse(time.RFC1123, h.Get(keyLastMod))
	if err!=nil {
		log.Errorf(cxt, err.Error())
	}

	item := &memcache.Item{
		Key: url,
		Value: []byte(fmt.Sprintf("%s|%s", h.Get(keyEtag), t.Format(time.RFC1123))),
	}

	if err:= memcache.Set(ctx, item); err !=nil{
		log.Errorf(ctx, err.Error())
	}

}

// Do a HTTP Get request with conditional (e-tag) values. Caching time in memory
func conditionalGet(r *http.Request, url string) (body []byte, err error) {
	// preparing a new request.
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// set the Headers into the request (if exists headers in cache).
	headers, _ := getCacheHeaders(r, url)
	if headers != nil {
		req.Header.Set(_http.HeaderETag, headers[keyEtag])
		req.Header.Set(_http.HeaderLastMod, headers[keyLastMod])
	}

	// do Request.
	resp, err := _http.Get(req)
	if err!=nil{
		return nil, err
	}

	// check if content is OK 200 and this was modified, so set new values in Cache.
	if resp.StatusCode == http.StatusOK{
		setCacheHeaders(r, resp.Header, url)
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Defer if body is not Nil...
	if resp.Body != nil {
		defer resp.Body.close()
	}

	return body, nil
}
