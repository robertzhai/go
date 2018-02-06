package main 

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"encoding/json"
)

func main() {

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err )
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s : %v \n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
		var jsonData map[string]interface{}
		json.Unmarshal([]byte(b), &jsonData)
		// fmt.Fprintf(os.Stdout, "%v\n", jsonData)
		fmt.Print("---------\n")
		fmt.Println(jsonData)
		fmt.Println(jsonData["errno"])

		fmt.Println(jsonData["data"])

	}
}