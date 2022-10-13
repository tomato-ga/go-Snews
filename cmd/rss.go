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
	fmt.Println("■  " + feed.Title + "\n")
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
	fmt.Println("■  " + feed4.Title + "\n")
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
	fmt.Println("■  " + feed5.Title + "\n")
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
	fmt.Println("■  " + feed2.Title + "\n")
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
	fmt.Println("■  " + feed3.Title + "\n")
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

	feed6, err := gofeed.NewParser().ParseURL("https://qiita.com/tags/go/feed")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("■  " + feed6.Title + "\n")
	for _, item := range feed6.Items {
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

	feed7, err := gofeed.NewParser().ParseURL("https://qiita.com/tags/python/feed")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("■  " + feed7.Title + "\n")
	for _, item := range feed7.Items {
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

	feed10, err := gofeed.NewParser().ParseURL("https://assets.wor.jp/rss/rdf/nikkei/news.rdf")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("■  " + feed10.Title + "\n")
	for _, item := range feed10.Items {
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

	feed8, err := gofeed.NewParser().ParseURL("https://assets.wor.jp/rss/rdf/reuters/top.rdf")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("■  " + feed8.Title + "\n")
	for _, item := range feed8.Items {
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

	feed9, err := gofeed.NewParser().ParseURL("https://assets.wor.jp/rss/rdf/bloomberg/top.rdf")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("■  " + feed9.Title + "\n")
	for _, item := range feed9.Items {
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

	feed11, err := gofeed.NewParser().ParseURL("https://news.yahoo.co.jp/rss/categories/business.xml")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("■  " + feed11.Title + "\n")
	for _, item := range feed11.Items {
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

	feed_datascience1, err := gofeed.NewParser().ParseURL("https://datawokagaku.com/feed/")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("■ データサイエンス  " + feed_datascience1.Title + "\n")
	for _, item := range feed_datascience1.Items {
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

	feed_datascience2, err := gofeed.NewParser().ParseURL("https://toukei-lab.com/feed")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("■ データサイエンス  " + feed_datascience2.Title + "\n")
	for _, item := range feed_datascience2.Items {
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

	feed_datascience3, err := gofeed.NewParser().ParseURL("https://brainsnacks.org/feed/")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("■ データサイエンス  " + feed_datascience3.Title + "\n")
	for _, item := range feed_datascience3.Items {
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

	feed_datascience4, err := gofeed.NewParser().ParseURL("https://analysis-navi.com/?feed=rss2")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("■ データサイエンス  " + feed_datascience4.Title + "\n")
	for _, item := range feed_datascience4.Items {
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

	feed_datascience5, err := gofeed.NewParser().ParseURL("https://qiita.com/tags/%E6%A9%9F%E6%A2%B0%E5%AD%A6%E7%BF%92/feed")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("■ データサイエンス  " + feed_datascience5.Title + "\n")
	for _, item := range feed_datascience5.Items {
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
