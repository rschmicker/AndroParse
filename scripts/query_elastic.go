package main

import (
	"github.com/olivere/elastic"
	"golang.org/x/net/context"
	"io"
	"os"
)

func main() {
	client, err := elastic.NewClient()
        if err != nil {
                panic(err)
        }

	ctx := context.Background()
	f, err := os.Create("./dump.json")
	if err != nil {
                panic(err)
	}
	scroll := client.Scroll("apks").Type("_doc").Size(100)
	for {
		results, err := scroll.Do(ctx)
		if err == io.EOF {
			break// all results retrieved
		}
		for _, hit := range results.Hits.Hits[:len(results.Hits.Hits)-1] {
			f.Write(*hit.Source)
		}
	}
}
