package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Data struct {
	Data      []Crypto `json:"data"`
	Timestamp int      `json: "timestamp"`
}

type Crypto struct {
	Id                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json: "symbol"`
	Name              string `json: "name"`
	Supply            string `json: "supply"`
	MaxSupply         string `json: "maxSupply"`
	MarketCapUsd      string `json: "marketCapUsd"`
	VolumeUsd24Hr     string `json: "volumeUsd24Hr"`
	PriceUsd          string `json: "priceUsd"`
	ChangePercent24Hr string `json: "changePercent24Hr"`
	Vwap24Hr          string `json: "vwap24Hr"`
}

func (d *Crypto) Info() string {
	return fmt.Sprintf("[ID] %s | [RANK] %s | [SYMBOL] %s | [PRICE] %s", d.Id, d.Rank, d.Symbol, d.PriceUsd)
}

type myOwnLogging struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l *myOwnLogging) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.next.RoundTrip(r)
}
func main() {
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("REDIRECT")
			fmt.Println(req.Response.Status)
			return nil
		},
		Transport: &myOwnLogging{
			logger: os.Stdout,
			next:   http.DefaultTransport,
		},
	}
	resp, err := client.Get("http://api.coincap.io/v2/assets")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	var dt Data

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(body, &dt); err != nil {
		log.Fatal(err)
	}

	for _, crypto := range dt.Data {
		fmt.Println(crypto.Info())
	}

}
