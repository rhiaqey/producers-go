package main

import (
	"github.com/rhiaqey/common-go/pkg/producer"
	"github.com/rhiaqey/producers-go/pkg/producers"
)

func main() {
	producer.Run(&producers.Ticker{}, producers.TickerSettings{})
}
