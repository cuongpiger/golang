package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const timeoutMilliseconds = 5000

type barrierResp struct {
	Err  error
	Resp string
}

func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) * time.Millisecond),
	}

	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	res.Resp = string(byt)
	out <- res
}

func barrier(endpoints ...string) {
	requestNumber := len(endpoints)
	in := make(chan barrierResp, requestNumber)
	defer close(in)

	responses := make([]barrierResp, requestNumber)
	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}

	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("Error: ", resp.Err)
			hasError = true
		}

		responses[i] = resp
	}

	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}
}

func captureBarrierOutput(endpoints ...string) string {
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	out := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	barrier(endpoints...)
	writer.Close()
	temp := <-out

	return temp
}

func main() {
	endpoints := []string{"https://catfact.ninja/fact", "https://api.coindesk.com/v1/bpi/currentprice.json"}
	result := captureBarrierOutput(endpoints...)
	println(result)

	endpoints2 := []string{"https://catfact.ninja/fact", "https://api.coindesk.com/v1/bpi/currentprices.json"} // the second endpoint is incorrect
	result2 := captureBarrierOutput(endpoints2...)
	println(result2)

	endpoints3 := []string{"https://catfact.ninja/fact", "https://api.coindesk.com/v1/bpi/currentprices.json"} // the second endpoint is incorrect
	result3 := captureBarrierOutput(endpoints3...)
	println(result3)
}
