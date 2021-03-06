package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/metarsit/exchange"
	"gopkg.in/yaml.v2"
)

func main() {
	var id, symbol, key, secret string

	flag.StringVar(&id, "id", "", "Order ID")
	flag.StringVar(&symbol, "symbol", "eosusdt", "Name of a market")
	flag.StringVar(&key, "api", "", "API Key Generated by Crypto Exchange")
	flag.StringVar(&secret, "secret", "", "Secret Key Generated by Crypto Exchange")
	flag.Parse()

	if id == "" {
		flag.Usage()
		log.Fatal("ID cannot be empty")
	}

	api, err := exchange.NewUserAPI(key, secret)
	if err != nil {
		log.Fatalf("Unable to create UserAPI Instance: %s", err.Error())
	}

	resp, err := api.ShowOrder(id, symbol)
	if err != nil {
		log.Fatalf("Unable to retrieve Account Balance: %s", err.Error())
	}

	switch resp.Code {
	case "0":
	case "22":
		log.Fatalf("[%s] Order number %s does not exist", resp.Code, id)
	default:
		log.Fatalf("[%s] API Error %s", resp.Code, resp.Message)
	}

	var data exchange.Orders
	json.Unmarshal(*resp.Data, &data)

	yamlFormat, err := yaml.Marshal(data)
	if err != nil {
		log.Fatalf("Enable to parse into YAML format: %s | %v", err.Error(), data)
	}
	fmt.Print(string(yamlFormat))
}
