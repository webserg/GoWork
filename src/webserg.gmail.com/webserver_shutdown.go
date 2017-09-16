package webserg_gmail_com

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/", homePage)
	http.ListenAndServe("localhost:8080", nil)
}
func shutdown(res http.ResponseWriter, req *http.Request) {
	fmt.Print("exit")
	os.Exit(0)
}
func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage.")
}
