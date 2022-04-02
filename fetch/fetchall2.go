package main


import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"bufio"
)


func main() {
	var count int
	start := time.Now()
	ch := make(chan string)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		go fetch(input.Text(), ch)
		count++
	}

	for count > 0 {
		fmt.Println(<-ch)
		count--
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}


func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}
