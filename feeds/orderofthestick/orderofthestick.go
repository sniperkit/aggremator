package orderofthestick

import (
	"github.com/danielheath/aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl:          "http://www.giantitp.com/comics/oots.rss",
		FeedSample:       Sample,
		MailCategory:     "Comics.OrderOfTheStick",
		Selector:         feeds.ParentSelector(feeds.CssSelector("body [src^=\"/comic\"]")), // attribute prefix selector
		SupportTheArtist: "http://www.giantitp.com/Shop.html",
	}
)