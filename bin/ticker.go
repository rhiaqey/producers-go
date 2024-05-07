package main

import (
	"betagon/pkg/producers"
	"github.com/rhiaqey/common-go/pkg/producer"
)

func main() {
	producer.Run(&producers.Ticker{}, producers.TickerSettings{})
}
