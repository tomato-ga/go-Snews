package cmd

import (
	"fmt"
	"os"

	"github.com/mmcdole/gofeed"
)

var (
	counts int = 0
)

func RSS() {
	feed, err := gofeed.NewParser().ParseURL("https://rss.itmedia.co.jp/rss/2.0/news_bursts.xml")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, item := range feed.Items {
		if item != nil {
			counts++
		}

		if counts >= 4 {
			counts = 0
			break
		}
		fmt.Println(item.Title)
		fmt.Println("->", item.Link+"\n")

	}

	feed4, err := gofeed.NewParser().ParseURL("https://rss.itmedia.co.jp/rss/2.0/business.xml")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, item := range feed4.Items {
		if item != nil {
			counts++
		}

		if counts >= 4 {
			counts = 0
			break
		}
		fmt.Println(item.Title)
		fmt.Println("->", item.Link+"\n")

	}

	feed5, err := gofeed.NewParser().ParseURL("http://feeds.japan.cnet.com/rss/cnet/all.rdf")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, item := range feed5.Items {
		if item != nil {
			counts++
		}

		if counts >= 4 {
			counts = 0
			break
		}
		fmt.Println(item.Title)
		fmt.Println("->", item.Link+"\n")

	}

	feed2, err := gofeed.NewParser().ParseURL("https://kabumatome.doorblog.jp/index.rdf")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, item := range feed2.Items {
		if item != nil {
			counts++
		}

		if counts >= 4 {
			counts = 0
			break
		}
		fmt.Println(item.Title)
		fmt.Println("->", item.Link+"\n")

	}

	feed3, err := gofeed.NewParser().ParseURL("https://b.hatena.ne.jp/hotentry.rss")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, item := range feed3.Items {
		if item != nil {
			counts++
		}

		if counts >= 4 {
			counts = 0
			break
		}
		fmt.Println(item.Title)
		fmt.Println("->", item.Link+"\n")

	}

}
