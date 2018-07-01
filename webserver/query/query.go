package query

import(
        "github.com/olivere/elastic"
        "golang.org/x/net/context"
        "encoding/json"
        "io"
        "os"
        "strings"
        "log"
	"archive/zip"
	"golang.org/x/sync/errgroup"
)

func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func Query(filename string, fromArg string, toArg string, flds string) {
	client, err := elastic.NewClient()
        if err != nil {
                panic(err)
        }

	ctx := context.Background()

        fields := []string{}
        if len(flds) > 0 {
                fields = strings.Split(flds, ",")
        }

        rangeQuery := elastic.NewRangeQuery("Date")
        if len(fromArg) == 0 {
                rangeQuery = rangeQuery.From(nil)
        } else {
                rangeQuery = rangeQuery.From(fromArg)
        }
        if len(toArg) == 0 {
                rangeQuery = rangeQuery.To(nil)
        } else {
                rangeQuery = rangeQuery.To(toArg)
        }

        log.Println("===============================")
        log.Println("From: " + fromArg)
        log.Println("To: " + toArg)
        log.Println("Fields: ")
        log.Println(fields)
	log.Println("File name: " + filename)
        log.Println("===============================")

	jsonFilename := strings.Split(filename, ".")[0] + ".json"
        f, err := os.Create("/iscsi/queries/" + jsonFilename)
	if err != nil {
                log.Println(err)
                return
        }

	log.Println("Created json file: " + jsonFilename)

        io.WriteString(f, "{\"data\":[")
        first := true
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		scroll := client.Scroll("apks").Type("apk").Query(rangeQuery).Size(100)
		for {
			results, err := scroll.Do(ctx)
			if err == io.EOF {
				break// all results retrieved
			} else if !first {
				io.WriteString(f, ",")
			}
			first = false
			if err != nil {
				return err
			}
				// Send the hits to the hits channel
			for _, hit := range results.Hits.Hits[:len(results.Hits.Hits)-1] {
				if len(fields) == 0 {
					f.Write(*hit.Source)
					io.WriteString(f, ",")
					continue
				}
				data := make(map[string]interface{})
				out := make(map[string]interface{})
				err := json.Unmarshal(*hit.Source, &data)
				if err != nil {
					return err
				}
				for key, val := range data {
					if StringInSlice(key, fields) {
						out[key] = val
					}
				}
				writer := []byte{}
				writer, err = json.Marshal(out)
				if err != nil {
					return err
				}
				f.Write(writer)
				io.WriteString(f, ",")
			}
			data := make(map[string]interface{})
			out := make(map[string]interface{})
			err = json.Unmarshal(*results.Hits.Hits[len(results.Hits.Hits)-1].Source, &data)
			if err != nil {
				return err
			}
			for key, val := range data {
				if StringInSlice(key, fields) {
					out[key] = val
				}
			}
			writer := []byte{}
			writer, err = json.Marshal(out)
			if err != nil {
				return err
			}
			f.Write(writer)
			log.Println("Wrote data...")
		}
		return nil
	})
        if err := g.Wait(); err != nil {
		panic(err)
	}
	io.WriteString(f, "]}")
	f.Chmod(os.FileMode(int(0644)))
	f.Close()
	log.Println("Finished writing to json file")
	newfile, err := os.Create("/iscsi/queries/" + filename)
	if err != nil {
		panic(err)
	}
	defer newfile.Close()
	zipWriter := zip.NewWriter(newfile)
	defer zipWriter.Close()
	zipfile, err := os.Open("/iscsi/queries/" + jsonFilename)
        if err != nil {
		panic(err)
        }
	info, err := zipfile.Stat()
        if err != nil {
		panic(err)
        }
	header, err := zip.FileInfoHeader(info)
        if err != nil {
		panic(err)
        }
	header.Method = zip.Deflate
	writer, err := zipWriter.CreateHeader(header)
        if err != nil {
		panic(err)
        }
        _, err = io.Copy(writer, zipfile)
        if err != nil {
		panic(err)
        }
	zipfile.Close()
	err = os.Remove("/iscsi/queries/" + jsonFilename)
	if err != nil {
		panic(err)
	}
	log.Println("Finished zipping contents")
}
