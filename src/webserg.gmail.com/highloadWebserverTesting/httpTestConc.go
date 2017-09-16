// Fetchall fetches URLs in parallel and reports their times and sizes.
package highloadWebserverTesting

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	url := "http://localhost/users/"
	for j := 1; j <= 100123; j++ {
		url := fmt.Sprintf("%s%d",url,j)
		go getUser(url, ch) // start a goroutine
	}
	for j := 1; j <= 100123; j++ {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func getUser(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources	
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	if resp.StatusCode != 200 {
		ch <- fmt.Sprintf("while reading %s: %d", url, resp.StatusCode)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
