// Fetchall fetches URLs in parallel and reports their times and sizes.
package highloadWebserverTesting

import (
	"os"
	"encoding/json"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestArray(t *testing.T) {
	runTestAddNewUser()
}

func TestGetUsers(t *testing.T) {
	run()
}

func run() {
	start := time.Now()
	ch := make(chan string)
	url := "http://localhost/users/"
	for j := 1; j <= 3000; j++ {
		url := fmt.Sprintf("%s%d", url, j)
		go getUser(url, ch) // start a goroutine
	}
	for j := 1; j <= 10010; j++ {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func runTestAddNewUser() {
	start := time.Now()
	ch := make(chan string)
	url := "http://localhost/users/new"
	go addUser(url, ch)                             // start a goroutin
	fmt.Println(<-ch) 
	sendPostNewUser()                              // receive from channel ch
	go getUser("http://localhost/users/100200", ch) // start a goroutine
	fmt.Println(<-ch)
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
	return
}

func addUser(url string, ch chan<- string) {
	var jsonStr = []byte(`{
        "id": 100200,
        "email": "foobar@mail.ru",
        "first_name": "Маша",
        "last_name": "Пушкина",
        "gender": "f",
        "birth_date": 365299200
    }`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Charset", "utf-8")
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	ch <- string(body)
	return
}

type User struct{
	id      uint64
	email      string
	first_name      string
	last_name      string
	gender      string
    birth_date uint64
}

func sendPostNewUser(){
	

	u := User{id: 100200,
		email: "foobar@mail.ru",
		first_name: "Маша",
		last_name: "Пушкина",
		gender: "f",
		birth_date: 365299200}
    b := new(bytes.Buffer)
    json.NewEncoder(b).Encode(u)
    res, _ := http.Post("http://localhost/users/new", "application/json; charset=utf-8", b)
    io.Copy(os.Stdout, res.Body)
}

//!-
