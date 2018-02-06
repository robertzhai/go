package main 

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main(){
	start := time.Now()
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err )
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println(len(b))
	}
	fmt.Printf("in main %.2fs elapsed\n", time.Since(start).Seconds())

}