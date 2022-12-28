package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/nats-io/nats.go"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	client := redis.NewClient(&redis.Options{Addr: "redis:6379"})

	nc, err := nats.Connect("nats")

	if err != nil {
		fmt.Println(err)
		return
	}

	nc.Subscribe("shop.pricing.exchange-rates", func(m *nats.Msg) {
		var rates map[string]float64
		err = json.Unmarshal(m.Data, &rates)

		if err != nil {
			fmt.Println(err)
			return
		}

		resp, err := http.Get("http://catalog:8000/products")

		if err != nil {
			fmt.Println(err)
			return
		}

		defer resp.Body.Close()
		productsData, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err)
			return
		}

		var products []map[string]interface{}
		err = json.Unmarshal(productsData, &products)

		if err != nil {
			fmt.Println(err)
			return
		}

		ctx := context.Background()
		for _, product := range products {
			price := product["price"].(float64)
			id := int(product["id"].(float64))

			err := client.MSet(
				ctx,
				fmt.Sprintf("%d:EUR", id), strconv.Itoa(int(price)),
				fmt.Sprintf("%d:GBP", id), strconv.Itoa(int(price*rates["GBP"])),
				fmt.Sprintf("%d:CHF", id), strconv.Itoa(int(price*rates["CHF"])),
				fmt.Sprintf("%d:SEK", id), strconv.Itoa(int(price*rates["SEK"])),
				fmt.Sprintf("%d:NOK", id), strconv.Itoa(int(price*rates["NOK"])),
				fmt.Sprintf("%d:DKK", id), strconv.Itoa(int(price*rates["DKK"])),
				fmt.Sprintf("%d:USD", id), strconv.Itoa(int(price*rates["USD"])),
				fmt.Sprintf("%d:CAD", id), strconv.Itoa(int(price*rates["CAD"])),
				fmt.Sprintf("%d:AUD", id), strconv.Itoa(int(price*rates["AUD"])),
				fmt.Sprintf("%d:JPY", id), strconv.Itoa(int(price*rates["JPY"]/100)),
				fmt.Sprintf("%d:INR", id), strconv.Itoa(int(price*rates["INR"])),
				fmt.Sprintf("%d:SGD", id), strconv.Itoa(int(price*rates["SGD"])),
				fmt.Sprintf("%d:HKD", id), strconv.Itoa(int(price*rates["HKD"])),
				fmt.Sprintf("%d:CNY", id), strconv.Itoa(int(price*rates["CNY"])),
			).Err()

			if err != nil {
				fmt.Println(err)
				return
			}
		}
	})

	wg.Wait()
}
