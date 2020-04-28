package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func ESMain() {
	fmt.Println("in ESMain")
}

func ESSearch(qs string) {
	log.SetFlags(0)

	// Initialize a client with the default settings.
	//
	// An `ELASTICSEARCH_URL` environment variable will be used when exported.

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://172.17.0.38:9200",
		},
	}
	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 3. Search for the indexed documents
	//
	// Build the request body.

	var buf bytes.Buffer
	// query := map[string]interface{}{
	// 	"query": map[string]interface{}{
	// 		"query_string": map[string]interface{}{
	// 			"query": "172.*",
	// 		},
	// 	},
	// }

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query": qs,
			},
		},
	}
	/* 暂时没法又发现为什么不能使用 */
	// query := `{"query":{"bool":{"must":{"query_string":{"query":"172.*"}}}}}`

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	fmt.Println("bufString:=", buf.String())

	// // Perform the search request.
	resp, _ := es.Search(
		es.Search.WithBody(&buf),
		es.Search.WithIndex("hawkeye"),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer resp.Body.Close()

	// resultByte, _ := ioutil.ReadAll(res.Body)
	// result := string(resultByte)
	// fmt.Println("resultByte = ", result)

	if resp.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				resp.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var esResult ElasticSearchResult
	err = json.Unmarshal(data, &esResult)
	if err != nil {
		log.Fatal(err)
	}

	if esResult.Hits.Total.Value == 0 {
		return
	}

	for _, hit := range esResult.Hits.Hits {
		// im.qq.com A 1.1.1.1 enable
		re := fmt.Sprintf("%s %s %s %s", hit.Record.Fdsn, hit.Record.Type, hit.Record.Value, hit.Record.Status)

		fmt.Println(re)

	}
}
