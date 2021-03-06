package orderofthestick

import (
	"aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl:          "http://www.giantitp.com/comics/oots.rss",
		FeedSample:       Sample,
		MailCategory:     "Comics",
		Selector:         feeds.CssSelector(":haschild([src^=\"/comic\"])"), // attribute prefix selector
		SupportTheArtist: "http://www.giantitp.com/Shop.html",
	}
)
