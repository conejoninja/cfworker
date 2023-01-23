package main

import (
	"net/http"
	"strconv"

	"github.com/syumai/workers/cloudflare"

	"github.com/syumai/workers"
)

var (
	countStr string
	count    int32
	countKey string
)

func main() {
	countStr = "0"
	count = 0
	countKey = "count"

	kv, _ := cloudflare.NewKVNamespace("COUNTER")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		countStr, _ = kv.GetString(countKey, nil)

		count, _ = func() (int32, error) {
			i, err := strconv.Atoi(countStr)
			return int32(i), err
		}()
		countStr = strconv.Itoa(int(count + 1))
		kv.PutString(countKey, countStr, nil)

		w.Write([]byte(countStr))
	})

	workers.Serve(nil)
}
