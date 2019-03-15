package rss

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"net/http"
)

// Get content of a Item of a specific key from memcached.
func getItem(r *http.Request, key string) (item *memcache.Item, err error) {
	// New context...
	cxt := appengine.NewContext(r)
	// Getting key from cache...
	item, err = memcache.Get(cxt, key)
	if err != nil {
		return nil, err
	}
	return item, nil

}

// Set item in memcached.
// v = value to stored, k = Unique key for identified item in memcache.
func setItem(r *http.Request, k string, v string){
	cxt := appengine.NewContext(r)
	item := &memcache.Item{
		Key: url,
		Value: []byte(v),
	}
	if err := memcache.Set(cxt, item); err !=nil {
		log.Errorf(ctx, err.Error())
	}
}