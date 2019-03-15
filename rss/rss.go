/*
Hearbeat is like a Cron. Is a task that executes every X time. But also can to have
a strategy and learn about what's the best moment for execute again a consult...
*/
package rss

import (
	"fmt"
	"time"
)

func StartBeat() {
	ticker := time.NewTicker(2 * time.Minute)
	go func() {
		for t := range ticker.C {
			readFeeds() // Reading Feeds... Here we go!! :-)
			fmt.Println("Hearbeat-->", t)
		}
	}()

}
